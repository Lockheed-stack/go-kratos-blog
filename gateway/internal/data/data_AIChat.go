package data

import (
	"bytes"
	"context"
	"gateway/api/chat"
	"gateway/internal/biz"
	"io"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type gatewayAIChatRepo struct {
	data *Data
	log  *log.Helper
}

func NewGatewayAIChatRepo(data *Data, logger log.Logger) biz.GatewayAIChatRepo {
	return &gatewayAIChatRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *gatewayAIChatRepo) GRPC_AIChatStreamGetResponse(req *chat.AIChatRequest, ch chan *chat.AIChatReply) {

	client := chat.NewChatClient(r.data.ConnGRPC_ai_chat)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	stream, err := client.ServerStreamAIChat(ctx, req)
	if err != nil {
		r.log.Error(err)
		ch <- &chat.AIChatReply{
			Msg: err.Error(),
		}
		ch <- nil
		return
	}

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err != io.EOF {
				ch <- &chat.AIChatReply{
					Msg: err.Error(),
				}
			}
			ch <- nil
			return
		}
		ch <- reply
	}
}

func (r *gatewayAIChatRepo) GRPC_AIPainting(req *chat.AIPaintRequest) (*chat.AIPaintReply, error) {
	client := chat.NewChatClient(r.data.ConnGRPC_ai_chat)
	resp, err := client.AIPaint(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (r *gatewayAIChatRepo) GRPC_AISummarizationStreamGetResponse(req_body io.ReadCloser, blogID string, ch chan *chat.AISummarizationReply) {

	// fast path
	val, err := GetAISummarizationRedis(r.data.Redis_cli, blogID)
	if err == nil {
		ch <- val
		ch <- nil
		return
	}

	// slow path
	req := &chat.AISummarizationRequest{}
	var b = make([]byte, 0, 4096)
	buf := bytes.NewBuffer(b)
	buf.ReadFrom(req_body)
	req.ArticleText = buf.Bytes()
	defer req_body.Close()

	client := chat.NewChatClient(r.data.ConnGRPC_ai_chat)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	stream, err := client.AISummarization(ctx, req)
	if err != nil {
		r.log.Error(err)
		ch <- &chat.AISummarizationReply{
			TextAbstract: err.Error(),
		}
		ch <- nil
		return
	}

	var builder strings.Builder
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err != io.EOF {
				ch <- &chat.AISummarizationReply{
					TextAbstract: err.Error(),
				}
			}
			ch <- nil
			// store summarization into redis
			go SetAISummarizationRedis(r.data.Redis_cli, blogID, builder.String())
			return
		}
		ch <- reply
		builder.WriteString(reply.TextAbstract)
	}
}
