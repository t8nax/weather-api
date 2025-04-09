package port

import (
	"time"

	"github.com/t8nax/weather-api/internal/entity"
)

type WeatherService interface {
	GetCurrent(location string) (entity.Weather, error)
	GetHourly(location string, date time.Time) ([]entity.Weather, error)
}
