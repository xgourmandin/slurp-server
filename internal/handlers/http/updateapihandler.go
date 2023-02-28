package http

import (
	"github.com/gofiber/fiber/v2"
	"slurp-server/internal/core/ports"
)

func UpdateApiHandler(c *fiber.Ctx) error {
	body := new(ports.ApiConfiguration)
	c.BodyParser(&body)
	return nil
}
