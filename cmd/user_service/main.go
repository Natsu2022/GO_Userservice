package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"OSCIRA.com/service/user_service/internal/config"
	"OSCIRA.com/service/user_service/internal/router"
	"OSCIRA.com/service/user_service/pkg/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg := config.LoadConfig()

	app := fiber.New(fiber.Config{
		AppName:      "user_service",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	})

	app.Use(recover.New())
	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", //!!! ก่อน Producttion ปรับตรงนี้ด้วย
		AllowMethods: "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Use(limiter.New(limiter.Config{
		Max:        60,
		Expiration: 1 * time.Minute,
	}))

	dbConn, err := db.ConnectPostgres(cfg)
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	router.SetupRoutes(app, dbConn)

	go func() {
		log.Println("Starting server on port:", cfg.AppPort)
		if err := app.Listen(":" + cfg.AppPort); err != nil {
			log.Printf("Server stopped: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = app.Shutdown()

	dbConn.Close()

	log.Println("Shutdown complete.")
}
