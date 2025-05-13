package router

import (
	"github.com/decrypt6969/auth-service/internal/handler"
	"github.com/decrypt6969/auth-service/internal/repository"
	"github.com/decrypt6969/auth-service/internal/service"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	authHandler := handler.NewAuthHandler(userService)

	auth := app.Group("/auth")
	auth.Post("/register", authHandler.Register)
}
