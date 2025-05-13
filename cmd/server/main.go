package main

import (
	"log"
	"os"

	"github.com/decrypt6969/auth-service/internal/config"
	"github.com/decrypt6969/auth-service/internal/db"
	"github.com/decrypt6969/auth-service/internal/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	db.ConnectPostgres()

	app := fiber.New()

	router.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
