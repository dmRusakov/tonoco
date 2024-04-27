package model_test

import (
	"context"
	"fmt"
	"github.com/dmRusakov/tonoco/internal/appInit"
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/domain/product_category/model"
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

// test data
var testItems = []*entity.ProductCategory{
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9901",
		Url:              "island",
		Name:             "Island range hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        0,
		Prime:            true,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9902",
		Url:              "wall",
		Name:             "Wall range hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        1,
		Prime:            true,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9903",
		Url:              "ait-loop",
		Name:             "Air loop range hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        2,
		Prime:            true,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9904",
		Url:              "built-in",
		Name:             "Built-in range hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        3,
		Prime:            true,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9905",
		Url:              "under-cabinet",
		Name:             "Under Cabinet range hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        4,
		Prime:            true,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9906",
		Url:              "accessories",
		Name:             "Accessories",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        5,
		Prime:            true,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9907",
		Url:              "black",
		Name:             "Black range hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        6,
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9908",
		Url:              "white",
		Name:             "White range hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        7,
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9909",
		Url:              "wood",
		Name:             "Wood range hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        8,
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9910",
		Url:              "stainless-steel",
		Name:             "Stainless Steel range hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        9,
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9911",
		Url:              "glass",
		Name:             "Glass range hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        10,
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9912",
		Url:              "perimeter-filter",
		Name:             "Perimeter Filter range hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        11,
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9913",
		Url:              "murano",
		Name:             "Murano range hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        12,
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9914",
		Url:              "ductless",
		Name:             "Ductless range hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        13,
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9915",
		Url:              "ducted",
		Name:             "Ducted range hoods",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        14,
		Prime:            false,
		Active:           true,
	},
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9999",
		Url:              "discontinued",
		Name:             "Discontinued",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        15,
		Prime:            false,
		Active:           false,
	},
}
var newTestItems = []*entity.ProductCategory{
	{
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9961",
		Url:              "test1-url",
		Name:             "Test 1",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        16,
		Prime:            false,
		Active:           false,
	}, {
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9962",
		Url:              "test2-url",
		Name:             "Test 2",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        17,
		Prime:            false,
		Active:           false,
	}, {
		Url:              "test3-url",
		Name:             "Test 3",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        18,
		Prime:            false,
		Active:           false,
	}, {
		Url:              "test4-url",
		Name:             "Test 4",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        19,
		Prime:            false,
		Active:           false,
	}, {
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9963",
		Url:              "test5-url",
		Name:             "Test 5",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        20,
		Prime:            false,
		Active:           false,
	}, {
		ID:               "1f484cda-c00e-4ed8-a325-9c5e035f9966",
		Url:              "test6-url",
		Name:             "Test 6",
		ShortDescription: "Some text",
		Description:      "Some text",
		SortOrder:        21,
		Prime:            false,
		Active:           false,
	},
}

// Test Product Category
func TestProductCategory(t *testing.T) {
	t.Run("clearTestData", clearTestData)
	defer t.Run("clearTestData", clearTestData)

	t.Run("list", list)
	t.Run("get", get)
	t.Run("getByUrl", getByUrl)
	t.Run("createWithId", createWithId)
	t.Run("createWithoutId", createWithoutId)
	t.Run("update", update)
	t.Run("patch", patch)
	t.Run("updatedAt", updatedAt)
	t.Run("tableIndexCount", tableIndexCount)
	t.Run("maxSortOrder", maxSortOrder)
	t.Run("del", del)
}

// clear test data
func clearTestData(t *testing.T) {
	// Create a storage with real database client
	storage := initStorage(t)

	// get list data from the table
	all, err := storage.List(initContext(), &entity.ProductCategoryFilter{})

	// check if there is an error
	assert.NoError(t, err)

	// go thought list data and del it if is not in the testProductStatuses
	for _, v := range all {
		found := false
		for _, tv := range testItems {
			if v.ID == tv.ID {
				found = true
				break
			}
		}

		if !found {
			err = storage.Delete(initContext(), &v.ID)
			assert.NoError(t, err)
		}
	}

	// go thought list testProductStatuses and create or update them
	for _, v := range testItems {
		// get the product status by the ID
		ps, err := storage.Get(initContext(), &v.ID, nil)

		// check if there is an error
		assert.NoError(t, err)

		// if the item not found create it
		if ps == nil {
			_, err = storage.Create(initContext(), v)
			assert.NoError(t, err)
		} else {
			err = storage.Update(initContext(), v)
			assert.NoError(t, err)
		}
	}
}

// test list
func list(t *testing.T) {
	// Create a storage with real database client
	storage := initStorage(t)

	// test varietals
	isPrime := true
	page1 := uint64(1)
	perPage3 := uint64(3)
	perPage100 := uint64(100)
	searchAir := "air"
	searchSteel := "steel"

	// Define the test cases
	testCases := []struct {
		name     string
		filter   *entity.ProductCategoryFilter
		expected []*entity.ProductCategory
	}{
		{
			name:     "Get List",
			filter:   &entity.ProductCategoryFilter{},
			expected: testItems[:10],
		}, {
			name: "Get 3 products categories 02",
			filter: &entity.ProductCategoryFilter{
				PerPage: &perPage3,
				Page:    &page1,
			},
			expected: testItems[:3],
		}, {
			name: "Get Prime products categories 03",
			filter: &entity.ProductCategoryFilter{
				Prime: &isPrime,
			},
			expected: testItems[0:6],
		}, {
			name: "Get Active products categories 04",
			filter: &entity.ProductCategoryFilter{
				Active:  &isPrime,
				Page:    &page1,
				PerPage: &perPage100,
			},
			expected: testItems[:15],
		}, {
			name: "Get with name like `air` products categories 05",
			filter: &entity.ProductCategoryFilter{
				Search: &searchAir,
			},
			expected: testItems[2:3],
		}, {
			name: "Get with name like `steel` products categories 06",
			filter: &entity.ProductCategoryFilter{
				Search: &searchSteel,
			},
			expected: testItems[9:10],
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call List method
			result, err := storage.List(initContext(), (*entity.ProductCategoryFilter)(tc.filter))

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
		get  *entity.ProductCategory
	}{
		{
			name: "Get 01",
			id:   testItems[0].ID,
			get:  testItems[0],
		}, {
			name: "Get 02",
			id:   testItems[1].ID,
			get:  testItems[1],
		}, {
			name: "Get 03",
			id:   testItems[2].ID,
			get:  testItems[2],
		}, {
			name: "Get 04",
			id:   testItems[3].ID,
			get:  testItems[3],
		}, {
			name: "Get 05",
			id:   testItems[4].ID,
			get:  testItems[4],
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := storage.Get(initContext(), &tc.id, nil)

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

// test get by url
func getByUrl(t *testing.T) {
	// Create a storage with real database client
	storage := initStorage(t)

	// Define the test cases
	testCases := []struct {
		name string
		url  string
		get  *entity.ProductCategory
	}{
		{
			name: "Get by url 01",
			url:  testItems[0].Url,
			get:  testItems[0],
		}, {
			name: "Get by url 02",
			url:  testItems[1].Url,
			get:  testItems[1],
		}, {
			name: "Get by url 03",
			url:  testItems[2].Url,
			get:  testItems[2],
		}, {
			name: "Get by url 04",
			url:  testItems[3].Url,
			get:  testItems[3],
		}, {
			name: "Get by url 05",
			url:  testItems[4].Url,
			get:  testItems[4],
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := storage.Get(initContext(), nil, &tc.url)

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
		sent *entity.ProductCategory
		get  *entity.ProductCategory
	}{
		{
			name: "Create with id 1",
			sent: newTestItems[0],
			get:  newTestItems[0],
		}, {
			name: "Create with id 2",
			sent: newTestItems[1],
			get:  newTestItems[1],
		}, {
			name: "Create with id 3",
			sent: newTestItems[4],
			get:  newTestItems[4],
		}, {
			name: "Create with id 4",
			sent: newTestItems[5],
			get:  newTestItems[5],
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the Create method
			id, err := storage.Create(initContext(), tc.sent)
			assert.NoError(t, err)

			// Call the Get method
			result, err := storage.Get(initContext(), id, nil)
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
		sent *entity.ProductCategory
		get  *entity.ProductCategory
	}{
		{
			name: "Create without id - 1",
			sent: newTestItems[2],
			get:  newTestItems[2],
		}, {
			name: "Create without id - 2",
			sent: newTestItems[3],
			get:  newTestItems[3],
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the Create method
			id, err := storage.Create(initContext(), tc.sent)
			assert.NoError(t, err)
			tc.get.ID = *id

			// Call the Get method
			result, err := storage.Get(initContext(), id, nil)
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

	// create new variable from testItems[0]
	testItemNew := newTestItems[0]
	testItemNew.Name = fmt.Sprintf("%s Test", testItemNew.Name)
	testItemNew.Url = fmt.Sprintf("%s-test", testItemNew.Url)

	// Define the test cases
	testCases := []struct {
		name   string
		sent   *entity.ProductCategory
		update *entity.ProductCategory
	}{
		{
			name:   "Update 01",
			sent:   testItemNew,
			update: testItemNew,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call Update method
			err := storage.Update(initContext(), tc.sent)
			assert.NoError(t, err)

			// Call the Get method
			result, err := storage.Get(initContext(), &tc.update.ID, nil)
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
func patch(t *testing.T) {
	// Create a storage with real database client
	storage := initStorage(t)

	// create new variable from testItems[0]
	testItemNew := newTestItems[1]
	testItemNew.Name = fmt.Sprintf("%s Test", testItemNew.Name)
	testItemNew.Url = fmt.Sprintf("%s-test", testItemNew.Url)

	// Define the test cases
	testCases := []struct {
		name   string
		id     string
		fields map[string]interface{}
		get    *entity.ProductCategory
	}{
		{
			name: "Patch 01",
			id:   newTestItems[1].ID,
			fields: map[string]interface{}{
				"Name":  testItemNew.Name,
				"Url":   testItemNew.Url,
				"Prime": testItemNew.Prime,
			},
			get: testItemNew,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the Patch method
			err := storage.Patch(initContext(), &tc.id, &tc.fields)
			assert.NoError(t, err)

			// Call the Get method
			result, err := storage.Get(initContext(), &tc.id, nil)
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

// test UpdatedAt
func updatedAt(t *testing.T) {
	// Create a storage with real database client
	storage := initStorage(t)

	// Define the test cases
	testCases := []struct {
		name string
		id   string
		sent *entity.ProductCategory
	}{
		{
			name: "Updated at 01",
			id:   newTestItems[0].ID,
			sent: newTestItems[0],
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		// Call the UpdatedAt method
		updatedAtBefore, err := storage.UpdatedAt(initContext(), &tc.id)
		assert.NoError(t, err)

		// Call the Update method
		err = storage.Update(initContext(), tc.sent)
		assert.NoError(t, err)

		// Call the UpdatedAt method
		updatedAtAfter, err := storage.UpdatedAt(initContext(), &tc.id)
		assert.NoError(t, err)

		// Assert that the updatedAtAfter is greater than updatedAtBefore
		assert.NotEqual(t, updatedAtBefore, updatedAtAfter)
	}
}

func tableIndexCount(t *testing.T) {
	// Create a storage with real database client
	storage := initStorage(t)

	// Define the test cases
	testCases := []struct {
		name  string
		id    string
		patch map[string]interface{}
	}{
		{
			name: "Table Updated 01",
			id:   newTestItems[5].ID,
			patch: map[string]interface{}{
				"Name": fmt.Sprintf("%s - patched", newTestItems[5].Name),
			},
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		// Call the TableIndexCount method
		tableUpdatedBefore, err := storage.TableIndexCount(context.Background())
		assert.NoError(t, err)
		assert.NotEmpty(t, tableUpdatedBefore)

		// Call the Patch method
		err = storage.Patch(initContext(), &tc.id, &tc.patch)
		assert.NoError(t, err)

		// Call the TableIndexCount method
		tableUpdatedAfter, err := storage.TableIndexCount(context.Background())
		assert.NoError(t, err)
		assert.NotEmpty(t, tableUpdatedAfter)
	}
}

// max sort order
func maxSortOrder(t *testing.T) {
	// Create a storage with real database client
	storage := initStorage(t)

	// Call the MaxSortOrder method
	sortOrder, err := storage.MaxSortOrder(context.Background())
	assert.NoError(t, err)

	// Assert that the result is not empty
	assert.NotEmpty(t, sortOrder)
}

// test del
func del(t *testing.T) {
	// Create a storage with real database client
	storage := initStorage(t)

	// Define the test cases
	testCases := []struct {
		name string
		id   string
	}{
		{
			name: "Delete 01",
			id:   newTestItems[0].ID,
		}, {
			name: "Delete 02",
			id:   newTestItems[1].ID,
		}, {
			name: "Delete 03",
			id:   newTestItems[2].ID,
		}, {
			name: "Delete 04",
			id:   newTestItems[3].ID,
		}, {
			name: "Delete 05",
			id:   newTestItems[4].ID,
		}, {
			name: "Delete 06",
			id:   newTestItems[5].ID,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the Delete method
			err := storage.Delete(initContext(), &tc.id)

			// Assert that there was no error
			assert.NoError(t, err)

			// Call the Get method
			_, err = storage.Get(initContext(), &tc.id, nil)
			assert.Error(t, err)
		})
	}

}

// initStorage will create a storage with real database client
func initStorage(t *testing.T) *model.Model {
	// Create a real database client
	cfg := config.GetConfig(context.Background())
	app := appInit.NewAppInit(initContext(), cfg)
	err := app.SqlDBInit()
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	pgClient := app.SqlDB

	// Initialize Storage
	return model.NewStorage(pgClient)
}

// init context
func initContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "user_id", "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1")
	return ctx
}
