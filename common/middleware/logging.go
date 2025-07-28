package middleware

import (
	"github.com/gin-gonic/gin"
)

type LoggerI interface {
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Sync() error
}

func LoggingMiddleWare(log LoggerI) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Process request
		c.Next()

		// Log after request completion
		log.Infof("URL: %v, body: %v",
			c.Request.URL,
			c.Request.GetBody,
		)

		// Log errors if any
		if len(c.Errors) > 0 {
			log.Errorf("Request errors: %v", c.Errors)
		}
	}
}
