package handlers

import (
	_ "CarSaleAd-Web-Api/api/dto"
	_ "CarSaleAd-Web-Api/api/helper"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/services"

	"github.com/gin-gonic/gin"
)

type PersianYearHandler struct {
	service *services.PersianYearService
}

func NewPersianYearHandler(cfg *config.Config) *PersianYearHandler {
	return &PersianYearHandler{
		service: services.NewPersianYearService(cfg),
	}
}

// CreatePersianYear godoc
// @Summary Create a PersianYear
// @Description Create a PersianYear
// @Tags PersianYear
// @Accept json
// @Produces json
// @Param Request body dto.CreatePersianYearRequest true "Create a PersianYear"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.PersianYearResponse} "PersianYear response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/years/ [post]
// @Security AuthBearer
func (h *PersianYearHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)

}

// UpdatePersianYear godoc
// @Summary Update a PersianYear
// @Description Update a PersianYear
// @Tags PersianYear
// @Accept json
// @Produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdatePersianYearRequest true "Update a PersianYear"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PersianYearResponse} "PersianYear response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/years/{id} [put]
// @Security AuthBearer
func (h *PersianYearHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)

}

// DeletePersianYear godoc
// @Summary Delete a PersianYear
// @Description Delete a PersianYear
// @Tags PersianYear
// @Accept json
// @Produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/years/{id} [delete]
// @Security AuthBearer
func (h *PersianYearHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)

}

// GetPersianYear godoc
// @Summary Get a PersianYear
// @Description Get a PersianYear
// @Tags PersianYear
// @Accept json
// @Produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PersianYearResponse} "PersianYear Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/years/{id} [get]
// @Security AuthBearer
func (h *PersianYearHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)

}

// GetPersianYears godoc
// @Summary Get PersianYears
// @Description Get PersianYears
// @Tags PersianYear
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.PersianYearResponse]} "PersianYear Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/years/get-by-filter [post]
// @Security AuthBearer
func (h *PersianYearHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)

}
