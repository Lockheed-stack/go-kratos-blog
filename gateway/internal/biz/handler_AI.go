package biz

import (
	"gateway/api/AI_Cloudflare"
	"io"

	"github.com/gin-gonic/gin"
)

type GatewayAIRepo interface {
	GRPC_Stream_AISummarization(http_req_body io.ReadCloser, blogID_key string, ch chan *AI_Cloudflare.AISummarizationReply)
}

type GatewayAIUsecase struct {
	ai_repo GatewayAIRepo
}

func NewGatewayAIUsecase(ai_repo GatewayAIRepo) *GatewayAIUsecase {
	return &GatewayAIUsecase{
		ai_repo: ai_repo,
	}
}

func (u *GatewayAIUsecase) GetAISummarization(c *gin.Context) {

	// AI summarization
	ch := make(chan *AI_Cloudflare.AISummarizationReply)
	defer close(ch)
	c.Header("Content-Type", "text/event-stream")

	go u.ai_repo.GRPC_Stream_AISummarization(c.Request.Body, "", ch)

	for resp := range ch {
		if resp != nil {
			c.Writer.WriteString(resp.TextAbstract)
			c.Writer.Flush()
		} else {
			return
		}
	}
}
