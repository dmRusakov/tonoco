package model

import (
	"context"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"strconv"
)

// Update is a method on the ProductModel struct that updates a product in the database by its ID.
func (repo *ProductModel) Update(ctx context.Context, id string, product *ProductStorage) error {
	// build query
	statement := repo.qb.
		Update(table).
		Set("sku", product.SKU).
		Set("name", product.Name).
		Set("short_description", product.ShortDescription).
		Set("description", product.Description).
		Set("sort_order", product.SortOrder).
		Set("status_id", product.StatusID).
		Set("slug", product.Slug).
		Set("regular_price", product.RegularPrice).
		Set("sale_price", product.SalePrice).
		Set("factory_price", product.FactoryPrice).
		Set("is_taxable", product.IsTaxable).
		Set("quantity", product.Quantity).
		Set("return_to_stock_date", product.ReturnToStockDate).
		Set("is_track_stock", product.IsTrackStock).
		Set("shipping_class_id", product.ShippingClassID).
		Set("shipping_weight", product.ShippingWeight).
		Set("shipping_width", product.ShippingWidth).
		Set("shipping_height", product.ShippingHeight).
		Set("shipping_length", product.ShippingLength).
		Set("seo_title", product.SeoTitle).
		Set("seo_description", product.SeoDescription).
		Set("gtin", product.GTIN).
		Set("google_product_category", product.GoogleProductCategory).
		Set("google_product_type", product.GoogleProductType).
		Set("updated_at", product.UpdatedAt).
		Set("updated_by", product.UpdatedBy).
		Where("id = ?", id)

	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return err
	}

	tracing.SpanEvent(ctx, "Update Product")
	tracing.TraceVal(ctx, "SQL", query)
	for i, arg := range args {
		argStr, ok := arg.(string)
		if !ok {
			// arg is not of type string, handle the error or continue to next iteration
			continue
		}
		tracing.TraceVal(ctx, "Arg"+strconv.Itoa(i), argStr)
	}

	// execute the query
	_, err = repo.client.Exec(ctx, query, args...)
	if err != nil {
		err = psql.ErrExec(err)
		tracing.Error(ctx, err)

		return err
	}

	return nil
}
