package usecase

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/t8nax/weather-api/internal/application/port"
	"github.com/t8nax/weather-api/internal/entity"
)

type WeatherUCase interface {
	GetCurrent(location string) (entity.Weather, error)
	GetHourly(location string, day time.Time) ([]entity.Weather, error)
}

type weatherUCase struct {
	service port.WeatherService
	cache   port.WeatherCache
	log     logrus.FieldLogger
}

func NewWeather(service port.WeatherService, cache port.WeatherCache, log logrus.FieldLogger) WeatherUCase {
	return &weatherUCase{service, cache, log}
}

func (uc *weatherUCase) GetCurrent(location string) (entity.Weather, error) {
	weather, err := uc.service.GetCurrent(location)

	if err != nil {
		uc.log.WithField("location", location).WithError(err).Error("Unable to get weather")
		return entity.Weather{}, err
	}

	return weather, nil
}

func (uc *weatherUCase) GetHourly(location string, day time.Time) ([]entity.Weather, error) {
	weather, err := uc.service.GetHourly(location, day)

	if err != nil {
		uc.log.WithFields(logrus.Fields{
			"location": location,
			"day":      day,
		}).WithError(err).Error("Unable to get weather")

		return nil, err
	}

	return weather, nil
}
