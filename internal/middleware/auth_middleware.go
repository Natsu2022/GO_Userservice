package middleware

import (
	"log"
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
		}else{
			log.Println("step1 sid: " + sid)
		}

		sessionID, err := uuid.Parse(sid)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "invalid session")
		}else{
			log.Println("step2 Session: ", sessionID)
		} 

		// load session
		sess, err := sessionRepo.GetSession(c.Context(), sessionID)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "step3 session not found")
		}else{
			log.Println("sessionRepo found session.")
		}

		// expired?
		if time.Now().After(sess.ExpiresAt) {
			_ = sessionRepo.DeleteSession(c.Context(), sessionID)
			return fiber.NewError(fiber.StatusUnauthorized, "step4 session expired")
		}

		// store userID in context for next handler
		c.Locals("userID", sess.UserID)

		return c.Next()
	}
}
