package shop_page

import (
	"github.com/dmRusakov/tonoco/internal/config"
	currency_service "github.com/dmRusakov/tonoco/internal/domain/currency/service"
	file_service "github.com/dmRusakov/tonoco/internal/domain/file/service"
	folder_service "github.com/dmRusakov/tonoco/internal/domain/folder/service"
	image_service "github.com/dmRusakov/tonoco/internal/domain/image/service"
	price_service "github.com/dmRusakov/tonoco/internal/domain/price/service"
	price_type_service "github.com/dmRusakov/tonoco/internal/domain/price_type/service"
	product_image_service "github.com/dmRusakov/tonoco/internal/domain/product_image/service"
	product_info_service "github.com/dmRusakov/tonoco/internal/domain/product_info/service"
	shop_page_service "github.com/dmRusakov/tonoco/internal/domain/shop/service"
	shop_tag_type_service "github.com/dmRusakov/tonoco/internal/domain/shop_tag_type/service"
	stock_quantity_service "github.com/dmRusakov/tonoco/internal/domain/stock_quantity/service"
	store_service "github.com/dmRusakov/tonoco/internal/domain/store/service"
	tag_service "github.com/dmRusakov/tonoco/internal/domain/tag/service"
	tag_select_service "github.com/dmRusakov/tonoco/internal/domain/tag_select/service"
	tag_type_service "github.com/dmRusakov/tonoco/internal/domain/tag_type/service"
	text_service "github.com/dmRusakov/tonoco/internal/domain/text/service"
	warehouse_service "github.com/dmRusakov/tonoco/internal/domain/warehouse/service"
	"github.com/dmRusakov/tonoco/internal/entity/pages"
	"github.com/google/uuid"
	"sync"
)

type UseCase struct {
	// config
	cfg *config.Config

	// services
	currency      *currency_service.Service
	file          *file_service.Service
	folder        *folder_service.Service
	price         *price_service.Service
	priceType     *price_type_service.Service
	productInfo   *product_info_service.Service
	tag           *tag_service.Service
	tagType       *tag_type_service.Service
	tagSelect     *tag_select_service.Service
	text          *text_service.Service
	stockQuantity *stock_quantity_service.Service
	store         *store_service.Service
	shop          *shop_page_service.Service
	shopTagType   *shop_tag_type_service.Service
	warehouse     *warehouse_service.Service
	image         *image_service.Service
	productImage  *product_image_service.Service

	// cache
	shopPageCache       map[string]*pages.Shop
	shopPageFilterCache map[uuid.UUID]*pages.ShopPageFilter
	gridItemCache       map[uuid.UUID]*pages.ProductGridItem
	itemIdsCache        map[uuid.UUID]struct {
		ids   *[]uuid.UUID
		count *uint64
	}

	// mutex
	mu sync.Mutex
}

func NewUseCase(
	cfg *config.Config,
	currencyService *currency_service.Service,
	fileService *file_service.Service,
	folderService *folder_service.Service,
	priceService *price_service.Service,
	priceTypeService *price_type_service.Service,
	productInfoService *product_info_service.Service,
	tagService *tag_service.Service,
	tagTypeService *tag_type_service.Service,
	tagSelect *tag_select_service.Service,
	text *text_service.Service,
	stockQuantity *stock_quantity_service.Service,
	storeService *store_service.Service,
	shopPageService *shop_page_service.Service,
	shopTagType *shop_tag_type_service.Service,
	warehouseService *warehouse_service.Service,
	imageService *image_service.Service,
	productImageService *product_image_service.Service,
) *UseCase {
	return &UseCase{
		// config
		cfg: cfg,

		// services
		currency:      currencyService,
		file:          fileService,
		folder:        folderService,
		price:         priceService,
		priceType:     priceTypeService,
		productInfo:   productInfoService,
		tag:           tagService,
		tagType:       tagTypeService,
		tagSelect:     tagSelect,
		text:          text,
		stockQuantity: stockQuantity,
		store:         storeService,
		shop:          shopPageService,
		shopTagType:   shopTagType,
		warehouse:     warehouseService,
		image:         imageService,
		productImage:  productImageService,

		// cache
		shopPageCache:       make(map[string]*pages.Shop),
		shopPageFilterCache: make(map[uuid.UUID]*pages.ShopPageFilter),
		gridItemCache:       make(map[uuid.UUID]*pages.ProductGridItem),
		itemIdsCache: make(map[uuid.UUID]struct {
			ids   *[]uuid.UUID
			count *uint64
		}),

		// mutex
		mu: sync.Mutex{},
	}
}
