package biz

import (
	"context"
	"strings"
)

type AICloudflareRepo interface {
	StreamTextGeneration(ctx context.Context, messages *CloudflareAITextGenerationRequest, ch chan *CloudflareAITextGenerationReply, model string) error
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

func (uc *AICloudflareUsecase) StreamGetAISummarization(text_bytes []byte, ch chan *CloudflareAITextGenerationReply) {
	var builder strings.Builder
	builder.WriteString("对这篇markdown格式的文章进行摘要：\n")
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

	err := uc.repo.StreamTextGeneration(context.Background(), req, ch, "@cf/meta/llama-3.3-70b-instruct-fp8-fast")
	if err != nil {
		ch <- &CloudflareAITextGenerationReply{
			Response: err.Error(),
		}
	}
	ch <- nil
}
