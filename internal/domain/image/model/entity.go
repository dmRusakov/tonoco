package model

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dmRusakov/tonoco/pkg/common/errors"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
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
		m.mapFieldToDBColumn("Title"),
		m.mapFieldToDBColumn("AltText"),
		m.mapFieldToDBColumn("OriginPath"),
		m.mapFieldToDBColumn("FullPath"),
		m.mapFieldToDBColumn("LargePath"),
		m.mapFieldToDBColumn("MediumPath"),
		m.mapFieldToDBColumn("GridPath"),
		m.mapFieldToDBColumn("ThumbnailPath"),
		m.mapFieldToDBColumn("SortOrder"),
		m.mapFieldToDBColumn("IsWebp"),
		m.mapFieldToDBColumn("ImageType"),
		m.mapFieldToDBColumn("CreatedAt"),
		m.mapFieldToDBColumn("CreatedBy"),
		m.mapFieldToDBColumn("UpdatedAt"),
		m.mapFieldToDBColumn("UpdatedBy"),
	).From(m.table + " p")
}

// fillInFilter
func (m *Model) fillInFilter(statement sq.SelectBuilder, filter *Filter) sq.SelectBuilder {
	// Ids
	if filter.Ids != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Id"): *filter.Ids})
	}

	// IsWebp
	if filter.IsWebp != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("IsWebp"): *filter.IsWebp})
	}

	// Search
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

// makeGetStatement - make get statement by filter
func (m *Model) makeGetStatement(filter *Filter) sq.SelectBuilder {
	return m.fillInFilter(m.makeStatement(), filter)
}

func (m *Model) makeStatementByFilter(filter *Filter) sq.SelectBuilder {
	// OrderBy
	if filter.OrderBy == nil {
		filter.OrderBy = entity.StringPtr("SortOrder")
	}

	// OrderDir
	if filter.OrderDir == nil {
		filter.OrderDir = entity.StringPtr("ASC")
	}

	// PerPage
	if filter.PerPage == nil {
		if filter.Page == nil {
			filter.PerPage = entity.Uint64Ptr(999999999999999999)
		} else {
			filter.PerPage = entity.Uint64Ptr(10)
		}
	}

	// Page
	if filter.Page == nil {
		filter.Page = entity.Uint64Ptr(1)
	}

	// Build query
	statement := m.makeGetStatement(filter)

	// Add OrderBy, OrderDir, Page, Limit and return
	return statement.OrderBy(m.mapFieldToDBColumn(*filter.OrderBy) + " " + *filter.OrderDir).
		Offset((*filter.Page - 1) * *filter.PerPage).Limit(*filter.PerPage)
}

func (m *Model) makeCountStatementByFilter(filter *Filter) sq.SelectBuilder {
	return m.fillInFilter(m.qb.Select("COUNT(*)").From(m.table), filter)
}

func (m *Model) scanOneRow(ctx context.Context, rows sq.RowScanner) (*Item, error) {
	var id, title, altText, originPath, fullPath, largePath, mediumPath, gridPath, thumbnailPath, imageType, createdBy, updatedBy sql.NullString
	var isWebp sql.NullBool
	var sortOrder sql.NullInt64
	var createdAt, updatedAt sql.NullTime

	err := rows.Scan(
		&id,
		&title,
		&altText,
		&originPath,
		&fullPath,
		&largePath,
		&mediumPath,
		&gridPath,
		&thumbnailPath,
		&sortOrder,
		&isWebp,
		&imageType,
		&createdAt,
		&createdBy,
		&updatedAt,
		&updatedBy,
	)

	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, errors.AddCode(err, "853456")
	}

	var item = Item{}

	if id.Valid {
		item.Id = uuid.MustParse(id.String)
	}

	if title.Valid {
		item.Title = title.String
	}

	if altText.Valid {
		item.AltText = altText.String
	}

	if originPath.Valid {
		item.OriginPath = originPath.String
	}

	if fullPath.Valid {
		item.FullPath = fullPath.String
	}

	if largePath.Valid {
		item.LargePath = largePath.String
	}

	if mediumPath.Valid {
		item.MediumPath = mediumPath.String
	}

	if gridPath.Valid {
		item.GridPath = gridPath.String
	}

	if thumbnailPath.Valid {
		item.ThumbnailPath = thumbnailPath.String
	}

	if sortOrder.Valid {
		item.SortOrder = uint64(sortOrder.Int64)
	}

	if isWebp.Valid {
		item.IsWebp = isWebp.Bool
	}

	if imageType.Valid {
		item.ImageType = imageType.String
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
	// get user_id from context
	by := ctx.Value("user_id").(string)

	// if Id is not set, generate a new UUID
	if item.Id == uuid.Nil {
		item.Id = uuid.New()
	}

	// set Id to context
	ctx = context.WithValue(ctx, "itemId", item.Id)

	insertItem := m.qb.Insert(m.table).Columns(
		m.mapFieldToDBColumn("Id"),
		m.mapFieldToDBColumn("Title"),
		m.mapFieldToDBColumn("AltText"),
		m.mapFieldToDBColumn("OriginPath"),
		m.mapFieldToDBColumn("FullPath"),
		m.mapFieldToDBColumn("LargePath"),
		m.mapFieldToDBColumn("MediumPath"),
		m.mapFieldToDBColumn("GridPath"),
		m.mapFieldToDBColumn("ThumbnailPath"),
		m.mapFieldToDBColumn("SortOrder"),
		m.mapFieldToDBColumn("IsWebp"),
		m.mapFieldToDBColumn("ImageType"),
		m.mapFieldToDBColumn("ProductID"),
		m.mapFieldToDBColumn("CreatedAt"),
		m.mapFieldToDBColumn("CreatedBy"),
		m.mapFieldToDBColumn("UpdatedAt"),
		m.mapFieldToDBColumn("UpdatedBy"),
	).Values(
		item.Id,
		item.Title,
		item.AltText,
		item.OriginPath,
		item.FullPath,
		item.LargePath,
		item.MediumPath,
		item.GridPath,
		item.ThumbnailPath,
		item.SortOrder,
		item.IsWebp,
		item.ImageType,
		"NOW()",
		by,
		"NOW()",
		by,
	)

	// get itemId from context
	return &insertItem, &item.Id
}

func (m *Model) makeUpdateStatement(ctx context.Context, item *Item) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	return m.qb.Update(m.table).
		Set(m.mapFieldToDBColumn("Title"), item.Title).
		Set(m.mapFieldToDBColumn("AltText"), item.AltText).
		Set(m.mapFieldToDBColumn("OriginPath"), item.OriginPath).
		Set(m.mapFieldToDBColumn("FullPath"), item.FullPath).
		Set(m.mapFieldToDBColumn("LargePath"), item.LargePath).
		Set(m.mapFieldToDBColumn("MediumPath"), item.MediumPath).
		Set(m.mapFieldToDBColumn("GridPath"), item.GridPath).
		Set(m.mapFieldToDBColumn("ThumbnailPath"), item.ThumbnailPath).
		Set(m.mapFieldToDBColumn("SortOrder"), item.SortOrder).
		Set(m.mapFieldToDBColumn("IsWebp"), item.IsWebp).
		Set(m.mapFieldToDBColumn("ImageType"), item.ImageType).
		Set(m.mapFieldToDBColumn("UpdatedAt"), "NOW()").
		Set(m.mapFieldToDBColumn("UpdatedBy"), by)
}

func (m *Model) makePatchStatement(ctx context.Context, id *uuid.UUID, fields *map[string]interface{}) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	statement := m.qb.Update(m.table).Where("id = ?", id)

	for field, value := range *fields {
		statement = statement.Set(m.mapFieldToDBColumn(field), value)
	}

	return statement.Set(m.mapFieldToDBColumn("UpdatedAt"), "NOW()").Set(m.mapFieldToDBColumn("UpdatedBy"), by)
}
