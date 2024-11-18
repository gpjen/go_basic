package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {

	router.Get("/users", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Hello USER, World!",
		})
	})

	router.Get("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"id": id,
		})
	})
}
