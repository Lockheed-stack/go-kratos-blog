package biz

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GatewayUploadRepo interface {
	Local_UploadFile(multipart.File, int64) (string, error)
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
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": err.Error(),
		})
		return
	}

	fileSize := fileHeader.Size

	url, err := u.repo.Local_UploadFile(file, fileSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": url,
	})
}
