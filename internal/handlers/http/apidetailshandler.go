package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"log"
)

type ApiDetailsHandler struct {
}

func (h ApiDetailsHandler) HandleApiDetails(c *fiber.Ctx) error {
	apiName := utils.CopyString(c.Params("name"))
	log.Println(apiName)
	return nil
}
