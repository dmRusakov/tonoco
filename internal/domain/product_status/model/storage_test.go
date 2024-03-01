package model_test

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/appInit"
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/domain/product_status/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"testing"
)

var productStatusNewWithId = &model.ProductStatus{
	ID:     "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a00",
	Name:   "New",
	Url:    "new",
	Active: false,
}

var productStatusNewWithoutId = &model.ProductStatus{
	Name:   "New2",
	Url:    "new2",
	Active: false,
}

var productStatus11 = model.ProductStatus{
	ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
	Name:      "Public",
	Url:       "public",
	SortOrder: 1,
	Active:    true,
}

var productStatus12 = model.ProductStatus{
	ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12",
	Name:      "Privet",
	Url:       "private",
	SortOrder: 2,
	Active:    true,
}

var productStatus13 = model.ProductStatus{
	ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13",
	Name:      "Out of stock",
	Url:       "out-of-stock",
	SortOrder: 3,
	Active:    true,
}

var productStatus14 = model.ProductStatus{
	ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a14",
	Name:      "Discontinued",
	Url:       "discontinued",
	Active:    true,
	SortOrder: 4,
}

var productStatus15 = model.ProductStatus{
	ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a15",
	Name:      "Archived",
	Url:       "archived",
	SortOrder: 5,
	Active:    true,
}

// test get
func TestProductStatusGet(t *testing.T) {
	// Create a real database client
	pgClient := initDB(t)

	// Initialize a new instance of the model
	storage := model.NewProductStatusStorage(pgClient)

	// Define the test cases
	testCases := []struct {
		name     string
		id       string
		expected *model.ProductStatus
	}{
		{
			name:     "Get product status by ID",
			id:       productStatus11.ID,
			expected: &productStatus11,
		}, {
			name:     "Get product status by ID",
			id:       productStatus12.ID,
			expected: &productStatus12,
		}, {
			name:     "Get product status by ID",
			id:       productStatus13.ID,
			expected: &productStatus13,
		}, {
			name:     "Get product status by ID",
			id:       productStatus14.ID,
			expected: &productStatus14,
		}, {
			name:     "Get product status by ID",
			id:       productStatus15.ID,
			expected: &productStatus15,
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
			assert.Equal(t, tc.expected.Url, result.Url)
			assert.Equal(t, tc.expected.Active, result.Active)
		})
	}

}

// test update
func TestProductStatusUpdate(t *testing.T) {
	// Create a real database client
	pgClient := initDB(t)

	// Initialize a new instance of the model
	storage := model.NewProductStatusStorage(pgClient)

	// create new variable from productStatus01
	productStatus11New := productStatus11
	productStatus11New.Name = "PublicNew"
	productStatus11New.Url = "public-new"
	productStatus11New.Active = false

	// Define the test cases
	testCases := []struct {
		name string
		id   string
		sent *model.ProductStatus
		get  *model.ProductStatus
	}{
		{
			name: "Update product status by ID",
			id:   productStatus11.ID,
			sent: &productStatus11New,
			get:  &productStatus11New,
		}, {
			name: "Update product status by ID",
			id:   productStatus12.ID,
			sent: &productStatus12,
			get:  &productStatus12,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := storage.Update(context.Background(), tc.sent, "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1")

			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the result matches the expected value
			assert.Equal(t, tc.get.ID, result.ID)
			assert.Equal(t, tc.get.Name, result.Name)
			assert.Equal(t, tc.get.Url, result.Url)
			assert.Equal(t, tc.get.Active, result.Active)
		})
	}
}

// test patch
func TestProductStatusPatch(t *testing.T) {
	// Create a real database client
	pgClient := initDB(t)

	// Initialize a new instance of the model
	storage := model.NewProductStatusStorage(pgClient)

	// create new variable from productStatus01
	productStatus11New := productStatus11
	productStatus11New.Name = "PublicNew"
	productStatus11New.Url = "public-new"
	productStatus11New.Active = false

	// Define the test cases
	testCases := []struct {
		name   string
		id     string
		fields map[string]interface{}
		get    *model.ProductStatus
	}{
		{
			name: "Patch product status by ID",
			id:   productStatus11.ID,
			fields: map[string]interface{}{
				"Name":   productStatus11New.Name,
				"Url":    productStatus11New.Url,
				"Active": productStatus11New.Active,
			},
			get: &productStatus11New,
		}, {
			name: "Patch product status by ID",
			id:   productStatus11.ID,
			fields: map[string]interface{}{
				"Name":   productStatus11.Name,
				"Url":    productStatus11.Url,
				"Active": productStatus11.Active,
			},
			get: &productStatus11,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := storage.Patch(context.Background(), tc.id, tc.fields, "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1")

			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the result matches the expected value
			assert.Equal(t, tc.get.ID, result.ID)
			assert.Equal(t, tc.get.Name, result.Name)
			assert.Equal(t, tc.get.Url, result.Url)
			assert.Equal(t, tc.get.Active, result.Active)
		})

	}
}

// test create with id
func TestProductStatusCreateWithId(t *testing.T) {
	// Create a real database client
	pgClient := initDB(t)

	// Initialize a new instance of the model
	storage := model.NewProductStatusStorage(pgClient)

	// Define the test cases
	testCases := []struct {
		name     string
		sent     *model.ProductStatus
		expected *model.ProductStatus
	}{
		{
			name:     "Create product status with ID",
			sent:     productStatusNewWithId,
			expected: productStatusNewWithId,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := storage.Create(context.Background(), tc.sent, "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1")

			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the result matches the expected value
			assert.Equal(t, tc.expected.ID, result.ID)
			assert.Equal(t, tc.expected.Name, result.Name)
			assert.Equal(t, tc.expected.Url, result.Url)
			assert.Equal(t, tc.expected.Active, result.Active)
		})
	}

}

// test create without id
func TestProductStatusCreateWithoutId(t *testing.T) {
	// Create a real database client
	pgClient := initDB(t)

	// Initialize a new instance of the model
	storage := model.NewProductStatusStorage(pgClient)

	// Define the test cases
	testCases := []struct {
		name string
		sent *model.ProductStatus
		get  *model.ProductStatus
	}{
		{
			name: "Create product status without ID",
			sent: productStatusNewWithoutId,
			get:  productStatusNewWithoutId,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := storage.Create(context.Background(), tc.sent, "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1")

			// update ID
			tc.get.ID = result.ID

			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the result matches the expected value
			assert.Equal(t, tc.get.Name, result.Name)
			assert.Equal(t, tc.get.Url, result.Url)
			assert.Equal(t, tc.get.Active, result.Active)
		})
	}
}

// test all
func TestProductStatusAll(t *testing.T) {
	// Create a real database client
	pgClient := initDB(t)

	// Initialize a new instance of the model
	storage := model.NewProductStatusStorage(pgClient)
	isActive := true
	sortOrder := "SortOrder"
	sortOrderDesc := "desc"
	page := uint64(1)
	perPage := uint64(3)

	// Define the test cases
	testCases := []struct {
		name     string
		filter   *model.Filter
		expected []*model.ProductStatus
	}{
		{
			name:     "Get all product status",
			filter:   &model.Filter{},
			expected: []*model.ProductStatus{&productStatus11, &productStatus12, &productStatus13, &productStatus14, &productStatus15, productStatusNewWithId, productStatusNewWithoutId},
		}, {
			name: "Get active product status",
			filter: &model.Filter{
				Active: &isActive,
			},
			expected: []*model.ProductStatus{&productStatus11, &productStatus12, &productStatus13, &productStatus14, &productStatus15},
		}, {
			name: "Get product status by sort order desc",
			filter: &model.Filter{
				SortBy:    &sortOrder,
				SortOrder: &sortOrderDesc,
			},
			expected: []*model.ProductStatus{productStatusNewWithoutId, productStatusNewWithId, &productStatus15, &productStatus14, &productStatus13, &productStatus12, &productStatus11},
		}, {
			name: "Get product status with page and per page",
			filter: &model.Filter{
				Page:    &page,
				PerPage: &perPage,
			},
			expected: []*model.ProductStatus{&productStatus11, &productStatus12, &productStatus13},
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := storage.All(context.Background(), tc.filter)

			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the result matches the expected value
			assert.Equal(t, len(tc.expected), len(result))
			for i, v := range result {
				assert.Equal(t, tc.expected[i].ID, v.ID)
				assert.Equal(t, tc.expected[i].Name, v.Name)
				assert.Equal(t, tc.expected[i].Url, v.Url)
				assert.Equal(t, tc.expected[i].Active, v.Active)
			}
		})
	}

}

// test delete
func TestProductStatusDelete(t *testing.T) {
	// Create a real database client
	pgClient := initDB(t)

	// Initialize a new instance of the model
	storage := model.NewProductStatusStorage(pgClient)

	// Define the test cases
	testCases := []struct {
		name string
		id   string
	}{
		{
			name: "Delete product status by ID",
			id:   productStatusNewWithId.ID,
		}, {
			name: "Delete product status by ID",
			id:   productStatusNewWithoutId.ID,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := storage.Delete(context.Background(), tc.id)

			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the result matches the expected value
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
