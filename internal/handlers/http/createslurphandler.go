package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"log"
)

type CreateSlurpHandler struct {
}

func (h CreateSlurpHandler) HandleCreateSlurp(c *fiber.Ctx) error {
	apiName := utils.CopyString(c.Params("name"))
	log.Println(apiName)
	return nil
}
