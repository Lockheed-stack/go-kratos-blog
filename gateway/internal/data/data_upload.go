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

/*
Note: if the fileName is empty string, it will be uploaded to image bucket;
else it will be updated to blog(article) bucket.
*/
func (r *gatewayUploadRepo) UploadFile(file multipart.File, fileSize int64, fileName string) (string, error) {
	putPolicy := storage.PutPolicy{}

	if fileName == "" {
		putPolicy.Scope = r.data.Qiniu_Bucket_Img
		putPolicy.FsizeLimit = 1024 * 150 // max image size:150kb
	} else {
		putPolicy.Scope = r.data.Qiniu_Bucket_Article + ":" + fileName
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

	url := ""
	if fileName == "" { // upload image
		err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
		if err != nil {
			return "", err
		}
		url = r.data.cdnImg + ret.Key
	} else { // upload markdown
		err := formUploader.Put(context.Background(), &ret, upToken, fileName, file, fileSize, &putExtra)
		if err != nil {
			return "", err
		}
		url = r.data.cdnArticle + ret.Key
	}

	// r.log.Info("hash: ", ret.Hash)

	return url, nil
}
