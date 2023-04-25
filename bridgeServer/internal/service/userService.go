package service

import (
	"bridgeServer/internal/model"
	"bridgeServer/internal/repository"
	"log"
)

type UserService struct {
	repositories *repository.Repositories
}

func NewUserService(repositories *repository.Repositories) *UserService {
	return &UserService{repositories: repositories}
}

func (s *UserService) SaveUser(user *model.User) {
	log.Println(user.ID)
	log.Println(user.CurrentChat)
	s.repositories.UserRepository.SaveUser(user)
}
