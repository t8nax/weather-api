package service

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/t8nax/weather-api/internal/application/port"
	"github.com/t8nax/weather-api/internal/common"
	"github.com/t8nax/weather-api/internal/entity"
	presenters "github.com/t8nax/weather-api/internal/interfaces/presenters/service"
	http_client "github.com/t8nax/weather-api/pkg/http"
)

type visualCrossingService struct {
	api_key string
	log     logrus.FieldLogger
}

func NewVisualCrossingService(api_key string, log logrus.FieldLogger) port.WeatherService {
	return &visualCrossingService{api_key, log}
}

func (vcs *visualCrossingService) GetCurrent(location string) (entity.Weather, error) {
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=metric&key=%s", location, vcs.api_key)
	client := http_client.NewHttpClient(vcs.log)

	var resp presenters.VCrossingResponse
	err := client.Get(url, &resp)

	if err != nil {
		return entity.Weather{}, convertErr(err)
	}

	weather := presenters.FromVCrossingCurrentConditions(resp.CurrentConditions)
	weather.Location = location

	return weather, nil
}

func (vcs *visualCrossingService) GetHourly(location string, date time.Time) ([]entity.Weather, error) {
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s/%s?unitGroup=metric&include=hours&key=%s",
		location,
		date.Format("2006-01-02"),
		vcs.api_key)

	client := http_client.NewHttpClient(vcs.log)
	var resp presenters.VCrossingResponse
	err := client.Get(url, &resp)

	if err != nil {
		vcs.log.WithFields(logrus.Fields{
			"location": location,
			"date":     date.Format("2006-01-02"),
		}).WithError(err).Error("Failed to get hourly weather")

		return nil, convertErr(err)
	}

	if len(resp.Days) == 0 {
		return nil, common.NewAppError(common.CodeServiceError, errors.New("unable to find days in response"))
	}

	hours := resp.Days[0].Hours

	if len(hours) == 0 {
		return nil, common.NewAppError(common.CodeServiceError, errors.New("unable to find hours in response"))
	}

	weather := make([]entity.Weather, len(hours))

	for i, hour := range hours {
		weather[i], err = presenters.FromVCrossingHour(hour)

		if err != nil {
			return nil, common.NewAppError(common.CodeServiceError, err)
		}
	}

	return weather, nil
}

func (vcs *visualCrossingService) GetDaily(location string, date time.Time) (entity.Weather, error) {
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s/%s?unitGroup=metric&key=%s",
		location,
		date.Format("2006-01-02"),
		vcs.api_key)

	client := http_client.NewHttpClient(vcs.log)
	var resp presenters.VCrossingResponse
	err := client.Get(url, &resp)

	if err != nil {
		vcs.log.WithFields(logrus.Fields{
			"location": location,
			"date":     date.Format("2006-01-02"),
		}).WithError(err).Error("Failed to get daily weather")

		return entity.Weather{}, convertErr(err)
	}

	if len(resp.Days) == 0 {
		return entity.Weather{}, common.NewAppError(common.CodeServiceError, errors.New("unable to find days in response"))
	}

	return presenters.FromVCrossingDay(resp.Days[0]), nil
}

func convertErr(err error) error {
	const invalidLocationMsg = "Bad API Request:Invalid location parameter value."
	const invalidYearMsg = "Bad API Request:Invalid year requested. Years must be between 1950 and 2050"

	var httpErr *http_client.HttpClientError

	if errors.As(err, &httpErr) && httpErr.Status == http.StatusBadRequest && httpErr.Body == invalidLocationMsg {
		if httpErr.Body == invalidLocationMsg {
			return common.NewAppError(common.CodeInvalidLocation, errors.New("invalid location parameter"))
		}

		if httpErr.Body == invalidYearMsg {
			return common.NewAppError(common.CodeInvalidDate, errors.New("invalid date"))
		}
	}

	return err
}
