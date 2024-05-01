package routers

import (
	"CarSaleAd-Web-Api/api/handlers"
	"CarSaleAd-Web-Api/config"

	"github.com/gin-gonic/gin"
)

func User(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUsersHandler(cfg)

	r.POST("/send-otp", h.SendOtp)
}
