package config

type Config struct {
	AppId      string `env:"APP_ID" env-required:"true"`
	AppName    string `env:"APP_NAME" env-required:"true"`
	AppVersion string `env:"APP_VERSION" env-required:"true"`

	AppIsProd          bool   `env:"APP_IS_PROD" env-required:"true"`
	AppIsDebug         bool   `env:"APP_IS_DEBUG" env-required:"true"`
	AppWebPort         string `env:"APP_WEB_PORT" env-required:"true"`
	AppDefaultPage     uint64 `env:"APP_DEFAULT_PAGE" env-required:"true"`
	AppDefaultPerPAge  uint64 `env:"APP_DEFAULT_PER_PAGE" env-required:"true"`
	AppDefaultLanguage string `env:"APP_LANGUISH" env-required:"true"`

	ShopPageUrl string `env:"APP_SHOP_PAGE_URL" env-required:"true"`

	StoreUrl        string   `env:"STORE_STORE" env-required:"true"`
	StoreWarehouses []string `env:"STORE_WAREHOUSES" env-required:"true"`
	StoreCurrency   string   `env:"STORE_CURRENCY" env-required:"true"`

	Image ImageConfig

	CacheStorage    CacheStorage
	DataStorage     DataStorage
	ProductListener ServerProductListener
}

type ImageConfig struct {
	CompressionQuality int  `env:"APP_IMAGE_COMPRESSION_QUALITY" env-required:"true"`
	FullWidth          uint `env:"APP_IMAGE_FULL_WIDTH" env-required:"true"`
	FullHeight         uint `env:"APP_IMAGE_FULL_HEIGHT" env-required:"true"`
	MediumWidth        uint `env:"APP_IMAGE_MEDIUM_WIDTH" env-required:"true"`
	MediumHeight       uint `env:"APP_IMAGE_MEDIUM_HEIGHT" env-required:"true"`
	GridWidth          uint `env:"APP_IMAGE_GRID_WIDTH" env-required:"true"`
	GridHeight         uint `env:"APP_IMAGE_GRID_HEIGHT" env-required:"true"`
	ThumbWidth         uint `env:"APP_IMAGE_THUMB_WIDTH" env-required:"true"`
	ThumbHeight        uint `env:"APP_IMAGE_THUMB_HEIGHT" env-required:"true"`
}
