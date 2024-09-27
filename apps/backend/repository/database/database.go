package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}
