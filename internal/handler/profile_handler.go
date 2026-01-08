package handler

import (
	"OSCIRA.com/service/user_service/internal/constants"
	"OSCIRA.com/service/user_service/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *UserHandler) GetMyProfile(c *fiber.Ctx) error {
	userID, ok := c.Locals(constants.ContextUserID).(uuid.UUID)
	if !ok {
		return fiber.NewError(401, "unauthorized")
	}

	profile, err := h.userService.GetMyProfile(c.Context(), userID)
	if err != nil {
		return fiber.NewError(404, "profile not found")
	}

	return c.JSON(profile)
}

func (h *UserHandler) UpdateMyProfile(c *fiber.Ctx) error {
	userID := c.Locals(constants.ContextUserID).(uuid.UUID)

	var req model.UpdateProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.userService.UpdateMyProfile(c.Context(), userID, req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "updated"})
}
