package routes

import (
	"clean_arsitektur/internal/delivery/http/handler"
	"clean_arsitektur/internal/repository"
	"clean_arsitektur/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app appRoutes) {

	userRepo := repository.NewUserRepo()
	userUsecase := usecase.NewUserUsecase(userRepo)
	handler := handler.NewUserHandler(userUsecase)

	app.api.Get("/users/:id", handler.GetUser)
	app.api.Get("/users", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

}
