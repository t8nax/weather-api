package service

import (
	"fmt"
	"math"
	"time"

	"github.com/t8nax/weather-api/internal/entity"
)

type VCrossingResponse struct {
	Days              []VCrossingDay             `json:"days"`
	CurrentConditions VCrossingCurrentConditions `json:"currentConditions"`
}

type VCrossingDay struct {
	Descripton string          `json:"description"`
	Temp       float64         `json:"temp"`
	TempMax    float64         `json:"tempmax"`
	TempMin    float64         `json:"tempmin"`
	Humidity   float64         `json:"humidity"`
	CloudCover float64         `json:"cloudcover"`
	WindSpeed  float64         `json:"windspeed"`
	Hours      []VCrossingHour `json:"hours"`
}

type VCrossingHour struct {
	Time       string  `json:"datetime"`
	Temp       float64 `json:"temp"`
	Humidity   float64 `json:"humidity"`
	CloudCover float64 `json:"cloudcover"`
	WindSpeed  float64 `json:"windspeed"`
}

type VCrossingCurrentConditions struct {
	Timestamp  int64   `json:"datetimeEpoch"`
	Temp       float64 `json:"temp"`
	Humidity   float64 `json:"humidity"`
	CloudCover float64 `json:"cloudcover"`
	WindSpeed  float64 `json:"windspeed"`
}

func FromVCrossingDay(day VCrossingDay) entity.Weather {
	return entity.Weather{
		Descripton: day.Descripton,
		Temp:       int(math.Round(day.Temp)),
		TempMax:    int(math.Round(day.TempMax)),
		TempMin:    int(math.Round(day.TempMin)),
		Humidity:   int(math.Round(day.Humidity)),
		Cloudy:     int(math.Round(day.CloudCover)),
		Wind:       int(math.Round(day.WindSpeed)),
	}
}

func FromVCrossingHour(hour VCrossingHour) (entity.Weather, error) {
	t, err := time.Parse("15:04:05", hour.Time)

	if err != nil {
		return entity.Weather{}, fmt.Errorf("unable to parse time: %s", hour.Time)
	}

	return entity.Weather{
		Temp:     int(math.Round(hour.Temp)),
		Humidity: int(math.Round(hour.Humidity)),
		Cloudy:   int(math.Round(hour.CloudCover)),
		Wind:     int(math.Round(hour.WindSpeed)),
		DateTime: t,
	}, nil
}

func FromVCrossingCurrentConditions(conditions VCrossingCurrentConditions) entity.Weather {
	return entity.Weather{
		Temp:     int(math.Round(conditions.Temp)),
		Humidity: int(math.Round(conditions.Humidity)),
		Cloudy:   int(math.Round(conditions.CloudCover)),
		Wind:     int(math.Round(conditions.WindSpeed)),
		DateTime: time.Unix(conditions.Timestamp, 0),
	}
}
