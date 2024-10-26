package model

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/pkg/common/errors"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"github.com/dmRusakov/tonoco/pkg/utils/pointer"
	"github.com/google/uuid"
	"reflect"
)

func (m *Model) mapFieldToDBColumn(field string) string {
	m.mu.Lock()
	defer m.mu.Unlock()

	// check if field is in the cash
	if dbField, ok := m.dbFieldCash[field]; ok {
		return dbField
	}

	// get field from struct
	typeOf := reflect.TypeOf(Item{})
	byName, _ := typeOf.FieldByName(field)
	dbField := byName.Tag.Get("db")

	// set field to cash
	m.dbFieldCash[field] = dbField

	// done
	return dbField
}

func (m *Model) makeStatement() sq.SelectBuilder {
	return m.qb.Select(
		m.mapFieldToDBColumn("Id"),
		m.mapFieldToDBColumn("FileName"),
		m.mapFieldToDBColumn("Extension"),
		m.mapFieldToDBColumn("IsCompressed"),
		m.mapFieldToDBColumn("IsWebp"),
		m.mapFieldToDBColumn("FolderId"),
		m.mapFieldToDBColumn("SortOrder"),
		m.mapFieldToDBColumn("Title"),
		m.mapFieldToDBColumn("AltText"),
		m.mapFieldToDBColumn("CopyRight"),
		m.mapFieldToDBColumn("Creator"),
		m.mapFieldToDBColumn("Rating"),
		m.mapFieldToDBColumn("OriginPath"),
		m.mapFieldToDBColumn("CreatedAt"),
		m.mapFieldToDBColumn("CreatedBy"),
		m.mapFieldToDBColumn("UpdatedAt"),
		m.mapFieldToDBColumn("UpdatedBy"),
	).From(m.table + " p")
}

func (m *Model) fillInFilter(statement sq.SelectBuilder, filter *Filter) sq.SelectBuilder {
	if filter.Ids != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Id"): *filter.Ids})
	}

	if filter.IsWebp != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("IsWebp"): *filter.IsWebp})
	}

	if filter.Search != nil {
		statement = statement.Where(
			sq.Or{
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Title")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("AltText")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
			},
		)
	}

	return statement
}

func (m *Model) makeGetStatement(filter *Filter) sq.SelectBuilder {
	return m.fillInFilter(m.makeStatement(), filter)
}

func (m *Model) makeStatementByFilter(filter *Filter) sq.SelectBuilder {
	if filter.OrderBy == nil {
		filter.OrderBy = pointer.StringPtr("SortOrder")
	}

	if filter.OrderDir == nil {
		filter.OrderDir = pointer.StringPtr("ASC")
	}

	if filter.PerPage == nil {
		if filter.Page == nil {
			filter.PerPage = pointer.Uint64Ptr(999999999999999999)
		} else {
			filter.PerPage = pointer.Uint64Ptr(10)
		}
	}

	if filter.Page == nil {
		filter.Page = pointer.Uint64Ptr(1)
	}

	statement := m.makeGetStatement(filter)

	return statement.OrderBy(m.mapFieldToDBColumn(*filter.OrderBy) + " " + *filter.OrderDir).
		Offset((*filter.Page - 1) * *filter.PerPage).Limit(*filter.PerPage)
}

func (m *Model) makeCountStatementByFilter(filter *Filter) sq.SelectBuilder {
	return m.fillInFilter(m.qb.Select("COUNT(*)").From(m.table), filter)
}

func (m *Model) scanOneRow(ctx context.Context, rows sq.RowScanner) (*Item, error) {
	var id, fileName, extension, title, altText, copyRight, creator, originPath sql.NullString
	var isCompressed, isWebp sql.NullBool
	var folderId sql.NullString
	var sortOrder sql.NullInt64
	var rating sql.NullFloat64
	var createdAt, updatedAt sql.NullTime
	var createdBy, updatedBy sql.NullString

	err := rows.Scan(
		&id,
		&fileName,
		&extension,
		&isCompressed,
		&isWebp,
		&folderId,
		&sortOrder,
		&title,
		&altText,
		&copyRight,
		&creator,
		&rating,
		&originPath,
		&createdAt,
		&createdBy,
		&updatedAt,
		&updatedBy,
	)

	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, errors.AddCode(err, "616229")
	}

	var item = Item{}

	if id.Valid {
		item.Id = uuid.MustParse(id.String)
	}

	if fileName.Valid {
		item.FileName = fileName.String
	}

	if extension.Valid {
		item.Extension = extension.String
	}

	if isCompressed.Valid {
		item.IsCompressed = isCompressed.Bool
	}

	if isWebp.Valid {
		item.IsWebp = isWebp.Bool
	}

	if folderId.Valid {
		item.FolderId = uuid.MustParse(folderId.String)
	}

	if sortOrder.Valid {
		item.SortOrder = uint64(sortOrder.Int64)
	}

	if title.Valid {
		item.Title = title.String
	}

	if altText.Valid {
		item.AltText = altText.String
	}

	if copyRight.Valid {
		item.CopyRight = copyRight.String
	}

	if creator.Valid {
		item.Creator = creator.String
	}

	if rating.Valid {
		item.Rating = float32(rating.Float64)
	}

	if originPath.Valid {
		item.OriginPath = originPath.String
	}

	if createdAt.Valid {
		item.CreatedAt = createdAt.Time
	}

	if createdBy.Valid {
		item.CreatedBy = uuid.MustParse(createdBy.String)
	}

	if updatedAt.Valid {
		item.UpdatedAt = updatedAt.Time
	}

	if updatedBy.Valid {
		item.UpdatedBy = uuid.MustParse(updatedBy.String)
	}

	return &item, nil
}

func (m *Model) scanCountRow(ctx context.Context, rows sq.RowScanner) (*uint64, error) {
	var count uint64

	err := rows.Scan(&count)
	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, err
	}

	return &count, nil
}

func (m *Model) makeInsertStatement(ctx context.Context, item *Item) (*sq.InsertBuilder, *uuid.UUID) {
	by := ctx.Value("user_id").(string)

	if item.Id == uuid.Nil {
		item.Id = uuid.New()
	}

	ctx = context.WithValue(ctx, "itemId", item.Id)

	insertItem := m.qb.Insert(m.table).Columns(
		m.mapFieldToDBColumn("Id"),
		m.mapFieldToDBColumn("FileName"),
		m.mapFieldToDBColumn("Extension"),
		m.mapFieldToDBColumn("IsCompressed"),
		m.mapFieldToDBColumn("IsWebp"),
		m.mapFieldToDBColumn("FolderId"),
		m.mapFieldToDBColumn("SortOrder"),
		m.mapFieldToDBColumn("Title"),
		m.mapFieldToDBColumn("AltText"),
		m.mapFieldToDBColumn("CopyRight"),
		m.mapFieldToDBColumn("Creator"),
		m.mapFieldToDBColumn("Rating"),
		m.mapFieldToDBColumn("OriginPath"),
		m.mapFieldToDBColumn("CreatedAt"),
		m.mapFieldToDBColumn("CreatedBy"),
		m.mapFieldToDBColumn("UpdatedAt"),
		m.mapFieldToDBColumn("UpdatedBy"),
	).Values(
		item.Id,
		item.FileName,
		item.Extension,
		item.IsCompressed,
		item.IsWebp,
		item.FolderId,
		item.SortOrder,
		item.Title,
		item.AltText,
		item.CopyRight,
		item.Creator,
		item.Rating,
		item.OriginPath,
		"NOW()",
		by,
		"NOW()",
		by,
	)

	return &insertItem, &item.Id
}

func (m *Model) makeUpdateStatement(ctx context.Context, item *Item) sq.UpdateBuilder {
	by := ctx.Value("user_id").(string)

	return m.qb.Update(m.table).
		Set(m.mapFieldToDBColumn("FileName"), item.FileName).
		Set(m.mapFieldToDBColumn("Extension"), item.Extension).
		Set(m.mapFieldToDBColumn("IsCompressed"), item.IsCompressed).
		Set(m.mapFieldToDBColumn("IsWebp"), item.IsWebp).
		Set(m.mapFieldToDBColumn("FolderId"), item.FolderId).
		Set(m.mapFieldToDBColumn("SortOrder"), item.SortOrder).
		Set(m.mapFieldToDBColumn("Title"), item.Title).
		Set(m.mapFieldToDBColumn("AltText"), item.AltText).
		Set(m.mapFieldToDBColumn("CopyRight"), item.CopyRight).
		Set(m.mapFieldToDBColumn("Creator"), item.Creator).
		Set(m.mapFieldToDBColumn("Rating"), item.Rating).
		Set(m.mapFieldToDBColumn("OriginPath"), item.OriginPath).
		Set(m.mapFieldToDBColumn("UpdatedAt"), "NOW()").
		Set(m.mapFieldToDBColumn("UpdatedBy"), by)
}

func (m *Model) makePatchStatement(ctx context.Context, id *uuid.UUID, fields *map[string]interface{}) sq.UpdateBuilder {
	by := ctx.Value("user_id").(string)

	statement := m.qb.Update(m.table).Where("id = ?", id)

	for field, value := range *fields {
		statement = statement.Set(m.mapFieldToDBColumn(field), value)
	}

	return statement.Set(m.mapFieldToDBColumn("UpdatedAt"), "NOW()").Set(m.mapFieldToDBColumn("UpdatedBy"), by)
}
