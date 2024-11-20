package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
	Index int    `json:"index,omitempty"`
}

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	validate := validator.New()

	// Custom tag untuk validasi khusus
	validate.RegisterValidation("unique", uniqueValidator)
	validate.RegisterValidation("uppercase", uppercaseValidator)

	return &Validator{validate: validate}
}

// ValidateStruct melakukan validasi struktur dengan detail error
func (v *Validator) ValidateStruct(s interface{}) []ValidationError {
	var errors []ValidationError

	err := v.validate.Struct(s)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrors {
				errors = append(errors, processValidationError(e, s)...)
			}
		}
	}

	return errors
}

// processValidationError menghasilkan detail error yang komprehensif
func processValidationError(e validator.FieldError, s interface{}) []ValidationError {
	var errors []ValidationError

	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	field, _ := val.Type().FieldByName(e.StructField())
	jsonTag := field.Tag.Get("json")
	if jsonTag == "" {
		jsonTag = strings.ToLower(e.StructField())
	}

	// Untuk slice/array, berikan informasi index
	if strings.Contains(e.Kind().String(), "slice") || strings.Contains(e.Kind().String(), "array") {
		sliceValue := reflect.ValueOf(e.Value())
		for i := 0; i < sliceValue.Len(); i++ {
			errors = append(errors, ValidationError{
				Field: jsonTag,
				Error: fmt.Sprintf("%s at index %d", e.Translate(nil), i),
				Index: i,
			})
		}
	} else {
		errors = append(errors, ValidationError{
			Field: jsonTag,
			Error: e.Translate(nil),
		})
	}

	return errors
}

// Validator kustom untuk memastikan keunikan dalam slice
func uniqueValidator(fl validator.FieldLevel) bool {
	field := fl.Field()
	if field.Kind() != reflect.Slice {
		return true
	}

	unique := make(map[interface{}]bool)
	for i := 0; i < field.Len(); i++ {
		val := field.Index(i).Interface()
		if unique[val] {
			return false
		}
		unique[val] = true
	}
	return true
}

// Validator kustom untuk memastikan huruf besar
func uppercaseValidator(fl validator.FieldLevel) bool {
	str, ok := fl.Field().Interface().(string)
	if !ok {
		return true
	}
	return str == strings.ToUpper(str)
}

// Middleware Fiber untuk validasi
func (v *Validator) FiberValidate(s interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := c.BodyParser(s); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		validationErrors := v.ValidateStruct(s)
		if len(validationErrors) > 0 {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"errors": validationErrors,
			})
		}

		c.Locals("validated_data", s)
		return c.Next()
	}
}
