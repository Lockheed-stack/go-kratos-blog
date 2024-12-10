package service

import (
	pb "AIChat/api/chat"
	"AIChat/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/errors"
)

type ChatService struct {
	pb.UnimplementedChatServer
	uc *biz.AIChatUsecase
}

func NewChatService(uc *biz.AIChatUsecase) *ChatService {
	return &ChatService{
		uc: uc,
	}
}

func (s *ChatService) ServerStreamAIChat(req *pb.AIChatRequest, conn pb.Chat_ServerStreamAIChatServer) error {
	ch := make(chan *biz.AIChatRespond)
	defer close(ch)
	go s.uc.StreamGetAIChatRespond(req, ch)

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

func (s *ChatService) AIPaint(ctx context.Context, req *pb.AIPaintRequest) (*pb.AIPaintReply, error) {
	resp := &pb.AIPaintReply{}
	bin, err := s.uc.CloudflareAIPaintingRespond(req)
	if err != nil {
		kratos_err := err.(*errors.Error)
		resp.Msg = kratos_err.Message + ": " + kratos_err.Reason
	} else {
		resp.ImgBinary = bin
		resp.Msg = "OK"
	}
	return resp, nil
}

func (s *ChatService) AISummarization(req *pb.AISummarizationRequest, conn pb.Chat_AISummarizationServer) error {

	ch := make(chan *biz.AIChatRespond)
	defer close(ch)
	go s.uc.StreamGetAISummarization(req, ch)

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
