package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/t8nax/weather-api/internal/common"
	"github.com/t8nax/weather-api/internal/interfaces/presenters/api"
)

const (
	unexpectedErrMsg = "Oooops! An unexpected error occurred on the server"
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

		var appErr *common.AppError
		if !errors.As(err, &appErr) {
			return ctx.Status(fiber.StatusInternalServerError).JSON(&api.ErrorResponse{
				Error: unexpectedErrMsg,
			})
		}

		status := toStatusCode(appErr.Code)
		msg := toMessage(appErr.Code)

		return ctx.Status(status).JSON(&api.ErrorResponse{
			Error: msg,
		})
	}
}

func toMessage(code common.AppErrorCode) string {
	switch code {
	case common.CodeInvalidLocation:
		return "Invalid location parameter"
	case common.CodeInvalidDate:
		return "Invalid date parameter"
	case common.CodeServiceError:
		return "Weather service is unavailable. Please try later"
	default:
		return unexpectedErrMsg
	}
}

func toStatusCode(code common.AppErrorCode) int {
	switch code {
	case common.CodeInvalidLocation, common.CodeInvalidDate:
		return fiber.StatusBadRequest
	default:
		return fiber.StatusInternalServerError
	}
}
