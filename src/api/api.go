package api

import (
	"CarSaleAd-Web-Api/api/routers"
	"CarSaleAd-Web-Api/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	//r1 := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())

	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		routers.Health(health)
		
		test_router := v1.Group("/test")
		routers.TestRouter(test_router)

	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
