package services

import (
	"CarSaleAd-Web-Api/api/dto"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/constants"
	"CarSaleAd-Web-Api/data/db"
	"CarSaleAd-Web-Api/data/models"
	"CarSaleAd-Web-Api/pkg/logging"
	"context"
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type CountryService struct {
	database *gorm.DB
	logger   logging.Logger
}

func NewCountryService(cfg *config.Config) *CountryService {
	return &CountryService{
		database: db.GetDb(),
		logger:   logging.NewLogger(cfg),
	}
}

// Create
func (s *CountryService) Create(ctx context.Context, req *dto.CreateUpdateCountryRequest) (*dto.CountryResponse, error) {
	country := models.Country{Name: req.Name}
	country.CreatedBy = int(ctx.Value(constants.UserIdKey).(float64))
	country.CreatedAt = time.Now().UTC()

	tx := s.database.WithContext(ctx).Begin()
	err := tx.Create(&country).Error

	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Insert, err.Error(), nil)
		return nil, err
	}
	tx.Commit()

	dto := &dto.CountryResponse{Name: country.Name, Id: country.Id}

	return dto, nil

}

//update

func (s *CountryService) Update(ctx context.Context, id int, req *dto.CreateUpdateCountryRequest) (*dto.CountryResponse, error) {
	updateMap := map[string]interface{}{
		"Name":        req.Name,
		"modified_by": sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true},
		"modified_at": sql.NullTime{Time: time.Now().UTC(), Valid: true},
	}

	tx := s.database.WithContext(ctx).Begin()

	err := tx.
		Model(&models.Country{}).
		Where("id = ?", id).
		Updates(updateMap).
		Error

	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
		return nil, err
	}

	country := models.Country{}

	err = tx.
		Model(&models.Country{}).
		Where("id = ? AND deleted_by is null", id).
		First(&country).Error

	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return nil, err
	}

	tx.Commit()
	dto := &dto.CountryResponse{Name: country.Name, Id: country.Id}

	return dto, nil

}

//Delete

func (s *CountryService) Delete(ctx context.Context, id int) error {
	tx := s.database.WithContext(ctx).Begin()
	deleteMap := map[string]interface{}{
		"deleted_by": sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true},
		"deleted_at": sql.NullTime{Time: time.Now().UTC(), Valid: true},
	}

	if err := tx.
		Model(&models.Country{}).
		Where("id = ?", id).
		Updates(deleteMap).
		Error; err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Delete, err.Error(), nil)
		return err
	}
	tx.Commit()
	return nil

}

// GetById
func (s *CountryService) GetById(ctx context.Context, id int) (*dto.CountryResponse, error) {
	country := &models.Country{}

	err := s.database.
		Where("id = ?", id).
		First(&country).
		Error

	if err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return nil, err
	}

	dto := &dto.CountryResponse{Name: country.Name, Id: country.Id}
	return dto, nil
}
