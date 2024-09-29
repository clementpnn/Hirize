package router

import (
	"backend/api/handler"
	"backend/config"
	repository "backend/repository/database"
	"backend/service"

	"github.com/jmoiron/sqlx"

	"github.com/gofiber/fiber/v3"
)

func User(app *fiber.App, db *sqlx.DB) {
	user := app.Group("/user")

	userHandler := handler.NewUserHandler(service.NewUserService(repository.NewUserRepository(db)), service.NewSecurityService(), config.SessionStore)

	user.Post("/create", userHandler.CreateUser)
	user.Post("/login", userHandler.LoginUser)
	user.Get("/check-session", userHandler.CheckSession)
}
