package usecase

import (
	"time"

	"github.com/t8nax/weather-api/internal/application/port"
	"github.com/t8nax/weather-api/internal/entity"
)

type WeatherUCase interface {
	GetCurrent(location string) (entity.Weather, error)
	GetHourly(location string, day time.Time) ([]entity.Weather, error)
	GetDaily(location string, date time.Time) (entity.Weather, error)
}

type weatherUCase struct {
	service port.WeatherService
	cache   port.WeatherCache
}

func NewWeather(service port.WeatherService, cache port.WeatherCache) WeatherUCase {
	return &weatherUCase{service, cache}
}

func (uc *weatherUCase) GetCurrent(location string) (entity.Weather, error) {
	weather, err := uc.service.GetCurrent(location)

	if err != nil {
		return entity.Weather{}, err
	}

	return weather, nil
}

func (uc *weatherUCase) GetHourly(location string, date time.Time) ([]entity.Weather, error) {
	weather, err := uc.service.GetHourly(location, date)

	if err != nil {
		return nil, err
	}

	return weather, nil
}

func (uc *weatherUCase) GetDaily(location string, date time.Time) (entity.Weather, error) {
	weather, err := uc.service.GetDaily(location, date)

	if err != nil {
		return entity.Weather{}, err
	}

	return weather, nil
}
