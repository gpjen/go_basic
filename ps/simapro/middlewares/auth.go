package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gpjen/simapro/app/dto"
	"github.com/gpjen/simapro/utils"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		r := utils.ResponseHandler{}

		validToken, ok := c.Locals("valid_token").(bool)
		if !ok || !validToken {
			userCtx, ok := c.Locals("user").(*dto.UserContext)
			if !ok || userCtx.ID == 0 {
				return r.Unauthorized(c, []string{"Unauthorized: You must be logged in to access this resource"})
			}
		}

		return c.Next()
	}
}

func SetupUserContext(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		userCtx := &dto.UserContext{
			ID:           0,
			Email:        "Unknown",
			SelectedRole: 0,
			Roles:        []uint{},
		}

		token := c.Cookies("access_token")

		if token == "" {
			authHeader := c.Get("Authorization")
			if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
				token = strings.TrimSpace(authHeader[7:])
			}
		}

		if token != "" {
			verifyToken, err := utils.VerifyToken(token)
			if err == nil && verifyToken.Valid {
				if claims, ok := verifyToken.Claims.(jwt.MapClaims); ok {
					idFloat, idExists := claims["id"].(float64)
					email, emailExists := claims["email"].(string)
					roleFloat, roleExists := claims["selected_role"].(float64)
					rolesInterface, rolesExists := claims["roles"].([]interface{})

					if idExists && emailExists && roleExists && rolesExists {
						roles := make([]uint, 0)
						for _, role := range rolesInterface {
							if roleFloat, ok := role.(float64); ok {
								roles = append(roles, uint(roleFloat))
							}
						}

						userCtx.ID = uint(idFloat)
						userCtx.Email = email
						userCtx.SelectedRole = uint(roleFloat)
						userCtx.Roles = roles

						c.Locals("valid_token", true)
					}
				}
			}
		}

		c.Locals("user", userCtx)
		return c.Next()
	})
}
