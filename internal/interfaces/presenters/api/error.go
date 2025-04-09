package api

import "errors"

var (
	ErrInvalidLocation = errors.New("Invalid location parameter")
	ErrInvalidDate     = errors.New("Ivalid date parameter")
	ErrInternalError = errors.New("Oooops! An unexpected error occurred on the server")
)

type ErrorResponse struct {
	Error string `json:"error"`
}
