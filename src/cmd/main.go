package main

import (
	"CarSaleAd-Web-Api/api"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/data/cache"
	"CarSaleAd-Web-Api/data/db"
	"CarSaleAd-Web-Api/data/db/migrations"
	"CarSaleAd-Web-Api/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.StartUp, err.Error(), nil)
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.StartUp, err.Error(), nil)
	}
	migrations.Up_1()

	api.InitServer(cfg)

}
