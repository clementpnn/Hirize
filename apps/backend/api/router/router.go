package router

import (
	"github.com/jmoiron/sqlx"

	"github.com/gofiber/fiber/v3"
)

func Router(app *fiber.App, db *sqlx.DB) {
	User(app, db)

	// app.Get("/csrf-token", func(c fiber.Ctx) error {
	// 	log.Println("CSRF")
	// 	token := c.Locals("csrf")
	// 	if token == nil {
	// 		log.Println("CSRF Token non trouvé")
	// 	}
	// 	return c.JSON(fiber.Map{
	// 		"csrfToken": token,
	// 	})
	// })

}
