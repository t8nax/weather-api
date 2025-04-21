package controller

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/t8nax/weather-api/internal/application/usecase"
	"github.com/t8nax/weather-api/internal/common"
	"github.com/t8nax/weather-api/internal/interfaces/presenters/api"
)

type WeatherController interface {
	GetCurrentWeather() fiber.Handler
	GetHourlyWeather() fiber.Handler
	GetDailyWeather() fiber.Handler
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
			return common.NewAppError(common.CodeInvalidLocation, errors.New("invalid location from request"))
		}

		weather, err := c.uCase.GetCurrent(location)

		if err != nil {
			return err
		}

		response := api.ToCurrentWeather(weather)

		return ctx.JSON(response)
	}
}

func (c *weatherController) GetHourlyWeather() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		location := ctx.Query("location")

		if err := checkLocation(location); err != nil {
			return err
		}

		date, err := parseDate(ctx.Query("date"))

		if err != nil {
			return err
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

func (c *weatherController) GetDailyWeather() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		location := ctx.Query("location")

		if err := checkLocation(location); err != nil {
			return err
		}

		date, err := parseDate(ctx.Query("date"))

		if err != nil {
			return err
		}

		weather, err := c.uCase.GetDaily(location, date)

		if err != nil {
			return err
		}

		response := api.ToDailyWeather(weather)

		return ctx.JSON(response)
	}
}

func checkLocation(location string) error {
	if location == "" {
		return common.NewAppError(common.CodeInvalidLocation, fmt.Errorf("invalid location: %s", location))
	}

	return nil
}

func parseDate(dateStr string) (time.Time, error) {
	date, err := time.Parse("2006-01-02", dateStr)

	if err != nil {
		return time.Time{}, common.NewAppError(common.CodeInvalidDate, errors.New("invalid year from request"))
	}

	return date, nil
}
