package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/t8nax/weather-api/internal/application/usecase"
	"github.com/t8nax/weather-api/internal/interfaces/presenters/api"
)

type WeatherController interface {
	GetCurrentWeather() fiber.Handler
	GetHourlyWeather() fiber.Handler
}

type weatherController struct {
	uCase usecase.WeatherUCase
	log   logrus.FieldLogger
}

func NewWeatherController(service usecase.WeatherUCase, log logrus.FieldLogger) WeatherController {
	return &weatherController{service, log}
}

func (c *weatherController) GetCurrentWeather() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		location := ctx.Query("location")

		if location == "" {
			return api.ErrInvalidLocation
		}

		weather, err := c.uCase.GetCurrent(location)

		if err != nil {
			return err
		}

		response := api.ToDailyWeather(weather)

		return ctx.JSON(response)
	}
}

func (c *weatherController) GetHourlyWeather() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		location := ctx.Query("location")

		if location == "" {
			return api.ErrInvalidLocation
		}

		dateParam := ctx.Query("date")
		date, err := time.Parse("2006-01-02", dateParam)

		if err != nil {
			return api.ErrInvalidDate
		}

		weathers, err := c.uCase.GetHourly(location, date)

		if err != nil {
			return err
		}

		response := make([]api.HourlyWeather, len(weathers))

		for i, weather := range weathers {
			response[i] = api.ToHourlyWeather(weather)
		}

		return ctx.JSON(response)
	}
}
