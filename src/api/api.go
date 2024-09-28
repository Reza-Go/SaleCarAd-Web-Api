package api

import (
	middlewares "CarSaleAd-Web-Api/api/midllewares"
	"CarSaleAd-Web-Api/api/routers"
	validation "CarSaleAd-Web-Api/api/validations"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/docs"
	"CarSaleAd-Web-Api/pkg/logging"
	"CarSaleAd-Web-Api/pkg/metrics"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var logger = logging.NewLogger(config.GetConfig())

func InitServer(cfg *config.Config) {

	r := gin.New()
	//r1 := gin.Default()

	//Prometheus
	RegisterPrometheus()

	//validations
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
		if err != nil {
			logger.Error(logging.Validation, logging.StartUp, err.Error(), nil)
		}
		err = val.RegisterValidation("password", validation.PasswordValidator, true)
		if err != nil {
			logger.Error(logging.Validation, logging.StartUp, err.Error(), nil)
		}
	}

	//middlewares
	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(middlewares.Prometheus())
	r.Use(gin.Logger(), gin.Recovery() /*middlewares.TestMiddleware()*/, middlewares.LimitByRequest())

	//Routes
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		routers.Health(health)

		test_router := v1.Group("/test")
		routers.TestRouter(test_router)

		user_router := v1.Group("/users")
		routers.User(user_router, cfg)

		countries := v1.Group("/countries", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.Country(countries, cfg)

		cities := v1.Group("/cities", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.City(cities, cfg)

		files := v1.Group("/files", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.File(files, cfg)

		propertyCategories := v1.Group("/property-categories", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.PropertyCategory(propertyCategories, cfg)

		properties := v1.Group("/properties", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.Property(properties, cfg)

		companies := v1.Group("/companies", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.Company(companies, cfg)

		carTypes := v1.Group("/car-types", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.CarType(carTypes, cfg)

		gearboxes := v1.Group("/gearboxes", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.Gearbox(gearboxes, cfg)

		carModels := v1.Group("/car-models", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin", "default"}))
		routers.CarModel(carModels, cfg)

		colors := v1.Group("/colors", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.Color(colors, cfg)

		carModelColors := v1.Group("/car-model-colors", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.CarModelColor(carModelColors, cfg)

		years := v1.Group("/years", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.Year(years, cfg)

		carModelYears := v1.Group("/car-model-years", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.CarModelYear(carModelYears, cfg)

		carModelPriceHistories := v1.Group("/car-model-price-histories", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.CarModelPriceHistory(carModelPriceHistories, cfg)

		carModelImages := v1.Group("/car-model-images", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.CarModelImage(carModelImages, cfg)

		carModelProperties := v1.Group("/car-model-properties", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.CarModelProperty(carModelProperties, cfg)

		carModelComments := v1.Group("/car-model-comments", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin", "default"}))
		routers.CarModelComment(carModelComments, cfg)

		r.Static("/static", "./uploads")

		r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	}
	//Swagger
	RegisterSwagger(r, cfg)

	err := r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
	if err != nil {
		logger.Error(logging.General, logging.StartUp, err.Error(), nil)
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "Car Sale Ad Web Api"
	docs.SwaggerInfo.Description = "Sell your Car"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.ExternalPort)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func RegisterPrometheus() {
	err := prometheus.Register(metrics.DbCall)
	if err != nil {
		logger.Error(logging.Prometheus, logging.StartUp, err.Error(), nil)
	}

	err = prometheus.Register(metrics.HttpDuration)
	if err != nil {
		logger.Error(logging.Prometheus, logging.StartUp, err.Error(), nil)
	}
}
