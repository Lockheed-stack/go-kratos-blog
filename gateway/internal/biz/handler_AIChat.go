package biz

import (
	"gateway/api/chat"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GatewayAIChatRepo interface {
	GRPC_AIChatStreamGetResponse(req *chat.AIChatRequest, ch chan *chat.AIChatReply)
	GRPC_AIPainting(req *chat.AIPaintRequest) (*chat.AIPaintReply, error)
}

type GatewayAIChatUsecase struct {
	repo GatewayAIChatRepo
}

func NewGatewayAIChatUsecase(repo GatewayAIChatRepo) *GatewayAIChatUsecase {
	return &GatewayAIChatUsecase{
		repo: repo,
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
	go u.repo.GRPC_AIChatStreamGetResponse(req, ch)
	for {
		select {
		case resp := <-ch:
			{
				if resp != nil {
					c.Writer.WriteString(resp.Msg)
					c.Writer.Flush()
				} else {
					return
				}
			}
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

	resp, err := u.repo.GRPC_AIPainting(req)
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
