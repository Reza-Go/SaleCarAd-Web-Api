package routers

import (
	"CarSaleAd-Web-Api/api/handlers"
	middlewares "CarSaleAd-Web-Api/api/midllewares"
	"CarSaleAd-Web-Api/config"

	"github.com/gin-gonic/gin"
)

func User(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUsersHandler(cfg)

	r.POST("/send-otp", middlewares.OtpLimiter(cfg), h.SendOtp)
	r.POST("/login-by-username", h.LoginByUsername)
	r.POST("/register-by-username", h.RegisterByUsername)
	r.POST("/login-by-mobile", h.RegisterLoginByMobileNumber)
}
