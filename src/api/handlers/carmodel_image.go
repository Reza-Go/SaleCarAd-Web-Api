package handlers

import (
	_ "CarSaleAd-Web-Api/api/dto"
	_ "CarSaleAd-Web-Api/api/helper"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/services"

	"github.com/gin-gonic/gin"
)

type CarModelImageHandler struct {
	service *services.CarModelImageService
}

func NewCarModelImageHandler(cfg *config.Config) *CarModelImageHandler {
	return &CarModelImageHandler{
		service: services.NewCarModelImageService(cfg),
	}
}

// CreateCarModelImage godoc
// @Summary Create a CarModelImage
// @Description Create a CarModelImage
// @Tags CarModelImage
// @Accept json
// @Produces json
// @Param Request body dto.CreateCarModelImageRequest true "Create a CarModelImage"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelImageResponse} "CarModelImage response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-images/ [post]
// @Security AuthBearer
func (h *CarModelImageHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)

}

// UpdateCarModelImage godoc
// @Summary Update a CarModelImage
// @Description Update a CarModelImage
// @Tags CarModelImage
// @Accept json
// @Produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelImageRequest true "Update a CarModelImage"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelImageResponse} "CarModelImage response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-images/{id} [put]
// @Security AuthBearer
func (h *CarModelImageHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)

}

// DeleteCarModelImage godoc
// @Summary Delete a CarModelImage
// @Description Delete a CarModelImage
// @Tags CarModelImage
// @Accept json
// @Produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-images/{id} [delete]
// @Security AuthBearer
func (h *CarModelImageHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)

}

// GetCarModelImage godoc
// @Summary Get a CarModelImage
// @Description Get a CarModelImage
// @Tags CarModelImage
// @Accept json
// @Produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelImageResponse} "CarModelImage Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-images/{id} [get]
// @Security AuthBearer
func (h *CarModelImageHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)

}

// GetCarModelImages godoc
// @Summary Get CarModelImages
// @Description Get CarModelImages
// @Tags CarModelImage
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelImageResponse]} "CarModelImage Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-images/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelImageHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)

}
