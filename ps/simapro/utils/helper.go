package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gpjen/simapro/app/dto"
)

func StringToUint(str string) uint {
	i, _ := strconv.ParseUint(str, 10, 0)
	return uint(i)
}

func GetUserSession(c *fiber.Ctx) dto.UserContext {
	userLocal := c.Locals("user")

	if userLocal == nil {
		return dto.UserContext{
			Email: "Unknown",
		}
	}

	userDTO, ok := userLocal.(*dto.UserContext)
	if !ok {
		return dto.UserContext{
			ID:    0,
			Email: "Unknown",
		}
	}

	return *userDTO
}

func StringToBool(str string) bool {
	b, _ := strconv.ParseBool(str)
	return b
}
