package routes

import "github.com/gofiber/fiber/v2"

type userRoutes struct {
	router fiber.Router
	// handler
}

func NewUserRoutes(router fiber.Router) *userRoutes {
	// call dependencies
	return &userRoutes{
		router: router,
	}
}

func (r *userRoutes) Install() {
	users := r.router.Group("/users")

	users.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})
}
