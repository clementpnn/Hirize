package main

import (
	"backend/api/router"
	"backend/config"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("ENV")
	if env != "production" {
		if err := godotenv.Load("../../.env"); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	db := config.InitDB()
	config.InitSession()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL")},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
	}))
	// app.Use(csrf.New(csrf.Config{
	// 	CookieName:     "csrf_token",
	// 	CookieSecure:   false, // Mettre sur true en production (HTTPS)
	// 	CookieHTTPOnly: true,
	// 	CookieSameSite: "Strict",
	// 	KeyLookup:      "cookie:csrf_token",
	// }))

	router.Router(app, db)

	app.Listen(":3000")
}
