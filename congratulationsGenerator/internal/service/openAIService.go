package service

import (
	"congratulationsGenerator/utils"
	"context"
	"fmt"
	"github.com/e1esm/protobuf/bridge_to_API/gen_proto"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
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
	builder.WriteString(fmt.Sprintf("Поздравь %s с днем рождения\n", name))
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
		utils.Logger.Error("Couldn't have gotten response from ChatGPT", zap.String("error", err.Error()))
	}
	utils.Logger.Info("ChatGPT response", zap.String("response", resp.Choices[0].Message.Content))
	return resp.Choices[0].Message.Content
}
