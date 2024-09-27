package service

import (
	"golang.org/x/crypto/bcrypt"
)

type SecurityService struct{}

func NewSecurityService() *SecurityService {
	return &SecurityService{}
}

func (s *SecurityService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (s *SecurityService) CheckPasswordHash(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
