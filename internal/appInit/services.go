package appInit

import (
	"github.com/dmRusakov/tonoco/internal/config"
	currency_model "github.com/dmRusakov/tonoco/internal/domain/currency/model"
	currency_service "github.com/dmRusakov/tonoco/internal/domain/currency/service"
	file_model "github.com/dmRusakov/tonoco/internal/domain/file/model"
	file_service "github.com/dmRusakov/tonoco/internal/domain/file/service"
	folder_model "github.com/dmRusakov/tonoco/internal/domain/folder/model"
	folder_service "github.com/dmRusakov/tonoco/internal/domain/folder/service"
	image_model "github.com/dmRusakov/tonoco/internal/domain/image/model"
	image_service "github.com/dmRusakov/tonoco/internal/domain/image/service"
	price_model "github.com/dmRusakov/tonoco/internal/domain/price/model"
	price_service "github.com/dmRusakov/tonoco/internal/domain/price/service"
	price_type_model "github.com/dmRusakov/tonoco/internal/domain/price_type/model"
	price_type_service "github.com/dmRusakov/tonoco/internal/domain/price_type/service"
	product_image_model "github.com/dmRusakov/tonoco/internal/domain/product_image/model"
	product_image_service "github.com/dmRusakov/tonoco/internal/domain/product_image/service"
	product_info_model "github.com/dmRusakov/tonoco/internal/domain/product_info/model"
	product_info_service "github.com/dmRusakov/tonoco/internal/domain/product_info/service"
	stock_quantity_model "github.com/dmRusakov/tonoco/internal/domain/stock_quantity/model"
	stock_quantity_service "github.com/dmRusakov/tonoco/internal/domain/stock_quantity/service"
	store_model "github.com/dmRusakov/tonoco/internal/domain/store/model"
	store_service "github.com/dmRusakov/tonoco/internal/domain/store/service"
	tag_model "github.com/dmRusakov/tonoco/internal/domain/tag/model"
	tag_service "github.com/dmRusakov/tonoco/internal/domain/tag/service"
	tag_select_model "github.com/dmRusakov/tonoco/internal/domain/tag_select/model"
	tag_select_service "github.com/dmRusakov/tonoco/internal/domain/tag_select/service"
	tag_type_model "github.com/dmRusakov/tonoco/internal/domain/tag_type/model"
	tag_type_service "github.com/dmRusakov/tonoco/internal/domain/tag_type/service"
	warehouse_model "github.com/dmRusakov/tonoco/internal/domain/warehouse/model"
	warehouse_service "github.com/dmRusakov/tonoco/internal/domain/warehouse/service"
	"github.com/dmRusakov/tonoco/internal/entity"
	"sync"
)

type Services struct {
	Currency            *currency_service.Service
	File                *file_service.Service
	Folder              *folder_service.Service
	ImageService        *image_service.Service
	Price               *price_service.Service
	PriceType           *price_type_service.Service
	ProductImageService *product_image_service.Service
	ProductInfo         *product_info_service.Service
	StockQuantity       *stock_quantity_service.Service
	Store               *store_service.Service
	Tag                 *tag_service.Service
	TagSelect           *tag_select_service.Service
	TagType             *tag_type_service.Service
	Warehouse           *warehouse_service.Service
}

func (a *App) ServicesInit(cfg *config.Config) error {
	services := &Services{}
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	// file
	wg.Add(1)
	go func() {
		defer wg.Done()
		fileStorage := file_model.NewStorage(a.SqlDB)
		mu.Lock()
		services.File = file_service.NewService(fileStorage)
		mu.Unlock()
	}()

	// folder
	wg.Add(1)
	go func() {
		defer wg.Done()
		folderStorage := folder_model.NewStorage(a.SqlDB)
		mu.Lock()
		services.Folder = folder_service.NewService(folderStorage)
		mu.Unlock()
	}()

	// image
	wg.Add(1)
	go func() {
		defer wg.Done()
		imageStorage := image_model.NewStorage(a.SqlDB)
		mu.Lock()
		services.ImageService = image_service.NewService(imageStorage)
		mu.Unlock()
	}()

	// price
	wg.Add(1)
	go func() {
		defer wg.Done()
		priceStorage := price_model.NewStorage(a.SqlDB)
		mu.Lock()
		services.Price = price_service.NewService(priceStorage)
		mu.Unlock()
	}()

	// price type
	wg.Add(1)
	go func() {
		defer wg.Done()
		priceTypeStorage := price_type_model.NewStorage(a.SqlDB)
		mu.Lock()
		services.PriceType = price_type_service.NewService(priceTypeStorage)
		mu.Unlock()
	}()

	// product image
	wg.Add(1)
	go func() {
		defer wg.Done()
		productImageStorage := product_image_model.NewStorage(a.SqlDB)
		mu.Lock()
		services.ProductImageService = product_image_service.NewService(productImageStorage)
		mu.Unlock()
	}()

	// product info
	wg.Add(1)
	go func() {
		defer wg.Done()
		productInfoStorage := product_info_model.NewStorage(a.SqlDB)
		mu.Lock()
		services.ProductInfo = product_info_service.NewService(productInfoStorage)
		mu.Unlock()
	}()

	// stock quantity
	wg.Add(1)
	go func() {
		defer wg.Done()
		stockQuantityStorage := stock_quantity_model.NewStorage(a.SqlDB)
		mu.Lock()
		services.StockQuantity = stock_quantity_service.NewService(stockQuantityStorage)
		mu.Unlock()
	}()

	// tag
	wg.Add(1)
	go func() {
		defer wg.Done()
		tagStorage := tag_model.NewStorage(a.SqlDB)
		mu.Lock()
		services.Tag = tag_service.NewService(tagStorage)
		mu.Unlock()
	}()

	// tag select
	wg.Add(1)
	go func() {
		defer wg.Done()
		tagSelectStorage := tag_select_model.NewStorage(a.SqlDB)
		mu.Lock()
		services.TagSelect = tag_select_service.NewService(tagSelectStorage)
		mu.Unlock()
	}()

	// tag type
	wg.Add(1)
	go func() {
		defer wg.Done()
		tagTypeStorage := tag_type_model.NewStorage(a.SqlDB)
		mu.Lock()
		services.TagType = tag_type_service.NewService(tagTypeStorage)
		mu.Unlock()
	}()

	// warehouse
	wg.Add(1)
	go func() {
		defer wg.Done()
		warehouseStorage := warehouse_model.NewStorage(a.SqlDB)
		mu.Lock()
		services.Warehouse = warehouse_service.NewService(warehouseStorage)
		mu.Unlock()
	}()

	// store
	wg.Add(1)
	go func() {
		defer wg.Done()
		storeStorage := store_model.NewStorage(a.SqlDB)
		var err error
		mu.Lock()
		services.Store, err = store_service.NewService(storeStorage, cfg)
		mu.Unlock()

		if err != nil {
			return
		}

		// currency
		wg.Add(1)
		go func(defaultStore *entity.Store) {
			defer wg.Done()
			currencyStorage := currency_model.NewStorage(a.SqlDB)
			mu.Lock()
			services.Currency = currency_service.NewService(currencyStorage, defaultStore)
			mu.Unlock()
		}(services.Store.DefaultStore)
	}()

	wg.Wait()

	a.Services = services

	return nil
}
