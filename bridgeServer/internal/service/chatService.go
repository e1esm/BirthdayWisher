package service

import "bridgeServer/internal/repository"

type ChatService struct {
	userRepositories *repository.UserRepository
}

func NewChatService(userRepository *repository.UserRepository) *ChatService {
	return &ChatService{userRepositories: userRepository}
}
