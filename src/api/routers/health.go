package routers

import (
	"CarSaleAd-Web-Api/api/handlers"

	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()

	r.GET("/", handler.Health)
	r.POST("/:id", handler.HealthPostById)
}
