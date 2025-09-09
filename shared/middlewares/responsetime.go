package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// ResponseTimeGin logs the request duration and adds X-Response-Time header (Gin version).
// Usage: r.Use(middleware.ResponseTimeGin())
func ResponseTimeGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next() // process request

		dur := time.Since(start)
		// Add the response time header (may be ignored if headers already sent)
		c.Writer.Header().Set("X-Response-Time", dur.String())

		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		log.Printf("[%s] %s %d %s", method, path, status, dur)
	}
}
