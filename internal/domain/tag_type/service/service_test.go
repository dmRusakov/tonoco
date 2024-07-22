package service_test

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dmRusakov/tonoco/internal/appInit"
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/domain/tag_type/model"
	"github.com/dmRusakov/tonoco/internal/domain/tag_type/service"
	"github.com/dmRusakov/tonoco/internal/entity"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"testing"
)

func TestTagType(t *testing.T) {
	t.Run("testItemValidations", testItemValidations)
	t.Run("testValidateFields", testValidateFields)
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
	t.Run("del", del)
	t.Run("maxSortOrder", maxSortOrder)
}

func testItemValidations(t *testing.T) {
	t.Parallel()

	srv, _ := initStorage(t)

	// Define the test cases
	testCases := []struct {
		name     string
		item     service.Item
		fields   map[string]interface{}
		countErr int
		err      []entity.Error
	}{
		{
			name: "Good Item",
			item: service.Item{
				Id:               "a0eebc99-9c0b-4ef8-bb6d-6bb9bd381001",
				Name:             "Category",
				Url:              "category",
				ShortDescription: "",
				Description:      "",
				Required:         true,
				Active:           true,
				Prime:            true,
				ListItem:         true,
				Filter:           true,
				SortOrder:        0,
				Type:             "select",
				Prefix:           "",
				Suffix:           "",
				CreatedBy:        "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
				UpdatedBy:        "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
			},
			fields:   nil,
			countErr: 0,
			err:      nil,
		}, {
			name: "Item without name",
			item: service.Item{
				Url:              "category",
				ShortDescription: "",
				Description:      "",
				Required:         true,
				Active:           true,
				Prime:            true,
				ListItem:         true,
				Filter:           true,
				SortOrder:        0,
				Type:             "select",
				Prefix:           "",
				Suffix:           "",
				CreatedBy:        "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
				UpdatedBy:        "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
			},
			fields:   nil,
			countErr: 1,
			err:      []entity.Error{service.NoName},
		}, {
			name: "Item without url",
			item: service.Item{
				Name:             "Category",
				ShortDescription: "",
				Description:      "",
				Required:         true,
				Active:           true,
				Prime:            true,
				ListItem:         true,
				Filter:           true,
				SortOrder:        0,
				Type:             "select",
				Prefix:           "",
				Suffix:           "",
				CreatedBy:        "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
				UpdatedBy:        "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
			},
			fields:   nil,
			countErr: 1,
			err:      []entity.Error{service.NoUrl},
		}, {
			name: "Item without type",
			item: service.Item{
				Name:             "Category",
				Url:              "category",
				ShortDescription: "",
				Description:      "",
				Required:         true,
				Active:           true,
				Prime:            true,
				ListItem:         true,
				Filter:           true,
				SortOrder:        0,
				Prefix:           "",
				Suffix:           "",
				CreatedBy:        "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
				UpdatedBy:        "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
			},
			fields:   nil,
			countErr: 1,
			err:      []entity.Error{service.NoType},
		}, {
			name: "Item with wrong type",
			item: service.Item{
				Name:             "Category",
				Url:              "category",
				ShortDescription: "",
				Description:      "",
				Required:         true,
				Active:           true,
				Prime:            true,
				ListItem:         true,
				Filter:           true,
				SortOrder:        0,
				Type:             "wrong",
				Prefix:           "",
				Suffix:           "",
				CreatedBy:        "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
				UpdatedBy:        "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
			},
			fields:   nil,
			countErr: 1,
			err:      []entity.Error{service.InvalidType},
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validations := srv.Validate(&tc.item, &tc.fields)
			assert.Equal(t, tc.countErr, len(validations))
			assert.Equal(t, validations, tc.err)
		})
	}
}

func testValidateFields(t *testing.T) {
	t.Parallel()

	srv, _ := initStorage(t)

	// Define the test cases
	testCases := []struct {
		name     string
		fields   map[string]interface{}
		countErr int
		err      []entity.Error
	}{
		{
			name: "Good Fields",
			fields: map[string]interface{}{
				"Name": "name",
				"Url":  "url",
			},
			countErr: 0,
			err:      nil,
		}, {
			name: "Fields with empty name",
			fields: map[string]interface{}{
				"Name": "",
				"Url":  "jh",
			},
			countErr: 1,
			err:      []entity.Error{service.NoName},
		}, {
			name: "Fields with empty url",
			fields: map[string]interface{}{
				"Name": "name",
				"Url":  "",
			},
			countErr: 1,
			err:      []entity.Error{service.NoUrl},
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validations := srv.Validate(nil, &tc.fields)
			assert.Equal(t, tc.countErr, len(validations))
			assert.Equal(t, validations, tc.err)
		})
	}
}

// Create test data for the test cases
func clearTestData(t *testing.T) {
	// Create a storage with real database client
	_, cfg := initStorage(t)

	// Connect to the database and execute SQL commands
	dbConn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DataStorage.Host, cfg.DataStorage.User, cfg.DataStorage.Password, cfg.DataStorage.DB, cfg.DataStorage.Port)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Failed to close database: %v", err)
		}
	}(db)

	sqlBytes, err := ioutil.ReadFile("/home/dm/Apps/tonoco/config/dbData/DataStorage/init/14_tag_type.sql")
	if err != nil {
		log.Fatalf("Failed to read SQL file: %v", err)
	}

	_, err = db.Exec(string(sqlBytes))
	if err != nil {
		log.Fatalf("Failed to execute SQL migrations: %v", err)
	}
}

// test list
func list(t *testing.T) {
	// Create a storage with real database client
	storage, _ := initStorage(t)

	// test varietals
	isActive := true
	perPage3 := uint64(3)
	page1 := uint64(1)
	searchText := "status"

	// Define the test cases
	tests := []struct {
		name   string
		filter *service.Filter
		count  int
	}{
		{
			name:   "GetModel list",
			filter: &service.Filter{},
			count:  10,
		}, {
			name:   "GetModel list active",
			filter: &service.Filter{Active: &isActive},
			count:  10,
		}, {
			name:   "GetModel with name like 'status'",
			filter: &service.Filter{Search: &searchText},
			count:  1,
		}, {
			name:   "GetModel with perPage 3",
			filter: &service.Filter{PerPage: &perPage3},
			count:  3,
		}, {
			name:   "GetModel with perPage 3 and page 1",
			filter: &service.Filter{PerPage: &perPage3, Page: &page1},
			count:  3,
		},
	}

	// Run the test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Call List method
			result, _, err := storage.List(initContext(), tc.filter)
			assert.NoError(t, err)

			// Assert that the result length is equal to the expected length
			assert.Len(t, *result, tc.count)
		})
	}
}

// test get
func get(t *testing.T) {
	// Create a storage with real database client
	storage, _ := initStorage(t)

	// Define the test cases
	testCases := []struct {
		name string
		id   string
		get  *service.Item
	}{
		{
			name: "GetModel by Id",
			id:   "a0eebc99-9c0b-4ef8-bb6d-6bb9bd381001",
			get: &service.Item{
				Id:               "a0eebc99-9c0b-4ef8-bb6d-6bb9bd381001",
				Name:             "Category",
				Url:              "category",
				ShortDescription: "",
				Description:      "",
				Required:         true,
				Active:           true,
				Prime:            true,
				ListItem:         true,
				Filter:           true,
				Type:             "select",
				Prefix:           "",
				Suffix:           "",
				CreatedBy:        "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
				UpdatedBy:        "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
			},
		}, {
			name: "GetModel by Id",
			id:   "a0eebc99-9c0b-4ef8-bb6d-6bb9bd381002",
			get: &service.Item{
				Id:               "a0eebc99-9c0b-4ef8-bb6d-6bb9bd381002",
				Name:             "Status",
				Url:              "status",
				ShortDescription: "",
				Description:      "",
				Required:         true,
				Active:           true,
				Prime:            true,
				ListItem:         true,
				Filter:           true,
				Type:             "select",
				Prefix:           "",
				Suffix:           "",
				CreatedBy:        "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
				UpdatedBy:        "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
			},
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call GetModel method
			result, err := storage.Get(initContext(), &tc.id, nil)
			assert.NoError(t, err)

			// Assert that the result is equal to the expected
			assert.Equal(t, tc.get.Id, result.Id)
			assert.Equal(t, tc.get.Name, result.Name)
			assert.Equal(t, tc.get.Url, result.Url)
			assert.Equal(t, tc.get.ShortDescription, result.ShortDescription)
			assert.Equal(t, tc.get.Description, result.Description)
			assert.Equal(t, tc.get.Required, result.Required)
			assert.Equal(t, tc.get.Active, result.Active)
			assert.Equal(t, tc.get.Prime, result.Prime)
			assert.Equal(t, tc.get.ListItem, result.ListItem)
			assert.Equal(t, tc.get.Filter, result.Filter)
			assert.Equal(t, tc.get.Type, result.Type)
			assert.Equal(t, tc.get.Prefix, result.Prefix)
			assert.Equal(t, tc.get.Suffix, result.Suffix)
		})
	}
}

// test get by url
func getByUrl(t *testing.T) {
	// Create a storage with real database client
	storage, _ := initStorage(t)

	// Define the test cases
	testCases := []struct {
		name string
		url  string
		get  *service.Item
	}{
		{
			name: "GetModel by URL",
			url:  "shipping-class",
			get: &service.Item{
				Id:               "a0eebc99-9c0b-4ef8-bb6d-6bb9bd381003",
				Name:             "Shipping Class",
				Url:              "shipping-class",
				ShortDescription: "",
				Description:      "",
				Required:         true,
				Active:           true,
				Prime:            true,
				ListItem:         true,
				Filter:           true,
				Type:             "select",
				Prefix:           "",
				Suffix:           "",
				CreatedBy:        "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
				UpdatedBy:        "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
			},
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call GetModel method
			result, err := storage.Get(initContext(), nil, &tc.url)
			assert.NoError(t, err)

			// Assert that the result is equal to the expected
			assert.Equal(t, tc.get.Id, result.Id)
			assert.Equal(t, tc.get.Name, result.Name)
			assert.Equal(t, tc.get.Url, result.Url)
			assert.Equal(t, tc.get.ShortDescription, result.ShortDescription)
			assert.Equal(t, tc.get.Description, result.Description)
			assert.Equal(t, tc.get.Required, result.Required)
			assert.Equal(t, tc.get.Active, result.Active)
			assert.Equal(t, tc.get.Prime, result.Prime)
			assert.Equal(t, tc.get.ListItem, result.ListItem)
			assert.Equal(t, tc.get.Filter, result.Filter)
			assert.Equal(t, tc.get.Type, result.Type)
			assert.Equal(t, tc.get.Prefix, result.Prefix)
			assert.Equal(t, tc.get.Suffix, result.Suffix)
		})
	}
}

// test create with id
func createWithId(t *testing.T) {
	// Create a storage with real database client
	storage, _ := initStorage(t)

	// Define the test cases
	testCases := []struct {
		name string
		item *service.Item
	}{
		{
			name: "Create with Id",
			item: newTestItems[0],
		}, {
			name: "Create with Id",
			item: newTestItems[1],
		}, {
			name: "Create with Id",
			item: newTestItems[2],
		}, {
			name: "Create with Id",
			item: newTestItems[3],
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call Create method
			id, err := storage.Create(initContext(), tc.item)
			assert.NoError(t, err)

			// get the product status by the Id
			result, err := storage.Get(initContext(), id, nil)
			assert.NoError(t, err)

			// Assert that the result is equal to the expected
			assert.Equal(t, tc.item.Id, result.Id)
			assert.Equal(t, tc.item.Name, result.Name)
			assert.Equal(t, tc.item.Url, result.Url)
			assert.Equal(t, tc.item.ShortDescription, result.ShortDescription)
			assert.Equal(t, tc.item.Description, result.Description)
			assert.Equal(t, tc.item.Required, result.Required)
			assert.Equal(t, tc.item.Active, result.Active)
			assert.Equal(t, tc.item.Prime, result.Prime)
			assert.Equal(t, tc.item.ListItem, result.ListItem)
			assert.Equal(t, tc.item.Filter, result.Filter)
		})
	}

}

// test create without id
func createWithoutId(t *testing.T) {
	// Create a storage with real database client
	storage, _ := initStorage(t)

	// Define the test cases
	testCases := []struct {
		name string
		item *service.Item
	}{
		{
			name: "Create without Id",
			item: newTestItems[4],
		}, {
			name: "Create without Id",
			item: newTestItems[5],
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call Create method
			id, err := storage.Create(initContext(), tc.item)
			assert.NoError(t, err)

			// get the product status by the Id
			result, err := storage.Get(initContext(), id, nil)
			assert.NoError(t, err)

			// update Id
			tc.item.Id = result.Id

			// Assert that the result is equal to the expected
			assert.Equal(t, tc.item.Id, result.ID)
			assert.Equal(t, tc.item.Name, result.Name)
			assert.Equal(t, tc.item.Url, result.Url)
			assert.Equal(t, tc.item.ShortDescription, result.ShortDescription)
			assert.Equal(t, tc.item.Description, result.Description)
			assert.Equal(t, tc.item.Required, result.Required)
			assert.Equal(t, tc.item.Active, result.Active)
			assert.Equal(t, tc.item.Prime, result.Prime)
			assert.Equal(t, tc.item.ListItem, result.ListItem)
			assert.Equal(t, tc.item.Filter, result.Filter)
		})
	}
}

// test update
func update(t *testing.T) {
	// Create a storage with real database client
	storage, _ := initStorage(t)

	// create new variable from the testItems[0]
	testItemNew := *newTestItems[0]
	testItemNew.Name = fmt.Sprintf("%s - updated", testItemNew.Name)
	testItemNew.Url = fmt.Sprintf("%s - updated", testItemNew.Url)

	// Define the test cases
	testCases := []struct {
		name string
		sent *service.Item
		get  *service.Item
	}{
		{
			name: "Update",
			sent: &testItemNew,
			get:  &testItemNew,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call Update method
			err := storage.Update(initContext(), tc.sent)
			assert.NoError(t, err)

			// get the product status by the Id
			result, err := storage.Get(initContext(), &tc.get.ID, nil)
			assert.NoError(t, err)

			// Assert that the result is equal to the expected
			assert.Equal(t, tc.get.ID, result.ID)
			assert.Equal(t, tc.get.Name, result.Name)
			assert.Equal(t, tc.get.Url, result.Url)
			assert.Equal(t, tc.get.ShortDescription, result.ShortDescription)
			assert.Equal(t, tc.get.Description, result.Description)
			assert.Equal(t, tc.get.Required, result.Required)
			assert.Equal(t, tc.get.Active, result.Active)
			assert.Equal(t, tc.get.Prime, result.Prime)
			assert.Equal(t, tc.get.ListItem, result.ListItem)
			assert.Equal(t, tc.get.Filter, result.Filter)
		})
	}
}

// test patch
func patch(t *testing.T) {
	// Create a storage with real database client
	storage, _ := initStorage(t)

	// create new variable from the testItems[0]
	testItemNew := *newTestItems[1]
	testItemNew.Name = fmt.Sprintf("%s - patched", testItemNew.Name)
	testItemNew.Url = fmt.Sprintf("%s - patched", testItemNew.Url)

	// Define the test cases
	testCases := []struct {
		name   string
		id     string
		fields map[string]interface{}
		get    *service.Item
	}{
		{
			name: "Patch",
			fields: map[string]interface{}{
				"Name": testItemNew.Name,
				"Url":  testItemNew.Url,
			},
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call Patch method
			err := storage.Patch(initContext(), &testItemNew.ID, &tc.fields)
			assert.NoError(t, err)

			// get item by Id
			result, err := storage.Get(initContext(), &testItemNew.ID, nil)
			assert.NoError(t, err)

			// Assert that the result is equal to the expected
			assert.Equal(t, testItemNew.ID, result.ID)
			assert.Equal(t, testItemNew.Name, result.Name)
			assert.Equal(t, testItemNew.Url, result.Url)
			assert.Equal(t, testItemNew.ShortDescription, result.ShortDescription)
			assert.Equal(t, testItemNew.Description, result.Description)
			assert.Equal(t, testItemNew.Required, result.Required)
			assert.Equal(t, testItemNew.Active, result.Active)
			assert.Equal(t, testItemNew.Prime, result.Prime)
			assert.Equal(t, testItemNew.ListItem, result.ListItem)
			assert.Equal(t, testItemNew.Filter, result.Filter)
		})
	}
}

// test updated at
func updatedAt(t *testing.T) {
	// Create a storage with real database client
	storage, _ := initStorage(t)

	// Define the test cases
	testCases := []struct {
		name string
		id   string
		sent *service.Item
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

// max sort order
func maxSortOrder(t *testing.T) {
	// Create a storage with real database client
	storage, _ := initStorage(t)

	// Call the MaxSortOrder method
	sortOrder, err := storage.MaxSortOrder(context.Background())
	assert.NoError(t, err)

	// Assert that the result is not empty
	assert.NotEmpty(t, sortOrder)
}

// test del
func del(t *testing.T) {
	// Create a storage with real database client
	storage, _ := initStorage(t)

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
		assert.NoError(t, err)

		// Call the GetModel method
		_, err = storage.Get(initContext(), &tc.id, nil)
		assert.Error(t, err)
	}
}

func initStorage(t *testing.T) (*service.Service, *config.Config) {
	// Create a real database client
	cfg := config.GetConfig(context.Background())
	app := appInit.NewAppInit(initContext(), cfg)
	err := app.SqlDBInit()
	assert.NoError(t, err)

	pgClient := app.SqlDB

	// Initialize Storage
	storage := model.NewStorage(pgClient)

	return service.NewService(storage), cfg
}

func initContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "user_id", "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1")
	return ctx
}

var newTestItems = []*service.Item{
	{
		ID:               "a0eebc99-9c0b-4ef8-bb6d-6bb9bd381004",
		Name:             "Test item 1",
		Url:              "test-item-1",
		ShortDescription: "",
		Description:      "",
		Required:         false,
		Active:           false,
		Prime:            false,
		ListItem:         false,
		Filter:           false,
	}, {
		ID:               "a0eebc99-9c0b-4ef8-bb6d-6bb9bd381005",
		Name:             "Test item 2",
		Url:              "test-item-2",
		ShortDescription: "",
		Description:      "",
		Required:         false,
		Active:           false,
		Prime:            false,
		ListItem:         false,
		Filter:           false,
	}, {
		ID:               "a0eebc99-9c0b-4ef8-bb6d-6bb9bd381006",
		Name:             "Test item 3",
		Url:              "test-item-3",
		ShortDescription: "",
		Description:      "",
		Required:         false,
		Active:           false,
		Prime:            false,
		ListItem:         false,
		Filter:           false,
	}, {
		ID:               "a0eebc99-9c0b-4ef8-bb6d-6bb9bd381007",
		Name:             "Test item 4",
		Url:              "test-item-4",
		ShortDescription: "",
		Description:      "",
		Required:         false,
		Active:           false,
		Prime:            false,
		ListItem:         false,
		Filter:           false,
	}, {
		Name:             "Test item 5",
		Url:              "test-item-5",
		ShortDescription: "",
		Description:      "",
		Required:         false,
		Active:           false,
		Prime:            false,
		ListItem:         false,
		Filter:           false,
	}, {
		Name:             "Test item 6",
		Url:              "test-item-6",
		ShortDescription: "",
		Description:      "",
		Required:         false,
		Active:           false,
		Prime:            false,
		ListItem:         false,
		Filter:           false,
	},
}
