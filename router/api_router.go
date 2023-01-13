package router

import (
	"bookManage/controller"

	"github.com/gin-gonic/gin"
)

func LoadAPIRouter(r *gin.Engine) {
	r.POST("/register", controller.RegisterHandler)
	r.POST("/login", controller.LoginHandler)

	v1 := r.Group("/api/v1")
	v1.POST("book", controller.CreateBookHandler)
	v1.GET("booklist", controller.GetBookListHandler)
	v1.GET("booklist/:id", controller.GetBookDetailHandler)
}
