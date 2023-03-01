package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"slurp-server/internal/core/usecases"
)

type CreateSlurpHandler struct {
	SlurpUc usecases.SlurpUseCase
}

func (h CreateSlurpHandler) HandleCreateSlurp(c *fiber.Ctx) error {
	apiName := utils.CopyString(c.Params("name"))
	err := h.SlurpUc.CreateSlurp(apiName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ApiError{Message: err.Error()})
	}
	return nil
}
