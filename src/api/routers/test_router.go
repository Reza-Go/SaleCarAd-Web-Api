package routers

import (
	"CarSaleAd-Web-Api/api/handlers"

	"github.com/gin-gonic/gin"
)

func TestRouter(r *gin.RouterGroup) {
	h := handlers.NewTestHandler()

	r.GET("/", h.Test)
	r.GET("/users", h.Users)
	r.GET("/user/:id", h.UserById)
	r.GET("/user/get-user-by-username/:username", h.UserByUsername)
	r.GET("/user/:id/accounts", h.Accounts)
	r.POST("/add-user", h.AddUser)

	r.GET("/binder/header1", h.HeaderBinder1)
	r.GET("/binder/header2", h.HeaderBinder2)

	r.GET("/binder/query1", h.QueryBinder1)
	r.GET("/binder/query2", h.QueryBinder2)

	r.GET("/binder/uri/:id/:name", h.UriBinder)
	r.POST("/binder/body", h.BodyBinder)
	r.GET("/binder/form", h.FormBinder)
	r.GET("/binder/file", h.FileBinder)

}
