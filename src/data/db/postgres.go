package db

import (
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/pkg/logging"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var logger = logging.NewLogger(config.GetConfig())
var dbClient *gorm.DB

func InitDb(cfg *config.Config) error {
	var err error
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password,
		cfg.Postgres.DbName, cfg.Postgres.SSLMode)
	dbClient, err = gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDb, _ := dbClient.DB()

	err = sqlDb.Ping()
	if err != nil {
		return err
	}

	sqlDb.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDb.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)

	logger.Info(logging.Postgres, logging.StartUp, "DB connection established", nil)
	return nil

}

func GetDb() *gorm.DB {
	return dbClient
}

func CloseDb() {
	con, _ := dbClient.DB()
	con.Close()
}
