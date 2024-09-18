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

	db, err := config.OpenDB()
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Use(cors.New())

	router.Router(app, db)

	app.Listen(":3000")
}
