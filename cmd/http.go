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

	api := app.Group("/api")
	api.Get("/", http.ApiListHandler)
	api.Post("/", http.ValidateApiConfig, http.CreateApiHandler)
	api.Put("/", http.ValidateApiConfig, http.UpdateApiHandler)
	api.Get("/:name", http.ApiDetailsHandler)

	slurp := app.Group("/slurp")
	slurp.Post("/:name", http.CreateSlurpHandler)

	log.Fatal(app.Listen(":3000"))
}
