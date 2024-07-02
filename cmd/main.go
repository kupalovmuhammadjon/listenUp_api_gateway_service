package main

import "api_gateway/api"

func main() {
	router := api.NewRouter()
	router.Run()
}
