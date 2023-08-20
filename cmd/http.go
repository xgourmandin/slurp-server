package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"os"
	"slurp-server/internal/core/ports"
	"slurp-server/internal/core/usecases"
	"slurp-server/internal/handlers/http"
	"slurp-server/internal/handlers/repositories"
)

func main() {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())

	var apiCrud usecases.ApiCrud
	if os.Getenv("STORAGE_TYPE") == "firestore" {
		apiCrud = usecases.ApiCrud{Repo: repositories.NewFirestoreRepository(os.Getenv("PROJECT_ID"), getEnv("COLLECTION_PREFIX", "slurp-"))}
	} else {
		apiCrud = usecases.ApiCrud{Repo: repositories.NewInMemoryRepository()}
	}

	apiListHandler := http.ApiListHandler{Crud: apiCrud}
	createApiHandler := http.CreateApiHandler{Crud: apiCrud}
	updateApiHandler := http.UpdateApiHandler{Crud: apiCrud}
	deleteApiHandler := http.DeleteApiHandler{Crud: apiCrud}
	apiDetailHandler := http.ApiDetailsHandler{Crud: apiCrud}
	apiPauseHandler := http.PauseSlurpHandler{Crud: apiCrud}

	api := app.Group("/api")
	api.Get("/", apiListHandler.HandleApiList)
	api.Post("/", http.ValidateApiConfig, createApiHandler.HandleCreateApi)
	api.Put("/:name", http.ValidateApiConfig, updateApiHandler.HandleUpdateApi)
	api.Get("/:name", apiDetailHandler.HandleApiDetails)
	api.Delete("/:name", deleteApiHandler.HandleDeleteApi)
	api.Post("/:name/pause", apiPauseHandler.HandlePauseSlurp)
	api.Post("/:name/unpause", apiPauseHandler.HandleUnpauseSlurp)

	var historyCrud ports.HistoryCrud
	var historyRepository ports.ApiHistoryRepository
	if os.Getenv("STORAGE_TYPE") == "firestore" {
		historyRepository = repositories.NewFirestoreHistoryRepository(os.Getenv("PROJECT_ID"), getEnv("COLLECTION_PREFIX", "slurp-"))
		historyCrud = usecases.HistoryCrud{Repo: historyRepository}
	} else {
		historyRepository = repositories.NewInMemoryHistoryRepository()
		historyCrud = usecases.HistoryCrud{Repo: historyRepository}
	}

	listHistoryHandler := http.ListHistoryHandler{Crud: historyCrud}

	history := app.Group("/history")
	history.Get("/", listHistoryHandler.HandleListHistories)

	createSlurpUsecase := usecases.SlurpUseCase{
		ApiCrud:    apiCrud,
		ApiHistory: historyRepository,
	}

	createSlurpHandler := http.CreateSlurpHandler{SlurpUc: createSlurpUsecase}

	slurp := app.Group("/slurp")
	slurp.Post("/:name", createSlurpHandler.HandleCreateSlurp)

	log.Fatal(app.Listen(":3000"))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
