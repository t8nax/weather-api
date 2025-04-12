package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
	_ "github.com/t8nax/weather-api/docs"
	"github.com/t8nax/weather-api/internal/infrastructure/registry"
	"github.com/t8nax/weather-api/internal/infrastructure/router"
	"github.com/t8nax/weather-api/internal/interfaces/adapter/middleware"
)

func main() {
	log := logrus.New()
	log.SetLevel(logrus.TraceLevel)

	api_key, exists := os.LookupEnv("WEATHER_SERVICE_API_KEY")

	if !exists {
		log.Fatal("Could not find weather service API key in environment variables")
	}

	app := fiber.New()
	app.Use(recover.New())
	app.Use(middleware.LogMiddleware(log))
	app.Use(middleware.ErrorMiddleware(log))

	registry := registry.NewRegistry(api_key, log)
	app = router.NewRouter(app, registry.NewAppController())

	err := app.Listen(":8080")

	if err != nil {
		log.Fatal(err)
	}
}
