package config

import (
	"bridgeServer/internal/model"
	"bridgeServer/internal/service"
	"bridgeServer/utils"
	"context"
	bot_to_server_proto "github.com/e1esm/protobuf/bot_to_server/gen_proto"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"sync"
	"time"
)

type Server struct {
	userService *service.UserService
	gptService  *service.GPTService
	bot_to_server_proto.CongratulationServiceServer
	config *Config
}

func NewServer(userService *service.UserService, gptService *service.GPTService, config *Config) *Server {
	return &Server{userService: userService, gptService: gptService, config: config}
}

func (s *Server) SaveUserInfo(ctx context.Context, req *bot_to_server_proto.UserRequest) (*emptypb.Empty, error) {
	chat := model.NewChat(req.ChatRequest.ChatID, req.ChatRequest.ChatID)
	date, err := time.Parse(time.DateOnly, req.Date)
	if err != nil {
		utils.Logger.Error("Time parse error", zap.String("error", err.Error()))
		return new(emptypb.Empty), err
	}
	localization, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		utils.Logger.Error("Localization loading error", zap.String("error", err.Error()))
		return new(emptypb.Empty), err
	}
	dateInMoscow := date.In(localization)
	user := model.NewUser(req.UserID, dateInMoscow, []model.Chat{*chat}, req.Username)
	s.userService.SaveUser(user)
	return new(emptypb.Empty), nil
}

func (s *Server) GetDataForCongratulations(req *emptypb.Empty, server bot_to_server_proto.CongratulationService_GetDataForCongratulationsServer) error {
	users := s.userService.GetUsersWithBirthdayToday()
	wg := new(sync.WaitGroup)
	for _, user := range users {
		wg.Add(1)
		go func(wg *sync.WaitGroup, user model.User) {
			chats := make([]*bot_to_server_proto.ChatRequest, 0)
			for _, chat := range user.CurrentChat {
				chats = append(chats, &bot_to_server_proto.ChatRequest{ChatID: chat.ID})
			}
			congratulationSentence := s.gptService.GetCongratulation(user.Username)
			res := &bot_to_server_proto.CongratulationResponse{Username: user.Username, UserID: user.ID, ChatIDs: chats, CongratulationSentence: congratulationSentence}
			if err := server.Send(res); err != nil {
				utils.Logger.Error("Couldn't have sent a message to user", zap.String("user", res.Username))
				log.Println("Couldn't have sent a message.")
			}
			wg.Done()
		}(wg, user)
		time.Sleep(time.Second * 21)
	}
	wg.Wait()
	return nil
}

func (s *Server) GetSoonBirthdays(ctx context.Context, req *bot_to_server_proto.ChatRequest) (*bot_to_server_proto.ChatBirthdaysResponse, error) {
	return s.userService.GetUsersWithBirthdaySoon(req.ChatID), nil
}
