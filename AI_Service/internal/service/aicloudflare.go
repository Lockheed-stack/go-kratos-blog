package service

import (
	pb "AI_Service/api/AI_Cloudflare"
	"AI_Service/internal/biz"
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
