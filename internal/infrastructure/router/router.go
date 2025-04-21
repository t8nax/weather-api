package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/t8nax/weather-api/docs"
	"github.com/t8nax/weather-api/internal/interfaces/adapter/controller"
)

func NewRouter(app *fiber.App, ctl *controller.AppController) *fiber.App {
	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")
	GetCurrentWeather(api, ctl)
	GetHourlyWeather(api, ctl)
	GetDailyWeather(api, ctl)

	return app
}

// GetWeatherInfo Getting Weather for now
//
// @Summary        Getting Weather Info for now
// @Tags            Weather
// @Accept            json
// @Produce        json
// @Param            location  query  string    true    "Location (ex. 'Moscow')"
// @Success        200                {string}    string
// @Router            /api/weather/current [get]
func GetCurrentWeather(api fiber.Router, ctl *controller.AppController) fiber.Router {
	return api.Get("/weather/current", ctl.WeatherController.GetCurrentWeather())
}

// GetWeatherInfo Getting Hourly Weather Info
//
// @Summary        Getting Hourly Weather Info
// @Tags           Weather
// @Accept         json
// @Produce        json
// @Param          location  query  string    true    "Location (ex. 'Moscow')"
// @Param          date  query  string    true    "Date (ex. '2000-05-30')"
// @Success        200                {string}    string
// @Router         /api/weather/hourly [get]
func GetHourlyWeather(api fiber.Router, ctl *controller.AppController) fiber.Router {
	return api.Get("/weather/dailt", ctl.WeatherController.GetHourlyWeather())
}

// GetWeatherInfo Getting Daily Weather Info
//
// @Summary        Getting Daily Weather Info
// @Tags           Weather
// @Accept         json
// @Produce        json
// @Param          location  query  string    true    "Location (ex. 'Moscow')"
// @Param          date  query  string    true    "Date (ex. '2000-05-30')"
// @Success        200                {string}    string
// @Router         /api/weather/daily [get]
func GetDailyWeather(api fiber.Router, ctl *controller.AppController) fiber.Router {
	return api.Get("/weather/daily", ctl.WeatherController.GetDailyWeather())
}