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
	externalSrv port.WeatherService
	cache       port.WeatherCache
	log         logrus.FieldLogger
}

func NewWeather(externalSrv port.WeatherService, cache port.WeatherCache, log logrus.FieldLogger) WeatherUCase {
	return &weatherUCase{externalSrv, cache, log}
}

func (uc weatherUCase) GetCurrent(location string) (entity.Weather, error) {
	// weather, err := s.cache.Get(location)

	// if err != nil {
	// 	s.log.WithField("location", location).
	// 		WithError(err).
	// 		Warn("Failed to get weather from cache")
	// } else if weather == nil {
	// 	s.log.WithField("location", location).Debug("Weather not found in cache")
	// } else {
	// 	return weather, nil
	// }

	weather, err := uc.externalSrv.GetCurrent(location)

	if err != nil {
		uc.log.WithField("location", location).WithError(err).Error("Unable to get weather")
		return entity.Weather{}, err
	}

	// if err = s.cache.Set(weather); err != nil {
	// 	s.log.WithField("location", location).
	// 		WithError(err).
	// 		Warn("Failed to set weather in cache")
	// }

	return weather, err
}

func (uc weatherUCase) GetHourly(location string, day time.Time) ([]entity.Weather, error) {
	weather, err := uc.externalSrv.GetHourly(location, day)

	if err != nil {
		uc.log.WithFields(logrus.Fields{
			"location": location,
			"day":     day,
		}).WithError(err).Error("Unable to get weather")

		return nil, err
	}

	return weather, err
}
