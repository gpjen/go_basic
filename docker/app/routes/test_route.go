package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func TestRoutes(router fiber.Router) {
	router.Get("/test", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Hello TEST, World! tets",
		})
	})
}
