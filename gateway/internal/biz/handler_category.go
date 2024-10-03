package biz

import (
	"gateway/api/category"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GatewayCategoryRepo interface {
	GRPC_CreateOneCategory(*category.CreateCategoryRequest) (*category.CreateCategoryReply, error)
	GRPC_ListCategory(*category.ListCategoryRequest) (*category.ListCategoryReply, error)
	GRPC_DeleteOneCategory(*category.DeleteCategoryRequest) (*category.DeleteCategoryReply, error)
	GRPC_UpdateOneCategory(*category.UpdateCategoryRequest) (*category.UpdateCategoryReply, error)
}

type GatewayCategoryUsecase struct {
	repo GatewayCategoryRepo
}

func NewGatewayCategoryUsecase(repo GatewayCategoryRepo) *GatewayCategoryUsecase {
	return &GatewayCategoryUsecase{
		repo: repo,
	}
}

func (u *GatewayCategoryUsecase) CreateOneCategory(c *gin.Context) {
	req := &category.CreateCategoryRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}
	resp, err := u.repo.GRPC_CreateOneCategory(req)
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

func (u *GatewayCategoryUsecase) GetCategoryList(c *gin.Context) {
	pageSize, err1 := strconv.Atoi(c.Query("PageSize"))
	pageNum, err2 := strconv.Atoi(c.Query("PageNum"))

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Invalid Query",
		})
		return
	}
	req := &category.ListCategoryRequest{
		PageSize: uint32(pageSize),
		PageNum:  uint32(pageNum),
	}
	resp, err := u.repo.GRPC_ListCategory(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": resp.CategoryArray,
	})
}

func (u *GatewayCategoryUsecase) UpdateCategory(c *gin.Context) {
	req := &category.UpdateCategoryRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}

	resp, err := u.repo.GRPC_UpdateOneCategory(req)
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

func (u *GatewayCategoryUsecase) DeleteCategory(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Query("categoryID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Please check category id",
		})
		return
	}
	resp, err := u.repo.GRPC_DeleteOneCategory(&category.DeleteCategoryRequest{
		ID: uint64(categoryID),
	})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": "Please check category id",
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
