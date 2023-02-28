package http

import (
	"github.com/gofiber/fiber/v2"
	"slurp-server/internal/core/ports"
)

type CreateApiHandler struct {
}

func (h CreateApiHandler) HandleCreateApi(c *fiber.Ctx) error {
	body := new(ports.ApiConfiguration)
	c.BodyParser(&body)
	return nil
}
