package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gpjen/simapro/app/dto"
	"github.com/gpjen/simapro/app/services"
	"github.com/gpjen/simapro/utils"
)

type AuthController struct {
	AuthService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService}
}

func (ctrl *AuthController) Login(c *fiber.Ctx) error {
	h := &utils.ResponseHandler{}
	var dto dto.UserLoginDTO

	if err := c.BodyParser(&dto); err != nil {
		return h.BadRequest(c, []string{err.Error()})
	}

	user, err := ctrl.AuthService.Login(dto)
	if err != nil {
		return h.BadRequest(c, []string{err.Error()})
	}

	utils.SetTokenCookie(c, user.Token)
	return h.Ok(c, user, "user logged in successfully", nil)
}

func (ctrl *AuthController) Logout(c *fiber.Ctx) error {
	h := &utils.ResponseHandler{}

	tokenStr := c.Cookies("access_token")

	if tokenStr == "" {
		return h.BadRequest(c, []string{"token not found"})
	}

	utils.DeleteTokenCookie(c)

	return h.Ok(c, nil, "user logged out successfully", nil)
}
