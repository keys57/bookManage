package mysql

import (
	"bookManage/model"
	"fmt"

	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 将变量定义在全局,是因为在其他包中需要使用
var DB *gorm.DB

func InitMysql() {
	dsn := "root:root@tcp(127.0.0.1:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(gmysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("初始化mysql失败", err)
	}

	DB = db

	if err := DB.AutoMigrate(model.User{}, model.Book{}); err != nil {
		fmt.Println("自动构建表失败", err)
	}
}
