package config

import (
	"os"

	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/gofiber/storage/redis/v3"
)

var SessionStore *session.Store

func InitSession() {
	redisStore := redis.New(redis.Config{
		Host:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Database: 0,
	})

	SessionStore = session.New(session.Config{
		Storage:        redisStore,
		CookieHTTPOnly: true,
		CookieSecure:   true,
		CookieDomain:   os.Getenv("SESSION_DOMAIN"),
	})
}
