package model_test

import (
	"context"
	"fmt"
	"github.com/dmRusakov/tonoco/internal/appInit"
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/domain/shipping_class/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testItems = []*model.Item{
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
		Name:      "Freight",
		Url:       "freight",
		SortOrder: 0,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12",
		Name:      "Ground",
		Url:       "ground",
		SortOrder: 1,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13",
		Name:      "Ground - small",
		Url:       "ground-small",
		SortOrder: 2,
		Active:    true,
	},
}
var newTestItems = []*model.Item{
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a00",
		Name:      "New",
		Url:       "new",
		SortOrder: 3,
		Active:    false,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a01",
		Name:      "New1",
		Url:       "new1",
		SortOrder: 4,
		Active:    false,
	}, {
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a02",
		Name:      "New2",
		Url:       "new2",
		SortOrder: 5,
		Active:    false,
	}, {
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a03",
		Name:      "New3",
		Url:       "new3",
		SortOrder: 6,
		Active:    false,
	},
	{
		Name:      "New4",
		Url:       "new4",
		SortOrder: 7,
		Active:    false,
	}, {
		Name:      "New5",
		Url:       "new5",
		SortOrder: 8,
		Active:    false,
	},
}

// Test Product Status
func TestShippingClass(t *testing.T) {
	// prepare data
	t.Run("clearTestData", clearTestData)
	defer t.Run("clearTestData", clearTestData)

	// run tests
	t.Run("list", list)
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

	// get list data from the table
	all, err := storage.List(initContext(), &model.Filter{})

	// check if there is an error
	assert.NoError(t, err)

	// go thought list data and del it if is not in the testItems
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

	// Add this in your clearTestData function
	for i, v := range testItems {
		if v.SortOrder == 0 {
			v.SortOrder = uint64(i + 1)
		}
		// get the product status by the ID
		ps, err := storage.Get(initContext(), &v.ID, nil)

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

// test list
func list(t *testing.T) {
	// Create a storage with real database client
	storage := initStorage(t)

	// test varietals
	isActive := true
	perPage1 := uint64(1)
	perPage2 := uint64(2)
	page2 := uint64(2)
	searchSmall := "small"
	twoIDs := []string{testItems[0].ID, testItems[1].ID}
	twoUrls := []string{testItems[1].Url, testItems[2].Url}

	// Define the test cases
	tests := []struct {
		name     string
		filter   *model.Filter
		expected []*model.Item
	}{
		{
			name:     "Get list",
			filter:   &model.Filter{},
			expected: testItems,
		}, {
			name:     "Get list with 2 IDs",
			filter:   &model.Filter{IDs: &twoIDs, PerPage: &perPage2},
			expected: testItems[:2],
		}, {
			name: "Get list with 2 Urls",
			filter: &model.Filter{
				Urls:   &twoUrls,
				Active: &isActive,
			},
			expected: testItems[1:3],
		},
		{
			name:     "Get list active",
			filter:   &model.Filter{Active: &isActive},
			expected: testItems,
		}, {
			name:     "Get with name like 'small'",
			filter:   &model.Filter{Search: &searchSmall},
			expected: testItems[2:3],
		}, {
			name:     "Get with perPage 2",
			filter:   &model.Filter{PerPage: &perPage1},
			expected: testItems[:1],
		}, {
			name:     "Get with perPage 2 and page 2",
			filter:   &model.Filter{PerPage: &perPage2, Page: &page2},
			expected: testItems[2:3],
		},
	}

	// Run the test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Call List method
			result, err := storage.List(initContext(), tc.filter)

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
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call Get method
			id := tc.id
			result, err := storage.Get(initContext(), &id, nil)

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
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call Get method
			url := tc.url
			result, err := storage.Get(initContext(), nil, &url)

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
		get  *model.Item
	}{
		{
			name: "Create with ID",
			get:  newTestItems[0],
		}, {
			name: "Create with ID",
			get:  newTestItems[1],
		}, {
			name: "Create with ID",
			get:  newTestItems[2],
		}, {
			name: "Create with ID",
			get:  newTestItems[3],
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call Create method
			result, err := storage.Create(initContext(), tc.get)
			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the result is equal to the expected
			assert.Equal(t, tc.get.ID, result.ID)
			assert.Equal(t, tc.get.Name, result.Name)
			assert.Equal(t, tc.get.Url, result.Url)
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
			assert.Equal(t, tc.get.ID, result.ID)
			assert.Equal(t, tc.get.Name, result.Name)
			assert.Equal(t, tc.get.Url, result.Url)
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
			result, err := storage.Patch(initContext(), &testProductCategoryNew.ID, &tc.fields)

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
		id := tc.id

		// Call the UpdatedAt method
		updatedAtBefore, err := storage.UpdatedAt(initContext(), &id)

		// Assert that there was no error
		assert.NoError(t, err)

		// Call the Update method
		_, err = storage.Update(initContext(), tc.sent)

		// Assert that there was no error
		assert.NoError(t, err)

		// Call the UpdatedAt method
		updatedAtAfter, err := storage.UpdatedAt(initContext(), &id)

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

	// Call the TableIndexCount method
	tableIndexCount, err := storage.TableIndexCount(context.Background())

	// Assert that there was no error
	assert.NoError(t, err)

	// Assert that the result is not empty
	assert.NotEmpty(t, tableIndexCount)
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
		err := storage.Delete(initContext(), &tc.id)

		// Assert that there was no error
		assert.NoError(t, err)

		// Call the Get method
		id := tc.id

		_, err = storage.Get(initContext(), &id, nil)
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
