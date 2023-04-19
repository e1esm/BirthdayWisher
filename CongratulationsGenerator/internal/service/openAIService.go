package service

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
)

type OpenAIService struct {
	client *openai.Client
}

func NewOpenAIService(client *openai.Client) *OpenAIService {
	return &OpenAIService{client: client}
}

func (s *OpenAIService) Query(name string) {
	resp, err := s.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: name},
			},
		},
	)
	if err != nil {
		log.Fatalf("ChatCompletion error: %v", err)
	}
	fmt.Println(resp.Choices[0].Message.Content)

}

//(tg://user?id=?")
