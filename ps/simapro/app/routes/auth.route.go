package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gpjen/simapro/app/controllers"
	"github.com/gpjen/simapro/app/dto"
	"github.com/gpjen/simapro/app/repositories"
	"github.com/gpjen/simapro/app/services"
	"github.com/gpjen/simapro/middlewares"
)

func AuthRoutes(route fiber.Router) {
	userRepository := repositories.NewUserRepository()
	userService := services.NewAuthService(userRepository)
	userController := controllers.NewAuthController(userService)

	route.Post("/login", middlewares.ValidateRequest(&dto.UserLoginDTO{}), userController.Login)
	route.Get("/logout", userController.Logout)
}
