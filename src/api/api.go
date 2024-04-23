package api

import (
	middlewares "CarSaleAd-Web-Api/api/midllewares"
	"CarSaleAd-Web-Api/api/routers"
	validation "CarSaleAd-Web-Api/api/validations"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/docs"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {

	r := gin.New()
	//r1 := gin.Default()

	//validations
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
		val.RegisterValidation("password", validation.PasswordValidator, true)
	}

	//middlewares
	r.Use(gin.Logger(), gin.Recovery() /*middlewares.TestMiddleware()*/, middlewares.LimitByRequest())

	//Routes
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		routers.Health(health)

		test_router := v1.Group("/test")
		routers.TestRouter(test_router)

	}
	//Swagger
	RegisterSwagger(r, cfg)

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "Car Sale Ad Web Api"
	docs.SwaggerInfo.Description = "Sell your Car"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
