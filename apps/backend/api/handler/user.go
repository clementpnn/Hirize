package handler

import (
	"backend/domain/entitie"
	"backend/domain/port"
	"backend/service"
	"log"

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
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not create user",
			"error":   err.Error(),
		})
	}

	hashedPassword, err := h.securityService.HashPassword(userForm.Password)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create user",
			"error":   err.Error(),
		})
	}

	userForm.Password = hashedPassword

	user, err := h.userService.CreateUser(userForm)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create user",
			"error":   err.Error(),
		})
	}

	sess, err := h.sessionStore.Get(c)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create user",
			"error":   err.Error(),
		})
	}

	sess.Set("userID", user.ID)
	if err := sess.Save(); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create user",
			"error":   err.Error(),
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
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not login user",
			"error":   err.Error(),
		})
	}

	userFound, err := h.userService.FindUserByEmail(userForm.Email)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Could not login user",
			"error":   err.Error(),
		})
	}

	err = h.securityService.CheckPasswordHash(userForm.Password, userFound.Password)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Could not login user",
			"error":   err.Error(),
		})
	}

	sess, err := h.sessionStore.Get(c)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not login user",
			"error":   err.Error(),
		})
	}

	sess.Set("userID", userFound.ID)
	if err := sess.Save(); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not login user",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User logged in",
		"data":    nil,
	})
}
