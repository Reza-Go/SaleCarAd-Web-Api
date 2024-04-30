package common

import (
	"CarSaleAd-Web-Api/config"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func GenerateOtp() string {
	cfg := config.GetConfig()
	rand.Seed(time.Now().UnixNano())
	min := int(math.Pow(10, float64(cfg.Otp.Digits-1)))   //10^(d-1)  100000
	max := int(math.Pow(10, float64(cfg.Otp.Digits)) - 1) //(10^d)-1  999999

	var num = rand.Intn(max-min) + min
	return strconv.Itoa(num)
}
