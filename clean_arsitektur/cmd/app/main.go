package main

import (
	"clean_arsitektur/internal/delivery/http/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.InitializeRoutes(app)

	app.Listen(":3000")
}
