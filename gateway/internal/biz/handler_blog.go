package biz

import (
	"gateway/api/articles"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GatewayBlogRepo interface {
	GRPC_CreateOneBlog(*articles.CreateArticlesRequest) (*articles.CreateArticlesReply, error)
	GRPC_GetBlogsInSameCategory(*articles.GetArticlesInSameCategoryRequest) (*articles.GetArticlesInSameCategoryReply, error)
	GRPC_GetBlogsByCidAndUid(*articles.GetArticlesByCidAndUidRequest) (*articles.GetArticlesByCidAndUidReply, error)
	GRPC_GetSingleBlog(*articles.GetSingleArticleRequest) (*articles.GetSingleArticleReply, error)
	GRPC_UpdateBlog(*articles.UpdateArticlesRequest) (*articles.UpdateArticlesReply, error)
	GRPC_DeleteBlog(*articles.DeleteArticlesRequest) (*articles.DeleteArticlesReply, error)
}

type GatewayBlogUsecase struct {
	repo GatewayBlogRepo
}

func NewGatewayBlogUsecase(repo GatewayBlogRepo) *GatewayBlogUsecase {
	return &GatewayBlogUsecase{
		repo: repo,
	}
}

func (u *GatewayBlogUsecase) CreateOneBlog(c *gin.Context) {

	req := &articles.CreateArticlesRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}
	req.Uid = uint64(c.GetInt("request_userid"))

	resp, err := u.repo.GRPC_CreateOneBlog(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": resp.Msg,
	})
}

func (u *GatewayBlogUsecase) GetBlogsInSameCategory(c *gin.Context) {

	cid, isExist := c.Params.Get("cid")
	if !isExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Invalid Query",
		})
		return
	}

	pageSize, err1 := strconv.Atoi(c.Query("PageSize"))
	pageNum, err2 := strconv.Atoi(c.Query("PageNum"))
	cid_int, err3 := strconv.Atoi(cid)
	if err1 != nil || err2 != nil || err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Invalid Query",
		})
		return
	}
	req := &articles.GetArticlesInSameCategoryRequest{
		PageSize: uint32(pageSize),
		PageNum:  uint32(pageNum),
		CID:      uint64(cid_int),
	}
	resp, err := u.repo.GRPC_GetBlogsInSameCategory(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": resp.SelectedArticles,
		"total":  resp.Total,
	})
}
func (u *GatewayBlogUsecase) GetBlogsByCidAndUid(c *gin.Context) {
	cid, isExist_1 := c.Params.Get("cid")
	uid, isExist_2 := c.Params.Get("uid")
	if !isExist_1 || !isExist_2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Invalid Query",
		})
		return
	}
	pageSize, err1 := strconv.Atoi(c.Query("PageSize"))
	pageNum, err2 := strconv.Atoi(c.Query("PageNum"))
	cid_int, err3 := strconv.Atoi(cid)
	uid_int, err4 := strconv.Atoi(uid)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Invalid Query",
		})
		return
	}
	req := &articles.GetArticlesByCidAndUidRequest{
		PageSize: uint32(pageSize),
		PageNum:  uint32(pageNum),
		CID:      uint64(cid_int),
		UID:      uint64(uid_int),
	}
	resp, err := u.repo.GRPC_GetBlogsByCidAndUid(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": resp.SelectedArticles,
		"total":  resp.Total,
	})
}
func (u *GatewayBlogUsecase) GetOneBlog(c *gin.Context) {
	blogid, err := strconv.Atoi(c.Param("blogID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Please check blog id",
		})
		return
	}

	resp, err := u.repo.GRPC_GetSingleBlog(&articles.GetSingleArticleRequest{
		ArticleID: uint64(blogid),
	})

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": "Please check blog id",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": resp.Article,
	})
}

func (u *GatewayBlogUsecase) UpdateOneBlog(c *gin.Context) {

	req := &articles.UpdateArticlesRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}

	resp, err := u.repo.GRPC_UpdateBlog(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": resp.Msg,
	})
}

func (u *GatewayBlogUsecase) DeleteOneBlog(c *gin.Context) {
	blogid, err := strconv.Atoi(c.Query("blogID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Please check blog id",
		})
		return
	}
	resp, err := u.repo.GRPC_DeleteBlog(&articles.DeleteArticlesRequest{
		ArticleID: uint64(blogid),
	})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": "Please check blog id",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": resp.Msg,
	})
}
