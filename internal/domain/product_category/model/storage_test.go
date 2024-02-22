package model_test

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/appInit"
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/domain/product_category/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"testing"
)

var productCategoryNewWithId = model.ProductCategory{
	ID:               "1f484cda-c00e-4ed8-a325-9c5e035f0000",
	Name:             "New product category",
	Slug:             "new-product-category",
	ShortDescription: "Some text",
	Description:      "Some text",
	SortOrder:        1,
	Prime:            true,
	Active:           true,
}

var productCategoryNewWithoutId = model.ProductCategory{
	Name:             "New product category",
	Slug:             "new-product-category-01",
	ShortDescription: "Some text",
	Description:      "Some text",
	SortOrder:        1,
	Prime:            true,
	Active:           true,
}

var productCategory01 = model.ProductCategory{
	ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9901",
	Name:             "Island Range Hoods",
	Slug:             "island",
	ShortDescription: "Some text",
	Description:      "Some text",
	SortOrder:        1,
	Prime:            true,
	Active:           true,
}

var productCategory02 = model.ProductCategory{
	ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9902",
	Name:             "Wall range hoods",
	Slug:             "wall",
	ShortDescription: "Some text",
	Description:      "Some text",
	SortOrder:        2,
	Prime:            true,
	Active:           true,
}

var productCategory03 = model.ProductCategory{
	ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9903",
	Name:             "Air loop range hoods",
	Slug:             "ait-loop",
	ShortDescription: "Some text",
	Description:      "Some text",
	SortOrder:        3,
	Prime:            true,
	Active:           true,
}

var productCategory04 = model.ProductCategory{
	ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9904",
	Name:             "Built-in Range Hoods",
	Slug:             "built-in",
	ShortDescription: "Some text",
	Description:      "Some text",
	SortOrder:        4,
	Prime:            true,
	Active:           true,
}

// test get
func TestProductCategoryGet(t *testing.T) {
	// Create a real database client
	pgClient := initDB(t)

	// Initialize a new instance of the model
	storage := model.NewProductStorage(pgClient)

	// Define the test cases
	testCases := []struct {
		name     string
		id       string
		expected *model.ProductCategory
	}{
		{
			name:     "Get product category 01",
			id:       productCategory01.ID,
			expected: &productCategory01,
		}, {
			name:     "Get product category 02",
			id:       productCategory02.ID,
			expected: &productCategory02,
		}, {
			name:     "Get product category 03",
			id:       productCategory03.ID,
			expected: &productCategory03,
		}, {
			name:     "Get product category 04",
			id:       productCategory04.ID,
			expected: &productCategory04,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := storage.Get(context.Background(), tc.id)

			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the result matches the expected value
			assert.Equal(t, tc.expected.ID, result.ID)
			assert.Equal(t, tc.expected.Name, result.Name)
			assert.Equal(t, tc.expected.Slug, result.Slug)
			assert.Equal(t, tc.expected.ShortDescription, result.ShortDescription)
			assert.Equal(t, tc.expected.Description, result.Description)
			assert.Equal(t, tc.expected.Prime, result.Prime)
			assert.Equal(t, tc.expected.Active, result.Active)
		})
	}
}

func initDB(t *testing.T) *pgxpool.Pool {
	// Create a real database client
	cfg := config.GetConfig(context.Background())
	app := appInit.NewAppInit(context.Background(), cfg)
	err := app.SqlDBInit()
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	pgClient := app.SqlDB
	return pgClient
}
