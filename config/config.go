package config

import (
	"log"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
}

var Cfg Config
var once sync.Once

func GetConfig() Config {
	once.Do(func() {
		instance, err := loadConfig()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		Cfg = instance
	})
	return Cfg
}

func loadConfig() (Config, error) {
	err := godotenv.Load(filepath.Join("..", ".env"))

	return Config{}, err

}
