package config

import "github.com/dmRusakov/tonoco/pkg/redisdb"

type CacheStorage struct {
	Host      string `env:"PERMISSION_STORAGE_HOST" env-required:"true"`
	Port      string `env:"PERMISSION_STORAGE_PORT" env-required:"true"`
	Password  string `env:"PERMISSION_STORAGE_PASSWORD" env-required:"true"`
	MaxMemory string `env:"PERMISSION_STORAGE_MAXMEMORY" env-required:"true"`
}

func (cacheStorage *CacheStorage) ToRedisConfig() *redisdb.Config {
	return &redisdb.Config{
		Host:      cacheStorage.Host,
		Port:      cacheStorage.Port,
		Password:  cacheStorage.Password,
		MaxMemory: cacheStorage.MaxMemory,
	}
}
