package config

import (
	"context"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

var instance *Config
var once sync.Once

// GetConfig read config and returns a pointer to the Config struct
func GetConfig(ctx context.Context) *Config {
	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			helpText := "TonocoAdmin"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})

	return instance
}
