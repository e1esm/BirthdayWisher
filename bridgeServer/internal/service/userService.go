package service

import (
	"bridgeServer/internal/model"
	"bridgeServer/internal/repository"
	bot_to_server_proto "github.com/e1esm/protobuf/bot_to_server/gen_proto"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) SaveUser(user *model.User) {
	s.userRepository.SaveUser(user)
}

func (s *UserService) GetUsersWithBirthdayToday() []model.User {
	return s.userRepository.FindUsers()
}

func (s *UserService) GetUsersWithBirthdaySoon(chatId int64) *bot_to_server_proto.ChatBirthdaysResponse {
	return s.userRepository.SoonBirthdaysOfUsers(chatId)
}
