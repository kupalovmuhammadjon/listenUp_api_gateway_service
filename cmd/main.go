package main

import (
	"api_gateway/api"
	"api_gateway/config"
)

func main() {
	cfg := config.Load()
	
	router := api.NewRouter(cfg)
	router.Run(cfg.HTTP_PORT)
}
