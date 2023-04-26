package service

import (
	"BirthdayWisherBot/internal/models/bridge"
	"context"
	"github.com/e1esm/protobuf/bot_to_server/gen_proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
)

type ConnectorService interface {
}

type BridgeConnectorService struct {
	client gen_proto.CongratulationServiceClient
}

func NewBridgeConnectorService(client gen_proto.CongratulationServiceClient) *BridgeConnectorService {
	return &BridgeConnectorService{client: client}
}

func (s *BridgeConnectorService) SaveUser(user bridge.User) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := s.client.SaveUserInfo(ctx, &gen_proto.UserRequest{UserID: user.UserId, Date: user.Date, Username: user.Username,
		ChatRequest: &gen_proto.UserRequest_ChatRequest{ChatID: user.CurrentChat.ChatId}})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *BridgeConnectorService) DailyRetriever() ([]*gen_proto.CongratulationResponse, error) {
	messages := make([]*gen_proto.CongratulationResponse, 0, 10)
	message, err := s.client.GetDataForCongratulations(context.Background(), new(emptypb.Empty))
	if err != nil {
		log.Println("Couldn't have listened to server's stream")
		return messages, err
	}
	for {
		retrievedMessage, err := message.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Couldn't have retrieved message from protobuf")
			return messages, nil
		}
		messages = append(messages, retrievedMessage)
	}

	return messages, nil
}
