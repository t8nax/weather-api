package port

import "github.com/t8nax/weather-api/internal/entity"

type WeatherCache interface {
	Get(location string) (*entity.Weather, error)
	Set(weather *entity.Weather) error
}
