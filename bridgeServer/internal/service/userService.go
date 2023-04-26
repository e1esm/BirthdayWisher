package service

import (
	"bridgeServer/internal/model"
	"bridgeServer/internal/repository"
)

type UserService struct {
	repositories *repository.Repositories
}

func NewUserService(repositories *repository.Repositories) *UserService {
	return &UserService{repositories: repositories}
}

func (s *UserService) SaveUser(user *model.User) {
	s.repositories.UserRepository.SaveUser(user)
}

func (s *UserService) GetUsersWithBirthdayToday() []model.User {
	return s.repositories.UserRepository.FindUsers()
}
