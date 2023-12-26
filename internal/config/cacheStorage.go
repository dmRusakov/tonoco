package config

import "github.com/dmRusakov/tonoco/pkg/redisdb"

type cacheStorage struct {
	Host      string `env:"CACHE_STORAGE_HOST" env-required:"true"`
	Port      string `env:"CACHE_STORAGE_PORT" env-required:"true"`
	Password  string `env:"CACHE_STORAGE_PASSWORD" env-required:"true"`
	MaxMemory string `env:"CACHE_STORAGE_MAXMEMORY" env-required:"true"`
}

// ToRedisConfig - convert cacheStorage to redisdb.Config
func (cacheStorage *cacheStorage) ToRedisConfig() *redisdb.Config {
	return &redisdb.Config{
		Host:      cacheStorage.Host,
		Port:      cacheStorage.Port,
		Password:  cacheStorage.Password,
		MaxMemory: cacheStorage.MaxMemory,
	}
}
