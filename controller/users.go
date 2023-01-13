package controller

import (
	"bookManage/dao/mysql"
	"bookManage/model"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

// 用户注册
func RegisterHandler(c *gin.Context) {
	p := new(model.User)
	//参数校验 绑定
	if err := c.ShouldBind(p); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	//写入数据库
	mysql.DB.Create(p)
	c.JSON(200, gin.H{"msg": "success"})
}

// 登陆
func LoginHandler(c *gin.Context) {
	p := new(model.User)
	if err := c.ShouldBind(p); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	//判断用户名密码是否正确
	u := model.User{UserName: p.UserName, Password: p.Password}
	if rows := mysql.DB.Where(&u).First(&u).Row(); rows == nil {
		c.JSON(403, gin.H{"msg": "用户名或密码错误"})
		return
	}
	//随机生成字符串作为token
	token := uuid.New().String()
	mysql.DB.Model(u).Update("token", token)
	c.JSON(200, gin.H{"msg": "success", "token": token})
}
