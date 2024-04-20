package validation

import (
	"CarSaleAd-Web-Api/common"

	"github.com/go-playground/validator/v10"
)

func PasswordValidator(fld validator.FieldLevel) bool {

	value, ok := fld.Field().Interface().(string)

	if !ok {
		fld.Param()
		return false
	}

	return common.CheckPassword(value)
}
