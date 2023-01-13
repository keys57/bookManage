package router

import (
	"bookManage/cors"

	"github.com/gin-gonic/gin"
)

/*
加载其他路由文件中的路由
*/

// 初始化其他文件中的路由
func InitRouter() *gin.Engine {
	//1.初始化gin服务
	r := gin.Default()
	r.Use(cors.Cors())
	LoadAPIRouter(r)
	return r
}
