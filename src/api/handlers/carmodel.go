package handlers

import (
	_ "CarSaleAd-Web-Api/api/dto"
	_ "CarSaleAd-Web-Api/api/helper"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/services"

	"github.com/gin-gonic/gin"
)

type CarModelHandler struct {
	service *services.CarModelService
}

func NewCarModelHandler(cfg *config.Config) *CarModelHandler {
	return &CarModelHandler{
		service: services.NewCarModelService(cfg),
	}
}

// CreateCarModel godoc
// @Summary Create a CarModel
// @Description Create a CarModel
// @Tags CarModel
// @Accept json
// @Produces json
// @Param Request body dto.CreateCarModelRequest true "Create a CarModel"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelResponse} "CarModel response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-models/ [post]
// @Security AuthBearer
func (h *CarModelHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)

}

// UpdateCarModel godoc
// @Summary Update a CarModel
// @Description Update a CarModel
// @Tags CarModel
// @Accept json
// @Produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelRequest true "Update a CarModel"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelResponse} "CarModel response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-models/{id} [put]
// @Security AuthBearer
func (h *CarModelHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)

}

// DeleteCarModel godoc
// @Summary Delete a CarModel
// @Description Delete a CarModel
// @Tags CarModel
// @Accept json
// @Produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-models/{id} [delete]
// @Security AuthBearer
func (h *CarModelHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)

}

// GetCarModel godoc
// @Summary Get a CarModel
// @Description Get a CarModel
// @Tags CarModel
// @Accept json
// @Produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelResponse} "CarModel Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-models/{id} [get]
// @Security AuthBearer
func (h *CarModelHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)

}

// GetCarModels godoc
// @Summary Get CarModels
// @Description Get CarModels
// @Tags CarModel
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelResponse]} "CarModel Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-models/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)

}
