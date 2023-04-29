package handlers

import (
	"time"

	"br.com.charlesrodrigo/ms-person/helper/logger"
	"github.com/gin-gonic/gin"
)

func AddLogRequestAndResponse() gin.HandlerFunc {

	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		logger.InfoWithContext(c.Request.Context(), "Start request",
			"method", c.Request.Method,
			"path", path,
			"query", query,
			"ip", c.ClientIP(),
			"user-agent", c.Request.UserAgent())

		c.Next()
		end := time.Now()
		latency := end.Sub(start)

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, e := range c.Errors.Errors() {
				logger.ErrorWithContext(c.Request.Context(), e,
					"status", c.Writer.Status(),
					"method", c.Request.Method,
					"path", path,
					"query", query,
					"ip", c.ClientIP(),
					"user-agent", c.Request.UserAgent(),
					"latency", latency)
			}
		} else {
			logger.InfoWithContext(c.Request.Context(), "End request",
				"status", c.Writer.Status(),
				"method", c.Request.Method,
				"path", path,
				"query", query,
				"ip", c.ClientIP(),
				"user-agent", c.Request.UserAgent(),
				"latency", latency)
		}

	}
}
