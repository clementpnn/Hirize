package handler

import (
	"backend/domain/entitie"
	"backend/domain/port"
	"backend/service"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

type UserHandler struct {
	userService     port.UserService
	securityService port.SecurityService
	sessionStore    *session.Store
}

func NewUserHandler(userService *service.UserService, securityService *service.SecurityService, sessionStore *session.Store) *UserHandler {
	return &UserHandler{
		userService,
		securityService,
		sessionStore,
	}
}

func (h *UserHandler) CreateUser(c fiber.Ctx) error {
	var userForm entitie.UserForm
	if err := c.Bind().JSON(&userForm); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not create user",
			"data":    nil,
		})
	}

	hashedPassword, err := h.securityService.HashPassword(userForm.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create user",
			"data":    nil,
		})
	}

	userForm.Password = hashedPassword

	user, err := h.userService.CreateUser(userForm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create user",
			"data":    nil,
		})
	}

	sess, err := h.sessionStore.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create user",
			"data":    nil,
		})
	}

	sess.Set("userID", user.Email)
	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create user",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created",
		"data":    nil,
	})
}

func (h *UserHandler) LoginUser(c fiber.Ctx) error {
	var userForm entitie.UserForm
	if err := c.Bind().JSON(&userForm); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not login user",
			"data":    nil,
		})
	}

	userFound, err := h.userService.FindUserByEmail(userForm.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Could not login user",
			"data":    nil,
		})
	}

	err = h.securityService.CheckPasswordHash(userForm.Password, userFound.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Could not login user",
			"data":    nil,
		})
	}

	sess, err := h.sessionStore.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not login user",
			"data":    nil,
		})
	}

	sess.Set("userID", userFound.ID.String())
	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not login user",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User logged in",
		"data":    nil,
	})
}

func (h *UserHandler) CheckSession(c fiber.Ctx) error {
	sess, err := h.sessionStore.Get(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Could not check session",
			"data":    false,
		})
	}

	userID := sess.Get("userID")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Could not check session",
			"data":    false,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User logged in",
		"data":    true,
	})
}
