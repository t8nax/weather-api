package api

import (
	"github.com/t8nax/weather-api/internal/entity"
)

type CurrentWeather struct {
	Temp     int    `json:"temp"`
	Humidity int    `json:"humidity"`
	Cloudy   int    `json:"clody"`
	Wind     int    `json:"wind"`
	DateTime string `json:"datetime"`
}

func ToCurrentWeather(weather entity.Weather) CurrentWeather {
	return CurrentWeather{
		Temp:     weather.Temp,
		Humidity: weather.Humidity,
		Cloudy:   weather.Cloudy,
		Wind:     weather.Wind,
		DateTime: weather.DateTime.Format("2006-01-02 15:04:05"),
	}
}

type HourlyWeather struct {
	Time     string `json:"time"`
	Temp     int    `json:"temp"`
	Humidity int    `json:"humidity"`
	Cloudy   int    `json:"clody"`
	Wind     int    `json:"wind"`
}

func ToHourlyWeather(weather entity.Weather) HourlyWeather {
	return HourlyWeather{
		Time:     weather.DateTime.Format("15:04:05"),
		Temp:     weather.Temp,
		Humidity: weather.Humidity,
		Cloudy:   weather.Cloudy,
		Wind:     weather.Wind,
	}
}

type DailyWeather struct {
	Description string `json:"description"`
	Temp        int    `json:"temp"`
	TempMax     int    `json:"tempMax"`
	TempMin     int    `json:"tempMin"`
	Humidity    int    `json:"humidity"`
	Cloudy      int    `json:"clody"`
	Wind        int    `json:"wind"`
}

func ToDailyWeather(weather entity.Weather) DailyWeather {
	return DailyWeather{
		Description: weather.Description,
		Temp:        weather.Temp,
		TempMax:     weather.TempMax,
		TempMin:     weather.TempMin,
		Humidity:    weather.Humidity,
		Cloudy:      weather.Cloudy,
		Wind:        weather.Wind,
	}
}
