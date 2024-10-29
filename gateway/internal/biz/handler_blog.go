package biz

import (
	"gateway/api/articles"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type GatewayBlogRepo interface {
	GRPC_CreateOneBlog(*articles.CreateArticlesRequest) (*articles.CreateArticlesReply, error)
	GRPC_GetBlogsInSameCategory(*articles.GetArticlesInSameCategoryRequest) (*articles.GetArticlesInSameCategoryReply, error)
	GRPC_GetBlogsByCidAndUid(*articles.GetArticlesByCidAndUidRequest) (*articles.GetArticlesByCidAndUidReply, error)
	GRPC_GetBlogsForRecommend(*articles.GetRecommendArticlesRequest) (*articles.GetRecommendArticlesReply, error)
	GRPC_GetBlogsByRandom(*articles.GetRandomArticlesRequest) (*articles.GetRandomArticlesReply, error)
	GRPC_GetSingleBlog(*articles.GetSingleArticleRequest) (*articles.GetSingleArticleReply, error)
	GRPC_UpdateBlog(*articles.UpdateArticlesRequest) (*articles.UpdateArticlesReply, error)
	GRPC_DeleteBlog(*articles.DeleteArticlesRequest) (*articles.DeleteArticlesReply, error)
}

type GatewayBlogUsecase struct {
	repo        GatewayBlogRepo
	upload_repo GatewayUploadRepo
}

func NewGatewayBlogUsecase(repo GatewayBlogRepo, upload_repo GatewayUploadRepo) *GatewayBlogUsecase {
	return &GatewayBlogUsecase{
		repo:        repo,
		upload_repo: upload_repo,
	}
}

func (u *GatewayBlogUsecase) CreateOneBlog(c *gin.Context) {

	req := &articles.CreateArticlesRequest{}
	// if err := c.ShouldBindJSON(req); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"result": err.Error(),
	// 	})
	// 	return
	// }

	/* ------------ Processing blog info ---------------*/
	title := c.Request.FormValue("Title")
	cid := c.Request.FormValue("Cid")
	req.Desc = c.Request.FormValue("Desc")
	if title == "" || cid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "title or cid cannot be empty",
		})
		return
	} else {
		cid_num, err := strconv.Atoi(cid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"result": "invalid cid",
			})
			return
		}
		req.Title = title
		req.Cid = uint64(cid_num)
	}
	// blog cover
	if img := c.Request.FormValue("Img"); img != "" { // if the image is an url
		protocal := strings.Split(img, ":")[0]
		if protocal != "http" && protocal != "https" {
			c.JSON(http.StatusBadRequest, gin.H{
				"result": "Invalid Img",
			})
			return
		}
		req.Img = img
	} else { // the image is a file
		imgfile, fileheader, err := c.Request.FormFile("Img_blob")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"result": err.Error(),
			})
			return
		}
		url, err := u.upload_repo.Local_UploadFile(imgfile, fileheader.Size)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"result": err.Error(),
			})
			return
		}
		req.Img = url
	}
	// blog content
	if content := c.Request.FormValue("Content"); content != "" { // if the blog content is an url
		protocal := strings.Split(content, ":")[0]
		if protocal != "http" && protocal != "https" {
			c.JSON(http.StatusBadRequest, gin.H{
				"result": "Invalid content",
			})
			return
		}
		req.Content = content
	} else { // the blog content is a file
		blogfile, fileheader, err := c.Request.FormFile("Content_blob")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"result": err.Error(),
			})
			return
		}
		url, err := u.upload_repo.Local_UploadFile(blogfile, fileheader.Size)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"result": err.Error(),
			})
			return
		}
		req.Content = url
	}
	req.Uid = uint64(c.GetInt("request_userid"))
	/* ------------ Processing blog info ---------------*/

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
func (u *GatewayBlogUsecase) GetRecommendBlogs(c *gin.Context) {
	pageSize, err1 := strconv.Atoi(c.Query("PageSize"))
	pageNum, err2 := strconv.Atoi(c.Query("PageNum"))
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Invalid Query",
		})
		return
	}
	req := &articles.GetRecommendArticlesRequest{
		PageSize: uint32(pageSize),
		PageNum:  uint32(pageNum),
	}
	resp, err := u.repo.GRPC_GetBlogsForRecommend(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": resp.SelectedArticles,
	})
}
func (u *GatewayBlogUsecase) GetRandomBlogs(c *gin.Context) {
	str_count := c.Query("Count")
	var count int
	if str_count == "" {
		count = 6
	} else {
		tmp, err := strconv.Atoi(str_count)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"result": "Invalid Query",
			})
			return
		}
		count = tmp
	}

	req := &articles.GetRandomArticlesRequest{
		Count: uint32(count),
	}
	resp, err := u.repo.GRPC_GetBlogsByRandom(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": resp.SelectedArticles,
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

	info := &articles.DetailArticleInfo{}
	// if err := c.ShouldBindJSON(info); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"result": err.Error(),
	// 	})
	// 	return
	// }

	/*---------------- Processing blog info for update ----------------*/
	title := c.Request.FormValue("Title")
	cid := c.Request.FormValue("Cid")
	info.Desc = c.Request.FormValue("Desc")
	if title == "" || cid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "title or cid cannot be empty",
		})
		return
	} else {
		cid_num, err := strconv.Atoi(cid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"result": "invalid cid",
			})
			return
		}
		info.Title = title
		info.Cid = uint64(cid_num)
	}
	// blog cover
	if img := c.Request.FormValue("Img"); img != "" { // if the image is an url
		protocal := strings.Split(img, ":")[0]
		if protocal != "http" && protocal != "https" {
			c.JSON(http.StatusBadRequest, gin.H{
				"result": "Invalid Img",
			})
			return
		}
		info.Img = img
	} else { // the image is a file
		imgfile, fileheader, err := c.Request.FormFile("Img_blob")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"result": err.Error(),
			})
			return
		}
		url, err := u.upload_repo.Local_UploadFile(imgfile, fileheader.Size)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"result": err.Error(),
			})
			return
		}
		info.Img = url
	}
	// blog content
	if content := c.Request.FormValue("Content"); content != "" { // if the blog content is an url
		protocal := strings.Split(content, ":")[0]
		if protocal != "http" && protocal != "https" {
			c.JSON(http.StatusBadRequest, gin.H{
				"result": "Invalid content",
			})
			return
		}
		info.Content = content
	} else { // the blog content is a file
		blogfile, fileheader, err := c.Request.FormFile("Content_blob")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"result": err.Error(),
			})
			return
		}
		url, err := u.upload_repo.Local_UploadFile(blogfile, fileheader.Size)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"result": err.Error(),
			})
			return
		}
		info.Content = url
	}

	id_num, err := strconv.Atoi(c.Request.FormValue("ID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}
	info.ID = uint64(id_num)

	info.Uid = uint64(c.GetInt("request_userid"))
	/*---------------- Processing blog info for update ----------------*/

	req := &articles.UpdateArticlesRequest{
		ArticleInfo: info,
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
