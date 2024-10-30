package biz

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GatewayUploadRepo interface {
	UploadFile(file multipart.File, filesize int64, filename string) (string, error)
}

type GatewayUploadUsecase struct {
	repo GatewayUploadRepo
}

func NewGatewayUploadUsecase(repo GatewayUploadRepo) *GatewayUploadUsecase {
	return &GatewayUploadUsecase{
		repo: repo,
	}
}

func (u *GatewayUploadUsecase) Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "err 1:" + err.Error(),
		})
		return
	}
	name := c.Request.FormValue("title")
	url, err := u.repo.UploadFile(file, header.Size, name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "err 2:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": url,
	})
}
