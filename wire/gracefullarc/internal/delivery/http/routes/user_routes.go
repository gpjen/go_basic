package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type userRoutes struct {
	router fiber.Router
	// handler
}

func NewUserRoute(router fiber.Router) *userRoutes {
	return &userRoutes{router}
}

func (r *userRoutes) Install() {
	users := r.router.Group("/users")

	users.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

}
