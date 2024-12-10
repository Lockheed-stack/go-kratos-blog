package biz

import (
	"gateway/api/articles"
	"gateway/api/chat"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GatewayAIChatRepo interface {
	GRPC_AIChatStreamGetResponse(req *chat.AIChatRequest, ch chan *chat.AIChatReply)
	GRPC_AIPainting(req *chat.AIPaintRequest) (*chat.AIPaintReply, error)
	GRPC_AISummarizationStreamGetResponse(req_body io.ReadCloser, blogID_key string, ch chan *chat.AISummarizationReply)
}

type GatewayAIChatUsecase struct {
	ai_repo   GatewayAIChatRepo
	blog_repo GatewayBlogRepo
}

func NewGatewayAIChatUsecase(repo GatewayAIChatRepo, blog_repo GatewayBlogRepo) *GatewayAIChatUsecase {
	return &GatewayAIChatUsecase{
		ai_repo:   repo,
		blog_repo: blog_repo,
	}
}

func (u *GatewayAIChatUsecase) AIChatStreamGetResponse(c *gin.Context) {
	req := &chat.AIChatRequest{}

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}

	ch := make(chan *chat.AIChatReply)
	defer close(ch)

	c.Header("Content-Type", "text/event-stream")
	go u.ai_repo.GRPC_AIChatStreamGetResponse(req, ch)
	for resp := range ch {
		if resp != nil {
			c.Writer.WriteString(resp.Msg)
			c.Writer.Flush()
		} else {
			return
		}
	}

}

func (u *GatewayAIChatUsecase) AIPainting(c *gin.Context) {

	req := &chat.AIPaintRequest{}
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

func (u *GatewayAIChatUsecase) AISummarizationStreamGetResponse(c *gin.Context) {

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
	ch := make(chan *chat.AISummarizationReply)
	defer close(ch)
	c.Header("Content-Type", "text/event-stream")

	go u.ai_repo.GRPC_AISummarizationStreamGetResponse(c.Request.Body, id_str, ch)

	for resp := range ch {
		if resp != nil {
			c.Writer.WriteString(resp.TextAbstract)
			c.Writer.Flush()
		} else {
			return
		}
	}
}
