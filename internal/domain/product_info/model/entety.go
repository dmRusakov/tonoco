package model

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/internal/entity"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"github.com/google/uuid"
	"reflect"
)

type Item = entity.ProductInfo
type Filter = entity.ProductInfoFilter

func (m *Model) fieldMap(field string) string {
	typeOf := reflect.TypeOf(Item{})
	byName, _ := typeOf.FieldByName(field)
	return byName.Tag.Get("db")
}

// makeStatement
func (m *Model) makeStatement() sq.SelectBuilder {
	return m.qb.Select(
		m.fieldMap("ID"),
		m.fieldMap("SKU"),
		m.fieldMap("Name"),
		m.fieldMap("ShortDescription"),
		m.fieldMap("Description"),
		m.fieldMap("SortOrder"),
		m.fieldMap("Url"),
		m.fieldMap("IsTaxable"),
		m.fieldMap("IsTrackStock"),
		m.fieldMap("ShippingWeight"),
		m.fieldMap("ShippingWidth"),
		m.fieldMap("ShippingHeight"),
		m.fieldMap("ShippingLength"),
		m.fieldMap("SeoTitle"),
		m.fieldMap("SeoDescription"),
		m.fieldMap("GTIN"),
		m.fieldMap("GoogleProductCategory"),
		m.fieldMap("GoogleProductType"),
		m.fieldMap("CreatedAt"),
		m.fieldMap("CreatedBy"),
		m.fieldMap("UpdatedAt"),
		m.fieldMap("UpdatedBy"),
	).From(m.table + " p")
}

// make Get statement
func (m *Model) makeGetStatement(id *string, url *string) sq.SelectBuilder {
	// build query
	statement := m.makeStatement()

	// id
	if id != nil {
		statement = statement.Where(m.fieldMap("ID")+" = ?", *id)
	}

	// url
	if url != nil {
		statement = statement.Where(m.fieldMap("Url")+" = ?", *url)
	}

	return statement
}

// makeStatementByFilter
func (m *Model) makeStatementByFilter(filter *Filter) sq.SelectBuilder {
	// OrderBy
	if filter.OrderBy == nil {
		filter.OrderBy = new(string)
		*filter.OrderBy = "SortOrder"
	}

	// OrderDir
	if filter.OrderDir == nil {
		filter.OrderDir = new(string)
		*filter.OrderDir = "ASC"
	}

	// PerPage
	if filter.PerPage == nil {
		filter.PerPage = new(uint64)
		if filter.Page == nil {
			*filter.PerPage = 999999999999999999
		} else {
			*filter.PerPage = 10
		}
	}

	// Page
	if filter.Page == nil {
		filter.Page = new(uint64)
		*filter.Page = 1
	}

	// Build query
	statement := m.makeStatement()

	// Ids
	if filter.IDs != nil && len(*filter.IDs) > 0 {
		statement = statement.Where(sq.Eq{m.fieldMap("ID"): *filter.IDs})

		*filter.Page = 1
		if (*filter.PerPage) > uint64(len(*filter.IDs)) {
			*filter.PerPage = uint64(len(*filter.IDs))
		}
	}

	// Urls
	if filter.Urls != nil && len(*filter.Urls) > 0 {
		statement = statement.Where(sq.Eq{m.fieldMap("Url"): *filter.Urls})

		*filter.Page = 1
		if (*filter.PerPage) > uint64(len(*filter.Urls)) {
			*filter.PerPage = uint64(len(*filter.Urls))
		}
	}

	// SKUs
	if filter.SKUs != nil {
		countSKUs := len(*filter.SKUs)

		if countSKUs > 0 {
			statement = statement.Where(sq.Eq{m.fieldMap("SKU"): *filter.SKUs})
		}

		*filter.Page = 1
		if (*filter.PerPage) > uint64(countSKUs) {
			*filter.PerPage = uint64(countSKUs)
		}

	}

	// Search
	if filter.Search != nil {
		statement = statement.Where(
			sq.Or{
				sq.Expr("LOWER("+m.fieldMap("Name")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.fieldMap("Url")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.fieldMap("ShortDescription")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.fieldMap("Description")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.fieldMap("SeoTitle")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.fieldMap("SeoDescription")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
			},
		)
	}

	statement = statement.OrderBy(m.fieldMap(*filter.OrderBy) + " " + *filter.OrderDir).
		Offset((*filter.Page - 1) * *filter.PerPage).Limit(*filter.PerPage)

	// Add OrderBy, OrderDir, Page, Limit and return
	return statement
}

// scanOneRow
func (m *Model) scanOneRow(ctx context.Context, rows sq.RowScanner) (*Item, error) {
	var item = Item{}
	var id, sku, name, shortDescription, description, url, seoTitle, seoDescription, gtin, googleProductCategory, googleProductType, createdBy, updatedBy sql.NullString
	var sortOrder sql.NullInt64
	var isTaxable, isTrackStock sql.NullBool
	var shippingWeight, shippingWidth, shippingHeight, shippingLength sql.NullInt64
	var createdAt, updatedAt sql.NullTime

	err := rows.Scan(
		&id,
		&sku,
		&name,
		&shortDescription,
		&description,
		&sortOrder,
		&url,
		&isTaxable,
		&isTrackStock,
		&shippingWeight,
		&shippingWidth,
		&shippingHeight,
		&shippingLength,
		&seoTitle,
		&seoDescription,
		&gtin,
		&googleProductCategory,
		&googleProductType,
		&createdAt,
		&createdBy,
		&updatedAt,
		&updatedBy,
	)

	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, err
	}

	if id.Valid {
		item.ID = id.String
	}
	if sku.Valid {
		item.SKU = sku.String
	}
	if name.Valid {
		item.Name = name.String
	}
	if shortDescription.Valid {
		item.ShortDescription = shortDescription.String
	}
	if description.Valid {
		item.Description = description.String
	}
	if sortOrder.Valid {
		item.SortOrder = uint64(sortOrder.Int64)
	}
	if url.Valid {
		item.Url = url.String
	}
	if isTaxable.Valid {
		item.IsTaxable = isTaxable.Bool
	}
	if isTrackStock.Valid {
		item.IsTrackStock = isTrackStock.Bool
	}
	if shippingWeight.Valid {
		item.ShippingWeight = uint64(shippingWeight.Int64)
	}
	if shippingWidth.Valid {
		item.ShippingWidth = uint64(shippingWidth.Int64)
	}
	if shippingHeight.Valid {
		item.ShippingHeight = uint64(shippingHeight.Int64)
	}
	if shippingLength.Valid {
		item.ShippingLength = uint64(shippingLength.Int64)
	}
	if seoTitle.Valid {
		item.SeoTitle = seoTitle.String
	}
	if seoDescription.Valid {
		item.SeoDescription = seoDescription.String
	}
	if gtin.Valid {
		item.GTIN = gtin.String
	}
	if googleProductCategory.Valid {
		item.GoogleProductCategory = googleProductCategory.String
	}
	if googleProductType.Valid {
		item.GoogleProductType = googleProductType.String
	}
	if createdAt.Valid {
		item.CreatedAt = createdAt.Time
	}
	if createdBy.Valid {
		item.CreatedBy = createdBy.String
	}
	if updatedAt.Valid {
		item.UpdatedAt = updatedAt.Time
	}
	if updatedBy.Valid {
		item.UpdatedBy = updatedBy.String
	}

	return &item, nil
}

// makeInsertStatement
func (m *Model) makeInsertStatement(ctx context.Context, item *Item) (*sq.InsertBuilder, *string) {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	// if ID is not set, generate a new UUID
	if item.ID == "" {
		item.ID = uuid.New().String()
	}

	// set ID to context
	ctx = context.WithValue(ctx, "itemId", item.ID)

	insertItem := m.qb.Insert(m.table).Columns(
		m.fieldMap("ID"),
		m.fieldMap("SKU"),
		m.fieldMap("Name"),
		m.fieldMap("ShortDescription"),
		m.fieldMap("Description"),
		m.fieldMap("SortOrder"),
		m.fieldMap("Url"),
		m.fieldMap("IsTaxable"),
		m.fieldMap("IsTrackStock"),
		m.fieldMap("ShippingWeight"),
		m.fieldMap("ShippingWidth"),
		m.fieldMap("ShippingHeight"),
		m.fieldMap("ShippingLength"),
		m.fieldMap("SeoTitle"),
		m.fieldMap("SeoDescription"),
		m.fieldMap("GTIN"),
		m.fieldMap("GoogleProductCategory"),
		m.fieldMap("GoogleProductType"),
		m.fieldMap("CreatedAt"),
		m.fieldMap("CreatedBy"),
		m.fieldMap("UpdatedAt"),
		m.fieldMap("UpdatedBy"),
	).Values(
		item.ID,
		item.SKU,
		item.Name,
		item.ShortDescription,
		item.Description,
		item.SortOrder,
		item.Url,
		item.IsTaxable,
		item.IsTrackStock,
		item.ShippingWeight,
		item.ShippingWidth,
		item.ShippingHeight,
		item.ShippingLength,
		item.SeoTitle,
		item.SeoDescription,
		item.GTIN,
		item.GoogleProductCategory,
		item.GoogleProductType,
		"NOW()",
		by,
		"NOW()",
		by,
	)

	return &insertItem, &item.ID
}

// makeUpdateStatement
func (m *Model) makeUpdateStatement(ctx context.Context, item *Item) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	return m.qb.Update(m.table).
		Set(m.fieldMap("SKU"), item.SKU).
		Set(m.fieldMap("Name"), item.Name).
		Set(m.fieldMap("ShortDescription"), item.ShortDescription).
		Set(m.fieldMap("Description"), item.Description).
		Set(m.fieldMap("SortOrder"), item.SortOrder).
		Set(m.fieldMap("Url"), item.Url).
		Set(m.fieldMap("IsTaxable"), item.IsTaxable).
		Set(m.fieldMap("IsTrackStock"), item.IsTrackStock).
		Set(m.fieldMap("ShippingWeight"), item.ShippingWeight).
		Set(m.fieldMap("ShippingWidth"), item.ShippingWidth).
		Set(m.fieldMap("ShippingHeight"), item.ShippingHeight).
		Set(m.fieldMap("ShippingLength"), item.ShippingLength).
		Set(m.fieldMap("SeoTitle"), item.SeoTitle).
		Set(m.fieldMap("SeoDescription"), item.SeoDescription).
		Set(m.fieldMap("GTIN"), item.GTIN).
		Set(m.fieldMap("GoogleProductCategory"), item.GoogleProductCategory).
		Set(m.fieldMap("GoogleProductType"), item.GoogleProductType).
		Set(m.fieldMap("UpdatedAt"), "NOW()").
		Set(m.fieldMap("UpdatedBy"), by).
		Where(m.fieldMap("ID")+" = ?", item.ID)
}

// makePatchStatement
func (m *Model) makePatchStatement(ctx context.Context, id *string, fields *map[string]interface{}) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	statement := m.qb.Update(m.table).Where("id = ?", id)

	for field, value := range *fields {
		field = m.fieldMap(field)
		statement = statement.Set(field, value)
	}

	return statement.Set(m.fieldMap("UpdatedAt"), "NOW()").Set(m.fieldMap("UpdatedBy"), by)
}
