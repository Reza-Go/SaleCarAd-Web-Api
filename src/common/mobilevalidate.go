package common

import (
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/pkg/logging"
	"regexp"
)

var logger = logging.NewLogger(config.GetConfig())

const IranianMobileNumberPattern string = `^09(1[0-9]|2[0-9]|3[0-9]|9[0-9])[0-9]{7}$`

func IranianMobileNumberValidate(mobileNumber string) bool {

	res, err := regexp.MatchString(IranianMobileNumberPattern, mobileNumber)

	if err != nil {
		logger.Error(logging.Validation, logging.MobileValidation, err.Error(), nil)
	}
	return res
}
