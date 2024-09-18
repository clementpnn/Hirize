package service

import (
	"backend/domain/entitie"
	"backend/domain/port"
)

type UserService struct {
	userRepository port.UserRepository
}

func NewUserService(userRepository port.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) CreateUser(user entitie.User) error {
	return s.userRepository.CreateUser(user)
}

func (s *UserService) LoginUser(email string, password string) (entitie.User, error) {
	return s.userRepository.LoginUser(email, password)
}
