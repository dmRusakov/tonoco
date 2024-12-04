package model

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dmRusakov/tonoco/pkg/common/errors"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"github.com/dmRusakov/tonoco/pkg/utils/pointer"
	"github.com/google/uuid"
)

func (m *Model) mapFieldToDBColumn(field string) string {
	return m.dbField[field]
}

func (m *Model) makeStatement() sq.SelectBuilder {
	return m.qb.Select(
		m.mapFieldToDBColumn("Id"),
		m.mapFieldToDBColumn("Name"),
		m.mapFieldToDBColumn("SeoTitle"),
		m.mapFieldToDBColumn("ShortDescription"),
		m.mapFieldToDBColumn("Description"),
		m.mapFieldToDBColumn("Url"),
		m.mapFieldToDBColumn("ImageId"),
		m.mapFieldToDBColumn("HoverImageId"),
		m.mapFieldToDBColumn("Page"),
		m.mapFieldToDBColumn("PerPage"),
		m.mapFieldToDBColumn("SortOrder"),
		m.mapFieldToDBColumn("Active"),
		m.mapFieldToDBColumn("Prime"),
		m.mapFieldToDBColumn("CreatedAt"),
		m.mapFieldToDBColumn("CreatedBy"),
		m.mapFieldToDBColumn("UpdatedAt"),
		m.mapFieldToDBColumn("UpdatedBy"),
	).From(m.table)
}

func (m *Model) makeIdsStatement() sq.SelectBuilder {
	return m.qb.Select(
		m.mapFieldToDBColumn("Id"),
	).From(m.table)
}

func (m *Model) filterToStatement(statement sq.SelectBuilder, filter *Filter) sq.SelectBuilder {
	// Ids
	if filter.Ids != nil && len(*filter.Ids) > 0 {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Id"): *filter.Ids})
	}

	// Urls
	if filter.Urls != nil && len(*filter.Urls) > 0 {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Url"): *filter.Urls})
	}

	// Active
	if filter.Active != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Active"): *filter.Active})
	}

	// Prime
	if filter.Prime != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Prime"): *filter.Prime})
	}

	return statement
}

func (m *Model) makeGetStatement(filter *Filter) sq.SelectBuilder {
	// build query
	statement := m.makeStatement()

	// id
	if filter.Ids != nil {
		statement = statement.Where(m.mapFieldToDBColumn("Id")+" = ?", (*filter.Ids)[0])
	}

	// url
	if filter.Urls != nil {
		statement = statement.Where(m.mapFieldToDBColumn("Url")+" = ?", (*filter.Urls)[0])
	}

	return statement
}

func (m *Model) makeStatementByFilter(statement sq.SelectBuilder, filter *Filter) sq.SelectBuilder {

	entity.CheckDataPagination(filter.DataPagination)
	statement = m.filterToStatement(statement, filter)
	return statement.OrderBy(m.mapFieldToDBColumn(*filter.DataPagination.OrderBy) + " " + *filter.DataPagination.OrderDir).
		Offset((*filter.DataPagination.Page - 1) * *filter.DataPagination.PerPage).Limit(*filter.DataPagination.PerPage)
}

func (m *Model) makeCountStatementByFilter(filter *Filter) sq.SelectBuilder {
	statement := m.qb.Select("COUNT(*) as count").From(m.table + " p")
	return m.filterToStatement(statement, filter)
}

func (m *Model) scanRow(ctx context.Context, row sq.RowScanner) (*Item, error) {
	var (
		item                                                                                                = Item{}
		id, name, seoTitle, shortDescription, description, url, imageId, hoverImageId, createdBy, updatedBy sql.NullString
		page, perPage, sortOrder                                                                            sql.NullInt64
		active, prime                                                                                       sql.NullBool
		createdAt, updatedAt                                                                                sql.NullTime
	)

	err := row.Scan(&id, &name, &seoTitle, &shortDescription, &description, &url, &imageId, &hoverImageId, &page, &perPage, &sortOrder, &active, &prime, &createdAt, &createdBy, &updatedAt, &updatedBy)
	if err != nil {
		tracing.Error(ctx, psql.ErrScan(psql.ParsePgError(err)))
		return nil, errors.AddCode(err, "396647")
	}

	if id.Valid {
		item.Id = uuid.MustParse(id.String)
	}

	if name.Valid {
		item.Name = uuid.MustParse(name.String)
	}

	if seoTitle.Valid {
		item.SeoTitle = uuid.MustParse(seoTitle.String)
	}

	if shortDescription.Valid {
		item.ShortDescription = uuid.MustParse(shortDescription.String)
	}

	if description.Valid {
		item.Description = uuid.MustParse(description.String)
	}

	if url.Valid {
		item.Url = url.String
	}

	if imageId.Valid {
		item.ImageId = uuid.MustParse(imageId.String)
	}

	if hoverImageId.Valid {
		item.HoverImageId = uuid.MustParse(hoverImageId.String)
	}

	if page.Valid {
		item.Page = uint64(page.Int64)
	}

	if perPage.Valid {
		item.PerPage = uint64(perPage.Int64)
	}

	if sortOrder.Valid {
		item.SortOrder = int(sortOrder.Int64)
	}

	if active.Valid {
		item.Active = active.Bool
	}

	if prime.Valid {
		item.Prime = prime.Bool
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

func (m *Model) scanIdRow(ctx context.Context, row sq.RowScanner) (*uuid.UUID, error) {
	var id sql.NullString

	err := row.Scan(&id)
	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, errors.AddCode(err, "227912")
	}

	return pointer.StringToUUID(id.String), nil
}

func (m *Model) scanCountRow(ctx context.Context, row sq.RowScanner) (*uint64, error) {
	var count uint64

	err := row.Scan(&count)
	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, errors.AddCode(err, "491602")
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

	insertItem := m.qb.Insert(m.table).
		Columns(
			m.mapFieldToDBColumn("Id"),
			m.mapFieldToDBColumn("Name"),
			m.mapFieldToDBColumn("SeoTitle"),
			m.mapFieldToDBColumn("ShortDescription"),
			m.mapFieldToDBColumn("Description"),
			m.mapFieldToDBColumn("Url"),
			m.mapFieldToDBColumn("ImageId"),
			m.mapFieldToDBColumn("HoverImageId"),
			m.mapFieldToDBColumn("Page"),
			m.mapFieldToDBColumn("PerPage"),
			m.mapFieldToDBColumn("SortOrder"),
			m.mapFieldToDBColumn("Active"),
			m.mapFieldToDBColumn("Prime"),
			m.mapFieldToDBColumn("CreatedAt"),
			m.mapFieldToDBColumn("CreatedBy"),
			m.mapFieldToDBColumn("UpdatedAt"),
			m.mapFieldToDBColumn("UpdatedBy"),
		).
		Values(
			item.Id,
			item.Name,
			item.SeoTitle,
			item.ShortDescription,
			item.Description,
			item.Url,
			item.ImageId,
			item.HoverImageId,
			item.Page,
			item.PerPage,
			item.SortOrder,
			item.Active,
			item.Prime,
			"NOW()",
			by,
			"NOW()",
			by,
		)

	return &insertItem, &item.Id
}

func (m *Model) makeUpdateStatement(ctx context.Context, item *Item) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	return m.qb.Update(m.table).
		SetMap(map[string]interface{}{
			m.mapFieldToDBColumn("Name"):             item.Name,
			m.mapFieldToDBColumn("SeoTitle"):         item.SeoTitle,
			m.mapFieldToDBColumn("ShortDescription"): item.ShortDescription,
			m.mapFieldToDBColumn("Description"):      item.Description,
			m.mapFieldToDBColumn("Url"):              item.Url,
			m.mapFieldToDBColumn("ImageId"):          item.ImageId,
			m.mapFieldToDBColumn("HoverImageId"):     item.HoverImageId,
			m.mapFieldToDBColumn("Page"):             item.Page,
			m.mapFieldToDBColumn("PerPage"):          item.PerPage,
			m.mapFieldToDBColumn("SortOrder"):        item.SortOrder,
			m.mapFieldToDBColumn("Active"):           item.Active,
			m.mapFieldToDBColumn("Prime"):            item.Prime,
			m.mapFieldToDBColumn("UpdatedAt"):        "NOW()",
			m.mapFieldToDBColumn("UpdatedBy"):        by,
		}).Where(m.mapFieldToDBColumn("Id")+" = ?", item.Id)
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
