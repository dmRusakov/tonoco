package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/internal/entity"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"github.com/google/uuid"
)

type Item = entity.ProductInfo
type Filter = entity.ProductInfoFilter

// fieldMap
var fieldMap = map[string]string{
	"ID":                    "id",
	"SKU":                   "sku",
	"Name":                  "name",
	"ShortDescription":      "short_description",
	"Description":           "description",
	"SortOrder":             "sort_order",
	"StatusID":              "status_id",
	"Url":                   "url",
	"IsTaxable":             "is_taxable",
	"ReturnToStockDate":     "return_to_stock_date",
	"IsTrackStock":          "is_track_stock",
	"ShippingClassID":       "shipping_class_id",
	"ShippingWeight":        "shipping_weight",
	"ShippingWidth":         "shipping_width",
	"ShippingHeight":        "shipping_height",
	"ShippingLength":        "shipping_length",
	"SeoTitle":              "seo_title",
	"SeoDescription":        "seo_description",
	"GTIN":                  "gtin",
	"GoogleProductCategory": "google_product_category",
	"GoogleProductType":     "google_product_type",
	"CreatedAt":             "created_at",
	"CreatedBy":             "created_by",
	"UpdatedAt":             "updated_at",
	"UpdatedBy":             "updated_by",
}

// makeStatement
func (m *Model) makeStatement() sq.SelectBuilder {
	return m.qb.Select(
		fieldMap["ID"],
		fieldMap["SKU"],
		fieldMap["Name"],
		fieldMap["ShortDescription"],
		fieldMap["Description"],
		fieldMap["SortOrder"],
		fieldMap["StatusID"],
		fieldMap["Url"],
		fieldMap["IsTaxable"],
		fieldMap["IsTrackStock"],
		fieldMap["ShippingClassID"],
		fieldMap["ShippingWeight"],
		fieldMap["ShippingWidth"],
		fieldMap["ShippingHeight"],
		fieldMap["ShippingLength"],
		fieldMap["SeoTitle"],
		fieldMap["SeoDescription"],
		fieldMap["GTIN"],
		fieldMap["GoogleProductCategory"],
		fieldMap["GoogleProductType"],
		fieldMap["CreatedAt"],
		fieldMap["CreatedBy"],
		fieldMap["UpdatedAt"],
		fieldMap["UpdatedBy"],
	).From(m.table + " p")
}

// make Get statement
func (m *Model) makeGetStatement(id *string, url *string) sq.SelectBuilder {
	// build query
	statement := m.makeStatement()

	// id
	if id != nil {
		statement = statement.Where(fieldMap["ID"]+" = ?", *id)
	}

	// url
	if url != nil {
		statement = statement.Where(fieldMap["Url"]+" = ?", *url)
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

	// Page
	if filter.Page == nil {
		filter.Page = new(uint64)
		*filter.Page = 1
	}

	// PerPage
	if filter.PerPage == nil {
		filter.PerPage = new(uint64)
		*filter.PerPage = 10
	}

	// Build query
	statement := m.makeStatement()

	// Ids
	if filter.IDs != nil && len(*filter.IDs) > 0 {
		statement = statement.Where(sq.Eq{fieldMap["ID"]: *filter.IDs})

		*filter.Page = 1
		if (*filter.PerPage) > uint64(len(*filter.IDs)) {
			*filter.PerPage = uint64(len(*filter.IDs))
		}
	}

	// Urls
	if filter.Urls != nil && len(*filter.Urls) > 0 {
		statement = statement.Where(sq.Eq{fieldMap["Url"]: *filter.Urls})

		*filter.Page = 1
		if (*filter.PerPage) > uint64(len(*filter.Urls)) {
			*filter.PerPage = uint64(len(*filter.Urls))
		}
	}

	// SKUs
	if filter.SKUs != nil {
		countSKUs := len(*filter.SKUs)

		if countSKUs > 0 {
			statement = statement.Where(sq.Eq{fieldMap["SKU"]: *filter.SKUs})
		}

		*filter.Page = 1
		if (*filter.PerPage) > uint64(countSKUs) {
			*filter.PerPage = uint64(countSKUs)
		}

	}

	// StatusID
	if filter.StatusID != nil && len(*filter.StatusID) > 0 {
		statement = statement.Where(sq.Eq{fieldMap["StatusID"]: *filter.StatusID})
	}

	// Search
	if filter.Search != nil {
		statement = statement.Where(
			sq.Or{
				sq.Expr("LOWER("+fieldMap["Name"]+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+fieldMap["Url"]+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+fieldMap["ShortDescription"]+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+fieldMap["Description"]+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+fieldMap["SeoTitle"]+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+fieldMap["SeoDescription"]+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
			},
		)
	}

	statement = statement.OrderBy(fieldMap[*filter.OrderBy] + " " + *filter.OrderDir).
		Offset((*filter.Page - 1) * *filter.PerPage).Limit(*filter.PerPage)

	// Add OrderBy, OrderDir, Page, Limit and return
	return statement
}

// scanOneRow
func (m *Model) scanOneRow(ctx context.Context, rows sq.RowScanner) (*Item, error) {
	var item = &Item{}
	err := rows.Scan(
		&item.ID,
		&item.SKU,
		&item.Name,
		&item.ShortDescription,
		&item.Description,
		&item.SortOrder,
		&item.StatusID,
		&item.Url,
		&item.IsTaxable,
		&item.IsTrackStock,
		&item.ShippingClassID,
		&item.ShippingWeight,
		&item.ShippingWidth,
		&item.ShippingHeight,
		&item.ShippingLength,
		&item.SeoTitle,
		&item.SeoDescription,
		&item.GTIN,
		&item.GoogleProductCategory,
		&item.GoogleProductType,
		&item.CreatedAt,
		&item.CreatedBy,
		&item.UpdatedAt,
		&item.UpdatedBy,
	)

	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, err
	}

	return item, nil
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
		fieldMap["ID"],
		fieldMap["SKU"],
		fieldMap["Name"],
		fieldMap["ShortDescription"],
		fieldMap["Description"],
		fieldMap["SortOrder"],
		fieldMap["StatusID"],
		fieldMap["Url"],
		fieldMap["IsTaxable"],
		fieldMap["IsTrackStock"],
		fieldMap["ShippingClassID"],
		fieldMap["ShippingWeight"],
		fieldMap["ShippingWidth"],
		fieldMap["ShippingHeight"],
		fieldMap["ShippingLength"],
		fieldMap["SeoTitle"],
		fieldMap["SeoDescription"],
		fieldMap["GTIN"],
		fieldMap["GoogleProductCategory"],
		fieldMap["GoogleProductType"],
		fieldMap["CreatedAt"],
		fieldMap["CreatedBy"],
		fieldMap["UpdatedAt"],
		fieldMap["UpdatedBy"],
	).Values(
		item.ID,
		item.SKU,
		item.Name,
		item.ShortDescription,
		item.Description,
		item.SortOrder,
		item.StatusID,
		item.Url,
		item.IsTaxable,
		item.IsTrackStock,
		item.ShippingClassID,
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
		Set(fieldMap["Name"], item.Name).
		Set(fieldMap["SKU"], item.SKU).
		Set(fieldMap["ShortDescription"], item.ShortDescription).
		Set(fieldMap["Description"], item.Description).
		Set(fieldMap["SortOrder"], item.SortOrder).
		Set(fieldMap["StatusID"], item.StatusID).
		Set(fieldMap["Url"], item.Url).
		Set(fieldMap["IsTaxable"], item.IsTaxable).
		Set(fieldMap["IsTrackStock"], item.IsTrackStock).
		Set(fieldMap["ShippingClassID"], item.ShippingClassID).
		Set(fieldMap["ShippingWeight"], item.ShippingWeight).
		Set(fieldMap["ShippingWidth"], item.ShippingWidth).
		Set(fieldMap["ShippingHeight"], item.ShippingHeight).
		Set(fieldMap["ShippingLength"], item.ShippingLength).
		Set(fieldMap["SeoTitle"], item.SeoTitle).
		Set(fieldMap["SeoDescription"], item.SeoDescription).
		Set(fieldMap["GTIN"], item.GTIN).
		Set(fieldMap["GoogleProductCategory"], item.GoogleProductCategory).
		Set(fieldMap["GoogleProductType"], item.GoogleProductType).
		Set(fieldMap["UpdatedAt"], "NOW()").
		Set(fieldMap["UpdatedBy"], by).
		Where(fieldMap["ID"]+" = ?", item.ID)
}

// makePatchStatement
func (m *Model) makePatchStatement(ctx context.Context, id *string, fields *map[string]interface{}) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	statement := m.qb.Update(m.table).Where("id = ?", id)

	for field, value := range *fields {
		field = fieldMap[field]
		statement = statement.Set(field, value)
	}

	return statement.Set(fieldMap["UpdatedAt"], "NOW()").Set(fieldMap["UpdatedBy"], by)
}
