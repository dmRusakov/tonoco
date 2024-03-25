package model_test

import (
	"context"
	"fmt"
	"github.com/dmRusakov/tonoco/internal/appInit"
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/domain/specification/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testItems = []*model.Item{
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380107",
		Name:      "Mounting Type",
		Url:       "mounting-type",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 0,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380108",
		Name:      "Width",
		Url:       "width",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 1,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380109",
		Name:      "Depth",
		Url:       "depth",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 2,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380110",
		Name:      "Height",
		Url:       "height",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 3,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380111",
		Name:      "Recommended Range Width",
		Url:       "recommended-range-width",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 4,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380112",
		Name:      "Height Above Cooktop",
		Url:       "height-above-cooktop",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 5,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380113",
		Name:      "Color / Finish",
		Url:       "color-finish",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 6,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380114",
		Name:      "Design",
		Url:       "design",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 7,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380115",
		Name:      "Materials",
		Url:       "materials",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 8,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380116",
		Name:      "Lighting Type",
		Url:       "lighting-type",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 9,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380117",
		Name:      "# of Lights",
		Url:       "of-lights",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 10,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380118",
		Name:      "# of Speeds",
		Url:       "of-speeds",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 11,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380119",
		Name:      "Control Panel Type",
		Url:       "control-panel-type",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 12,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380120",
		Name:      "Filter Type",
		Url:       "filter-type",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 13,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380121",
		Name:      "Airflow",
		Url:       "airflow",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 14,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380122",
		Name:      "Blower Type",
		Url:       "blower-type",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 15,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380123",
		Name:      "Noise Level",
		Url:       "noise-level",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 16,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380124",
		Name:      "Duct Size",
		Url:       "duct-size",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 17,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380125",
		Name:      "Exhaust Type",
		Url:       "exhaust-type",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 18,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380126",
		Name:      "Power Requirements",
		Url:       "power-requirements",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 19,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380127",
		Name:      "Certifications",
		Url:       "certifications",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 20,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380128",
		Name:      "Warranty",
		Url:       "warranty",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 21,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380129",
		Name:      "Order Processing Time",
		Url:       "order-processing-time",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 22,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380130",
		Name:      "Shipping Speed",
		Url:       "shipping-speed",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 23,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380131",
		Name:      "Ships Via",
		Url:       "ships-via",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 24,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380132",
		Name:      "Country of Production",
		Url:       "country-of-production",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 25,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380133",
		Name:      "Filter - Width (Group)",
		Url:       "width-group",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 26,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380135",
		Name:      "Shipping Weight",
		Url:       "shipping-weight",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 27,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380136",
		Name:      "Brand",
		Url:       "brand",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 28,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380137",
		Name:      "Item Weight",
		Url:       "item-weight",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 29,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380138",
		Name:      "Diameter",
		Url:       "diameter",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 30,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380139",
		Name:      "Additional Lighting",
		Url:       "additional-lighting",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 31,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380142",
		Name:      "Filter – Color",
		Url:       "filter-color",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 32,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380143",
		Name:      "Filter – Material",
		Url:       "filter-material",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 33,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380144",
		Name:      "Filter – Exhaust Type",
		Url:       "filter-exhaust-type",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 34,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380145",
		Name:      "Filter – Design",
		Url:       "filter-design",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 35,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380146",
		Name:      "Length",
		Url:       "max-usable-length",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 36,
		Active:    true,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380147",
		Name:      "Filter – Accessories",
		Url:       "filter-accessories",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 37,
		Active:    true,
	},
}
var newTestItems = []*model.Item{
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380148",
		Name:      "Test 01",
		Url:       "test-01",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 39,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380149",
		Name:      "Test 02",
		Url:       "test-02",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 40,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380150",
		Name:      "Test 03",
		Url:       "test-03",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 41,
	},
	{
		ID:        "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380151",
		Name:      "Test 04",
		Url:       "test-04",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 42,
	},
	{
		Name:      "Test 05",
		Url:       "test-05",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 43,
	},
	{
		Name:      "Test 06",
		Url:       "test-06",
		Type:      "a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13",
		SortOrder: 44,
	},
}

// Test Product Status
func TestSpecifications(t *testing.T) {
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

	// get all data from the table
	all, err := storage.List(initContext(), &model.Filter{})

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
	page2 := uint64(2)
	perPage5 := uint64(5)
	perPage100 := uint64(100)
	search := "panel"

	// Define the test cases
	tests := []struct {
		name     string
		filter   *model.Filter
		expected []*model.Item
	}{
		{
			name:     "Get all",
			filter:   &model.Filter{},
			expected: testItems[0:10],
		}, {
			name: "Get all active",
			filter: &model.Filter{
				Active:  &isActive,
				PerPage: &perPage100,
			},
			expected: testItems,
		}, {
			name:     "Get with name in search",
			filter:   &model.Filter{Search: &search},
			expected: testItems[6:7],
		}, {
			name:     "Get page 2",
			filter:   &model.Filter{Page: &page2},
			expected: testItems[10:20],
		}, {
			name:     "Get per page 5",
			filter:   &model.Filter{PerPage: &perPage5},
			expected: testItems[:5],
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
			result, err := storage.Get(initContext(), &tc.id, nil)

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
			result, err := storage.Get(initContext(), nil, &tc.url)

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
		// Call the UpdatedAt method
		updatedAtBefore, err := storage.UpdatedAt(initContext(), &tc.id)

		// Assert that there was no error
		assert.NoError(t, err)

		// Call the Update method
		_, err = storage.Update(initContext(), tc.sent)

		// Assert that there was no error
		assert.NoError(t, err)

		// Call the UpdatedAt method
		updatedAtAfter, err := storage.UpdatedAt(initContext(), &tc.id)

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
	_, err := storage.TableIndexCount(context.Background())

	// Assert that there was no error
	assert.NoError(t, err)
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
		_, err = storage.Get(initContext(), &tc.id, nil)
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
