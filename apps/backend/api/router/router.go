package router

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
)

func Router(app *fiber.App, db *sql.DB) {
	User(app, db)
}
