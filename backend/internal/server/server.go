package server

import (
	"github.com/GenerateNU/cooked/backend/internal/errs"
	"github.com/GenerateNU/cooked/backend/internal/server/handlers"
	"github.com/GenerateNU/cooked/backend/internal/storage"

	go_json "github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Params struct {
	Storage storage.Storage
}

func InitApp(params Params) *fiber.App {
	app := setupApp()
	service := handlers.NewService(params.Storage)

	setupRoutes(app, service)

	return app
}

func setupRoutes(app *fiber.App, service *handlers.Service) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
	app.Get("/hello/:name", service.Hello)

	// TODO: register your routes here
}

func setupApp() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:  go_json.Marshal,
		JSONDecoder:  go_json.Unmarshal,
		ErrorHandler: errs.ErrorHandler,
	})
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}:${port} ${pid} ${locals:requestid} ${status} - ${latency} ${method} ${path}\n",
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	return app
}
