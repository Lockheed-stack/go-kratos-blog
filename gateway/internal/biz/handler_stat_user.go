package biz

import (
	"gateway/api/stat_user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GatewayStatUserRepo interface {
	GRPC_GetUserSevenDayStat(*stat_user.GetUserSevenDaysStatRequest) (*stat_user.GetUserSevenDaysStatReply, error)
	GRPC_SetUserTodayStatData(*stat_user.SetUserStatInfoRequest) (*stat_user.SetUserStatInfoReply, error)
}

type GatewayStatUserUsecase struct {
	repo GatewayStatUserRepo
}

func NewGatewayStatUserUsecase(repo GatewayStatUserRepo) *GatewayStatUserUsecase {
	return &GatewayStatUserUsecase{
		repo: repo,
	}
}

func (u *GatewayStatUserUsecase) GetUserSevenDaysStat(c *gin.Context) {

	userID, err := strconv.Atoi(c.Query("userID"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}

	req := &stat_user.GetUserSevenDaysStatRequest{
		Uid: uint64(userID),
	}

	resp, err := u.repo.GRPC_GetUserSevenDayStat(req)
	if err != nil {
		c.JSON(http.StatusGatewayTimeout, gin.H{
			"result": err.Error(),
		})
		return
	}
	if resp.Code != 200 {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": resp.SevenDaysData,
	})
}
