package service

import (
	"context"
	"github.com/e1esm/congr_proto"
	"github.com/sashabaranov/go-openai"
	"log"
)

type OpenAIService struct {
	Client *openai.Client
	congr_proto.CongratulationServiceServer
}

func NewOpenAIService(client *openai.Client) *OpenAIService {
	return &OpenAIService{Client: client}
}

func (s *OpenAIService) QueryForCongratulation(ctx context.Context, request *congr_proto.CongratulationRequest) (*congr_proto.CongratulationResponse, error) {
	sentence := s.QueryFromAI(request.Name)
	return &congr_proto.CongratulationResponse{CongratulationSentence: sentence}, nil
}

func (s *OpenAIService) QueryFromAI(name string) string {
	resp, err := s.Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Wish happy birthday to: " + name},
			},
		},
	)
	if err != nil {
		log.Fatalf("ChatCompletion error: %v", err)
	}
	return resp.Choices[0].Message.Content
}

//(tg://user?id=?")
