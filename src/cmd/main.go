package main

import (
	"CarSaleAd-Web-Api/api"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/data/cache"
	"CarSaleAd-Web-Api/data/db"
	"log"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		log.Fatal(err)
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		log.Fatal(err)
	}

	api.InitServer(cfg)

}
