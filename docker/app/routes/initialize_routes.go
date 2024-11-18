package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func authMiddleware(c *fiber.Ctx) error {
	// Implementasi middleware otentikasi
	log.Println("Middleware otentikasi dijalankan")
	return c.Next()
}

func InitializeRoutes(app *fiber.App) {

	// Inisialisasi middleware otentikasi
	// authMiddleware := middleware.AuthMiddleware()

	apiV1 := app.Group("/api/v1")
	protectedApiV1 := apiV1.Use(authMiddleware)

	// Public Routes
	TestRoutes(apiV1)

	// Protected Routes
	UserRoutes(protectedApiV1)
}
