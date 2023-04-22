package service

import "bridgeServer/internal/repository"

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}
