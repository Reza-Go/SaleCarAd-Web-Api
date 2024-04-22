package main

import (
	"CarSaleAd-Web-Api/api"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/data/cache"
)

func main() {
	cfg := config.GetConfig()
	cache.InitRedis(cfg)
	defer cache.CloseRedis()
	api.InitServer(cfg)

}
