package services

import (
	"CarSaleAd-Web-Api/api/dto"
	"CarSaleAd-Web-Api/common"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/data/db"
	"CarSaleAd-Web-Api/pkg/logging"

	"gorm.io/gorm"
)

type UserService struct {
	logger     logging.Logger
	cfg        *config.Config
	OtpService *OtpService
	database   *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	database := db.GetDb()
	logger := logging.NewLogger(cfg)
	return &UserService{
		cfg:        cfg,
		logger:     logger,
		database:   database,
		OtpService: NewOtpService(cfg),
	}
}

func (s *UserService) SendOtp(req *dto.GetOtpRequest) error {
	otp := common.GenerateOtp()
	err := s.OtpService.SetOtp(req.MobileNumber, otp)
	if err != nil {
		return err
	}
	return nil

}
