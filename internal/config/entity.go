package config

type Config struct {
	Name        string `env:"ADMIN_NAME" env-default:"MonkeysMoonAdmin"`
	IsProd      bool   `env:"ADMIN_IS_PROD" env-default:"true"`
	AuthStorage AuthStorage
}

type AuthStorage struct {
	Host      string `env:"AUTH_STORAGE_HOST" env-default:"AuthStorage"`
	Port      string `env:"AUTH_STORAGE_PORT" env-default:"6379"`
	Password  string `env:"AUTH_STORAGE_PASSWORD" required:"true"`
	MaxMemory string `env:"AUTH_STORAGE_MAX_MEMORY" env-default:"64mb"`
}
