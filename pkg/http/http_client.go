package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

type httpClient struct {
	log logrus.FieldLogger
}

func NewHttpClient(log logrus.FieldLogger) *httpClient {
	return &httpClient{log}
}

func (c *httpClient) Get(url string, entity any) error {
	c.log.WithField("url", url).Debug("Sending GET request")

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching info: %w", err)
	}

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return &HttpClientError{
			Status: resp.StatusCode,
			Err:    fmt.Errorf("unable to read response body: %w", err)}
	}

	body := string(bytes)

	c.log.WithFields(logrus.Fields{
		"url":    url,
		"status": resp.Status,
	}).Debug("Receiving response")

	if resp.StatusCode != http.StatusOK {
		return &HttpClientError{
			Status: resp.StatusCode,
			Err: errors.New("unable to get info"),
			Body: body,
		}
	}

	if err = json.Unmarshal(bytes, entity); err != nil {
		return &HttpClientError{
			Status: resp.StatusCode,
			Err: fmt.Errorf("error unmarshalling JSON: %w", err),
			Body: body,
		}
	}

	return nil
}
