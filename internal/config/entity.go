package config

type Config struct {
	Id      string `env:"Id" env-required:"true"`
	IsProd  bool   `env:"IS_PROD" env-required:"true"`
	IsDebug bool   `env:"IS_DEBUG" env-required:"true"`
	WebPort string `env:"WEB_PORT" env-required:"true"`

	StoreUrl string `env:"STORE_URL" env-required:"true"`

	CacheStorage    CacheStorage
	DataStorage     DataStorage
	ProductListener ServerProductListener
}
