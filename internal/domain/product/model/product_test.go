package model_test

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/appInit"
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/domain/product/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	// Create a real database client
	cfg := config.GetConfig(context.Background())
	app := appInit.NewAppInit(context.Background(), cfg)
	err := app.SqlDBInit()
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	pgClient := app.SqlDB

	// Initialize a new instance of ProductDAO with a real database client
	dao := model.NewProductStorage(pgClient)

	// Define the test cases
	testCases := []struct {
		name     string
		id       string
		expected *model.ProductStorage
	}{
		{
			name: "Test Case 1: Valid ID",
			id:   "a0eebc99-9c0b-4ef8-bb6d-6bb9bd382500",
			expected: &model.ProductStorage{
				ID:                    "a0eebc99-9c0b-4ef8-bb6d-6bb9bd382500",
				SKU:                   "IS48LUXOR",
				Name:                  "48″ Luxor Island Range Hood",
				SortOrder:             1,
				StatusID:              "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
				Slug:                  "48-luxor-island-range-hood",
				RegularPrice:          4000,
				SalePrice:             2695,
				FactoryPrice:          0,
				IsTaxable:             true,
				Quantity:              5,
				ReturnToStockDate:     time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
				IsTrackStock:          true,
				ShippingClassID:       "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
				ShippingWeight:        140,
				ShippingWidth:         40,
				ShippingHeight:        26,
				ShippingLength:        51,
				SeoTitle:              "Range Hood 48-inch Luxor Island by Futuro Futuro",
				GTIN:                  "022099539230",
				GoogleProductCategory: "684",
				GoogleProductType:     "Kitchen Range Hood > Island Mount > Glass > 48-inch",
			},
		}, {
			name: "Test Case 2: Valid ID",
			id:   "a0eebc99-9c0b-4ef8-bb6d-6bb9bd382501",
			expected: &model.ProductStorage{
				ID:                    "a0eebc99-9c0b-4ef8-bb6d-6bb9bd382501",
				SKU:                   "WL36MELARA",
				Name:                  "36″ Melara Wall Range Hood",
				SortOrder:             2,
				StatusID:              "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
				Slug:                  "36-melara-wall-range-hood",
				RegularPrice:          2050,
				SalePrice:             1395,
				FactoryPrice:          0,
				IsTaxable:             true,
				Quantity:              8,
				ReturnToStockDate:     time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
				IsTrackStock:          true,
				ShippingClassID:       "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
				ShippingWeight:        67,
				ShippingWidth:         40,
				ShippingHeight:        31,
				ShippingLength:        48,
				SeoTitle:              "Range Hood 36-inch Melara Wall-Mount by Futuro Futuro",
				GTIN:                  "022099539612",
				GoogleProductCategory: "684",
				GoogleProductType:     "Kitchen Range Hood > Wall Mount > Stainless Steel > 36-inch",
			},
		}, {
			name: "Test Case 3: Valid ID",
			id:   "a0eebc99-9c0b-4ef8-bb6d-6bb9bd382502",
			expected: &model.ProductStorage{
				ID:                    "a0eebc99-9c0b-4ef8-bb6d-6bb9bd382502",
				SKU:                   "IS48POSITANO",
				Name:                  "48″ Positano Island Range Hood",
				SortOrder:             3,
				StatusID:              "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
				Slug:                  "48-positano-island-range-hood",
				RegularPrice:          3400,
				SalePrice:             2295,
				FactoryPrice:          0,
				IsTaxable:             true,
				Quantity:              35,
				ReturnToStockDate:     time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
				IsTrackStock:          true,
				ShippingClassID:       "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
				ShippingWeight:        110,
				ShippingWidth:         40,
				ShippingHeight:        25,
				ShippingLength:        51,
				SeoTitle:              "Range Hood 48-inch Positano Island by Futuro Futuro",
				GTIN:                  "022099539247",
				GoogleProductCategory: "684",
				GoogleProductType:     "Kitchen Range Hood > Island Mount > Stainless Steel > 48-inch",
			},
		}, {
			name: "Test Case 4: Valid ID",
			id:   "a0eebc99-9c0b-4ef8-bb6d-6bb9bd325681",
			expected: &model.ProductStorage{
				ID:                    "a0eebc99-9c0b-4ef8-bb6d-6bb9bd325681",
				SKU:                   "WL46SLIDE",
				Name:                  "46″ Slide Insert Range Hood",
				SortOrder:             273,
				StatusID:              "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
				Slug:                  "46-inch-slide-insert-range-hood",
				RegularPrice:          2500,
				SalePrice:             1695,
				FactoryPrice:          0,
				IsTaxable:             true,
				Quantity:              24,
				ReturnToStockDate:     time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
				IsTrackStock:          true,
				ShippingClassID:       "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12",
				ShippingWeight:        54,
				ShippingWidth:         17,
				ShippingHeight:        16,
				ShippingLength:        49,
				SeoTitle:              "Range Hood 46-inch Slide Insert Built-in by Futuro Futuro",
				GTIN:                  "703168132001",
				GoogleProductCategory: "684",
				GoogleProductType:     "Kitchen Range Hood > Wall Mount > Stainless Steel > 46-inch",
			},
		}, {
			name: "Test Case 5: Valid ID",
			id:   "a0eebc99-9c0b-4ef8-bb6d-6bb9bd326730",
			expected: &model.ProductStorage{
				ID:                    "a0eebc99-9c0b-4ef8-bb6d-6bb9bd326730",
				SKU:                   "WL36POSITANO-BLK",
				Name:                  "36″ Positano Black Wall Range Hood",
				SortOrder:             274,
				StatusID:              "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
				Slug:                  "36-positano-black-wall-range-hood",
				RegularPrice:          2650,
				SalePrice:             1795,
				FactoryPrice:          0,
				IsTaxable:             true,
				Quantity:              10,
				ReturnToStockDate:     time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
				IsTrackStock:          true,
				ShippingClassID:       "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12",
				ShippingWeight:        65,
				ShippingWidth:         23,
				ShippingHeight:        20,
				ShippingLength:        39,
				SeoTitle:              "Range Hood 36-inch Positano Black Wall-Mount by Futuro Futuro",
				GTIN:                  "703168132100",
				GoogleProductCategory: "684",
				GoogleProductType:     "Kitchen Range Hood > Wall Mount > Black > 36-inch",
			},
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the Get method
			result, err := dao.Get(context.Background(), tc.id)

			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the result matches the expected value
			assert.Equal(t, tc.expected.ID, result.ID)
			assert.Equal(t, tc.expected.SKU, result.SKU)
			assert.Equal(t, tc.expected.Name, result.Name)
			assert.Equal(t, tc.expected.SortOrder, result.SortOrder)
			assert.Equal(t, tc.expected.StatusID, result.StatusID)
			assert.Equal(t, tc.expected.Slug, result.Slug)
			assert.Equal(t, tc.expected.RegularPrice, result.RegularPrice)
			assert.Equal(t, tc.expected.SalePrice, result.SalePrice)
			assert.Equal(t, tc.expected.FactoryPrice, result.FactoryPrice)
			assert.Equal(t, tc.expected.IsTaxable, result.IsTaxable)
			assert.Equal(t, tc.expected.Quantity, result.Quantity)
			assert.Equal(t, tc.expected.ReturnToStockDate, result.ReturnToStockDate)
			assert.Equal(t, tc.expected.IsTrackStock, result.IsTrackStock)
			assert.Equal(t, tc.expected.ShippingClassID, result.ShippingClassID)
			assert.Equal(t, tc.expected.ShippingWeight, result.ShippingWeight)
			assert.Equal(t, tc.expected.ShippingWidth, result.ShippingWidth)
			assert.Equal(t, tc.expected.ShippingHeight, result.ShippingHeight)
			assert.Equal(t, tc.expected.ShippingLength, result.ShippingLength)
			assert.Equal(t, tc.expected.SeoTitle, result.SeoTitle)
			assert.Equal(t, tc.expected.GTIN, result.GTIN)
			assert.Equal(t, tc.expected.GoogleProductCategory, result.GoogleProductCategory)
			assert.Equal(t, tc.expected.GoogleProductType, result.GoogleProductType)
		})
	}
}
