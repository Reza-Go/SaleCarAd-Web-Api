package services

import (
	"CarSaleAd-Web-Api/api/dto"
	"CarSaleAd-Web-Api/common"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/constants"
	"CarSaleAd-Web-Api/data/db"
	"CarSaleAd-Web-Api/data/models"
	"CarSaleAd-Web-Api/pkg/logging"
	"CarSaleAd-Web-Api/pkg/service_errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	logger       logging.Logger
	cfg          *config.Config
	OtpService   *OtpService
	tokenService *TokenService
	database     *gorm.DB
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

// Register by Username
func (s *UserService) RegisterByUsername(req dto.RegisterUserByUsernameRequest) error {
	u := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Email:     req.Email,
	}
	exists, err := s.existsByEmail(req.Email)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.EmailExists}
	}
	exists, err = s.existsByUsername(req.Username)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.UsernameExists}
	}
	//string >>>byte
	bp := []byte(req.Password)
	//hash passwrod
	hp, err := bcrypt.GenerateFromPassword(bp, bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error(logging.Internal, logging.HashPassword, err.Error(), nil)
		return err
	}
	u.Password = string(hp)
	roleId, err := s.getDefaultRole()
	if err != nil {
		s.logger.Error(logging.Postgres, logging.DefaultRoleNotFound, err.Error(), nil)
		return err
	}
	//start transaction
	tx := s.database.Begin()
	err = tx.Create(&u).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return err
	}
	err = tx.Create(&models.UserRole{RoleId: roleId, UserId: u.Id}).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return err
	}
	tx.Commit()
	return nil

}

// Register or login by mobile number
func (s *UserService) RegisterLoginByMobileNumber(req dto.RegisterLoginByMobileRequest) (*dto.TokenDetail, error) {
	err := s.OtpService.ValidateOtp(req.MobileNumber, req.Otp)
	if err != nil {
		return nil, err
	}
	exists, err := s.existsByMobileNumber(req.MobileNumber)
	if err != nil {
		return nil, err
	}

	u := models.User{MobileNumber: req.MobileNumber, Username: req.MobileNumber}
	//MObile exist >>>Login
	if exists {
		var user models.User
		err = s.database.
			Model(&models.User{}).
			Where("username = ?", u.Username).
			Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
				return tx.Preload("Role")

			}).Find(&user).Error
		if err != nil {
			return nil, err
		}
		tdto := tokenDto{UserId: user.Id, FirstName: user.FirstName, LastName: user.LastName,
			Email: user.Email, MobileNumber: user.MobileNumber}
		if len(*user.UserRoles) > 0 {
			for _, ur := range *user.UserRoles {
				tdto.Roles = append(tdto.Roles, ur.Role.Name)
			}
		}
		token, err := s.tokenService.GenerateToken(&tdto)
		if err != nil {
			return nil, err
		}
		return token, nil

		//Mobile Number does not exist >>>Register
		//Create Password
	}
	bp := []byte(common.GeneratePassword())
	hp, err := bcrypt.GenerateFromPassword(bp, bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error(logging.General, logging.HashPassword, err.Error(), nil)
		return nil, err
	}
	u.Password = string(hp)
	roleId, err := s.getDefaultRole()
	if err != nil {
		s.logger.Error(logging.Postgres, logging.DefaultRoleNotFound, err.Error(), nil)
		return nil, err
	}

	//Create User And UserRoles

	tx := s.database.Begin()
	err = tx.Create(&u).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return nil, err
	}
	err = tx.Create(&models.UserRole{RoleId: roleId, UserId: u.Id}).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return nil, err
	}
	tx.Commit()

	//Create TokenDto and Token

	var user models.User
	err = s.database.
		Model(&models.User{}).
		Where("username = ?", u.Username).
		Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Role")
		}).
		Find(&user).Error
	if err != nil {
		return nil, err
	}
	tdto := tokenDto{UserId: user.Id, FirstName: user.FirstName, LastName: user.LastName,
		MobileNumber: user.MobileNumber, Email: user.Email}

	if len(*user.UserRoles) > 0 {
		for _, ur := range *user.UserRoles {
			tdto.Roles = append(tdto.Roles, ur.Role.Name)
		}

	}

	token, err := s.tokenService.GenerateToken(&tdto)
	if err != nil {
		return nil, err
	}
	return token, nil

}

func (s *UserService) SendOtp(req *dto.GetOtpRequest) error {
	otp := common.GenerateOtp()
	err := s.OtpService.SetOtp(req.MobileNumber, otp)
	if err != nil {
		return err
	}
	return nil

}

func (s *UserService) existsByEmail(email string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("email = ? ", email).
		Find(&exists).Error; err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil

}

func (s *UserService) existsByUsername(username string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("username = ?", username).
		Find(&exists).Error; err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (s *UserService) existsByMobileNumber(mobileNumber string) (bool, error) {
	var exists bool
	if err := s.database.
		Model(&models.User{}).
		Select("count(*) > 0").
		Where("mobile_number = ?", mobileNumber).
		Find(&exists).Error; err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (s *UserService) getDefaultRole() (roleId int, err error) {
	if err = s.database.Model(&models.Role{}).
		Select("id").
		Where("name = ?", constants.DefaultRoleName).
		First(&roleId).Error; err != nil {
		return 0, err
	}
	return roleId, nil
}
