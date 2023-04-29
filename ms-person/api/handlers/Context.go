package handlers

import (
	"context"
	"fmt"
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

func AddRequestIdInRequestContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		if requestID := c.Writer.Header().Get(constants.REQUEST_ID); requestID != "" {
			ctx := context.WithValue(c.Request.Context(), constants.REQUEST_ID, requestID)
			c.Request = c.Request.WithContext(ctx)
		}
	}
}

func responseTimeout(c *gin.Context) {
	c.AbortWithError(http.StatusRequestTimeout, fmt.Errorf("Request Timeout"))
}
