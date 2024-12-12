package biz

import (
	pb "AI_Service/api/AI_Cloudflare"
	"context"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
)

type AICloudflareRepo interface {
	StreamTextGeneration(ctx context.Context, messages *CloudflareAITextGenerationRequest, ch chan *CloudflareAITextGenerationReply, model string) error
	ImageGeneration(messages *CloudflareAIPaintingRequest, model string) ([]byte, error)
}
type AICloudflareUsecase struct {
	repo AICloudflareRepo
}

func NewAICloudflareUsecase(repo AICloudflareRepo) *AICloudflareUsecase {
	return &AICloudflareUsecase{
		repo: repo,
	}
}

type CloudflareAITextGenerationMessages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type CloudflareAITextGenerationRequest struct {
	Messages    []CloudflareAITextGenerationMessages `json:"messages"`
	Stream      bool                                 `json:"stream"`
	Temperature float32                              `json:"temperature,omitempty"`
	TopK        uint32                               `json:"top_k,omitempty"`
	MaxTokens   uint32                               `json:"max_tokens,omitempty"` // The maximum number of tokens to generate in the response.
}
type CloudflareAITextGenerationReply struct {
	Response string `json:"response"`
}
type CloudflareAIPaintingRequest struct {
	Prompt   string  `json:"prompt"`
	Height   uint32  `json:"height"`
	Width    uint32  `json:"width"`
	Guidance float32 `json:"guidance"` // Controls how closely the generated image should adhere to the prompt; higher values make the image more aligned with the prompt
}
type CloudflareAIPaintingRespond struct {
	Bin []byte
}

func (uc *AICloudflareUsecase) StreamGetAISummarization(text_bytes []byte, ch chan *CloudflareAITextGenerationReply) {
	var builder strings.Builder
	builder.WriteString("对这篇 markdown 格式的文章进行简要的摘要：\n")
	builder.Write(text_bytes)
	MsgContents := make([]CloudflareAITextGenerationMessages, 2)
	MsgContents[0].Role = "system"
	MsgContents[0].Content = "your are a helpful assistant"
	MsgContents[1].Role = "user"
	MsgContents[1].Content = builder.String()

	req := &CloudflareAITextGenerationRequest{
		Messages:    MsgContents,
		Stream:      true,
		Temperature: 0.6,
		TopK:        30,
		MaxTokens:   2048,
	}

	err := uc.repo.StreamTextGeneration(context.Background(), req, ch, "llama-3.3-70b-instruct-fp8-fast")
	if err != nil {
		ch <- &CloudflareAITextGenerationReply{
			Response: err.Error(),
		}
	}
	ch <- nil
}

func (uc *AICloudflareUsecase) StreamGetAIChat(promptAndSetting *pb.AIChatRequest, ch chan *CloudflareAITextGenerationReply) {

	// Checking parameter validity
	if promptAndSetting.Temperature < 0 || promptAndSetting.Temperature > 5 {
		ch <- &CloudflareAITextGenerationReply{
			Response: "[ERROR]: The parameter 'temperature' out of range",
		}
		ch <- nil
		return
	}
	if promptAndSetting.TopK < 1 || promptAndSetting.TopK > 50 {
		ch <- &CloudflareAITextGenerationReply{
			Response: "[ERROR]: The parameter 'top_k' out of range",
		}
		ch <- nil
		return
	}

	// assembling the request structure
	MsgContents := make([]CloudflareAITextGenerationMessages, len(promptAndSetting.Msg))
	for i, v := range promptAndSetting.Msg {
		MsgContents[i].Role = v.Role
		MsgContents[i].Content = v.Content
	}

	req := &CloudflareAITextGenerationRequest{
		Messages:    MsgContents,
		Stream:      true,
		Temperature: promptAndSetting.Temperature,
		TopK:        promptAndSetting.TopK,
	}
	err := uc.repo.StreamTextGeneration(context.Background(), req, ch, promptAndSetting.ModelKind)
	if err != nil {
		ch <- &CloudflareAITextGenerationReply{
			Response: err.Error(),
		}
	}
	ch <- nil
}

func (uc *AICloudflareUsecase) AIPaint(promptAndSetting *pb.AIPaintRequest) ([]byte, error) {
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
	req := &CloudflareAIPaintingRequest{
		Prompt:   promptAndSetting.Prompt,
		Height:   promptAndSetting.Height,
		Width:    promptAndSetting.Width,
		Guidance: promptAndSetting.Guidance,
	}
	bin, err := uc.repo.ImageGeneration(req, promptAndSetting.ModelKind)
	if err != nil {
		return nil, errors.New(400, "Please adjust parameters", "Failed to generate images")
	}
	return bin, nil
}
