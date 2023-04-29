package service

import (
	"bridgeServer/utils"
	"context"
	"github.com/e1esm/protobuf/bridge_to_API/gen_proto"
	"go.uber.org/zap"
	"time"
)

type GPTService struct {
	client gen_proto.CongratulationServiceClient
}

func NewGPTService(client gen_proto.CongratulationServiceClient) *GPTService {
	return &GPTService{client: client}
}

func (s *GPTService) GetCongratulation(name string) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	res, err := s.client.QueryForCongratulation(ctx, &gen_proto.CongratulationRequest{Name: name})
	if err != nil {
		utils.Logger.Error("Didn't get result from querying for congratulations", zap.String("error", err.Error()))
		return ""
	}
	return res.CongratulationSentence
}
