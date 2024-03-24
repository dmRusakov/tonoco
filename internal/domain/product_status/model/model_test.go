package model_test

import (
	"context"
	"fmt"
	"github.com/dmRusakov/tonoco/internal/appInit"
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/domain/product_status/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testItems = []*model.Item{
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
		Name:      "Public",
		Url:       "public",
		SortOrder: 0,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12",
		Name:      "Privet",
		Url:       "private",
		SortOrder: 1,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13",
		Name:      "Out of stock",
		Url:       "out-of-stock",
		SortOrder: 2,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a14",
		Name:      "Discontinued",
		Url:       "discontinued",
		SortOrder: 3,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a15",
		Name:      "Archived",
		Url:       "archived",
		SortOrder: 4,
		Active:    true,
	},
}
var newTestItems = []*model.Item{
	{
		ID:     "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a00",
		Name:   "New",
		Url:    "new",
		Active: false,
	},
	{
		ID:     "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a01",
		Name:   "New1",
		Url:    "new1",
		Active: false,
	}, {
		ID:     "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a02",
		Name:   "New2",
		Url:    "new2",
		Active: false,
	}, {
		ID:     "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a03",
		Name:   "New3",
		Url:    "new3",
		Active: false,
	},
	{
		Name:   "New4",
		Url:    "new2",
		Active: false,
	}, {
		Name:   "New5",
		Url:    "new3",
		Active: false,
	},
}

// Test Product Status
func TestProductStatus(t *testing.T) {
	// prepare data
	t.Run("clearTestData", clearTestData)
	defer t.Run("clearTestData", clearTestData)

	// run tests
	t.Run("list", all)
	t.Run("get", get)
	t.Run("getByUrl", getByUrl)
	t.Run("createWithId", createWithId)
	t.Run("createWithoutId", createWithoutId)
	t.Run("update", update)
	t.Run("patch", patch)
	t.Run("updatedAt", updatedAt)
	t.Run("tableUpdated", tableUpdated)
	t.Run("del", del)
	t.Run("maxSortOrder", maxSortOrder)
}

// Create test data for the test cases
func clearTestData(t *testing.T) {
	// Create a storage with real database client
	storage := initStorage(t)

	// get all data from the table
	all, err := storage.All(initContext(), &model.Filter{})

	// check if there is an error
	assert.NoError(t, err)

	// go thought all data and del it if is not in the testItems
	for _, v := range all {
		found := false
		for _, tv := range testItems {
			if v.ID == tv.ID {
				found = true
				break
			}
		}

		if !found {
			err = storage.Delete(initContext(), v.ID)
			assert.NoError(t, err)
		}
	}

	// go thought all testItems and create or update them
	for i, v := range testItems {
		v.SortOrder = uint32(i + 1) // Ensure SortOrder is greater than 0
		// get the product status by the ID
		ps, err := storage.Get(initContext(), v.ID)

		// check if there is an error
		assert.NoError(t, err)

		// if the item not found create it
		if ps == nil {
			_, err = storage.Create(initContext(), v)
			assert.NoError(t, err)
		} else {
			_, err = storage.Update(initContext(), v)
			assert.NoError(t, err)
		}
	}
}

// test all
func all(t *testing.T) {
	// Create a storage with real database client
	storage := initStorage(t)

	// test varietals
	isActive := true
	perPage3 := uint64(3)
	page1 := uint64(1)
	searchOut := "out"

	// Define the test cases
	tests := []struct {
		name     string
		filter   *model.Filter
		expected []*model.Item
	}{
		{
			name:     "Get all",
			filter:   &model.Filter{},
			expected: testItems,
		}, {
			name:     "Get all active",
			filter:   &model.Filter{Active: &isActive},
			expected: testItems,
		}, {
			name:     "Get with name like 'out'",
			filter:   &model.Filter{Search: &searchOut},
			expected: testItems[2:3],
		}, {
			name:     "Get with perPage 3",
			filter:   &model.Filter{PerPage: &perPage3},
			expected: testItems[:3],
		}, {
			name:     "Get with perPage 3 and page 1",
			filter:   &model.Filter{PerPage: &perPage3, Page: &page1},
			expected: testItems[:3],
		},
	}

	// Run the test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Call List method
			result, err := storage.All(initContext(), tc.filter)

			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the result length is equal to the expected length
			assert.Len(t, result, len(tc.expected))

			// Assert that each element in the result is equal to the corresponding element in the expected slice
			for i, v := range tc.expected {
				assert.Equal(t, tc.expected[i].ID, v.ID)
				assert.Equal(t, tc.expected[i].Name, v.Name)
				assert.Equal(t, tc.expected[i].Url, v.Url)
				assert.Equal(t, tc.expected[i].SortOrder, v.SortOrder)
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
		get  *model.Item
	}{
		{
			name: "Get by ID",
			id:   testItems[0].ID,
			get:  testItems[0],
		}, {
			name: "Get by ID",
			id:   testItems[1].ID,
			get:  testItems[1],
		}, {
			name: "Get by ID",
			id:   testItems[2].ID,
			get:  testItems[2],
		}, {
			name: "Get by ID",
			id:   testItems[3].ID,
			get:  testItems[3],
		}, {
			name: "Get by ID",
			id:   testItems[4].ID,
			get:  testItems[4],
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call Get method
			result, err := storage.Get(initContext(), tc.id)

			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the result is equal to the expected
			assert.Equal(t, tc.get.ID, result.ID)
			assert.Equal(t, tc.get.Name, result.Name)
			assert.Equal(t, tc.get.Url, result.Url)
			assert.Equal(t, tc.get.SortOrder, result.SortOrder)
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
		get  *model.Item
	}{
		{
			name: "Get by URL",
			url:  testItems[0].Url,
			get:  testItems[0],
		}, {
			name: "Get by URL",
			url:  testItems[1].Url,
			get:  testItems[1],
		}, {
			name: "Get by URL",
			url:  testItems[2].Url,
			get:  testItems[2],
		}, {
			name: "Get by URL",
			url:  testItems[3].Url,
			get:  testItems[3],
		}, {
			name: "Get by URL",
			url:  testItems[4].Url,
			get:  testItems[4],
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call Get method
			result, err := storage.GetByURL(initContext(), tc.url)

			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the result is equal to the expected
			assert.Equal(t, tc.get.ID, result.ID)
			assert.Equal(t, tc.get.Name, result.Name)
			assert.Equal(t, tc.get.Url, result.Url)
			assert.Equal(t, tc.get.SortOrder, result.SortOrder)
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
		new  *model.Item
	}{
		{
			name: "Create with ID",
			new:  newTestItems[0],
		}, {
			name: "Create with ID",
			new:  newTestItems[1],
		}, {
			name: "Create with ID",
			new:  newTestItems[2],
		}, {
			name: "Create with ID",
			new:  newTestItems[3],
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call Create method
			result, err := storage.Create(initContext(), tc.new)

			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the result is equal to the expected
			assert.Equal(t, tc.new.ID, result.ID)
			assert.Equal(t, tc.new.Name, result.Name)
			assert.Equal(t, tc.new.Url, result.Url)
			assert.Equal(t, tc.new.SortOrder, result.SortOrder)
			assert.Equal(t, tc.new.Active, result.Active)
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
		sent *model.Item
		get  *model.Item
	}{
		{
			name: "Create without ID",
			sent: newTestItems[4],
			get:  newTestItems[4],
		}, {
			name: "Create without ID",
			sent: newTestItems[5],
			get:  newTestItems[5],
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call Create method
			result, err := storage.Create(initContext(), tc.sent)

			// update ID
			tc.get.ID = result.ID

			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the result is equal to the expected
			assert.Equal(t, tc.get.Name, result.Name)
			assert.Equal(t, tc.get.Url, result.Url)
			assert.Equal(t, tc.get.SortOrder, result.SortOrder)
			assert.Equal(t, tc.get.Active, result.Active)
		})
	}
}

// test update
func update(t *testing.T) {
	// Create a storage with real database client
	storage := initStorage(t)

	// create new variable from the testItems[0]
	testProductCategoryNew := *testItems[0]
	testProductCategoryNew.Name = fmt.Sprintf("%s - updated", testProductCategoryNew.Name)
	testProductCategoryNew.Url = fmt.Sprintf("%s - updated", testProductCategoryNew.Url)

	// Define the test cases
	testCases := []struct {
		name   string
		sent   *model.Item
		update *model.Item
	}{
		{
			name:   "Update",
			sent:   &testProductCategoryNew,
			update: &testProductCategoryNew,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call Update method
			result, err := storage.Update(initContext(), tc.sent)

			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the result is equal to the expected
			assert.Equal(t, tc.update.ID, result.ID)
			assert.Equal(t, tc.update.Name, result.Name)
			assert.Equal(t, tc.update.Url, result.Url)
			assert.Equal(t, tc.update.SortOrder, result.SortOrder)
			assert.Equal(t, tc.update.Active, result.Active)
		})
	}
}

// test patch
func patch(t *testing.T) {
	// Create a storage with real database client
	storage := initStorage(t)

	// create new variable from the testItems[0]
	testProductCategoryNew := *testItems[0]
	testProductCategoryNew.Name = fmt.Sprintf("%s - patched", testProductCategoryNew.Name)
	testProductCategoryNew.Url = fmt.Sprintf("%s - patched", testProductCategoryNew.Url)

	// Define the test cases
	testCases := []struct {
		name   string
		id     string
		fields map[string]interface{}
		get    *model.Item
	}{
		{
			name: "Patch",
			fields: map[string]interface{}{
				"Name": testProductCategoryNew.Name,
				"Url":  testProductCategoryNew.Url,
			},
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call Patch method
			result, err := storage.Patch(initContext(), testProductCategoryNew.ID, tc.fields)

			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the result is equal to the expected
			assert.Equal(t, testProductCategoryNew.ID, result.ID)
			assert.Equal(t, testProductCategoryNew.Name, result.Name)
			assert.Equal(t, testProductCategoryNew.Url, result.Url)
			assert.Equal(t, testProductCategoryNew.SortOrder, result.SortOrder)
			assert.Equal(t, testProductCategoryNew.Active, result.Active)
		})
	}
}

// test updated at
func updatedAt(t *testing.T) {
	// Create a storage with real database client
	storage := initStorage(t)

	// Define the test cases
	testCases := []struct {
		name string
		id   string
		sent *model.Item
	}{
		{
			name: "Updated at 01",
			id:   newTestItems[4].ID,
			sent: newTestItems[4],
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		// Call the UpdatedAt method
		updatedAtBefore, err := storage.UpdatedAt(initContext(), tc.id)

		// Assert that there was no error
		assert.NoError(t, err)

		// Call the Update method
		_, err = storage.Update(initContext(), tc.sent)

		// Assert that there was no error
		assert.NoError(t, err)

		// Call the UpdatedAt method
		updatedAtAfter, err := storage.UpdatedAt(initContext(), tc.id)

		// Assert that there was no error
		assert.NoError(t, err)

		// Assert that the updatedAtAfter is greater than updatedAtBefore
		assert.NotEqual(t, updatedAtBefore, updatedAtAfter)
	}
}

// test TableIndexCount
func tableUpdated(t *testing.T) {
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
		tableUpdatedBefore, err := storage.TableUpdated(context.Background())

		// Assert that there was no error
		assert.NoError(t, err)

		// Assert that the result is not empty
		assert.NotEmpty(t, tableUpdatedBefore)

		// Call the Patch method
		_, err = storage.Patch(initContext(), tc.id, tc.patch)

		// Assert that there was no error
		assert.NoError(t, err)

		// Call the TableIndexCount method
		tableUpdatedAfter, err := storage.TableUpdated(context.Background())

		// Assert that there was no error
		assert.NoError(t, err)

		// Assert that the result is not empty
		assert.NotEmpty(t, tableUpdatedAfter)

		// Assert that the tableUpdatedAfter is greater than tableUpdatedBefore
		assert.NotEqual(t, tableUpdatedBefore, tableUpdatedAfter)

	}
}

// max sort order
func maxSortOrder(t *testing.T) {
	// Create a storage with real database client
	storage := initStorage(t)

	// Call the MaxSortOrder method
	sortOrder, err := storage.MaxSortOrder(context.Background())

	// Assert that there was no error
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
		// Call the Delete method
		err := storage.Delete(initContext(), tc.id)

		// Assert that there was no error
		assert.NoError(t, err)

		// Call the Get method
		_, err = storage.Get(initContext(), tc.id)
		assert.Error(t, err)
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
