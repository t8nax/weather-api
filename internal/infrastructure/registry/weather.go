package registry

import (
	"github.com/t8nax/weather-api/internal/application/usecase"
	"github.com/t8nax/weather-api/internal/interfaces/adapter/controller"
	"github.com/t8nax/weather-api/internal/interfaces/adapter/service"
)

func (r *registry) NewWeatherController() controller.WeatherController {
	api := service.NewVisualCrossingService(r.api_key, r.log)
	srv := usecase.NewWeather(api, nil, r.log)

	return controller.NewWeatherController(srv, r.log)
}
