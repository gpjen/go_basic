package routes

import "github.com/gofiber/fiber/v2"

type routesGroup struct {
	app *fiber.App
}

func NewRouterGroup(app *fiber.App) *routesGroup {
	return &routesGroup{app}
}

func (r *routesGroup) Install() {
	api := r.app.Group("/api")

	v1 := api.Group("/v1")
	NewUserRoute(v1).Install()
}
