/**
 *@filename       imgUpload.go
 *@Description
 *@author          liyajun
 *@create          2022-12-23 18:00:42
 */

package utils

import (
	"context"
	"io"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)
var
// UpLoadFile
//
//	@Description: 图片上传  smms图床
//	@param file
//	@param fileSize
//	@return string
//	@return string
func UpLoadFile(file io.Reader, fileSize int64) (string, string) {
	putPolicy := storage.PutPolicy{
		Scope: setting.Conf.ImgUpload.Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadongZheJiang2,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return err.Error(), RECODE_IOERR
	}
	//图片访问链接
	url := setting.Conf.ImgUpload.QiNiuServer + ret.Key
	return url, RECODE_OK
}
