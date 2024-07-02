package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT                   string
	USER_SERVICE_PORT           string
	COLLABORATIONS_SERVICE_PORT string
	DISCOVERY_SERVICE_PORT      string
	PODCAST_SERVICE_PORT        string
}

func Load() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot load .env ", err)
	}

	cfg := Config{}

	cfg.HTTP_PORT = cast.ToString(coalesce("HTTP_PORT", ":8080"))
	cfg.USER_SERVICE_PORT = cast.ToString(coalesce("USER_SERVICE_PORT", ":8081"))
	cfg.COLLABORATIONS_SERVICE_PORT = cast.ToString(coalesce("COLLABORATIONS_SERVICE_PORT", ":8082"))
	cfg.DISCOVERY_SERVICE_PORT = cast.ToString(coalesce("DISCOVERY_SERVICE_PORT", ":8083"))
	cfg.PODCAST_SERVICE_PORT = cast.ToString(coalesce("PODCAST_SERVICE_PORT", ":8084"))

	return &cfg
}

func coalesce(key string, value interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return value
}
