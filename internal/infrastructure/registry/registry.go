package registry

import (
	"github.com/sirupsen/logrus"
	"github.com/t8nax/weather-api/internal/interfaces/adapter/controller"
)

type Registry interface {
	NewAppController() *controller.AppController
}

type registry struct {
	api_key string
	log logrus.FieldLogger
}

func NewRegistry(api_key string, log logrus.FieldLogger) Registry {
    return &registry{api_key, log}
}

func (r *registry) NewAppController() *controller.AppController {
	return &controller.AppController{
		WeatherController: r.NewWeatherController(),
	}
}
