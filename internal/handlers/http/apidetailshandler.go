package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"log"
	"slurp-server/internal/core/usecases"
)

type ApiDetailsHandler struct {
	Crud usecases.ApiCrud
}

func (h ApiDetailsHandler) HandleApiDetails(c *fiber.Ctx) error {
	apiName := utils.CopyString(c.Params("name"))
	api, err := h.Crud.GetApi(apiName)
	if err != nil {
		log.Println(err)
		c.Status(fiber.StatusInternalServerError).JSON(ApiError{Message: err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(api)
}
