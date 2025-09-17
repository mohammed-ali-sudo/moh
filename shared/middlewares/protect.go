package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

// ---- unchanged: VerifyToken ----

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	secret := []byte("jwtsecretstring")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure HMAC signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok && ve.Errors&jwt.ValidationErrorExpired != 0 {
			return nil, fmt.Errorf("token expired")
		}
		return nil, fmt.Errorf("token validation failed: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

// ---- gin-native middleware: ProtectGin ----

func ProtectGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Allow CORS preflight to pass through
		if c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "authorization header is missing",
			})
			return
		}

		// Expect "Bearer <token>"
		const prefix = "Bearer "
		if !strings.HasPrefix(authHeader, prefix) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "invalid authorization header format",
			})
			return
		}
		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, prefix))

		claims, err := VerifyToken(tokenString)
		if err != nil {
			msg := "invalid token"
			if strings.Contains(err.Error(), "token expired") {
				msg = "token has expired"
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": msg,
			})
			return
		}

		// Save claims in Gin context for handlers
		c.Set("user_claims", claims)

		c.Next()
	}
}

// Optional helper to fetch claims inside handlers.
func ClaimsFromContext(c *gin.Context) (jwt.MapClaims, bool) {
	v, ok := c.Get("user_claims")
	if !ok {
		return nil, false
	}
	claims, ok := v.(jwt.MapClaims)
	return claims, ok
}
