package entity

import (
	"time"
)

type Weather struct {
	Location    string
	Description string
	DateTime    time.Time
	Temp        int
	TempMax     int
	TempMin     int
	Humidity    int
	Cloudy      int
	Wind        int
}
