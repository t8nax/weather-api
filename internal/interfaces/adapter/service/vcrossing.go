package service

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/t8nax/weather-api/internal/application/port"
	"github.com/t8nax/weather-api/internal/entity"
	presenters "github.com/t8nax/weather-api/internal/interfaces/presenters/service"
	http_client "github.com/t8nax/weather-api/pkg/http"
)

var (
	ErrInvalidLocation = errors.New("Invalid location parameter")
	ErrInvalidYear     = errors.New("Invalid year. Year must be between 1950 and 2050")
)

type visualCrossingService struct {
	api_key string
	log     logrus.FieldLogger
}

func NewVisualCrossingService(api_key string, log logrus.FieldLogger) port.WeatherService {
	return &visualCrossingService{api_key, log}
}

func (r *visualCrossingService) GetCurrent(location string) (entity.Weather, error) {
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=metric&key=%s", location, r.api_key)
	client := http_client.NewHttpClient(r.log)

	var resp presenters.VCrossingResponse
	err := client.Get(url, &resp)

	if err != nil {
		return entity.Weather{}, checkErr(err)
	}

	if len(resp.Days) == 0 {
		return entity.Weather{}, errors.New("unable to find days in response")
	}

	entity := presenters.FromVCrossingDay(resp.Days[0])
	entity.Time = time.Now()
	entity.Location = location

	return entity, nil
}

func (r *visualCrossingService) GetHourly(location string, date time.Time) ([]entity.Weather, error) {
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s/%s?unitGroup=metric&include=hours&key=%s",
		location,
		date.Format("2006-01-02"),
		r.api_key)

	client := http_client.NewHttpClient(r.log)
	var resp presenters.VCrossingResponse
	err := client.Get(url, &resp)

	if err != nil {
		return nil, checkErr(err)
	}

	if len(resp.Days) == 0 {
		return nil, errors.New("unable to find days in response")
	}

	hours := resp.Days[0].Hours

	if len(hours) == 0 {
		return nil, errors.New("unable to find hours in response")
	}

	weather := make([]entity.Weather, len(hours))

	for i, hour := range hours {
		weather[i], err = presenters.FromVCrossingHour(hour)

		if err != nil {
			return nil, err
		}
	}

	return weather, nil
}

func checkErr(err error) error {
	const invalidLocationMsg = "Bad API Request:Invalid location parameter value."
	const invalidYearMsg = "Bad API Request:Invalid year requested. Years must be between 1950 and 2050"

	var httpErr *http_client.HttpClientError

	if errors.As(err, &httpErr) && httpErr.Status == http.StatusBadRequest && httpErr.Body == invalidLocationMsg {
		if httpErr.Body == invalidLocationMsg {
			return ErrInvalidLocation
		}

		if httpErr.Body == invalidYearMsg {
			return ErrInvalidYear
		}
	}

	return err
}
