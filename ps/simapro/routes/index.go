package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	"github.com/gpjen/simapro/app/routes"
	_ "github.com/gpjen/simapro/docs"
	"github.com/gpjen/simapro/middlewares"
)

func SetupRoutesApp(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)

	routes.AuthRoutes(app)

	apiV1 := app.Group("/api/v1")
	apiV1Protected := apiV1.Use(middlewares.AuthMiddleware())

	routes.MenuRoutes(apiV1Protected)
	routes.UserRoutes(apiV1Protected)
	routes.RoleRoutes(apiV1Protected)

	// 404 Route
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Route not found"})
	})
}
