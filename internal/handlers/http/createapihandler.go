package http

import (
	"github.com/gofiber/fiber/v2"
	"slurp-server/internal/core/ports"
	"slurp-server/internal/core/usecases"
)

type CreateApiHandler struct {
	Crud usecases.ApiCrud
}

func (h CreateApiHandler) HandleCreateApi(c *fiber.Ctx) error {
	body := new(ports.ApiConfiguration)
	c.BodyParser(&body)
	return h.Crud.CreateApi(*body)
}
