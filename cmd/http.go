package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"slurp-server/internal/core/usecases"
	"slurp-server/internal/handlers/http"
	"slurp-server/internal/handlers/repositories"
)

func main() {
	app := fiber.New()
	app.Use(recover.New())

	apiCrud := usecases.ApiCrud{Repo: repositories.NewInMemoryRepository()}

	apiListHandler := http.ApiListHandler{Crud: apiCrud}
	createApiHandler := http.CreateApiHandler{Crud: apiCrud}
	updateApiHandler := http.UpdateApiHandler{Crud: apiCrud}
	apiDetailHandler := http.ApiDetailsHandler{Crud: apiCrud}

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
