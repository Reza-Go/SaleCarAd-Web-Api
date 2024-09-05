package routers

import (
	"CarSaleAd-Web-Api/api/handlers"
	"CarSaleAd-Web-Api/config"

	"github.com/gin-gonic/gin"
)

func CarType(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCarTypeHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}
func Gearbox(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewGearboxHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}
func Company(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCompanyHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}

func CarModel(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCarModelHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}
