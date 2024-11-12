package biz

import (
	"gateway/api/users"
	"gateway/internal/middlewares"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GatewayUserRepo interface {
	GRPC_CreateUser(*users.CreateUsersRequest) (*users.CreateUsersReply, error)
	GRPC_DeleteUser(*users.DeleteUsersRequest) (*users.DeleteUsersReply, error)
	GRPC_AuthUser(*users.AuthUsersRequest) (*users.AuthUsersReply, error)
	GRPC_GetSelectedUsers(*users.GetSelectedUsersRequest) (*users.GetSelectedUsersReply, error)
	GRPC_GetUserStatisticsInfo(*users.GetStatisticsRequest) (*users.GetStatisticsReply, error)
	GRPC_UpdateUserPublicInfo(*users.UpdateUserPublicInfoRequest) (*users.UpdateUserPublicInfoReply, error)
	MaintainUserStatisticsInfo(uid uint64, ip string, pv uint32, is_newBlog bool) error
	TodayUserStatisticsInfo(uid uint64) (*users.StatisticsInfo, error)
}

type GatewayUserUsecase struct {
	repo GatewayUserRepo
}

func NewGatewayUserUsecase(repo GatewayUserRepo) *GatewayUserUsecase {
	return &GatewayUserUsecase{
		repo: repo,
	}
}

func (u *GatewayUserUsecase) CreateOneUser(c *gin.Context) {
	req := &users.CreateUsersRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}
	resp, err := u.repo.GRPC_CreateUser(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}
	if resp.Code != 200 {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": resp.Msg,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": resp.Msg,
	})
}

func (u *GatewayUserUsecase) DeleteOneUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Please check user id",
		})
		return
	}
	resp, err := u.repo.GRPC_DeleteUser(
		&users.DeleteUsersRequest{
			ID: uint64(userID),
		},
	)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": err.Error(),
		})
		return
	}
	if resp.Code != 200 {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": resp.Msg,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": resp.Msg,
	})
}

func (u *GatewayUserUsecase) AuthUser(c *gin.Context) {
	req := &users.AuthUsersRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}
	resp, err := u.repo.GRPC_AuthUser(req)
	if err != nil {
		c.JSON(http.StatusGatewayTimeout, gin.H{
			"result": err.Error(),
		})
		return
	}
	if resp.Code != 200 {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": resp.Msg,
		})
		return
	}

	token, err := middlewares.GenerateToken(req.UserName, resp.SelectedUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": resp.SelectedUser,
		"token":  token,
	})
}

func (u *GatewayUserUsecase) TokenCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
}

func (u *GatewayUserUsecase) GetPublicUsersInfo(c *gin.Context) {
	req := &users.GetSelectedUsersRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}

	resp, err := u.repo.GRPC_GetSelectedUsers(req)
	// fmt.Println(resp, err)
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
		"result": resp.SelectedUsers,
	})
}

func (u *GatewayUserUsecase) GetUserStatisticsInfo(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}

	req := &users.GetStatisticsRequest{
		ID: uint64(userID),
	}
	resp, err := u.repo.GRPC_GetUserStatisticsInfo(req)
	if err != nil {
		c.JSON(http.StatusGatewayTimeout, gin.H{
			"result": err.Error(),
		})
		return
	}
	if resp.Code != 200 {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": resp.Msg,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": resp.Info,
	})
}
func (u *GatewayUserUsecase) GetUserTodayStatistics(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("userID"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}
	resp, err := u.repo.TodayUserStatisticsInfo(uint64(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": resp,
	})
}

func (u *GatewayUserUsecase) UpdateUserPublicInfo(c *gin.Context) {
	uid := c.GetInt("request_userid")
	req := &users.UserPublicInfo{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}
	req.ID = uint64(uid)

	resp, err := u.repo.GRPC_UpdateUserPublicInfo(&users.UpdateUserPublicInfoRequest{Info: req})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	} else if resp.Code != 200 {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": resp.Msg,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": resp.Msg,
	})
}
