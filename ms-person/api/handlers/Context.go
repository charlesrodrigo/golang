package handlers

import (
	"net/http"

	"br.com.charlesrodrigo/ms-person/helper/constants"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func TimeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(constants.TIMEOUT_CONTEXT),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(responseTimeout),
	)
}

func responseTimeout(c *gin.Context) {
	c.String(http.StatusRequestTimeout, "timeout")
}
