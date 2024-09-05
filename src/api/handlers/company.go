package handlers

import (
	_ "CarSaleAd-Web-Api/api/dto"
	_ "CarSaleAd-Web-Api/api/helper"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/services"

	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	service *services.CompanyService
}

func NewCompanyHandler(cfg *config.Config) *CompanyHandler {
	return &CompanyHandler{
		service: services.NewCompanyService(cfg),
	}
}

// CreateCompany godoc
// @Summary Create a Company
// @Description Create a Company
// @Tags Company
// @Accept json
// @Produces json
// @Param Request body dto.CreateCompanyRequest true "Create a Company"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CompanyResponse} "Company response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/companies/ [post]
// @Security AuthBearer
func (h *CompanyHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)

}

// UpdateCompany godoc
// @Summary Update a Company
// @Description Update a Company
// @Tags Company
// @Accept json
// @Produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCompanyRequest true "Update a Company"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CompanyResponse} "Company response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/companies/{id} [put]
// @Security AuthBearer
func (h *CompanyHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)

}

// DeleteCompany godoc
// @Summary Delete a Company
// @Description Delete a Company
// @Tags Company
// @Accept json
// @Produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/companies/{id} [delete]
// @Security AuthBearer
func (h *CompanyHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)

}

// GetCompany godoc
// @Summary Get a Company
// @Description Get a Company
// @Tags Company
// @Accept json
// @Produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CompanyResponse} "Company Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/companies/{id} [get]
// @Security AuthBearer
func (h *CompanyHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)

}

// GetCompanies godoc
// @Summary Get Companies
// @Description Get Companies
// @Tags Company
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CompanyResponse]} "Company Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/companies/get-by-filter [post]
// @Security AuthBearer
func (h *CompanyHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)

}
