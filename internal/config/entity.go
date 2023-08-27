package config

type Config struct {
	IsProd bool `env:"IS_PROD"  env-default:"true"`
}
