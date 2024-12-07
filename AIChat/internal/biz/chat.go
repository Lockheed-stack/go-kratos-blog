package biz

import (
	pb "AIChat/api/chat"
	"context"

	"github.com/go-kratos/kratos/v2/errors"
)

type AIChatRepo interface {
	CloudflareStreamGetAIChat(ctx context.Context, messages *AIChatRequest, ch chan *AIChatRespond, modelKind string) error
	CloudflareGetAIPaintImg(messages *AIPaintingRequest, modelKind string) ([]byte, error)
}

type AIChatUsecase struct {
	repo AIChatRepo
}

type AIChatRequestContent struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type AIChatRequest struct {
	Messages    []AIChatRequestContent `json:"messages"`
	Stream      bool                   `json:"stream"`
	Temperature float32                `json:"temperature"`
	TopK        uint32                 `json:"top_k"`
}
type AIChatRespond struct {
	Response string `json:"response"`
}
type AIPaintingRequest struct {
	Prompt   string  `json:"prompt"`
	Height   uint32  `json:"height"`
	Width    uint32  `json:"width"`
	Guidance float32 `json:"guidance"` // Controls how closely the generated image should adhere to the prompt; higher values make the image more aligned with the prompt
}
type AIPaintingRespond struct {
	Bin []byte
}

func NewAIChatUsecase(repo AIChatRepo) *AIChatUsecase {
	return &AIChatUsecase{
		repo: repo,
	}
}

func (uc *AIChatUsecase) StreamGetAIChatRespond(promptAndSetting *pb.AIChatRequest, ch chan *AIChatRespond) {

	// Checking parameter validity
	if promptAndSetting.Temperature < 0 || promptAndSetting.Temperature > 5 {
		ch <- &AIChatRespond{
			Response: "[ERROR]: The parameter 'temperature' out of range",
		}
		ch <- nil
		return
	}
	if promptAndSetting.TopK < 1 || promptAndSetting.TopK > 50 {
		ch <- &AIChatRespond{
			Response: "[ERROR]: The parameter 'top_k' out of range",
		}
		ch <- nil
		return
	}

	// assembling the request structure
	MsgContents := make([]AIChatRequestContent, 1)
	MsgContents[0].Role = promptAndSetting.Msg.Role
	MsgContents[0].Content = promptAndSetting.Msg.Content
	req := &AIChatRequest{
		Messages:    MsgContents,
		Stream:      true,
		Temperature: promptAndSetting.Temperature,
		TopK:        promptAndSetting.TopK,
	}
	err := uc.repo.CloudflareStreamGetAIChat(context.Background(), req, ch, promptAndSetting.ModelKind)
	if err != nil {
		ch <- &AIChatRespond{
			Response: err.Error(),
		}
	}
	ch <- nil
}

func (uc *AIChatUsecase) CloudflareAIPaintingRespond(promptAndSetting *pb.AIPaintRequest) ([]byte, error) {
	// checking parameter validity
	if promptAndSetting.Height < 256 || promptAndSetting.Height > 2048 {
		return nil, errors.New(400, "height", "Invalid AI Painting parameters")
	}
	if promptAndSetting.Width < 256 || promptAndSetting.Width > 2048 {
		return nil, errors.New(400, "width", "Invalid AI Painting parameters")
	}
	if promptAndSetting.Guidance < 5 || promptAndSetting.Guidance > 20 {
		return nil, errors.New(400, "guidance", "Invalid AI Painting parameters")
	}
	// assembling the request struct
	req := &AIPaintingRequest{
		Prompt:   promptAndSetting.Prompt,
		Height:   promptAndSetting.Height,
		Width:    promptAndSetting.Width,
		Guidance: promptAndSetting.Guidance,
	}
	bin, err := uc.repo.CloudflareGetAIPaintImg(req, promptAndSetting.ModelKind)
	if err != nil {
		return nil, errors.New(400, "Please adjust parameters", "Failed to generate images")
	}
	return bin, nil
}
