package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gpjen/simapro/app/controllers"
	"github.com/gpjen/simapro/app/dto"
	"github.com/gpjen/simapro/app/repositories"
	"github.com/gpjen/simapro/app/services"
	"github.com/gpjen/simapro/middlewares"
)

func RoleRoutes(route fiber.Router) {
	roleRepository := repositories.NewRoleRepository()
	roleService := services.NewRoleService(roleRepository)
	roleController := controllers.NewRoleController(roleService)

	roleGroup := route.Group("/roles")

	roleGroup.Get("/", roleController.GetRoles)
	roleGroup.Get("/:id", roleController.GetRole)
	roleGroup.Post("/", middlewares.ValidateRequest(&dto.CreateRoleDTO{}), roleController.CreateRole)
	roleGroup.Patch("/:id", middlewares.ValidateRequest(&dto.UpdateRoleDTO{}), roleController.UpdateRole)
	roleGroup.Delete("/:id", roleController.DeleteRole)

	roleGroup.Get("/menu-permission/:id", roleController.GetRoleMenuPermision)
}
