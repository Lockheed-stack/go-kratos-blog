package data

import (
	"bytes"
	"context"
	"gateway/api/AI_Cloudflare"
	"gateway/internal/biz"
	"io"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type gatewayAIRepo struct {
	data *Data
	log  *log.Helper
}

func NewGatewayAIRepo(data *Data, logger log.Logger) biz.GatewayAIRepo {
	return &gatewayAIRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// GRPC_Stream_AISummarization implements biz.GatewayAIRepo.
func (r *gatewayAIRepo) GRPC_Stream_AISummarization(http_req_body io.ReadCloser, blogID_key string, ch chan *AI_Cloudflare.AISummarizationReply) {

	// fast path
	val, err := GetAISummarizationRedis(r.data.Redis_cli, blogID_key)
	if err == nil {
		ch <- &AI_Cloudflare.AISummarizationReply{
			TextAbstract: val,
		}
		ch <- nil
		return
	}

	// slow path
	grpc_req := &AI_Cloudflare.AISummarizationRequest{}
	var b = make([]byte, 0, 4096)
	buf := bytes.NewBuffer(b)
	buf.ReadFrom(http_req_body)
	grpc_req.ArticleText = buf.Bytes()
	defer http_req_body.Close()

	client := AI_Cloudflare.NewAICloudflareClient(r.data.ConnGRPC_ai)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	stream, err := client.StreamAISummarization(ctx, grpc_req)
	if err != nil {
		r.log.Error(err)
		ch <- &AI_Cloudflare.AISummarizationReply{
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
				ch <- &AI_Cloudflare.AISummarizationReply{
					TextAbstract: err.Error(),
				}
			} else {
				// store summarization into redis
				go SetAISummarizationRedis(r.data.Redis_cli, blogID_key, builder.String())
			}
			ch <- nil
			return
		}
		ch <- reply
		builder.WriteString(reply.TextAbstract)
	}
}

// GRPC_AIPainting implements biz.GatewayAIRepo.
func (r *gatewayAIRepo) GRPC_AIPainting(req *AI_Cloudflare.AIPaintRequest) (*AI_Cloudflare.AIPaintReply, error) {
	client := AI_Cloudflare.NewAICloudflareClient(r.data.ConnGRPC_ai)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	resp, err := client.AIPaint(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GRPC_Stream_AIChat implements biz.GatewayAIRepo.
func (r *gatewayAIRepo) GRPC_Stream_AIChat(req *AI_Cloudflare.AIChatRequest, ch chan *AI_Cloudflare.AIChatReply) {
	client := AI_Cloudflare.NewAICloudflareClient(r.data.ConnGRPC_ai)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	stream, err := client.StreamAIChat(ctx, req)
	if err != nil {
		r.log.Error(err)
		ch <- &AI_Cloudflare.AIChatReply{
			Msg: err.Error(),
		}
		ch <- nil
		return
	}

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err != io.EOF {
				ch <- &AI_Cloudflare.AIChatReply{
					Msg: err.Error(),
				}
			}
			ch <- nil
			return
		}
		ch <- reply
	}
}
