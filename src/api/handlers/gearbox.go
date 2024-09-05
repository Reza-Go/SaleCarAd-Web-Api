package handlers

import (
	_ "CarSaleAd-Web-Api/api/dto"
	_ "CarSaleAd-Web-Api/api/helper"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/services"

	"github.com/gin-gonic/gin"
)

type GearboxHandler struct {
	service *services.GearboxService
}

func NewGearboxHandler(cfg *config.Config) *GearboxHandler {
	return &GearboxHandler{
		service: services.NewGearboxService(cfg),
	}
}

// CreateGearbox godoc
// @Summary Create a Gearbox
// @Description Create a Gearbox
// @Tags Gearbox
// @Accept json
// @Produces json
// @Param Request body dto.CreateGearboxRequest true "Create a Gearbox"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.GearboxResponse} "Gearbox response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/gearboxes/ [post]
// @Security AuthBearer
func (h *GearboxHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)

}

// UpdateGearbox godoc
// @Summary Update a Gearbox
// @Description Update a Gearbox
// @Tags Gearbox
// @Accept json
// @Produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateGearboxRequest true "Update a Gearbox"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.GearboxResponse} "Gearbox response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/gearboxes/{id} [put]
// @Security AuthBearer
func (h *GearboxHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)

}

// DeleteGearbox godoc
// @Summary Delete a Gearbox
// @Description Delete a Gearbox
// @Tags Gearbox
// @Accept json
// @Produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/gearboxes/{id} [delete]
// @Security AuthBearer
func (h *GearboxHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)

}

// GetGearbox godoc
// @Summary Get a Gearbox
// @Description Get a Gearbox
// @Tags Gearbox
// @Accept json
// @Produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.GearboxResponse} "Gearbox Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/gearboxes/{id} [get]
// @Security AuthBearer
func (h *GearboxHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)

}

// GetGearboxes godoc
// @Summary Get Gearboxes
// @Description Get Gearboxes
// @Tags Gearbox
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.GearboxResponse]} "Gearbox Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/gearboxes/get-by-filter [post]
// @Security AuthBearer
func (h *GearboxHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)

}
