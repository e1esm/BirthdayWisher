package service

import "bridgeServer/internal/repository"

type UserService struct {
	repositories *repository.Repositories
}

func NewUserService(repositories *repository.Repositories) *UserService {
	return &UserService{repositories: repositories}
}
