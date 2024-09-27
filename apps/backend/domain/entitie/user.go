package entitie

import (
	"github.com/google/uuid"
)

type UserForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	*UserForm
	ID uuid.UUID `json:"id"`
}
