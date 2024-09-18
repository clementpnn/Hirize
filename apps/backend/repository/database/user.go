package repository

import (
	"backend/domain/entitie"
	"backend/domain/port"
	"database/sql"
)

func NewUserRepository(db *sql.DB) port.UserRepository {
	return &Repository{
		db,
	}
}

func (r *Repository) CreateUser(user entitie.User) error {
	stmt, err := r.db.Prepare("INSERT INTO users (email, password) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) LoginUser(email string, password string) (entitie.User, error) {
	var user entitie.User
	err := r.db.QueryRow("SELECT * FROM users WHERE email = $1 AND password = $2", email, password).Scan(&user.Email, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}
