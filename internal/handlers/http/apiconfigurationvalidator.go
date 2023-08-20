package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"slurp-server/internal/core/ports"
)

type IError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

var Validator = validator.New()

func ValidateApiConfig(c *fiber.Ctx) error {
	var errors []*IError
	body := new(ports.ApiConfiguration)
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	err = Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	return c.Next()
}
