package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func RequestLogger(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// generate requestId
		requestID := uuid.New().String()

		// inject into response header
		c.Set("X-Request-ID", requestID)

		// continue request
		err := c.Next()

		// calculate duration
		duration := time.Since(start)

		// log
		logger.Info("request completed",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.String("request_id", requestID),
			zap.Duration("duration", duration),
			zap.Int("status", c.Response().StatusCode()),
		)

		return err
	}
}
