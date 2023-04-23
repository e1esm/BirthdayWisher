package service

import (
	"BirthdayWisherBot/internal/models/bridge"
	"context"
	"github.com/e1esm/protobuf/bot_to_server/gen_proto"
)

type ConnectorService interface {
}

type BridgeConnectorService struct {
	client gen_proto.CongratulationServiceClient
}

func NewBridgeConnectorService(client gen_proto.CongratulationServiceClient) *BridgeConnectorService {
	return &BridgeConnectorService{client: client}
}

func (s *BridgeConnectorService) SaveUser(user bridge.User) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res, err := s.client.SaveUserInfo(ctx, &gen_proto.UserRequest{UserID: user.UserId, Date: user.Date,
		ChatRequest: &gen_proto.UserRequest_ChatRequest{ChatID: user.CurrentChat.ChatId}})
}
