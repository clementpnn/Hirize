package repository

import (
	"backend/domain/entitie"
	"backend/domain/port"

	"github.com/jmoiron/sqlx"
)

func NewUserRepository(db *sqlx.DB) port.UserRepository {
	return &Repository{
		db,
	}
}

func (r *Repository) CreateUser(userForm entitie.UserForm) (*entitie.User, error) {
	stmt, err := r.db.Preparex("INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id, email, password")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	user := &entitie.User{}
	err = stmt.QueryRowx(userForm.Email, userForm.Password).StructScan(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) FindUserByEmail(email string) (*entitie.User, error) {
	stmt, err := r.db.Preparex("SELECT id, email, password FROM users WHERE email = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	user := &entitie.User{}
	err = stmt.QueryRowx(email).StructScan(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
