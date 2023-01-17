package controller

import (
	"bookManage/conf"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpLoad(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	url, code := UpLoadFile(file, fileSize)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": 200,
		"url":     url,
		"AK":      conf.AccessKey,
		"SK":      conf.SecretKey,
		"b":       conf.Bucket,
	})
}
