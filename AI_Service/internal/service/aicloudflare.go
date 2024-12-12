package service

import (
	pb "AI_Service/api/AI_Cloudflare"
	"AI_Service/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/errors"
)

type AICloudflareService struct {
	pb.UnimplementedAICloudflareServer
	uc *biz.AICloudflareUsecase
}

func NewAICloudflareService(uc *biz.AICloudflareUsecase) *AICloudflareService {
	return &AICloudflareService{
		uc: uc,
	}
}

func (s *AICloudflareService) StreamAISummarization(req *pb.AISummarizationRequest, conn pb.AICloudflare_StreamAISummarizationServer) error {

	ch := make(chan *biz.CloudflareAITextGenerationReply)
	defer close(ch)
	go s.uc.StreamGetAISummarization(req.ArticleText, ch)

	for resp := range ch {
		if resp != nil {
			err := conn.Send(&pb.AISummarizationReply{
				TextAbstract: resp.Response,
			})
			if err != nil {
				return err
			}
		} else {
			return nil
		}
	}
	return nil
}

func (s *AICloudflareService) StreamAIChat(req *pb.AIChatRequest, conn pb.AICloudflare_StreamAIChatServer) error {
	ch := make(chan *biz.CloudflareAITextGenerationReply)
	defer close(ch)
	go s.uc.StreamGetAIChat(req, ch)

	for resp := range ch {
		if resp != nil {
			err := conn.Send(&pb.AIChatReply{
				Msg: resp.Response,
			})
			if err != nil {
				return err
			}
		} else {
			return nil
		}
	}

	return nil
}

func (s *AICloudflareService) AIPaint(ctx context.Context, req *pb.AIPaintRequest) (*pb.AIPaintReply, error) {
	resp := &pb.AIPaintReply{}
	bin, err := s.uc.AIPaint(req)
	if err != nil {
		kratos_err := err.(*errors.Error)
		resp.Msg = kratos_err.Message + ": " + kratos_err.Reason
	} else {
		resp.ImgBinary = bin
		resp.Msg = "OK"
	}
	return resp, nil
}
