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

var testProductCategories = []*model.ProductCategory{
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9901",
		Url:              "island",
		Name:             "Island Range Hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            true,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9902",
		Url:              "wall",
		Name:             "Wall range hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            true,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9903",
		Url:              "ait-loop",
		Name:             "Air loop range hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            true,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9904",
		Url:              "built-in",
		Name:             "Built-in Range Hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            true,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9905",
		Url:              "under-cabinet",
		Name:             "Under Cabinet Range hood",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            true,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9906",
		Url:              "accessories",
		Name:             "Accessories",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            true,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9907",
		Url:              "black",
		Name:             "Black Range hood",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9908",
		Url:              "white",
		Name:             "White Range hood",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9909",
		Url:              "wood",
		Name:             "Wood Range hood",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9910",
		Url:              "stainless-steel",
		Name:             "Stainless Steel Range hood",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9911",
		Url:              "glass",
		Name:             "Glass Range hood",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9912",
		Url:              "perimeter-filter",
		Name:             "Perimeter Filter Range Hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9913",
		Url:              "murano",
		Name:             "Murano Range Hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9914",
		Url:              "ductless-range-hoods",
		Name:             "Ductless Range hood",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9915",
		Url:              "ducted-range-hoods",
		Name:             "Ducted Range hood",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9999",
		Url:              "discontinued-range-hoods",
		Name:             "Discontinued",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            false,
		Active:           false,
	},
}

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

func TestProductCategory(t *testing.T) {
	t.Run("testProductCategoryGet", testProductCategoryGet)
	t.Run("testProductCategoryUpdate", testProductCategoryUpdate)
	t.Run("testProductCategoryPatch", testProductCategoryPatch)
	t.Run("testProductCategoryCreateWithId", testProductCategoryCreateWithId)
	t.Run("testProductCategoryCreateWithoutId", testProductCategoryCreateWithoutId)
	t.Run("testProductCategoryCreateWithoutId", testProductCategoryDelete)
	t.Run("testProductCategoryAll", testProductCategoryAll)
}

// test get
func testProductCategoryGet(t *testing.T) {
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
			id:   testProductCategories[0].ID,
			get:  testProductCategories[0],
		}, {
			name: "Get product category 02",
			id:   testProductCategories[1].ID,
			get:  testProductCategories[1],
		}, {
			name: "Get product category 03",
			id:   testProductCategories[2].ID,
			get:  testProductCategories[2],
		}, {
			name: "Get product category 04",
			id:   testProductCategories[3].ID,
			get:  testProductCategories[3],
		}, {
			name: "Get product category 05",
			id:   testProductCategories[4].ID,
			get:  testProductCategories[4],
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
func testProductCategoryUpdate(t *testing.T) {
	// Create a real database client
	pgClient := initDB(t)

	// Initialize a new instance of the model
	storage := model.NewProductCategoryStorage(pgClient)

	// create new variable from testProductCategories[0]
	testProductCategoryNew := testProductCategories[0]
	testProductCategoryNew.Name = "Island Range Hoods New"
	testProductCategoryNew.Url = "island-new"
	testProductCategoryNew.Active = false

	// Define the test cases
	testCases := []struct {
		name   string
		id     string
		sent   *model.ProductCategory
		update *model.ProductCategory
	}{
		{
			name:   "Update product category 01",
			id:     testProductCategories[0].ID,
			sent:   testProductCategoryNew,
			update: testProductCategoryNew,
		}, {
			name:   "Update product category 02",
			id:     testProductCategories[0].ID,
			sent:   testProductCategories[0],
			update: testProductCategories[0],
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
			assert.Equal(t, tc.update.ShortDescription, result.ShortDescription)
			assert.Equal(t, tc.update.Description, result.Description)
			assert.Equal(t, tc.update.Prime, result.Prime)
			assert.Equal(t, tc.update.Active, result.Active)
		})
	}
}

// test patch
func testProductCategoryPatch(t *testing.T) {
	// Create a real database client
	pgClient := initDB(t)

	// Initialize a new instance of the model
	storage := model.NewProductCategoryStorage(pgClient)

	// create new variable from testProductCategories[0]
	testProductCategoryNew := testProductCategories[0]
	testProductCategoryNew.Name = "Island Range Hoods New"
	testProductCategoryNew.Url = "island-new"
	testProductCategoryNew.Active = false

	// Define the test cases
	testCases := []struct {
		name   string
		id     string
		fields map[string]interface{}
		get    *model.ProductCategory
	}{
		{
			name: "Patch product category 01",
			id:   testProductCategories[0].ID,
			fields: map[string]interface{}{
				"Name":  testProductCategoryNew.Name,
				"Url":   testProductCategoryNew.Url,
				"Prime": testProductCategoryNew.Prime,
			},
			get: testProductCategoryNew,
		}, {
			name: "Patch product category 02",
			id:   testProductCategories[0].ID,
			fields: map[string]interface{}{
				"Name":  testProductCategories[0].Name,
				"Url":   testProductCategories[0].Url,
				"Prime": testProductCategories[0].Prime,
			},
			get: testProductCategories[0],
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
func testProductCategoryCreateWithId(t *testing.T) {
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
func testProductCategoryCreateWithoutId(t *testing.T) {
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
func testProductCategoryDelete(t *testing.T) {
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

// test all
func testProductCategoryAll(t *testing.T) {
	// Create a real database client
	pgClient := initDB(t)

	// Initialize a new instance of the model
	storage := model.NewProductCategoryStorage(pgClient)

	// test varietals
	isPrime := true
	page := uint64(1)
	perPage3 := uint64(3)

	// Define the test cases
	testCases := []struct {
		name     string
		filter   *model.Filter
		expected []*model.ProductCategory
	}{
		{
			name:     "Get 10 products categories 01",
			filter:   &model.Filter{},
			expected: testProductCategories[:10],
		}, {
			name: "Get 3 products categories 02",
			filter: &model.Filter{
				PerPage: &perPage3,
				Page:    &page,
			},
			expected: testProductCategories[:3],
		}, {
			name: "Get Prime products categories 03",
			filter: &model.Filter{
				Prime: &isPrime,
			},
			expected: testProductCategories[0:6],
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call All method
			result, err := storage.All(context.Background(), tc.filter)

			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the result matches the expected value
			assert.Equal(t, len(tc.expected), len(result))

			for i, v := range result {
				assert.Equal(t, tc.expected[i].ID, v.ID)
				assert.Equal(t, tc.expected[i].Name, v.Name)
				assert.Equal(t, tc.expected[i].Url, v.Url)
				assert.Equal(t, tc.expected[i].ShortDescription, v.ShortDescription)
				assert.Equal(t, tc.expected[i].Description, v.Description)
				assert.Equal(t, tc.expected[i].Prime, v.Prime)
				assert.Equal(t, tc.expected[i].Active, v.Active)
			}
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
