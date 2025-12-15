package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {

		// Default 500
		code := fiber.StatusInternalServerError

		// Fiber Error → ใช้ status code จาก error เดิม
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		// Log error (แบบ production)
		log.Printf("🔥 ERROR: %v", err)

		// JSON Response
		return c.Status(code).JSON(fiber.Map{
			"success": false,
			"error": fiber.Map{
				"message": err.Error(),
			},
		})
	}
}
