package model_test

import (
	"context"
	"fmt"
	"github.com/dmRusakov/tonoco/internal/appInit"
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/domain/product_category/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testProductCategories = []*model.ProductCategory{
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9901",
		Url:              "island",
		Name:             "Island range hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        1,
		Prime:            true,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9902",
		Url:              "wall",
		Name:             "Wall range hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        2,
		Prime:            true,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9903",
		Url:              "ait-loop",
		Name:             "Air loop range hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        3,
		Prime:            true,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9904",
		Url:              "built-in",
		Name:             "Built-in range hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        4,
		Prime:            true,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9905",
		Url:              "under-cabinet",
		Name:             "Under Cabinet range hood",
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
		Name:             "Black range hood",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9908",
		Url:              "white",
		Name:             "White range hood",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9909",
		Url:              "wood",
		Name:             "Wood range hood",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9910",
		Url:              "stainless-steel",
		Name:             "Stainless Steel range hood",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9911",
		Url:              "glass",
		Name:             "Glass range hood",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9912",
		Url:              "perimeter-filter",
		Name:             "Perimeter Filter range hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9913",
		Url:              "murano",
		Name:             "Murano range hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9914",
		Url:              "ductless-range-hoods",
		Name:             "Ductless range hood",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9915",
		Url:              "ducted-range-hoods",
		Name:             "Ducted range hood",
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

var testProductCategoriesNew = []*model.ProductCategory{
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e045f0001",
		Url:              "test1-url",
		Name:             "Test 1",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            false,
		Active:           false,
	}, {
		ID:               "1f484cda-c00e-4ed8-a325-9c5e045f0002",
		Url:              "test2-url",
		Name:             "Test 2",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            false,
		Active:           false,
	}, {
		Url:              "test3-url",
		Name:             "Test 3",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            false,
		Active:           false,
	}, {
		Url:              "test4-url",
		Name:             "Test 4",
		ShortDescription: "Some text",
		Description:      "Some text",
		Prime:            false,
		Active:           false,
	},
}

func TestProductCategory(t *testing.T) {
	t.Run("all", all)
	t.Run("get", get)
	t.Run("createWithId", createWithId)
	t.Run("createWithoutId", createWithoutId)
	t.Run("update", update)
	t.Run("patch", patch)
	t.Run("delete", delete)
}

// test all
func all(t *testing.T) {
	// Create a storage with real database client
	storage := initStorage(t)

	// test varietals
	isPrime := true
	page := uint64(1)
	perPage3 := uint64(3)
	perPage100 := uint64(100)

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
		}, {
			name: "Get Active products categories 04",
			filter: &model.Filter{
				Active:  &isPrime,
				Page:    &page,
				PerPage: &perPage100,
			},
			expected: testProductCategories[:15],
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

// test get
func get(t *testing.T) {
	// Create a storage with real database client
	storage := initStorage(t)

	// Define the test cases
	testCases := []struct {
		name string
		id   string
		get  *model.ProductCategory
	}{
		{
			name: "Get 01",
			id:   testProductCategories[0].ID,
			get:  testProductCategories[0],
		}, {
			name: "Get 02",
			id:   testProductCategories[1].ID,
			get:  testProductCategories[1],
		}, {
			name: "Get 03",
			id:   testProductCategories[2].ID,
			get:  testProductCategories[2],
		}, {
			name: "Get 04",
			id:   testProductCategories[3].ID,
			get:  testProductCategories[3],
		}, {
			name: "Get 05",
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

// test create with id
func createWithId(t *testing.T) {
	// Create a storage with real database client
	storage := initStorage(t)

	// Define the test cases
	testCases := []struct {
		name string
		sent *model.ProductCategory
		get  *model.ProductCategory
	}{
		{
			name: "Create with id 1",
			sent: testProductCategoriesNew[0],
			get:  testProductCategoriesNew[0],
		}, {
			name: "Create with id 2",
			sent: testProductCategoriesNew[1],
			get:  testProductCategoriesNew[1],
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
func createWithoutId(t *testing.T) {
	// Create a storage with real database client
	storage := initStorage(t)

	// Define the test cases
	testCases := []struct {
		name string
		sent *model.ProductCategory
		get  *model.ProductCategory
	}{
		{
			name: "Create without id - 1",
			sent: testProductCategoriesNew[2],
			get:  testProductCategoriesNew[2],
		}, {
			name: "Create without id - 2",
			sent: testProductCategoriesNew[3],
			get:  testProductCategoriesNew[3],
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

// test update
func update(t *testing.T) {
	// Create a storage with real database client
	storage := initStorage(t)

	// create new variable from testProductCategories[0]
	testProductCategoryNew := testProductCategoriesNew[0]
	testProductCategoryNew.Name = fmt.Sprintf("%s Test", testProductCategoryNew.Name)
	testProductCategoryNew.Url = fmt.Sprintf("%s-test", testProductCategoryNew.Url)

	// Define the test cases
	testCases := []struct {
		name   string
		id     string
		sent   *model.ProductCategory
		update *model.ProductCategory
	}{
		{
			name:   "Update 01",
			id:     testProductCategoryNew.ID,
			sent:   testProductCategoryNew,
			update: testProductCategoryNew,
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

	// rename back
	_, err := storage.Update(context.Background(), testProductCategories[0], "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1")
	assert.NoError(t, err)

	_, err = storage.Update(context.Background(), testProductCategories[1], "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1")
	assert.NoError(t, err)
}

// test patch
func patch(t *testing.T) {
	// Create a storage with real database client
	storage := initStorage(t)

	// create new variable from testProductCategories[0]
	testProductCategoryNew := testProductCategories[1]
	testProductCategoryNew.Name = fmt.Sprintf("%s Test", testProductCategoryNew.Name)
	testProductCategoryNew.Url = fmt.Sprintf("%s-test", testProductCategoryNew.Url)

	// Define the test cases
	testCases := []struct {
		name   string
		id     string
		fields map[string]interface{}
		get    *model.ProductCategory
	}{
		{
			name: "Patch 01",
			id:   testProductCategories[1].ID,
			fields: map[string]interface{}{
				"Name":  testProductCategoryNew.Name,
				"Url":   testProductCategoryNew.Url,
				"Prime": testProductCategoryNew.Prime,
			},
			get: testProductCategoryNew,
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

	// rename back
	_, err := storage.Patch(context.Background(), testProductCategories[0].ID, map[string]interface{}{
		"Name":  testProductCategories[0].Name,
		"Url":   testProductCategories[0].Url,
		"Prime": testProductCategories[0].Prime,
	}, "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1")
	assert.NoError(t, err)
}

// test delete
func delete(t *testing.T) {
	// Create a storage with real database client
	storage := initStorage(t)

	// Define the test cases
	testCases := []struct {
		name string
		id   string
	}{
		{
			name: "Delete 01",
			id:   testProductCategoriesNew[0].ID,
		}, {
			name: "Delete 02",
			id:   testProductCategoriesNew[1].ID,
		}, {
			name: "Delete 03",
			id:   testProductCategoriesNew[2].ID,
		}, {
			name: "Delete 04",
			id:   testProductCategoriesNew[3].ID,
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

// initStorage with real database client
func initStorage(t *testing.T) model.ProductCategoryStorage {
	// Create a real database client
	cfg := config.GetConfig(context.Background())
	app := appInit.NewAppInit(context.Background(), cfg)
	err := app.SqlDBInit()
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	pgClient := app.SqlDB

	// Initialize a new instance of the model
	storage := model.NewProductCategoryStorage(pgClient)

	return storage
}
