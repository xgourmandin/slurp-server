package http

import (
	"github.com/gofiber/fiber/v2"
	"slurp-server/internal/core/ports"
	"slurp-server/internal/core/usecases"
)

type UpdateApiHandler struct {
	Crud usecases.ApiCrud
}

func (h UpdateApiHandler) HandleUpdateApi(c *fiber.Ctx) error {
	body := new(ports.ApiConfiguration)
	c.BodyParser(&body)
	return h.Crud.UpdateApi(*body)
}
