package config

import "github.com/dmRusakov/tonoco/pkg/redisdb"

type CacheStorage struct {
	Host      string `env:"CACHE_STORAGE_HOST" env-required:"true"`
	Port      string `env:"CACHE_STORAGE_PORT" env-required:"true"`
	Password  string `env:"CACHE_STORAGE_PASSWORD" env-required:"true"`
	MaxMemory string `env:"CACHE_STORAGE_MAXMEMORY" env-required:"true"`
}

// ToRedisConfig - convert CacheStorage to redisdb.Config
func (config *CacheStorage) ToRedisConfig() *redisdb.Config {
	return &redisdb.Config{
		Host:      config.Host,
		Port:      config.Port,
		Password:  config.Password,
		MaxMemory: config.MaxMemory,
	}
}
