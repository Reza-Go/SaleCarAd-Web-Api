package handlers

import (
	"CarSaleAd-Web-Api/api/dto"
	"CarSaleAd-Web-Api/api/helper"
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	service *services.UserService
}

func NewUsersHandler(cfg *config.Config) *UsersHandler {
	service := services.NewUserService(cfg)
	return &UsersHandler{service: service}
}

// SendOtp godoc
// @Summary SendOtp to user
// @Description Test SendOtp
// @Tags User
// @Accept json
// @Produce json
// @Param Request body dto.GetOtpRequest true "GetOtpRequest"
// @Success 201 {object} helper.BaseHttpResponse "Done"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/send-otp [post]
func (h *UsersHandler) SendOtp(c *gin.Context) {
	req := new(dto.GetOtpRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	err = h.service.SendOtp(req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, 0))
}
