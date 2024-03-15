package api

import (
	"CarSaleAd-Web-Api/api/routers"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	r := gin.New()
	//r1 := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1/")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}

	r.Run(":5005")
}
