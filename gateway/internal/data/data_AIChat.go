package data

import (
	"context"
	"gateway/api/chat"
	"gateway/internal/biz"
	"io"

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
	stream, err := client.ServerStreamAIChat(context.Background(), req)
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
