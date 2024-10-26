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
		m.mapFieldToDBColumn("Sku"),
		m.mapFieldToDBColumn("Brand"),
		m.mapFieldToDBColumn("Name"),
		m.mapFieldToDBColumn("ShortDescription"),
		m.mapFieldToDBColumn("Description"),
		m.mapFieldToDBColumn("SortOrder"),
		m.mapFieldToDBColumn("Url"),
		m.mapFieldToDBColumn("IsTaxable"),
		m.mapFieldToDBColumn("IsTrackStock"),
		m.mapFieldToDBColumn("ShippingWeight"),
		m.mapFieldToDBColumn("ShippingWidth"),
		m.mapFieldToDBColumn("ShippingHeight"),
		m.mapFieldToDBColumn("ShippingLength"),
		m.mapFieldToDBColumn("SeoTitle"),
		m.mapFieldToDBColumn("SeoDescription"),
		m.mapFieldToDBColumn("GTIN"),
		m.mapFieldToDBColumn("GoogleProductCategory"),
		m.mapFieldToDBColumn("GoogleProductType"),
		m.mapFieldToDBColumn("CreatedAt"),
		m.mapFieldToDBColumn("CreatedBy"),
		m.mapFieldToDBColumn("UpdatedAt"),
		m.mapFieldToDBColumn("UpdatedBy"),
	).From(m.table + " p")
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

func (m *Model) makeStatementByFilter(filter *Filter) sq.SelectBuilder {
	// OrderBy
	if filter.OrderBy == nil {
		filter.OrderBy = pointer.StringPtr("SortOrder")
	}

	// OrderDir
	if filter.OrderDir == nil {
		filter.OrderDir = pointer.StringPtr("ASC")
	}

	// PerPage
	if filter.PerPage == nil {
		if filter.Page == nil {
			filter.PerPage = pointer.Uint64Ptr(999999999999999999)
		} else {
			filter.PerPage = pointer.Uint64Ptr(10)
		}
	}

	// Page
	if filter.Page == nil {
		filter.Page = pointer.Uint64Ptr(1)
	}

	// Build query
	statement := m.makeStatement()

	// Ids
	if filter.Ids != nil && len(*filter.Ids) > 0 {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Id"): *filter.Ids})

		*filter.Page = 1
		if (*filter.PerPage) > uint64(len(*filter.Ids)) {
			*filter.PerPage = uint64(len(*filter.Ids))
		}
	}

	// Urls
	if filter.Urls != nil && len(*filter.Urls) > 0 {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Url"): *filter.Urls})

		*filter.Page = 1
		if (*filter.PerPage) > uint64(len(*filter.Urls)) {
			*filter.PerPage = uint64(len(*filter.Urls))
		}
	}

	// Skus
	if filter.Skus != nil {
		countSkus := len(*filter.Skus)

		if countSkus > 0 {
			statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Sku"): *filter.Skus})
		}

		*filter.Page = 1
		if (*filter.PerPage) > uint64(countSkus) {
			*filter.PerPage = uint64(countSkus)
		}
	}

	// Brands - filter by brand
	if filter.Brands != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Brand"): *filter.Brands})
	}

	// Search
	if filter.Search != nil {
		statement = statement.Where(
			sq.Or{
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Brand")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Name")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Url")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("ShortDescription")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Description")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("SeoTitle")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("SeoDescription")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
			},
		)
	}

	statement = statement.OrderBy(m.mapFieldToDBColumn(*filter.OrderBy) + " " + *filter.OrderDir).
		Offset((*filter.Page - 1) * *filter.PerPage).Limit(*filter.PerPage)

	// Add OrderBy, OrderDir, Page, Limit and return
	return statement
}

func (m *Model) makeCountStatementByFilter(filter *Filter) sq.SelectBuilder {
	// Build query
	statement := m.qb.Select("COUNT(*) as count").From(m.table + " p")

	// Ids
	if filter.Ids != nil && len(*filter.Ids) > 0 {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Id"): *filter.Ids})
	}

	// Urls
	if filter.Urls != nil && len(*filter.Urls) > 0 {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Url"): *filter.Urls})
	}

	// Skus
	if filter.Skus != nil && len(*filter.Skus) > 0 {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Sku"): *filter.Skus})
	}

	// Search
	if filter.Search != nil {
		statement = statement.Where(
			sq.Or{
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Name")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Brand")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Url")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("ShortDescription")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Description")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("SeoTitle")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("SeoDescription")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
			},
		)
	}

	// return
	return statement
}

func (m *Model) scanOneRow(ctx context.Context, rows sq.RowScanner) (*Item, error) {
	var item = Item{}
	var id, sku, brand, name, shortDescription, description, url, seoTitle, seoDescription, gtin, googleProductCategory, googleProductType, createdBy, updatedBy sql.NullString
	var sortOrder sql.NullInt64
	var isTaxable, isTrackStock sql.NullBool
	var shippingWeight, shippingWidth, shippingHeight, shippingLength sql.NullInt64
	var createdAt, updatedAt sql.NullTime

	err := rows.Scan(
		&id,
		&sku,
		&brand,
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
		return nil, errors.AddCode(err, "752006")
	}

	if id.Valid {
		item.Id = uuid.MustParse(id.String)
	}
	if sku.Valid {
		item.Sku = sku.String
	}
	if brand.Valid {
		item.Brand = brand.String
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
		return nil, errors.AddCode(err, "187456")
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
		m.mapFieldToDBColumn("Sku"),
		m.mapFieldToDBColumn("Brand"),
		m.mapFieldToDBColumn("Name"),
		m.mapFieldToDBColumn("ShortDescription"),
		m.mapFieldToDBColumn("Description"),
		m.mapFieldToDBColumn("SortOrder"),
		m.mapFieldToDBColumn("Url"),
		m.mapFieldToDBColumn("IsTaxable"),
		m.mapFieldToDBColumn("IsTrackStock"),
		m.mapFieldToDBColumn("ShippingWeight"),
		m.mapFieldToDBColumn("ShippingWidth"),
		m.mapFieldToDBColumn("ShippingHeight"),
		m.mapFieldToDBColumn("ShippingLength"),
		m.mapFieldToDBColumn("SeoTitle"),
		m.mapFieldToDBColumn("SeoDescription"),
		m.mapFieldToDBColumn("GTIN"),
		m.mapFieldToDBColumn("GoogleProductCategory"),
		m.mapFieldToDBColumn("GoogleProductType"),
		m.mapFieldToDBColumn("CreatedAt"),
		m.mapFieldToDBColumn("CreatedBy"),
		m.mapFieldToDBColumn("UpdatedAt"),
		m.mapFieldToDBColumn("UpdatedBy"),
	).Values(
		item.Id,
		item.Sku,
		item.Brand,
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

	return &insertItem, &item.Id
}

func (m *Model) makeUpdateStatement(ctx context.Context, item *Item) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	return m.qb.Update(m.table).
		SetMap(map[string]interface{}{
			m.mapFieldToDBColumn("Sku"):                   item.Sku,
			m.mapFieldToDBColumn("Brand"):                 item.Brand,
			m.mapFieldToDBColumn("Name"):                  item.Name,
			m.mapFieldToDBColumn("ShortDescription"):      item.ShortDescription,
			m.mapFieldToDBColumn("Description"):           item.Description,
			m.mapFieldToDBColumn("SortOrder"):             item.SortOrder,
			m.mapFieldToDBColumn("Url"):                   item.Url,
			m.mapFieldToDBColumn("IsTaxable"):             item.IsTaxable,
			m.mapFieldToDBColumn("IsTrackStock"):          item.IsTrackStock,
			m.mapFieldToDBColumn("ShippingWeight"):        item.ShippingWeight,
			m.mapFieldToDBColumn("ShippingWidth"):         item.ShippingWidth,
			m.mapFieldToDBColumn("ShippingHeight"):        item.ShippingHeight,
			m.mapFieldToDBColumn("ShippingLength"):        item.ShippingLength,
			m.mapFieldToDBColumn("SeoTitle"):              item.SeoTitle,
			m.mapFieldToDBColumn("SeoDescription"):        item.SeoDescription,
			m.mapFieldToDBColumn("GTIN"):                  item.GTIN,
			m.mapFieldToDBColumn("GoogleProductCategory"): item.GoogleProductCategory,
			m.mapFieldToDBColumn("GoogleProductType"):     item.GoogleProductType,
			m.mapFieldToDBColumn("UpdatedAt"):             "NOW()",
			m.mapFieldToDBColumn("UpdatedBy"):             by,
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
