package routers

import (
	"CarSaleAd-Web-Api/api/handlers"
	"CarSaleAd-Web-Api/config"

	"github.com/gin-gonic/gin"
)

func File(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewFileHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}
