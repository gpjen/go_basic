package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gpjen/simapro/app/controllers"
	"github.com/gpjen/simapro/app/dto"
	"github.com/gpjen/simapro/app/repositories"
	"github.com/gpjen/simapro/app/services"
	"github.com/gpjen/simapro/middlewares"
)

func UserRoutes(route fiber.Router) {
	userRepository := repositories.NewUserRepository()
	roleRepository := repositories.NewRoleRepository()
	userService := services.NewUserService(userRepository, roleRepository)
	userController := controllers.NewUserController(userService)

	userGroup := route.Group("/users")

	userGroup.Get("/", userController.GetUsers)
	userGroup.Get("/:id", userController.GetUser)
	userGroup.Post("/", middlewares.ValidateRequest(&dto.CreateUserDTO{}), userController.CreateUser)
	userGroup.Patch("/:id", middlewares.ValidateRequest(&dto.UpdateUserDTO{}), userController.UpdateUser)
	userGroup.Delete("/:id", userController.DeleteUser)
}
