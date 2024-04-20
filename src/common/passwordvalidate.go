package common

import (
	"CarSaleAd-Web-Api/config"
	"unicode"
)

func CheckPassword(password string) bool {
	cfg := config.GetConfig()
	if len(password) < cfg.Password.MinLength {
		return false
	}
	if len(password) > cfg.Password.MaxLength {
		return false
	}
	if cfg.Password.IncludeChars && !HasLetter(password) {
		return false
	}
	if cfg.Password.IncludeDigits && !HasDigits(password) {
		return false
	}
	if cfg.Password.IncludeLowercase && !HasLower(password) {
		return false
	}
	if cfg.Password.IncludeUppercase && !HasUpper(password) {
		return false
	}
	return true
}

func HasLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}
func HasDigits(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}
func HasLower(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && unicode.IsLower(r) {
			return true
		}
	}
	return false
}
func HasUpper(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && unicode.IsUpper(r) {
			return true
		}
	}
	return false
}
