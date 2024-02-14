package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"strconv"
	"time"
)

const table = "public.product"

// ProductDAO is a struct that contains the SQL statement builder and the PostgreSQL client.
type ProductDAO struct {
	qb     sq.StatementBuilderType
	client psql.Client
}

// ProductStorage is a struct that maps to the 'public.product' table in the PostgreSQL database.
type ProductStorage struct {
	ID                    string    `db:"id"`
	SKU                   string    `db:"sku"`
	Name                  string    `db:"name"`
	ShortDescription      string    `db:"short_description"`
	Description           string    `db:"description"`
	SortOrder             int       `db:"sort_order"`
	StatusID              string    `db:"status_id"`
	Slug                  string    `db:"slug"`
	RegularPrice          float64   `db:"regular_price"`
	SalePrice             float64   `db:"sale_price"`
	FactoryPrice          float64   `db:"factory_price"`
	IsTaxable             bool      `db:"is_taxable"`
	Quantity              int       `db:"quantity"`
	ReturnToStockDate     time.Time `db:"return_to_stock_date"`
	IsTrackStock          bool      `db:"is_track_stock"`
	ShippingClassID       string    `db:"shipping_class_id"`
	ShippingWeight        float64   `db:"shipping_weight"`
	ShippingWidth         float64   `db:"shipping_width"`
	ShippingHeight        float64   `db:"shipping_height"`
	ShippingLength        float64   `db:"shipping_length"`
	SeoTitle              string    `db:"seo_title"`
	SeoDescription        string    `db:"seo_description"`
	GTIN                  string    `db:"gtin"`
	GoogleProductCategory string    `db:"google_product_category"`
	GoogleProductType     string    `db:"google_product_type"`
	CreatedAt             time.Time `db:"created_at"`
	CreatedBy             string    `db:"created_by"`
	UpdatedAt             time.Time `db:"updated_at"`
	UpdatedBy             string    `db:"updated_by"`
}

// NewProductStorage is a constructor function that returns a new instance of ProductDAO.
func NewProductStorage(client psql.Client) *ProductDAO {
	return &ProductDAO{
		qb:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client: client,
	}
}

// Get is a method on the ProductDAO struct that retrieves a product from the database by its ID.
func (repo *ProductDAO) Get(ctx context.Context, id string) (*ProductStorage, error) {
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
