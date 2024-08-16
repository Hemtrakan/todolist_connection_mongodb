package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const DefaultHeader = "Tracking-Id"

func NewTrackingId() string {
	return uuid.New().String()
}

// Gin Middleware with default header
func TrackingId() gin.HandlerFunc {
	return TrackingIdWithCustomizedHeader(DefaultHeader)
}

// Gin Middleware with cusomized header
func TrackingIdWithCustomizedHeader(head string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tId := c.GetHeader(head)
		if tId == "" {
			tId = NewTrackingId()
		}
		c.Header(head, tId)
		c.Request.Header.Set(head, tId)
		headerCorRelation := "X-Correlation-ID"
		correlation := c.GetHeader(headerCorRelation)
		if correlation != "" {
			c.Writer.Header().Set(headerCorRelation, correlation)
		}
		c.Next()
	}
}
