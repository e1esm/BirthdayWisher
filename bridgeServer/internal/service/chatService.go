package service

import "bridgeServer/internal/repository"

type ChatService struct {
	chatRepository repository.ChatRepository
}

func NewChatService(chatRepository repository.ChatRepository) *ChatService {
	return &ChatService{chatRepository: chatRepository}
}
