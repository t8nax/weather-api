package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func LogMiddleware(log *logrus.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()

		log.WithFields(logrus.Fields{
			"method": c.Method(),
			"path":   c.Path(),
			"status": c.Response().StatusCode(),
		}).Debug("HTTP Request completed")

		return err
	}
}
