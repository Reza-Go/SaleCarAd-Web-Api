package handlers

import (
	"CarSaleAd-Web-Api/api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type personData struct {
	FirstName    string `json:"first_name" binding:"required,alpha,min=4,max=10"`
	LastName     string `json:"last_name" binding:"required,alpha,min=6,max=20"`
	MobileNumber string `json:"mobile_number" binding:"required,mobile,min=11,max=11"`
}

type header struct {
	UserId  string
	Browser string
}

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{
	// 	"result": "test",
	// })
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("Test", true, 0))
}
func (h *TestHandler) Users(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("Users", true, 0))
}

func (h *TestHandler) UserById(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "UserById",
		"id":     id,
	}, true, 0))
}

func (h *TestHandler) UserByUsername(c *gin.Context) {
	username := c.Param("username")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result":   "UserByUsername",
		"username": username,
	}, true, 0))
}

func (h *TestHandler) Accounts(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "Accounts",
		"id":     id,
	}, true, 0))
}

func (h *TestHandler) AddUser(c *gin.Context) {

	c.JSON(http.StatusOK, helper.GenerateBaseResponse("AddUser", true, 0))
}

func (h *TestHandler) HeaderBinder1(c *gin.Context) {
	userId := c.GetHeader("UserId")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "HeaderBinder1",
		"userId": userId,
	}, true, 0))
}
func (h *TestHandler) HeaderBinder2(c *gin.Context) {
	header := header{}
	c.BindHeader(&header)
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "HeaderBinder2",
		"header": header,
	}, true, 0))

}
func (h *TestHandler) QueryBinder1(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "QueryBinder1",
		"id":     id,
		"name":   name,
	}, true, 0))
}

func (h *TestHandler) QueryBinder2(c *gin.Context) {
	ids := c.QueryArray("id")
	name := c.Query("name")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "QueryBinder2",
		"ids":    ids,
		"name":   name,
	}, true, 0))
}

func (h *TestHandler) UriBinder(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "UriBinder",
		"id":     id,
		"name":   name,
	}, true, 0))
}

func (h *TestHandler) BodyBinder(c *gin.Context) {
	p := personData{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "BodyBinder",
		"person": p,
	}, true, 0))
}

func (h *TestHandler) FormBinder(c *gin.Context) {
	p := personData{}
	c.ShouldBind(&p)
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "FormBinder",
		"person": p,
	}, true, 0))

}

func (h *TestHandler) FileBinder(c *gin.Context) {
	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, "file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			helper.GenerateBaseResponseWithError(nil, false, -1, err))

		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "FileBinder",
		"file":   file.Filename,
	}, true, 0))
}
