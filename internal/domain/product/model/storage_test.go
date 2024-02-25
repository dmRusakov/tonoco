package model_test

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/appInit"
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/domain/product/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// Product for test
var productNew1 = model.Product{
	ID:                    "f0eebc99-9c0b-4ef8-bb6d-6bb9bd382503",
	SKU:                   "TESTSKU",
	Name:                  "Test Product",
	ShortDescription:      "Test Short Description",
	Description:           "Test Description",
	StatusID:              "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
	Slug:                  "test-product",
	RegularPrice:          100,
	SalePrice:             50,
	FactoryPrice:          0,
	IsTaxable:             true,
	Quantity:              10,
	ReturnToStockDate:     time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
	IsTrackStock:          true,
	ShippingClassID:       "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
	ShippingWeight:        10,
	ShippingWidth:         10,
	ShippingHeight:        10,
	ShippingLength:        10,
	SeoTitle:              "Test Product",
	SeoDescription:        "Test Product",
	GTIN:                  "1234567890",
	GoogleProductCategory: "123",
	GoogleProductType:     "Test Product",
}

var productNew2 = model.Product{
	SKU:                   "TESTSKU2",
	Name:                  "Test Product 2",
	ShortDescription:      "Test Short Description 2",
	Description:           "Test Description 2",
	StatusID:              "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
	Slug:                  "test-product-2",
	RegularPrice:          200,
	SalePrice:             100,
	FactoryPrice:          0,
	IsTaxable:             true,
	Quantity:              20,
	ReturnToStockDate:     time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
	IsTrackStock:          true,
	ShippingClassID:       "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
	ShippingWeight:        20,
	ShippingWidth:         20,
	ShippingHeight:        20,
	ShippingLength:        20,
	SeoTitle:              "Test Product 2",
	SeoDescription:        "Test Product 2",
	GTIN:                  "1234567891",
	GoogleProductCategory: "124",
	GoogleProductType:     "Test Product 2",
}

var product2500 = model.Product{
	ID:                    "a0eebc99-9c0b-4ef8-bb6d-6bb9bd382500",
	SKU:                   "IS48LUXOR",
	Name:                  "48″ Luxor Island Range Hood",
	ShortDescription:      "Radiant, lustrous, and on the cutting edge of modernity, the Luxor Island Range Hood will make a striking statement in any kitchen.",
	Description:           "Radiant, lustrous, and on the cutting edge of modernity, the Futuro Futuro 48 inch Luxor Island Range Hood will make a striking statement in any kitchen. With its sleek, gleaming exterior and bright white perimeter light, this luxury range hood is reminiscent of an extraterrestrial spacecraft- and it has the advanced technology to boot. This unique range hood is composed of an illuminated glass perimeter that is lit entirely from within, and specialized stainless steel that prevents smudging or streaks and leaves only a glossy, finger-print free finish. This unique range hood possesses cutting-edge industrial technology that optimizes air filtration performance. This suspended range hood will keep your kitchen cool and contaminant free with its exclusive perimeter suction system, which redirects the airflow from the center of the unit to the mesh filters along the edges of the exhaust hood, providing a cleaner look, increasing suction, and cutting ambient noise. This professional style range hood further improves upon its superb performance by entrapping airborne grease and other culinary contaminants with its designer grease filters, which are concealed behind its illuminated glass panels to provide a clean, streamlined design. This unique range hood possesses four different airflow speeds, energy efficient fluorescent lights, and an extendable chimney that can be easily lengthened between 25 and 48 inches for perfect placement over any kitchen island range. This unique range hood makes regular maintenance easy with its automated reminders that indicate when it’s time to clean the grease filters in the dishwasher to maintain optimal performance. This Italian range hood also possesses a delayed shut-off function, which keeps the ventilation hood on even after you have finished cooking until the removal of odors, grease, and impurities from the air has been completed. Customers may purchase this kitchen range hood with the choice of a ducted or ductless exhaust vent, and all Futuro Futuro ventilation hoods come with a one year US standard and two year parts warranty. Show-stoppingly gorgeous and expertly crafted, the Italian made 36 inch Luxor island mount range hood is perfect for kitchen owners seeking a powerful kitchen accessory with exclusive style. sku: IS48LUXOR",
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
	SeoDescription:        "Futuro Futuro Luxor 48 Inch Island-mount Range Hood, Unique Illuminated Glass Body, Ultra-Quiet, with Blower",
	GTIN:                  "022099539230",
	GoogleProductCategory: "684",
	GoogleProductType:     "Kitchen Range Hood > Island Mount > Glass > 48-inch",
	CreatedAt:             time.Date(2024, 2, 14, 10, 56, 31, 167347, time.UTC),
	CreatedBy:             "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
	UpdatedAt:             time.Date(2024, 2, 14, 10, 56, 31, 167347, time.UTC),
	UpdatedBy:             "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
}

var product2501 = model.Product{
	ID:                    "a0eebc99-9c0b-4ef8-bb6d-6bb9bd382501",
	SKU:                   "WL36MELARA",
	Name:                  "36″ Melara Wall Range Hood",
	ShortDescription:      "Comfortable yet sophisticated, the 36″ Melara country style range hood is a beautifully crafted Italian kitchen accessory that will perfectly complement more traditionally styled kitchens",
	Description:           "Comfortable yet sophisticated, the Futuro Futuro 36 inch Melara Wall Range Hood is a beautifully crafted Italian range hood that will feel right at home in any traditional or country style kitchen. This steel range hood is composed of solid white oak and high quality steel that has been specially coated in a durable and unique enamel that can serve as either the finished exterior or as a base coat for another paint color. Kitchen owners can easily alter the exterior of this luxury range hood by using any kind of paint, including oil-based, water-based, or sprays, to coat the exhaust hood’s steel body to match or complement other colors in the kitchen- with no primer needed. The wooden accents on this traditional style range hood come unstained, so kitchen owners may perfectly match them to existing kitchen cabinetry. This modern range hood offers user-friendly features, including its three different airflow speeds, energy efficient LED lights, and an easy to use slider control. Robust and efficient, this powerful range hood uses cutting-edge technology to quickly remove steam, smoke, and odors from the air while maintaining whisper quiet noise levels, thanks to Futuro Futuro’s innovative blower and exclusive sound absorbing technology. This wood range hood also contains dishwasher safe, designer metal mesh filters that specifically target the capture of oil and grease, keeping it from collecting on your kitchen surfaces. Regular maintenance is straightforward with this country style range hood’s convenient features and functionality. This designer range hood has automated reminders that are programmed to routinely indicate when it’s time to clean the filters in the dishwasher to maintain optimal performance. Additionally, this unique range hood’s convenient delayed shut-off function keeps the exhaust hood on until all odors, grease, and impurities have been removed from the air, allowing users to focus their attention elsewhere after cooking. Customers may purchase this kitchen range hood with the choice of a ducted or ductless exhaust vent, dependent on their existing kitchen’s build. All Futuro Futuro ventilation hoods come with a one year US standard and two year parts warranty. Stately yet homey, the 36 inch Melara wall mount range hood is perfect for kitchen owners who prefer traditional styles but require a modern performance. sku: WL36MELARA",
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
	SeoDescription:        "Futuro Futuro Melara 36 Inch Wall-mount Range Hood, Classic Design, Painted Steel and Wood, Ultra-Quiet, with Blower",
	GTIN:                  "022099539612",
	GoogleProductCategory: "684",
	GoogleProductType:     "Kitchen Range Hood > Wall Mount > Stainless Steel > 36-inch",
	CreatedAt:             time.Date(2024, 2, 14, 11, 45, 11, 849586, time.UTC),
	CreatedBy:             "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
	UpdatedAt:             time.Date(2024, 2, 14, 11, 45, 11, 849586, time.UTC),
	UpdatedBy:             "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
}

var product2502 = model.Product{
	ID:                    "a0eebc99-9c0b-4ef8-bb6d-6bb9bd382502",
	SKU:                   "IS48POSITANO",
	Name:                  "48″ Positano Island Range Hood",
	ShortDescription:      "Futuro Futuro’s Positano Island Range Hoods are made with high-end stainless steel materials and premium craftsmanship with a simple design that will look great in any kitchen.",
	Description:           "Futuro Futuro’s Positano series of island range hoods have a simple design that will look great in any kitchen. These hoods with the highest grade of Stainless Steel, patented technology, and premium Italian craftsmanship. The Positano series puts an emphasis on form and function, ease of use, and simplicity. The Stainless Steel used in all Futuro Futuro products is of the highest quality for kitchen hoods, AISI 304. What makes AISI 304 Stainless Steel perfect for kitchen hoods is the high proportion of Chromium which gives it a high resistance to corrosion and rust. This helps the material stand up to the grime that is inevitable when regularly using a kitchen hood. The main piece of patented technology that you will notice immediately is the sound-absorbing chamber. By advancing technology in sound absorption for kitchen hoods Futuro Futuro is able to keep their products powerful yet still quiet for a residential kitchen. By using a powerful blower motor that would simply be too loud without their patented sound-absorbing chamber gives you the power of a commercial hood without all the noise. To make this simple hood even easier to use they added an illuminated electronic control panel to control every function of the hood. You can easily set a delayed shut off timer so the hood will continue to run for 15 minutes after shutting it off. This lets you turn the hood off as soon as your done and it will keep running to ensure the air in your kitchen is clean. These control panels even come equipped with a reminder for when it’s time to clean the designer filters that come standard! The Positano Island Range Hoods comes equipped with four LED lights that perfectly illuminate your cooking area. The four lights are positioned in the corners of the hood to create a central point of light over your stove top that melds out to your countertop. These LED lights are long lasting and energy efficient bulbs to keep your costs and maintenance as low as possible. The goal in designing the Positano series was to create a simple design that perfectly combined form and function. Through patented technology, premium Italian craftsmanship, and high quality materials Futuro Futuro created a hood that is built to last and will always look great. sku: IS48POSITANO",
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
	SeoDescription:        "Futuro Futuro Positano 48 Inch Island-mount Range Hood, Stainless Steel, Modern Style, LED, Ultra-Quiet, with Blower",
	GTIN:                  "022099539247",
	GoogleProductCategory: "684",
	GoogleProductType:     "Kitchen Range Hood > Island Mount > Stainless Steel > 48-inch",
	CreatedAt:             time.Date(2024, 2, 14, 11, 45, 11, 855960, time.UTC),
	CreatedBy:             "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
	UpdatedAt:             time.Date(2024, 2, 14, 11, 45, 11, 855960, time.UTC),
	UpdatedBy:             "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
}

var product5681 = model.Product{
	ID:                    "a0eebc99-9c0b-4ef8-bb6d-6bb9bd325681",
	SKU:                   "WL46SLIDE",
	Name:                  "46″ Slide Insert Range Hood",
	ShortDescription:      "Futuro Futuro’s 46-Inch Slide Insert Stainless Steel Range Hood is the perfect solution for fitting into a kitchen built around gorgeous custom cabinetry with a modern flair.",
	Description:           "Futuro Futuro’s 46-Inch Slide Insert Range Hood offers a sleek design perfect for fitting into a kitchen centered around custom cabinetry. This hood slides right into your custom cabinetry to bring a uniformed look across your kitchen. It also comes packed with high quality features such as a rotatable blower, the highest grade Stainless Steel, premium Italian craftsmanship, an electronic control panel, and heavy-duty chrome baffle filters. The highest grade of Stainless Steel, AISI 304, and the best Italian craftsmanship make for a hood that looks as perfectly as it works. AISI 304 Stainless Steel is ideal for kitchen hoods because it has a particularly high resistance to corrosion and rust that lets it stand up to the most brutal kitchen applications. This all makes for a hood that will look great when it’s delivered and continue to look great for years to come. By implementing a cutting edge sound absorbing chamber, Futuro Futuro keeps extremely powerful hoods quiet no matter how hard or long it’s running for. The Slide Insert Range Hood uses a 940 CFM rotatable blower with four adjustable speeds that are easily controlled with an electronic control panel. A powerful blower like this with the patented sound absorbing chamber keeps your kitchen fresh, clean, and quiet at all times. To make sure the Slide Insert series hoods are always easy to maintain, they use heavy-duty chrome baffle filters that are easy to remove and dishwasher safe. Baffle filtration is a time tested standard among high quality kitchen hoods that simply work. It also comes equipped with a built-in oil collector to ensure clean up is as easy as it can be. The futuristic electronic control panel that comes standard with these hoods have a filter-cleaning reminder to help keep your kitchen clean. It also has a 15-minute delayed shut off timer, easy access to four adjustable speeds, and a boost mode for when you’re deep frying or doing any messy cooking. To illuminate your cooking space, this hood comes with an LED light strip that is long lasting and energy efficient. An LED strip light will bring all the light you could need to your stovetop with a smooth, balanced illumination. Futuro Futuro has stayed at the forefront of kitchen ventilation by always using cutting edge technology, high quality materials, and fine Italian craftsmanship. sku: WL46SLIDE",
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
	SeoDescription:        "Futuro Futuro Slide 46 Inch Wall-mount / In-Cabinet Range Hood, Remote Control, Oil trap, LED, Ultra-Quiet, with Blower",
	GTIN:                  "703168132001",
	GoogleProductCategory: "684",
	GoogleProductType:     "Kitchen Range Hood > Wall Mount > Stainless Steel > 46-inch",
	CreatedBy:             "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
	UpdatedBy:             "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
}

var product6730 = model.Product{
	ID:                    "a0eebc99-9c0b-4ef8-bb6d-6bb9bd326730",
	SKU:                   "WL36POSITANO-BLK",
	Name:                  "36″ Positano Black Wall Range Hood",
	ShortDescription:      "Radiant, lustrous, and on the cutting edge of modernity, the Luxor Island Range Hood will make a striking statement in any kitchen.",
	Description:           "Radiant, lustrous, and on the cutting edge of modernity, the Futuro Futuro 48 inch Luxor Island Range Hood will make a striking statement in any kitchen. With its sleek, gleaming exterior and bright white perimeter light, this luxury range hood is reminiscent of an extraterrestrial spacecraft- and it has the advanced technology to boot. This unique range hood is composed of an illuminated glass perimeter that is lit entirely from within, and specialized stainless steel that prevents smudging or streaks and leaves only a glossy, finger-print free finish. This unique range hood possesses cutting-edge industrial technology that optimizes air filtration performance. This suspended range hood will keep your kitchen cool and contaminant free with its exclusive perimeter suction system, which redirects the airflow from the center of the unit to the mesh filters along the edges of the exhaust hood, providing a cleaner look, increasing suction, and cutting ambient noise. This professional style range hood further improves upon its superb performance by entrapping airborne grease and other culinary contaminants with its designer grease filters, which are concealed behind its illuminated glass panels to provide a clean, streamlined design. This unique range hood possesses four different airflow speeds, energy efficient fluorescent lights, and an extendable chimney that can be easily lengthened between 25 and 48 inches for perfect placement over any kitchen island range. This unique range hood makes regular maintenance easy with its automated reminders that indicate when it’s time to clean the grease filters in the dishwasher to maintain optimal performance. Additionally, this Italian range hood possesses a delayed shut-off function, which keeps the ventilation hood on even after you have finished cooking until the removal of odors, grease, and impurities from the air has been completed. Customers may purchase this kitchen range hood with the choice of a ducted or ductless exhaust vent, and all Futuro Futuro ventilation hoods come with a one year US standard and two year parts warranty. Show-stoppingly gorgeous and expertly crafted, the Italian made 36 inch Luxor island mount range hood is perfect for kitchen owners seeking a powerful kitchen accessory with exclusive style. sku: IS48LUXOR",
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
	SeoDescription:        "Futuro Futuro Positano Black 36 Inch Wall-mount Range Hood, Modern Slim Stainless Steel Design, LED, Ultra-Quiet, with Blower",
	GTIN:                  "703168132100",
	GoogleProductCategory: "684",
	GoogleProductType:     "Kitchen Range Hood > Wall Mount > Black > 36-inch",
	CreatedAt:             time.Date(2024, 2, 14, 11, 45, 13, 559147, time.UTC),
	CreatedBy:             "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
	UpdatedAt:             time.Date(2024, 2, 14, 11, 45, 13, 559147, time.UTC),
	UpdatedBy:             "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1",
}

// test get
func TestGet(t *testing.T) {
	// Create a real database client
	pgClient := initDB(t)

	// Initialize a new instance of ProductModel with a real database client
	storage := model.NewProductStorage(pgClient)

	// Define the test cases
	testCases := []struct {
		name     string
		id       string
		expected *model.Product
	}{
		{
			name:     "Test Case 1: Valid ID",
			id:       "a0eebc99-9c0b-4ef8-bb6d-6bb9bd382500",
			expected: &product2500,
		}, {
			name:     "Test Case 2: Valid ID",
			id:       "a0eebc99-9c0b-4ef8-bb6d-6bb9bd382501",
			expected: &product2501,
		}, {
			name:     "Test Case 3: Valid ID",
			id:       "a0eebc99-9c0b-4ef8-bb6d-6bb9bd382502",
			expected: &product2502,
		}, {
			name:     "Test Case 4: Valid ID",
			id:       "a0eebc99-9c0b-4ef8-bb6d-6bb9bd325681",
			expected: &product5681,
		}, {
			name:     "Test Case 5: Valid ID",
			id:       "a0eebc99-9c0b-4ef8-bb6d-6bb9bd326730",
			expected: &product6730,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the Get method
			result, err := storage.Get(context.Background(), tc.id)

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

// test update
func TestUpdate(t *testing.T) {
	// Create a real database client
	pgClient := initDB(t)

	// Initialize a new instance of ProductModel with a real database client
	storage := model.NewProductStorage(pgClient)

	// create new variable from product2500 and modify it
	testProduct1 := product2500
	testProduct1.Name = "48″ Luxor Island Range Hood Updated"
	testProduct1.FactoryPrice = 2000
	testProduct1.IsTaxable = false

	// Define the test cases
	testCases := []struct {
		name     string
		id       string
		sent     *model.Product
		expected *model.Product
	}{
		{
			name:     "Test Case 1: Valid ID",
			id:       "a0eebc99-9c0b-4ef8-bb6d-6bb9bd382500",
			sent:     &testProduct1,
			expected: &testProduct1,
		}, {
			name:     "Test Case 2: Valid ID",
			id:       "a0eebc99-9c0b-4ef8-bb6d-6bb9bd382500",
			sent:     &product2500,
			expected: &product2500,
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
			assert.NotEqual(t, tc.expected.UpdatedAt, result.UpdatedAt)
		})
	}
}

// test patch
func TestPatch(t *testing.T) {
	// Create a real database client
	pgClient := initDB(t)

	// Initialize a new instance of ProductModel with a real database client
	storage := model.NewProductStorage(pgClient)

	// Prepare the test Product
	testProduct1 := product2500
	testProduct1.Name = "48″ Luxor Island Range Hood Updated"
	testProduct1.FactoryPrice = float32(2000.32)
	testProduct1.IsTaxable = false

	// Define the test cases
	testCases := []struct {
		name     string
		id       string
		fields   map[string]interface{}
		expected *model.Product
	}{
		{
			name: "Test Case 1: Valid ID",
			id:   "a0eebc99-9c0b-4ef8-bb6d-6bb9bd382500",
			fields: map[string]interface{}{
				"Name":         testProduct1.Name,
				"FactoryPrice": testProduct1.FactoryPrice,
				"IsTaxable":    testProduct1.IsTaxable,
			},
			expected: &testProduct1,
		}, {
			name: "Test Case 2: Valid ID",
			id:   "a0eebc99-9c0b-4ef8-bb6d-6bb9bd382500",
			fields: map[string]interface{}{
				"Name":         "48″ Luxor Island Range Hood",
				"FactoryPrice": float32(0),
				"IsTaxable":    true,
			},
			expected: &product2500,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the Patch method
			result, err := storage.Patch(context.Background(), tc.id, tc.fields, "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1")

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
			assert.NotEqual(t, tc.expected.UpdatedAt, result.UpdatedAt)
		})
	}
}

// test create with id
func TestCreateWithID(t *testing.T) {
	// Create a real database client
	pgClient := initDB(t)

	// Initialize a new instance of ProductModel with a real database client
	storage := model.NewProductStorage(pgClient)

	// Define the test cases
	testCases := []struct {
		name     string
		sent     *model.Product
		expected *model.Product
	}{
		{
			name:     "Test Case 1: Valid Product",
			sent:     &productNew1,
			expected: &productNew1,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the Create method
			result, err := storage.Create(context.Background(), tc.sent, "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1")

			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the result matches the expected value
			assert.Equal(t, tc.expected.ID, result.ID)
			assert.Equal(t, tc.expected.SKU, result.SKU)
			assert.Equal(t, tc.expected.Name, result.Name)
			assert.Equal(t, tc.expected.ShortDescription, result.ShortDescription)
			assert.Equal(t, tc.expected.Description, result.Description)
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

// test create without id
func TestCreateWithoutID(t *testing.T) {
	// Create a real database client
	pgClient := initDB(t)

	// Initialize a new instance of ProductModel with a real database client
	storage := model.NewProductStorage(pgClient)

	// Define the test cases
	testCases := []struct {
		name     string
		sent     *model.Product
		expected *model.Product
	}{
		{
			name:     "Test Case 1: Valid Product",
			sent:     &productNew2,
			expected: &productNew2,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the Create method
			result, err := storage.Create(context.Background(), tc.sent, "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1")

			// update ID
			productNew2.ID = result.ID

			// Assert that there was no error
			assert.NoError(t, err)

			// Assert that the result matches the expected value
			assert.Equal(t, tc.expected.SKU, result.SKU)
			assert.Equal(t, tc.expected.Name, result.Name)
			assert.Equal(t, tc.expected.ShortDescription, result.ShortDescription)
			assert.Equal(t, tc.expected.Description, result.Description)
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

// test delete
func TestDelete(t *testing.T) {
	// Create a real database client
	pgClient := initDB(t)

	// Initialize a new instance of ProductModel with a real database client
	storage := model.NewProductStorage(pgClient)

	// Define the test cases
	testCases := []struct {
		name    string
		id      string
		success bool
	}{
		{
			name:    "Test Case 1: Valid ID",
			id:      productNew1.ID,
			success: true,
		},
		{
			name:    "Test Case 2: Invalid ID",
			id:      productNew2.ID,
			success: true,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the Delete method
			err := storage.Delete(context.Background(), tc.id)

			// Assert that there was no error
			if tc.success {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
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
