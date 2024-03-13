package model

import (
	"context"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"github.com/google/uuid"
	"strconv"
)

// Create is a method of ProductModel that creates a new product in the database.
// It takes a context, a pointer to a Product, and a string representing the creator of the product.
// It returns a pointer to the created Product and an error if there was one.
func (repo *ProductModel) Create(ctx context.Context, product *Product, by string) (*Product, error) {
	// if ID is not set, generate a new UUID
	if product.ID == "" {
		product.ID = uuid.New().String()
	}

	// build query
	statement := repo.qb.Insert(repo.table).
		Columns(
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
		Values(
			product.ID,
			product.SKU,
			product.Name,
			product.ShortDescription,
			product.Description,
			product.SortOrder,
			product.StatusID,
			product.Url,
			product.RegularPrice,
			product.SalePrice,
			product.FactoryPrice,
			product.IsTaxable,
			product.Quantity,
			product.ReturnToStockDate,
			product.IsTrackStock,
			product.ShippingClassID,
			product.ShippingWeight,
			product.ShippingWidth,
			product.ShippingHeight,
			product.ShippingLength,
			product.SeoTitle,
			product.SeoDescription,
			product.GTIN,
			product.GoogleProductCategory,
			product.GoogleProductType,
			"NOW()",
			by,
			"NOW()",
			by,
		)

	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return nil, err
	}

	// add tracing for the insert product query
	tracing.SpanEvent(ctx, "Insert Product query")
	tracing.TraceVal(ctx, "sql", query)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	// Execute the query
	cmd, execErr := repo.client.Exec(ctx, query, args...)
	if execErr != nil {
		execErr = psql.ErrDoQuery(execErr)
		tracing.Error(ctx, execErr)

		return nil, execErr
	}

	// Check if any rows were affected
	if cmd.RowsAffected() == 0 {
		return nil, psql.ErrNothingInserted
	}

	// Return the created product
	return repo.Get(ctx, product.ID)
}
