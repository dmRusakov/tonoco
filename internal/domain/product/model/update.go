package model

import (
	"context"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"strconv"
)

// Update is a method on the ProductModel struct that updates a Product in the database by its ID.
func (repo *ProductModel) Update(ctx context.Context, product *Product, by string) (*Product, error) {
	// build query
	statement := repo.qb.
		Update(repo.table).
		Set(fieldMap["SKU"], product.SKU).
		Set(fieldMap["Name"], product.Name).
		Set(fieldMap["ShortDescription"], product.ShortDescription).
		Set(fieldMap["Description"], product.Description).
		Set(fieldMap["SortOrder"], product.SortOrder).
		Set(fieldMap["StatusID"], product.StatusID).
		Set(fieldMap["Slug"], product.Slug).
		Set(fieldMap["RegularPrice"], product.RegularPrice).
		Set(fieldMap["SalePrice"], product.SalePrice).
		Set(fieldMap["FactoryPrice"], product.FactoryPrice).
		Set(fieldMap["IsTaxable"], product.IsTaxable).
		Set(fieldMap["Quantity"], product.Quantity).
		Set(fieldMap["ReturnToStockDate"], product.ReturnToStockDate).
		Set(fieldMap["IsTrackStock"], product.IsTrackStock).
		Set(fieldMap["ShippingClassID"], product.ShippingClassID).
		Set(fieldMap["ShippingWeight"], product.ShippingWeight).
		Set(fieldMap["ShippingWidth"], product.ShippingWidth).
		Set(fieldMap["ShippingHeight"], product.ShippingHeight).
		Set(fieldMap["ShippingLength"], product.ShippingLength).
		Set(fieldMap["SeoTitle"], product.SeoTitle).
		Set(fieldMap["SeoDescription"], product.SeoDescription).
		Set(fieldMap["GTIN"], product.GTIN).
		Set(fieldMap["GoogleProductCategory"], product.GoogleProductCategory).
		Set(fieldMap["GoogleProductType"], product.GoogleProductType).
		Set(fieldMap["UpdatedAt"], "NOW()").
		Set(fieldMap["UpdatedBy"], by).
		Where(fieldMap["ID"]+" = ?", product.ID)

	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return nil, err
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

		return nil, err
	}

	return repo.Get(ctx, product.ID)
}
