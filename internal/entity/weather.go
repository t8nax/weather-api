package entity

import (
	"time"
)

type Weather struct {
	Location   string
	Descripton string
	Time       time.Time
	Temp       int
	TempMax    int
	TempMin    int
	Humidity   int
	Cloudy     int
	Wind       int
}
