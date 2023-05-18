package service

import (
	"BirthdayWisherBot/internal/models/bridge"
	"BirthdayWisherBot/utils"
	"context"
	"fmt"
	"github.com/e1esm/protobuf/bot_to_server/gen_proto"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"strconv"
)

type ConnectorService interface {
}

type BridgeConnectorService struct {
	client gen_proto.CongratulationServiceClient
}

func NewBridgeConnectorService(client gen_proto.CongratulationServiceClient) *BridgeConnectorService {
	return &BridgeConnectorService{client: client}
}

func (s *BridgeConnectorService) DeleteUser(userID, chatID int64) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	deleteResp, err := s.client.DeleteUser(ctx, &gen_proto.DeleteRequest{UserID: userID, ChatID: chatID})
	if err != nil {
		utils.Logger.Error(deleteResp.ErrorDescription)
		return err
	}
	return nil
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

func (s *BridgeConnectorService) GetChatStatistics(chatID int64) (tgbotapi.FileBytes, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	response, _ := s.client.GetStatistics(ctx, &gen_proto.ChatRequest{ChatID: chatID})
	receivedFile := tgbotapi.FileBytes{
		Name:  fmt.Sprintf("Chat-%d.pdf", chatID),
		Bytes: response.Data,
	}
	return receivedFile, nil
}

func (s *BridgeConnectorService) CreateInstanceToBeDelivered(update tgbotapi.Update) error {
	chat := bridge.NewChat(update.FromChat().ID)

	utils.RWapInstance.Mutex.RLock()
	v, _ := utils.RWapInstance.UserStateConfigs[update.SentFrom().ID]
	utils.RWapInstance.Mutex.RUnlock()
	year, month, day := transformCurrentDate(v.Year, v.Month, v.Day)
	date := fmt.Sprintf("%s-%s-%s", year, month, day)

	user := bridge.NewUser(
		v.UserID,
		date,
		*chat,
		update.SentFrom().FirstName+update.SentFrom().LastName,
	)
	if err := s.SaveUser(*user); err != nil {
		utils.Logger.Error("Couldn't have saved a user")
		return err
	}
	return nil
}

func transformCurrentDate(year, month, day int) (string, string, string) {
	utils.Logger.Info("", zap.Ints("Date", []int{year, month, day}))
	yearStr := strconv.Itoa(year)
	var monthStr string
	if month < 10 {
		monthStr = fmt.Sprintf("0%d", month)
	} else {
		monthStr = strconv.Itoa(month)
	}
	var dayStr string
	if day < 10 {
		dayStr = fmt.Sprintf("0%d", day)
	} else {
		dayStr = strconv.Itoa(day)
	}

	return yearStr, monthStr, dayStr
}
