package service_test

import (
	"github.com/dmRusakov/tonoco/internal/appInit"
	"github.com/dmRusakov/tonoco/internal/domain/product_status/service"
	"github.com/stretchr/testify/assert"

	"context"
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/domain/product_status/model"
	"testing"
)

func TestProductStatusModal(t *testing.T) {
	t.Run("testItemValidations", testItemValidations)
	t.Run("testValidateFields", testValidateFields)
}

func testItemValidations(t *testing.T) {
	t.Parallel()

	goodItem := &model.Item{
		Name:   "name",
		Url:    "url",
		Active: false,
	}

	itemWithoutName := &model.Item{
		Url:    "url",
		Active: false,
	}

	itemWithoutUrl := &model.Item{
		Name:   "name",
		Active: false,
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

}

func testValidateFields(t *testing.T) {
	t.Parallel()

	goodFields := map[string]interface{}{
		"Name": "name",
		"Url":  "url",
	}

	fieldsWithEmptyNameAndUrl := map[string]interface{}{
		"Name": "",
		"Url":  "jh",
	}

	srv := initStorage(t)

	// test fields
	validations := srv.Validate(nil, &goodFields)
	assert.Equal(t, 0, len(validations))

	// test fields with empty name
	validations2 := srv.Validate(nil, &fieldsWithEmptyNameAndUrl)
	assert.Equal(t, 1, len(validations2))
	assert.Equal(t, service.NoName, validations2[0])
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
