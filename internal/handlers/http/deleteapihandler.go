package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"slurp-server/internal/core/usecases"
)

type DeleteApiHandler struct {
	Crud usecases.ApiCrud
}

func (h DeleteApiHandler) HandleDeleteApi(c *fiber.Ctx) error {
	apiName := utils.CopyString(c.Params("name"))
	return h.Crud.DeleteApi(apiName)
}
