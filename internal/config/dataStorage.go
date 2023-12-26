package config

import "github.com/dmRusakov/tonoco/pkg/postgresql"

type dataStorage struct {
	Host     string `env:"DATA_STORAGE_HOST" env-required:"true"`
	Port     string `env:"DATA_STORAGE_PORT" env-required:"true"`
	User     string `env:"DATA_STORAGE_USER" env-required:"true"`
	Password string `env:"DATA_STORAGE_PASSWORD" env-required:"true"`
	DB       string `env:"DATA_STORAGE_DB" env-required:"true"`
}

// ToPostgreSQLConfig - convert DataStorage to PostgreSQLConfig
func (dataStorage *dataStorage) ToPostgreSQLConfig() *postgresql.Config {
	return &postgresql.Config{
		Host:     dataStorage.Host,
		Port:     dataStorage.Port,
		User:     dataStorage.User,
		Password: dataStorage.Password,
		DB:       dataStorage.DB,
	}
}
