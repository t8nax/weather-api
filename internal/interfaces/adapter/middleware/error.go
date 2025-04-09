package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/t8nax/weather-api/internal/interfaces/adapter/service"
	"github.com/t8nax/weather-api/internal/interfaces/presenters/api"
)

func ErrorMiddleware(log logrus.FieldLogger) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := ctx.Next()

		if err == nil {
			return nil
		}

		log.WithFields(logrus.Fields{
			"method": ctx.Method(),
			"path":   ctx.Path(),
			"error":  err.Error(),
		}).Error("HTTP Request failed")

		if _, ok := err.(*fiber.Error); ok {
			return err
		}

		if isBadRequest(err) {
			return ctx.Status(fiber.StatusBadRequest).JSON(&api.ErrorResponse{
				Error: err.Error(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(&api.ErrorResponse{
			Error: api.ErrInternalError.Error(),
		})
	}
}

func isBadRequest(err error) bool {
	return errors.Is(err, service.ErrInvalidLocation) || errors.Is(err, api.ErrInvalidLocation) ||
		errors.Is(err, service.ErrInvalidYear) || errors.Is(err, api.ErrInvalidDate)
}
