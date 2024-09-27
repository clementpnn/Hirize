package port

import (
	"backend/domain/entitie"
)

type UserService interface {
	CreateUser(userForm entitie.UserForm) (*entitie.User, error)
	FindUserByEmail(email string) (*entitie.User, error)
}

type UserRepository interface {
	CreateUser(userForm entitie.UserForm) (*entitie.User, error)
	FindUserByEmail(email string) (*entitie.User, error)
}
