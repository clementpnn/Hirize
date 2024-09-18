package router

import (
	"backend/api/handler"
	repository "backend/repository/database"
	"backend/service"
	"database/sql"

	"github.com/gofiber/fiber/v3"
)

func User(app *fiber.App, db *sql.DB) {
	user := app.Group("/user")

	userHandler := handler.NewUserHandler(service.NewUserService(repository.NewUserRepository(db)))

	user.Post("/create", userHandler.CreateUser)
	user.Post("/login", userHandler.LoginUser)
}
