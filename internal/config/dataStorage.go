package config

import "github.com/dmRusakov/tonoco/pkg/postgresql"

type DataStorage struct {
	Host     string `env:"DATA_STORAGE_HOST" env-required:"true"`
	Port     string `env:"DATA_STORAGE_PORT" env-required:"true"`
	User     string `env:"DATA_STORAGE_USER" env-required:"true"`
	Password string `env:"DATA_STORAGE_PASSWORD" env-required:"true"`
	DB       string `env:"DATA_STORAGE_DB" env-required:"true"`
}

// ToPostgreSQLConfig - convert DataStorage to PostgreSQLConfig
func (config *DataStorage) ToPostgreSQLConfig() *postgresql.Config {
	return &postgresql.Config{
		Host:     config.Host,
		Port:     config.Port,
		User:     config.User,
		Password: config.Password,
		DB:       config.DB,
	}
}
