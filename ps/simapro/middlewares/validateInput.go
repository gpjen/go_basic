package middlewares

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/gpjen/simapro/utils"
)

func ValidateRequest(dtoType interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {

		baseType := reflect.TypeOf(dtoType)

		if baseType.Kind() == reflect.Ptr {
			baseType = baseType.Elem()
		}

		dtoPtr := reflect.New(baseType).Interface()

		if err := c.BodyParser(dtoPtr); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Invalid request body",
				"errors":  err.Error(),
			})
		}

		dtoValue := reflect.ValueOf(dtoPtr).Elem()

		if err := utils.ValidateStruct(dtoValue); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"success": false,
				"message": "Validation failed",
				"errors":  err,
			})
		}

		c.Locals(utils.ValidatedDTOKey, dtoValue.Interface())

		return c.Next()
	}
}
