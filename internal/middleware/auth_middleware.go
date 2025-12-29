package middleware

import (
	"time"

	"OSCIRA.com/service/user_service/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AuthGuard(sessionRepo repository.SessionRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {

		sid := c.Cookies("session_id")
		if sid == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "missing session")
		}

		sessionID, err := uuid.Parse(sid)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "invalid session")
		} 

		// load session
		sess, err := sessionRepo.GetSession(c.Context(), sessionID)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "session not found")
		}

		// expired?
		if time.Now().After(sess.ExpiresAt) {
			_ = sessionRepo.DeleteSession(c.Context(), sessionID)
			return fiber.NewError(fiber.StatusUnauthorized, "session expired")
		}

		// store userID in context for next handler
		c.Locals("userID", sess.UserID)

		return c.Next()
	}
}
