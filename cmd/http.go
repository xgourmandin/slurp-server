package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"slurp-server/internal/handlers/http"
)

func main() {
	app := fiber.New()
	app.Use(recover.New())

	apiListHandler := http.ApiListHandler{}
	createApiHandler := http.CreateApiHandler{}
	updateApiHandler := http.UpdateApiHandler{}
	apiDetailHandler := http.ApiDetailsHandler{}

	api := app.Group("/api")
	api.Get("/", apiListHandler.HandleApiList)
	api.Post("/", http.ValidateApiConfig, createApiHandler.HandleCreateApi)
	api.Put("/", http.ValidateApiConfig, updateApiHandler.HandleUpdateApi)
	api.Get("/:name", apiDetailHandler.HandleApiDetails)

	createSlurpHandler := http.CreateSlurpHandler{}

	slurp := app.Group("/slurp")
	slurp.Post("/:name", createSlurpHandler.HandleCreateSlurp)

	log.Fatal(app.Listen(":3000"))
}
