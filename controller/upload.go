package controller

import (
	"bookManage/conf"
	"context"
	"mime/multipart"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

var AccessKey = conf.AccessKey
var SecretKey = conf.SecretKey
var Bucket = conf.Bucket
var ImgUrl = conf.QiniuServer

func UpLoadFile(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)
	region, _ := storage.GetRegion(AccessKey, Bucket)
	cfg := storage.Config{
		Region:        region,
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: true,
		UseHTTPS:      true,
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return err.Error(), 200
	}
	url := ImgUrl + ret.Key
	return url, 2
}
