package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"slurp-server/internal/core/ports"
	"slurp-server/internal/core/usecases"
)

type UpdateApiHandler struct {
	Crud usecases.ApiCrud
}

func (h UpdateApiHandler) HandleUpdateApi(c *fiber.Ctx) error {
	apiName := utils.CopyString(c.Params("name"))
	body := new(ports.ApiConfiguration)
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}
	return h.Crud.UpdateApi(*body, apiName)
}
