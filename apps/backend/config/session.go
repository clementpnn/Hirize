package config

import (
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/gofiber/storage/redis/v3"
)

var SessionStore *session.Store

func InitSession() {
	redisStore := redis.New(redis.Config{
		Host:     "localhost",
		Port:     6379,
		Username: "default",
		Password: "mysecretpassword",
		Database: 0,
	})

	SessionStore = session.New(session.Config{
		Storage: redisStore,
	})
}
