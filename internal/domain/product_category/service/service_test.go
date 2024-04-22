package service_test

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/appInit"
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/domain/product_category/service"
	"github.com/stretchr/testify/assert"

	"github.com/dmRusakov/tonoco/internal/domain/product_category/model"
	"testing"
)

func TestProductStatusModal(t *testing.T) {
	t.Run("testItemValidations", testItemValidations)
	t.Run("testValidateFields", testValidateFields)
}

func testItemValidations(t *testing.T) {
	t.Parallel()

	goodItem := &model.Item{
		Name:             "name",
		Url:              "url",
		ShortDescription: "short description",
		Description:      "description",
	}

	itemWithoutName := &model.Item{
		Url:              "url",
		ShortDescription: "short description",
		Description:      "description",
	}

	itemWithoutUrl := &model.Item{
		Name:             "name",
		ShortDescription: "short description",
		Description:      "description",
	}

	itemWithoutShortDescription := &model.Item{
		Name:        "name",
		Url:         "url",
		Description: "description",
	}

	itemWithoutDescription := &model.Item{
		Name:             "name",
		Url:              "url",
		ShortDescription: "short description",
	}

	srv := initStorage(t)

	// test good item
	validations := srv.Validate(goodItem, nil)
	assert.Equal(t, 0, len(validations))

	// test item without name
	validations = srv.Validate(itemWithoutName, nil)
	assert.Equal(t, 1, len(validations))
	assert.Equal(t, service.NoName, validations[0])

	// test item without url
	validations = srv.Validate(itemWithoutUrl, nil)
	assert.Equal(t, 1, len(validations))
	assert.Equal(t, service.NoUrl, validations[0])

	// test item without short description
	validations = srv.Validate(itemWithoutShortDescription, nil)
	assert.Equal(t, 1, len(validations))
	assert.Equal(t, service.NoShortDescription, validations[0])

	// test item without description
	validations = srv.Validate(itemWithoutDescription, nil)
	assert.Equal(t, 1, len(validations))
	assert.Equal(t, service.NoDescription, validations[0])
}

func testValidateFields(t *testing.T) {
	t.Parallel()

	goodFields := map[string]interface{}{
		"Name":             "name",
		"Url":              "url",
		"ShortDescription": "short description",
		"Description":      "description",
	}

	fieldsWithoutName := map[string]interface{}{
		"Name":             "",
		"Url":              "url",
		"ShortDescription": "short description",
		"Description":      "description",
	}

	fieldsWithoutUrl := map[string]interface{}{
		"Name":             "name",
		"Url":              "",
		"ShortDescription": "short description",
		"Description":      "description",
	}

	fieldsWithoutShortDescription := map[string]interface{}{
		"Name":             "name",
		"Url":              "url",
		"ShortDescription": "",
		"Description":      "description",
	}

	fieldsWithoutDescription := map[string]interface{}{
		"Name":             "name",
		"Url":              "url",
		"ShortDescription": "short description",
		"Description":      "",
	}

	srv := initStorage(t)

	// test good fields
	validations := srv.Validate(nil, &goodFields)
	assert.Equal(t, 0, len(validations))

	// test fields without name
	validations = srv.Validate(nil, &fieldsWithoutName)
	assert.Equal(t, 1, len(validations))
	assert.Equal(t, service.NoName, validations[0])

	// test fields without url
	validations = srv.Validate(nil, &fieldsWithoutUrl)
	assert.Equal(t, 1, len(validations))
	assert.Equal(t, service.NoUrl, validations[0])

	// test fields without short description
	validations = srv.Validate(nil, &fieldsWithoutShortDescription)
	assert.Equal(t, 1, len(validations))
	assert.Equal(t, service.NoShortDescription, validations[0])

	// test fields without description
	validations = srv.Validate(nil, &fieldsWithoutDescription)
	assert.Equal(t, 1, len(validations))
	assert.Equal(t, service.NoDescription, validations[0])
}

func initStorage(t *testing.T) *service.Service {
	// Create a real database client
	cfg := config.GetConfig(context.Background())
	app := appInit.NewAppInit(initContext(), cfg)
	err := app.SqlDBInit()
	assert.NoError(t, err)

	pgClient := app.SqlDB

	// Initialize Storage
	storage := model.NewStorage(pgClient)

	return service.NewService(storage)
}

func initContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "user_id", "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1")
	return ctx
}
