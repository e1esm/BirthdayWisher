package config

import (
	"bridgeServer/internal/service"
	bot_to_server_proto "github.com/e1esm/protobuf/bot_to_server/gen_proto"
)

type Server struct {
	userService *service.UserService
	chatService *service.ChatService
	bot_to_server_proto.CongratulationServiceServer
}

func NewServer(userService *service.UserService, chatService *service.ChatService) *Server {
	return &Server{userService: userService, chatService: chatService}
}
