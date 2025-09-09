package middleware

import (
	"github.com/gin-gonic/gin"
)

// SecurityHeader sets some basic security headers.
func SecurityHeaderGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-DNS-Prefetch-Control", "off")
		c.Writer.Header().Set("X-Frame-Options", "DENY")
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Next()
	}
}
