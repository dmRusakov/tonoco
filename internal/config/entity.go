package config

type Config struct {
	Id      string `env:"ID" env-required:"true"`
	IsProd  bool   `env:"IS_PROD" env-required:"true"`
	IsDebug bool   `env:"IS_DEBUG" env-required:"true"`
	WebPort string `env:"WEB_PORT" env-required:"true"`

	CacheStorage    CacheStorage
	DataStorage     DataStorage
	ProductListener ServerProductListener
}
