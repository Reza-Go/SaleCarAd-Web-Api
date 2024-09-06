package handlers

import (
	_ "CarSaleAd-Web-Api/api/dto"
	_ "CarSaleAd-Web-Api/api/helper"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/services"

	"github.com/gin-gonic/gin"
)

type CarModelYearHandler struct {
	service *services.CarModelYearService
}

func NewCarModelYearHandler(cfg *config.Config) *CarModelYearHandler {
	return &CarModelYearHandler{
		service: services.NewCarModelYearService(cfg),
	}
}

// CreateCarModelYear godoc
// @Summary Create a CarModelYear
// @Description Create a CarModelYear
// @Tags CarModelYear
// @Accept json
// @Produces json
// @Param Request body dto.CreateCarModelYearRequest true "Create a CarModelYear"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelYearResponse} "CarModelYear response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-years/ [post]
// @Security AuthBearer
func (h *CarModelYearHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)

}

// UpdateCarModelYear godoc
// @Summary Update a CarModelYear
// @Description Update a CarModelYear
// @Tags CarModelYear
// @Accept json
// @Produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelYearRequest true "Update a CarModelYear"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelYearResponse} "CarModelYear response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-years/{id} [put]
// @Security AuthBearer
func (h *CarModelYearHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)

}

// DeleteCarModelYear godoc
// @Summary Delete a CarModelYear
// @Description Delete a CarModelYear
// @Tags CarModelYear
// @Accept json
// @Produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-years/{id} [delete]
// @Security AuthBearer
func (h *CarModelYearHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)

}

// GetCarModelYear godoc
// @Summary Get a CarModelYear
// @Description Get a CarModelYear
// @Tags CarModelYear
// @Accept json
// @Produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelYearResponse} "CarModelYear Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-years/{id} [get]
// @Security AuthBearer
func (h *CarModelYearHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)

}

// GetCarModelYears godoc
// @Summary Get CarModelYears
// @Description Get CarModelYears
// @Tags CarModelYear
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelYearResponse]} "CarModelYear Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-years/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelYearHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)

}
