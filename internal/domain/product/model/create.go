package model

import (
	"context"
	"github.com/google/uuid"
)

func (repo *ProductModel) Create(ctx context.Context, product *Product, by string) (*Product, error) {
	// if ID is not set, generate a new UUID
	if product.ID == "" {
		product.ID = uuid.New().String()

		// check if the already exists
		_, err := repo.Get(ctx, product.ID)
		if err == nil {
			product.ID = uuid.New().String()
		}
	}

	// build query
	statement := repo.qb.Insert(repo.table).Columns(
		fieldMap["ID"], fieldMap["SKU"], fieldMap["Name"], fieldMap["ShortDescription"], fieldMap["Description"],
		fieldMap["SortOrder"], fieldMap["StatusID"], fieldMap["Slug"], fieldMap["RegularPrice"], fieldMap["SalePrice"],
		fieldMap["FactoryPrice"], fieldMap["IsTaxable"], fieldMap["Quantity"], fieldMap["ReturnToStockDate"],
		fieldMap["IsTrackStock"], fieldMap["ShippingClassID"], fieldMap["ShippingWeight"], fieldMap["ShippingWidth"],
		fieldMap["ShippingHeight"], fieldMap["ShippingLength"], fieldMap["SeoTitle"], fieldMap["SeoDescription"],
		fieldMap["GTIN"], fieldMap["GoogleProductCategory"], fieldMap["GoogleProductType"], fieldMap["CreatedAt"],
		fieldMap["CreatedBy"], fieldMap["UpdatedAt"], fieldMap["UpdatedBy"]).Values(product.ID,
		product.SKU, product.Name, product.ShortDescription, product.Description, product.SortOrder, product.StatusID, product.Slug,
		product.RegularPrice, product.SalePrice, product.FactoryPrice, product.IsTaxable, product.Quantity, product.ReturnToStockDate,
		product.IsTrackStock, product.ShippingClassID, product.ShippingWeight, product.ShippingWidth, product.ShippingHeight,
		product.ShippingLength, product.SeoTitle, product.SeoDescription, product.GTIN, product.GoogleProductCategory,
		product.GoogleProductType, "NOW()", by, "NOW()", by)

	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		return nil, err
	}

	// execute the query
	_, err = repo.client.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	// retrieve the created Product
	product, err = repo.Get(ctx, product.ID)
	if err != nil {
		return nil, err
	}

	return product, nil
}
