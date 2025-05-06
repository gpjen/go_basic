package utils

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/iancoleman/strcase"
)

var ValidatedDTOKey = "validatedDTO"

var validate *validator.Validate

func init() {
	validate = validator.New()
	registerCustomValidators()
}

func registerCustomValidators() {
	validate.RegisterValidation("ne_field", validateNotEqualToField)
}

func validateNotEqualToField(fl validator.FieldLevel) bool {
	fieldName := fl.Param()

	parent := fl.Parent()

	fieldToCompare := parent.FieldByName(fieldName)

	if !fieldToCompare.IsValid() {
		return true
	}

	currentField := fl.Field()

	if currentField.Kind() == reflect.Ptr && currentField.IsNil() {
		return true
	}

	var currentValue interface{}
	var compareValue interface{}

	if currentField.Kind() == reflect.Ptr {
		currentValue = currentField.Elem().Interface()
	} else {
		currentValue = currentField.Interface()
	}

	if fieldToCompare.Kind() == reflect.Ptr {
		if fieldToCompare.IsNil() {
			return true
		}
		compareValue = fieldToCompare.Elem().Interface()
	} else {
		compareValue = fieldToCompare.Interface()
	}

	return !reflect.DeepEqual(currentValue, compareValue)
}

func ValidateStruct(v reflect.Value) interface{} {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() == reflect.Struct {
		if err := validate.Struct(v.Interface()); err != nil {
			return processValidationErrors(err)
		}
		return nil
	}

	if v.Kind() == reflect.Slice {
		var allErrors []interface{}

		for i := 0; i < v.Len(); i++ {
			elemV := v.Index(i)

			if err := validate.Struct(elemV.Interface()); err != nil {
				elemErrors := processValidationErrors(err)
				allErrors = append(allErrors, fiber.Map{
					"index":  i,
					"errors": elemErrors,
				})
			}
		}

		if len(allErrors) > 0 {
			return allErrors
		}
		return nil
	}

	return nil
}

func processValidationErrors(err error) []fiber.Map {
	var validationErrors validator.ValidationErrors
	if !errors.As(err, &validationErrors) {
		return []fiber.Map{
			{
				"message": err.Error(),
			},
		}
	}

	var errors []fiber.Map
	for _, e := range validationErrors {
		var errorMessage string

		snakeField := strcase.ToSnake(e.Field())
		snakeParam := strcase.ToSnake(e.Param())

		switch e.Tag() {
		case "ne_field":
			errorMessage = fmt.Sprintf("%s must not be equal to %s", snakeParam, snakeField)
			errors = append(errors, fiber.Map{
				"field":   snakeParam,
				"value":   e.Value(),
				"tag":     e.Tag(),
				"message": errorMessage,
			})

			continue
		default:
			errorMessage = fmt.Sprintf("%s is invalid", snakeField)
		}

		errors = append(errors, fiber.Map{
			"field":   snakeField,
			"value":   e.Value(),
			"tag":     e.Tag(),
			"message": errorMessage,
		})
	}

	return errors
}

func GetValidatedDTO(c *fiber.Ctx) interface{} {
	dto := c.Locals(ValidatedDTOKey)
	return dto
}
