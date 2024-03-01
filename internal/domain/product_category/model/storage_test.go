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
	Url:              "new-product-category",
	ShortDescription: "Some text",
	Description:      "Some text",
	SortOrder:        1,
	Prime:            true,
	Active:           true,
}

var productCategoryNewWithoutId = model.ProductCategory{
	Name:             "New product category",
	Url:              "new-product-category-01",
	ShortDescription: "Some text",
	Description:      "Some text",
	SortOrder:        1,
	Prime:            true,
	Active:           true,
}

var productCategory01 = model.ProductCategory{
	ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9901",
	Name:             "Island Range Hoods",
	Url:              "island",
	ShortDescription: "Some text",
	Description:      "Some text",
	SortOrder:        1,
	Prime:            true,
	Active:           true,
}

var productCategory02 = model.ProductCategory{
	ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9902",
	Name:             "Wall range hoods",
	Url:              "wall",
	ShortDescription: "Some text",
	Description:      "Some text",
	SortOrder:        2,
	Prime:            true,
	Active:           true,
}

var productCategory03 = model.ProductCategory{
	ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9903",
	Name:             "Air loop range hoods",
	Url:              "ait-loop",
	ShortDescription: "Some text",
	Description:      "Some text",
	SortOrder:        3,
	Prime:            true,
	Active:           true,
}

var productCategory04 = model.ProductCategory{
	ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9904",
	Name:             "Built-in Range Hoods",
	Url:              "built-in",
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
	storage := model.NewProductCategoryStorage(pgClient)

	// Define the test cases
	testCases := []struct {
		name string
		id   string
		get  *model.ProductCategory
	}{
		{
			name: "Get product category 01",
			id:   productCategory01.ID,
			get:  &productCategory01,
		}, {
			name: "Get product category 02",
			id:   productCategory02.ID,
			get:  &productCategory02,
		}, {
			name: "Get product category 03",
			id:   productCategory03.ID,
			get:  &productCategory03,
		}, {
			name: "Get product category 04",
			id:   productCategory04.ID,
			get:  &productCategory04,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := storage.Get(context.Background(), tc.id)

			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the result matches the get value
			assert.Equal(t, tc.get.ID, result.ID)
			assert.Equal(t, tc.get.Name, result.Name)
			assert.Equal(t, tc.get.Url, result.Url)
			assert.Equal(t, tc.get.ShortDescription, result.ShortDescription)
			assert.Equal(t, tc.get.Description, result.Description)
			assert.Equal(t, tc.get.Prime, result.Prime)
			assert.Equal(t, tc.get.Active, result.Active)
		})
	}
}

// test update
func TestProductCategoryUpdate(t *testing.T) {
	// Create a real database client
	pgClient := initDB(t)

	// Initialize a new instance of the model
	storage := model.NewProductCategoryStorage(pgClient)

	// create new variable from productCategory01
	productCategory01New := productCategory01
	productCategory01New.Name = "Island Range Hoods New"
	productCategory01New.Url = "island-new"
	productCategory01New.Prime = false

	// Define the test cases
	testCases := []struct {
		name   string
		id     string
		sent   *model.ProductCategory
		update *model.ProductCategory
	}{
		{
			name:   "Update product category 01",
			id:     productCategory01.ID,
			sent:   &productCategory01New,
			update: &productCategory01New,
		}, {
			name:   "Update product category 02",
			id:     productCategory01.ID,
			sent:   &productCategory01,
			update: &productCategory01,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the Update method
			_, err := storage.Update(context.Background(), tc.sent, "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1")

			// Assert that there was no error
			assert.NoError(t, err)

			// Call the Get method
			result, err := storage.Get(context.Background(), tc.id)

			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the result matches the get value
			assert.Equal(t, tc.update.ID, result.ID)
			assert.Equal(t, tc.update.Name, result.Name)
			assert.Equal(t, tc.update.Url, result.Url)
			assert.Equal(t, tc.update.ShortDescription, result.ShortDescription)
			assert.Equal(t, tc.update.Description, result.Description)
			assert.Equal(t, tc.update.Prime, result.Prime)
			assert.Equal(t, tc.update.Active, result.Active)
		})
	}
}

// test patch
func TestProductCategoryPatch(t *testing.T) {
	// Create a real database client
	pgClient := initDB(t)

	// Initialize a new instance of the model
	storage := model.NewProductCategoryStorage(pgClient)

	// create new variable from productCategory01
	productCategory01New := productCategory01
	productCategory01New.Name = "Island Range Hoods New"
	productCategory01New.Url = "island-new"
	productCategory01New.Prime = false

	// Define the test cases
	testCases := []struct {
		name   string
		id     string
		fields map[string]interface{}
		get    *model.ProductCategory
	}{
		{
			name: "Patch product category 01",
			id:   productCategory01.ID,
			fields: map[string]interface{}{
				"Name":  productCategory01New.Name,
				"Url":   productCategory01New.Url,
				"Prime": productCategory01New.Prime,
			},
			get: &productCategory01New,
		}, {
			name: "Patch product category 02",
			id:   productCategory01.ID,
			fields: map[string]interface{}{
				"Name":  productCategory01.Name,
				"Url":   productCategory01.Url,
				"Prime": productCategory01.Prime,
			},
			get: &productCategory01,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the Patch method
			_, err := storage.Patch(context.Background(), tc.id, tc.fields, "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1")

			// Assert that there was no error
			assert.NoError(t, err)

			// Call the Get method
			result, err := storage.Get(context.Background(), tc.id)

			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the result matches the get value
			assert.Equal(t, tc.get.ID, result.ID)
			assert.Equal(t, tc.get.Name, result.Name)
			assert.Equal(t, tc.get.Url, result.Url)
			assert.Equal(t, tc.get.ShortDescription, result.ShortDescription)
			assert.Equal(t, tc.get.Description, result.Description)
			assert.Equal(t, tc.get.Prime, result.Prime)
			assert.Equal(t, tc.get.Active, result.Active)
		})
	}
}

// test create with id
func TestProductCategoryCreateWithId(t *testing.T) {
	// Create a real database client
	pgClient := initDB(t)

	// Initialize a new instance of the model
	storage := model.NewProductCategoryStorage(pgClient)

	// Define the test cases
	testCases := []struct {
		name string
		sent *model.ProductCategory
		get  *model.ProductCategory
	}{
		{
			name: "Create product category with id",
			sent: &productCategoryNewWithId,
			get:  &productCategoryNewWithId,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the Create method
			result, err := storage.Create(context.Background(), tc.sent, "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1")

			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the result matches the get value
			assert.Equal(t, tc.get.ID, result.ID)
			assert.Equal(t, tc.get.Name, result.Name)
			assert.Equal(t, tc.get.Url, result.Url)
			assert.Equal(t, tc.get.ShortDescription, result.ShortDescription)
			assert.Equal(t, tc.get.Description, result.Description)
			assert.Equal(t, tc.get.Prime, result.Prime)
			assert.Equal(t, tc.get.Active, result.Active)
		})
	}
}

// test create without id
func TestProductCategoryCreateWithoutId(t *testing.T) {
	// Create a real database client
	pgClient := initDB(t)

	// Initialize a new instance of the model
	storage := model.NewProductCategoryStorage(pgClient)

	// Define the test cases
	testCases := []struct {
		name string
		sent *model.ProductCategory
		get  *model.ProductCategory
	}{
		{
			name: "Create product category without id",
			sent: &productCategoryNewWithoutId,
			get:  &productCategoryNewWithoutId,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the Create method
			result, err := storage.Create(context.Background(), tc.sent, "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1")

			// update ID
			tc.get.ID = result.ID

			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the result matches the get value
			assert.Equal(t, tc.get.Name, result.Name)
			assert.Equal(t, tc.get.Url, result.Url)
			assert.Equal(t, tc.get.ShortDescription, result.ShortDescription)
			assert.Equal(t, tc.get.Description, result.Description)
			assert.Equal(t, tc.get.Prime, result.Prime)
			assert.Equal(t, tc.get.Active, result.Active)
		})
	}
}

// test delete
func TestProductCategoryDelete(t *testing.T) {
	// Create a real database client
	pgClient := initDB(t)

	// Initialize a new instance of the model
	storage := model.NewProductCategoryStorage(pgClient)

	// Define the test cases
	testCases := []struct {
		name string
		id   string
	}{
		{
			name: "Delete product category 01",
			id:   productCategoryNewWithId.ID,
		}, {
			name: "Delete product category 02",
			id:   productCategoryNewWithoutId.ID,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the Delete method
			err := storage.Delete(context.Background(), tc.id)

			// Assert that there was no error
			assert.NoError(t, err)

			// Call the Get method
			_, err = storage.Get(context.Background(), tc.id)
			assert.Error(t, err)
		})
	}

}

// initDB
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
