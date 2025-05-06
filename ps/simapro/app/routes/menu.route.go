package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gpjen/simapro/app/controllers"
	"github.com/gpjen/simapro/app/dto"
	"github.com/gpjen/simapro/app/repositories"
	"github.com/gpjen/simapro/app/services"
	"github.com/gpjen/simapro/middlewares"
)

func MenuRoutes(route fiber.Router) {
	menuRepository := repositories.NewMenuRepository()
	menuService := services.NewMenuService(menuRepository)
	menuController := controllers.NewMenuController(menuService)

	menuGroup := route.Group("/menus")

	menuGroup.Get("/", menuController.GetMenuHierarchy)
	menuGroup.Patch("/", middlewares.ValidateRequest(&dto.UpdateMenusDto{}), menuController.MenuUpdates)

}
