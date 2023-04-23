package service

import "bridgeServer/internal/repository"

type ChatService struct {
	repositories *repository.Repositories
}

func NewChatService(repositories *repository.Repositories) *ChatService {
	return &ChatService{repositories: repositories}
}
