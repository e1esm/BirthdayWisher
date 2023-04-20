package service

import (
	"context"
	"github.com/e1esm/protobuf/bridge_to_API/gen_proto"
	"github.com/sashabaranov/go-openai"
	"log"
	"strings"
)

type OpenAIService struct {
	Client *openai.Client
	gen_proto.CongratulationServiceServer
}

func NewOpenAIService(client *openai.Client) *OpenAIService {
	return &OpenAIService{Client: client}
}

func (s *OpenAIService) QueryForCongratulation(ctx context.Context, request *gen_proto.CongratulationRequest) (*gen_proto.CongratulationResponse, error) {
	sentence := s.QueryFromAI(request.Name)
	return &gen_proto.CongratulationResponse{CongratulationSentence: sentence}, nil
}

func (s *OpenAIService) QueryFromAI(name string) string {
	builder := strings.Builder{}
	builder.WriteString("Wish happy birthday to ")
	builder.WriteString(name)
	builder.WriteString(" as a friend would do")
	resp, err := s.Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: builder.String()},
			},
		},
	)
	if err != nil {
		log.Fatalf("ChatCompletion error: %v", err)
	}
	return resp.Choices[0].Message.Content
}

//(tg://user?id=?")
