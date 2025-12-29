package router

import (
	"github.com/gofiber/fiber/v2"

	"OSCIRA.com/service/user_service/internal/handler"
	"OSCIRA.com/service/user_service/internal/middleware"
	"OSCIRA.com/service/user_service/internal/repository"
	"OSCIRA.com/service/user_service/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(app *fiber.App, db *pgxpool.Pool) {

	userRepo := repository.NewUserRepository(db)
	sessionRepo := repository.NewSessionRepository(db)

	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo, sessionRepo)

	userHandler := handler.NewUserHandler(userService, authService)

	api := app.Group("/api/v1")

	users := api.Group("/users")
	users.Post("/register", userHandler.CreateUser)

	auth := users.Group("/auth")
	auth.Post("/login", userHandler.Login)
	auth.Post("/logout", userHandler.Logout)

	// Protected routes
	protected := auth.Group("/me", middleware.AuthGuard(sessionRepo))
	protected.Get("/", func(c *fiber.Ctx) error {
		userID := c.Locals("userID") // UUID ของ user
		return c.JSON(fiber.Map{
			"message": "You are authenticated",
			"user_id": userID,
		})
	})
}