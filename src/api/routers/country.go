package routers

import (
	"CarSaleAd-Web-Api/api/handlers"
	"CarSaleAd-Web-Api/config"

	"github.com/gin-gonic/gin"
)

func Country(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCountryHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
}
