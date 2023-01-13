package controller

import (
	"bookManage/dao/mysql"
	"bookManage/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 新增图书
func CreateBookHandler(c *gin.Context) {
	p := new(model.Book)
	if err := c.ShouldBind(p); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	mysql.DB.Create(p)
	c.JSON(200, gin.H{"msg": "success"})

}

// 查看数据列表
func GetBookListHandler(c *gin.Context) {
	books := []model.Book{}
	mysql.DB.Find(&books)
	c.JSON(200, gin.H{"books": books})
}

// 查看指定书籍
func GetBookDetailHandler(c *gin.Context) {
	bookIdStr := c.Param("id")
	bookId, _ := strconv.ParseInt(bookIdStr, 10, 64)
	book := model.Book{Id: bookId}
	mysql.DB.Find(&book)
	c.JSON(200, gin.H{"book": book})

}

// 修改数据
func UpdateBookhandler(c *gin.Context) {
	p := new(model.Book)
	if err := c.ShouldBind(p); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
}
