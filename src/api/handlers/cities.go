package handlers

import (
	"CarSaleAd-Web-Api/api/dto"
	"CarSaleAd-Web-Api/api/helper"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CityHandler struct {
	service *services.CityService
}

func NewCityHandler(cfg *config.Config) *CityHandler {
	return &CityHandler{
		service: services.NewCityService(cfg),
	}
}

// CreateCity godoc
// @Summary Create a City
// @Description Create a City
// @Tags City
// @Accept json
// @Produces json
// @Param Request body dto.CreateUpdateCityRequest true "Create a City"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CityResponse} "City response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/cities/ [post]
// @Security AuthBearer
func (h *CityHandler) Create(c *gin.Context) {
	req := dto.CreateUpdateCityRequest{}
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, 121, err))
		return
	}

	res, err := h.service.Create(c, &req)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, 0))

}

// UpdateCity godoc
// @Summary Update a City
// @Description Update a City
// @Tags City
// @Accept json
// @Produces json
// @Param id path int true "Id"
// @Param Request body dto.CreateUpdateCityRequest true "Update a City"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CityResponse} "City response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/cities/{id} [put]
// @Security AuthBearer
func (h *CityHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	req := dto.CreateUpdateCityRequest{}
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, 121, err))
		return
	}

	res, err := h.service.Update(c, id, &req)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))

}

// DeleteCity godoc
// @Summary Delete a City
// @Description Delete a City
// @Tags City
// @Accept json
// @Produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/cities/{id} [delete]
// @Security AuthBearer
func (h *CityHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, 121))
		return

	}

	err := h.service.Delete(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(nil, true, 0))

}

// GetCity godoc
// @Summary Get a City
// @Description Get a City
// @Tags City
// @Accept json
// @Produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CityResponse} "City Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/cities/{id} [get]
// @Security AuthBearer
func (h *CityHandler) GetById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, 121))
		return
	}

	res, err := h.service.GetById(c, id)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))

}

// GetCities godoc
// @Summary Get Cities
// @Description Get Cities
// @Tags City
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CityResponse]} "City Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/cities/get-by-filter [post]
// @Security AuthBearer
func (h *CityHandler) GetByFilter(c *gin.Context) {
	req := dto.PaginationInputWithFilter{}
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, 121, err))
		return
	}

	res, err := h.service.GetByFilter(c, &req)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))

}
