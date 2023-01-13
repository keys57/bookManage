package main

import (
	"bookManage/dao/mysql"
	"bookManage/router"
)

func main() {
	//初始化mysql
	mysql.InitMysql()
	//1.初始化gin服务
	r := router.InitRouter()
	//3.启动服务
	r.Run(":8000")
}
