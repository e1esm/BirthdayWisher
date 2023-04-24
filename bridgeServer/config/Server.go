package config

import (
	"bridgeServer/internal/model"
	"bridgeServer/internal/service"
	"context"
	bot_to_server_proto "github.com/e1esm/protobuf/bot_to_server/gen_proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"time"
)

const (
	layoutISO = "1970-01-01"
)

type Server struct {
	userService *service.UserService
	chatService *service.ChatService
	bot_to_server_proto.CongratulationServiceServer
}

func NewServer(userService *service.UserService, chatService *service.ChatService) *Server {
	return &Server{userService: userService, chatService: chatService}
}

func (s *Server) SaveUserInfo(ctx context.Context, req *bot_to_server_proto.UserRequest) (*emptypb.Empty, error) {
	chat := model.NewChat(req.ChatRequest.ChatID, req.UserID)
	date, err := time.Parse(time.DateOnly, req.Date)
	log.Println(date)
	if err != nil {
		return new(emptypb.Empty), err
	}
	user := model.NewUser(req.UserID, date, []model.Chat{*chat})
	s.userService.SaveUser(user)
	return new(emptypb.Empty), nil
}
