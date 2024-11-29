package handler

import (
	"gracefullac/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userUseCase *domain.UserUseCase
}

func NewUserHandler(app *fiber.App, userUseCase *domain.UserUseCase) {
	// handler := &userHandler{
	// 	userUseCase,
	// }

	app.Post("/user")
}
