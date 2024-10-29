package biz

import (
	"fmt"
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
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "err 1:" + err.Error(),
		})
		return
	}
	fmt.Println(form.File, form.Value, c.Request.FormValue("id"))

	// url, err := u.repo.Local_UploadFile(file, fileSize)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"result": "err 2:" + err.Error(),
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"result": "url",
	})
}
