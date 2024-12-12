package biz

import (
	"gateway/api/AI_Cloudflare"
	"gateway/api/articles"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GatewayAIRepo interface {
	GRPC_Stream_AISummarization(http_req_body io.ReadCloser, blogID_key string, ch chan *AI_Cloudflare.AISummarizationReply)
	GRPC_Stream_AIChat(req *AI_Cloudflare.AIChatRequest, ch chan *AI_Cloudflare.AIChatReply)
	GRPC_AIPainting(req *AI_Cloudflare.AIPaintRequest) (*AI_Cloudflare.AIPaintReply, error)
}

type GatewayAIUsecase struct {
	ai_repo   GatewayAIRepo
	blog_repo GatewayBlogRepo
}

func NewGatewayAIUsecase(ai_repo GatewayAIRepo, blog_repo GatewayBlogRepo) *GatewayAIUsecase {
	return &GatewayAIUsecase{
		ai_repo:   ai_repo,
		blog_repo: blog_repo,
	}
}

func (u *GatewayAIUsecase) GetAISummarization(c *gin.Context) {

	// check blog id
	id_str := c.Query("blogID")
	if id_str == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "blog ID required",
		})
		return
	}

	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Invalid request",
		})
		return
	}
	result, _ := u.blog_repo.GRPC_CheckBlogIsExistence(&articles.CheckExistenceOfBlogRequest{ArticleID: uint32(id)})
	if !result.Existence {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Invalid request",
		})
		return
	}

	// AI summarization
	ch := make(chan *AI_Cloudflare.AISummarizationReply)
	defer close(ch)
	c.Header("Content-Type", "text/event-stream")

	go u.ai_repo.GRPC_Stream_AISummarization(c.Request.Body, id_str, ch)

	for resp := range ch {
		if resp != nil {
			c.Writer.WriteString(resp.TextAbstract)
			c.Writer.Flush()
		} else {
			return
		}
	}
}

func (u *GatewayAIUsecase) AIChat(c *gin.Context) {
	req := &AI_Cloudflare.AIChatRequest{}

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}

	ch := make(chan *AI_Cloudflare.AIChatReply)
	defer close(ch)

	c.Header("Content-Type", "text/event-stream")
	go u.ai_repo.GRPC_Stream_AIChat(req, ch)
	for resp := range ch {
		if resp != nil {
			c.Writer.WriteString(resp.Msg)
			c.Writer.Flush()
		} else {
			return
		}
	}
}

func (u *GatewayAIUsecase) GetAIPainting(c *gin.Context) {

	req := &AI_Cloudflare.AIPaintRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}

	resp, err := u.ai_repo.GRPC_AIPainting(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": err.Error(),
		})
		return
	}
	if resp.Msg != "OK" {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": resp.Msg,
		})
		return
	}

	c.Header("Content-Type", "image/png")
	c.Writer.Write(resp.ImgBinary)
	c.Writer.Flush()
}
