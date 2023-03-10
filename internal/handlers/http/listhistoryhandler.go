package http

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"slurp-server/internal/core/ports"
)

type ListHistoryHandler struct {
	Crud ports.HistoryCrud
}

func (h ListHistoryHandler) HandleListHistories(c *fiber.Ctx) error {
	histories, err := h.Crud.ListHistory()
	if err != nil {
		log.Println(err)
		c.Status(fiber.StatusInternalServerError).JSON(ApiError{Message: err.Error()})
	}
	if len(*histories) == 0 {
		return c.Status(200).JSON([]string{})
	} else {
		return c.Status(200).JSON(histories)
	}
}
