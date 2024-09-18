package port

import (
	"backend/domain/entitie"
)

type UserService interface {
	CreateUser(user entitie.User) error
	LoginUser(email, password string) (entitie.User, error)
}

type UserRepository interface {
	CreateUser(user entitie.User) error
	LoginUser(email, password string) (entitie.User, error)
}
