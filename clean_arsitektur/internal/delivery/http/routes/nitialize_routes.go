package routes

import "github.com/gofiber/fiber/v2"

type appRoutes struct {
	api            fiber.Router
	authMiddleware fiber.Handler
}

func InitializeRoutes(app *fiber.App) {

	router := appRoutes{
		api:            app.Group("/api/v1"),
		authMiddleware: nil,
	}

	UserRoutes(router)

}
