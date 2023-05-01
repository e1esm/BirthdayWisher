package service

import (
	"BirthdayWisherBot/internal/models/bridge"
	"BirthdayWisherBot/utils"
	"context"
	"fmt"
	"github.com/e1esm/protobuf/bot_to_server/gen_proto"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
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
		ChatRequest: &gen_proto.ChatRequest{ChatID: user.CurrentChat.ChatId}})
	if err != nil {
		utils.Logger.Error("Received error from Protobuf message", zap.String("error", err.Error()), zap.String("user", user.Username))
		return err
	}
	return nil
}

func (s *BridgeConnectorService) DailyRetriever() ([]*gen_proto.CongratulationResponse, error) {
	messages := make([]*gen_proto.CongratulationResponse, 0, 10)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	message, err := s.client.GetDataForCongratulations(ctx, new(emptypb.Empty))
	
	if err != nil {
		utils.Logger.Error("Failed while receiving protobuf message's data from daily birthdays checker")
		return nil, err
	}
	for {
		retrievedMessage, err := message.Recv()
		if err == io.EOF {
			break
		}
		messages = append(messages, retrievedMessage)
	}
	return messages, nil
}

func (s *BridgeConnectorService) GetSoonBirthdays(chatID int64) (*gen_proto.ChatBirthdaysResponse, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	chatBirthdaysInfo, err := s.client.GetSoonBirthdays(ctx, &gen_proto.ChatRequest{ChatID: chatID})

	if err != nil {
		utils.Logger.Error("Error while checking today's birthdays of the chat's users", zap.Int64("chatID", chatID))
		return nil, fmt.Errorf("couldn't have gotten soon birthday list: %s", err)
	}
	return chatBirthdaysInfo, nil
}
