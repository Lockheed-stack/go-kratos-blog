package data

import (
	"context"
	"gateway/internal/biz"
	"mime/multipart"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type gatewayUploadRepo struct {
	data *Data
	log  *log.Helper
}

func NewGatewayUploadRepo(data *Data, logger log.Logger) biz.GatewayUploadRepo {
	return &gatewayUploadRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *gatewayUploadRepo) Local_UploadFile(file multipart.File, fileSize int64) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: r.data.Qiniu_Bucket,
	}
	mac := qbox.NewMac(r.data.Qiniu_AccessKey, r.data.Qiniu_SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:     &storage.ZoneXinjiapo,
		UseHTTPS: false,
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", err
	}
	url := r.data.WebHost + ret.Key
	return url, nil
}
