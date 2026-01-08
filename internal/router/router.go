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

	handler := handler.NewUserHandler(userService, authService)

	api := app.Group("/api/v1")

	api.Post("/users/register", handler.CreateUser)
	api.Post("/auth/login", handler.Login)
	api.Post("/auth/logout", handler.Logout)

	me := api.Group("/me", middleware.AuthGuard(sessionRepo))
	me.Get("/profile", handler.GetMyProfile)
	me.Patch("/profile", handler.UpdateMyProfile)
}
