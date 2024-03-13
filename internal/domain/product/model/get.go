package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"strconv"
)

// Get is a method on the ProductModel struct that retrieves a Product from the database by its ID.
func (repo *ProductModel) Get(ctx context.Context, id string) (*Product, error) {
	// build query
	statement := repo.qb.
		Select(
			fieldMap["ID"],
			fieldMap["SKU"],
			fieldMap["Name"],
			fieldMap["ShortDescription"],
			fieldMap["Description"],
			fieldMap["SortOrder"],
			fieldMap["StatusID"],
			fieldMap["Url"],
			fieldMap["RegularPrice"],
			fieldMap["SalePrice"],
			fieldMap["FactoryPrice"],
			fieldMap["IsTaxable"],
			fieldMap["Quantity"],
			fieldMap["ReturnToStockDate"],
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
		).
		From(repo.table + " p").
		Where(sq.Eq{fieldMap["ID"]: id})

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

	if !rows.Next() {
		return nil, psql.ErrNoRowForID(id)
	}

	// scan the result set into a slice of Item structs
	product := &Product{}
	if err = rows.Scan(
		&product.ID,
		&product.SKU,
		&product.Name,
		&product.ShortDescription,
		&product.Description,
		&product.SortOrder,
		&product.StatusID,
		&product.Url,
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

		return nil, err
	}

	return product, nil
}
