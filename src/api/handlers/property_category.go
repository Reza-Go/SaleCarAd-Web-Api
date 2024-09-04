package handlers

import (
	_ "CarSaleAd-Web-Api/api/dto"
	_ "CarSaleAd-Web-Api/api/helper"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/services"

	"github.com/gin-gonic/gin"
)

type PropertyCategoryHandler struct {
	service *services.PropertyCategoryService
}

func NewPropertyCategoryHandler(cfg *config.Config) *PropertyCategoryHandler {
	return &PropertyCategoryHandler{
		service: services.NewPropertyCategoryService(cfg),
	}
}

// CreatePropertyCategory godoc
// @Summary Create a PropertyCategory
// @Description Create a PropertyCategory
// @Tags PropertyCategory
// @Accept json
// @Produces json
// @Param Request body dto.CreatePropertyCategoryRequest true "Create a PropertyCategory"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.PropertyCategoryResponse} "PropertyCategory response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/property-categories/ [post]
// @Security AuthBearer
func (h *PropertyCategoryHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)

}

// UpdatePropertyCategory godoc
// @Summary Update a PropertyCategory
// @Description Update a PropertyCategory
// @Tags PropertyCategory
// @Accept json
// @Produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdatePropertyCategoryRequest true "Update a PropertyCategory"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PropertyCategoryResponse} "PropertyCategory response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/property-categories/{id} [put]
// @Security AuthBearer
func (h *PropertyCategoryHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)

}

// DeletePropertyCategory godoc
// @Summary Delete a PropertyCategory
// @Description Delete a PropertyCategory
// @Tags PropertyCategory
// @Accept json
// @Produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/property-categories/{id} [delete]
// @Security AuthBearer
func (h *PropertyCategoryHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)

}

// GetPropertyCategory godoc
// @Summary Get a PropertyCategory
// @Description Get a PropertyCategory
// @Tags PropertyCategory
// @Accept json
// @Produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PropertyCategoryResponse} "PropertyCategory Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/property-categories/{id} [get]
// @Security AuthBearer
func (h *PropertyCategoryHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)

}

// GetPropertyCategories godoc
// @Summary Get PropertyCategories
// @Description Get PropertyCategories
// @Tags PropertyCategory
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.PropertyCategoryResponse]} "PropertyCategory Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/property-categories/get-by-filter [post]
// @Security AuthBearer
func (h *PropertyCategoryHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)

}
