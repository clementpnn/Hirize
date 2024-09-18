package handler

import (
	"backend/domain/entitie"
	"backend/domain/port"
	"backend/service"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	userService port.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService,
	}
}

func (h *UserHandler) CreateUser(c fiber.Ctx) error {
	var user entitie.User
	if err := c.Bind().JSON(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing request body",
			"data":    nil,
		})
	}

	if err := h.userService.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating user",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
		"data":    nil,
	})
}

func (h *UserHandler) LoginUser(c fiber.Ctx) error {
	var user entitie.User
	if err := c.Bind().JSON(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing request body",
			"data":    nil,
		})
	}
	user, err := h.userService.LoginUser(user.Email, user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating user",
			"data":    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
		"data":    nil,
	})
}
