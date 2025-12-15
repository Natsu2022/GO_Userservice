package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"OSCIRA.com/service/user_service/internal/model"
	"OSCIRA.com/service/user_service/internal/service"
)

type UserHandler struct {
	userService service.UserService
	authService service.AuthService
}

func NewUserHandler(u service.UserService, a service.AuthService) *UserHandler {
	return &UserHandler{userService: u, authService: a}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req model.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON")
	}

	id, err := h.userService.CreateUser(c.Context(), req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(fiber.Map{
		"success": true,
		"user_id": id,
	})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req model.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON")
	}

	session, err := h.authService.Login(c.Context(), req, c.IP())
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	c.Cookie(&fiber.Cookie{
		Name:     "session_id",
		Value:    session.ID.String(),
		HTTPOnly: true,
		Secure:   false, // IMPORTANT สำหรับ localhost!
		SameSite: "Lax",
		Path:     "/",
		MaxAge:   86400,
	})

	return c.JSON(fiber.Map{
		"success": true,
		"message": "login successful",
	})
}

func (h *UserHandler) Logout(c *fiber.Ctx) error {
	cookie := c.Cookies("session_id")
	if cookie == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "no session")
	}

	sessionUUID, err := uuid.Parse(cookie)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid session")
	}

	if err := h.authService.Logout(c.Context(), sessionUUID); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to logout")
	}

	// Clear cookie
	c.Cookie(&fiber.Cookie{
		Name:     "session_id",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{
		"success": true,
		"message": "logged out",
	})
}
