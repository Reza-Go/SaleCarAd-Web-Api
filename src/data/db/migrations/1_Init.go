package migrations

import (
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/data/db"
	"CarSaleAd-Web-Api/data/models"
	"CarSaleAd-Web-Api/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

func Up_1() {
	database := db.GetDb()

	tables := []interface{}{}

	country := models.Country{}

	city := models.City{}

	if !database.Migrator().HasTable(country) {
		tables = append(tables, country)
	}
	if !database.Migrator().HasTable(city) {
		tables = append(tables, city)
	}
	database.Migrator().CreateTable(tables...)
	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)

}

func Down_1() {

}
