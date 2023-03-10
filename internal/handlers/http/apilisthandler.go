package http

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"slurp-server/internal/core/usecases"
)

type ApiListHandler struct {
	Crud usecases.ApiCrud
}

func (h ApiListHandler) HandleApiList(c *fiber.Ctx) error {
	apis, err := h.Crud.ListApi()
	if err != nil {
		log.Println(err)
		c.Status(fiber.StatusInternalServerError).JSON(ApiError{Message: err.Error()})
	}
	if len(*apis) == 0 {
		return c.Status(200).JSON([]string{})
	} else {
		return c.Status(200).JSON(apis)
	}

}
