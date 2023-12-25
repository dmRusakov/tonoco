package config

type Config struct {
	Name         string `env:"ADMIN_NAME" env-default:"TonocoAdmin"`
	IsProd       bool   `env:"ADMIN_IS_PROD" env-default:"true"`
	CacheStorage CacheStorage
}
