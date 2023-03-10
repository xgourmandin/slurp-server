package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"slurp-server/internal/core/ports"
)

type PauseSlurpHandler struct {
	Crud ports.ApiCrud
}

func (h PauseSlurpHandler) HandlePauseSlurp(c *fiber.Ctx) error {
	apiName := utils.CopyString(c.Params("name"))
	err := h.Crud.PauseApi(apiName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ApiError{Message: err.Error()})
	}
	return nil
}

func (h PauseSlurpHandler) HandleUnpauseSlurp(c *fiber.Ctx) error {
	apiName := utils.CopyString(c.Params("name"))
	err := h.Crud.UnPauseApi(apiName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ApiError{Message: err.Error()})
	}
	return nil
}
