package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"strconv"
)

// Get is a method on the ProductModel struct that retrieves a product from the database by its ID.
func (repo *ProductModel) Get(ctx context.Context, id string) (*ProductStorage, error) {
	// build query
	statement := repo.qb.
		Select(
			"id",
			"sku",
			"name",
			"short_description",
			"description",
			"sort_order",
			"status_id",
			"slug",
			"regular_price",
			"sale_price",
			"factory_price",
			"is_taxable",
			"quantity",
			"return_to_stock_date",
			"is_track_stock",
			"shipping_class_id",
			"shipping_weight",
			"shipping_width",
			"shipping_height",
			"shipping_length",
			"seo_title",
			"seo_description",
			"gtin",
			"google_product_category",
			"google_product_type",
			"created_at",
			"created_by",
			"updated_at",
			"updated_by",
		).
		From(table + " p").
		Where(sq.Eq{"id": id})

	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return nil, err
	}

	tracing.SpanEvent(ctx, "Select Product")
	tracing.TraceVal(ctx, "SQL", query)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	// execute the SQL query
	rows, err := repo.client.Query(ctx, query, args...)
	if err != nil {
		err = psql.ErrDoQuery(err)
		tracing.Error(ctx, err)

		return nil, err
	}

	defer rows.Close()

	// scan the result set into a slice of ProductStorage structs
	product := &ProductStorage{}
	for rows.Next() {
		if err = rows.Scan(
			&product.ID,
			&product.SKU,
			&product.Name,
			&product.ShortDescription,
			&product.Description,
			&product.SortOrder,
			&product.StatusID,
			&product.Slug,
			&product.RegularPrice,
			&product.SalePrice,
			&product.FactoryPrice,
			&product.IsTaxable,
			&product.Quantity,
			&product.ReturnToStockDate,
			&product.IsTrackStock,
			&product.ShippingClassID,
			&product.ShippingWeight,
			&product.ShippingWidth,
			&product.ShippingHeight,
			&product.ShippingLength,
			&product.SeoTitle,
			&product.SeoDescription,
			&product.GTIN,
			&product.GoogleProductCategory,
			&product.GoogleProductType,
			&product.CreatedAt,
			&product.CreatedBy,
			&product.UpdatedAt,
			&product.UpdatedBy,
		); err != nil {
			err = psql.ErrScan(psql.ParsePgError(err))
			tracing.Error(ctx, err)

			return product, err
		}
	}

	return product, nil
}
