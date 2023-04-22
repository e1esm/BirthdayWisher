package service

import (
	"context"
	"github.com/e1esm/protobuf/bridge_to_API/gen_proto"
	"log"
	"time"
)

type GPTService struct {
	client gen_proto.CongratulationServiceClient
}

func NewGPTService(client gen_proto.CongratulationServiceClient) *GPTService {
	return &GPTService{client: client}
}

func (s *GPTService) GetCongratulation(name string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	res, err := s.client.QueryForCongratulation(ctx, &gen_proto.CongratulationRequest{Name: name})
	if err != nil {
		log.Fatalf("Couldn't have gotten result from querying congratulation: %s", err)
	}
	log.Printf("%s\n", res.CongratulationSentence)
}
