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

func (s *UserService) CreateUser(userForm entitie.UserForm) (*entitie.User, error) {
	return s.userRepository.CreateUser(userForm)
}

func (s *UserService) FindUserByEmail(email string) (*entitie.User, error) {
	return s.userRepository.FindUserByEmail(email)
}
