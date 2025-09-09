package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
)

// XssValidatorGin sanitizes path, query params, and JSON bodies.
// Usage: r.Use(middleware.XssValidatorGin())
func XssValidatorGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ---- Sanitize Path ----
		c.Request.URL.Path = sanitizeString(c.Request.URL.Path)

		// ---- Sanitize Query Parameters ----
		sanitizedQuery := url.Values{}
		for key, values := range c.Request.URL.Query() {
			safeKey := sanitizeString(key)
			for _, v := range values {
				safeVal := sanitizeString(v)
				sanitizedQuery.Add(safeKey, safeVal)
			}
		}
		c.Request.URL.RawQuery = sanitizedQuery.Encode()

		// ---- Sanitize JSON Body (if present) ----
		if c.Request.Body != nil &&
			(strings.Contains(c.GetHeader("Content-Type"), "application/json")) {

			// Read body (it’s a one-shot reader)
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "invalid_body",
					"message": "failed to read request body",
				})
				return
			}
			// Restore empty body if nothing there
			if len(bodyBytes) == 0 {
				c.Request.Body = io.NopCloser(bytes.NewReader(bodyBytes))
			} else {
				var payload any
				if err := json.Unmarshal(bodyBytes, &payload); err != nil {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"error":   "invalid_json",
						"message": "invalid JSON body",
					})
					return
				}

				cleaned, err := clean(payload)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"error":   "sanitize_failed",
						"message": "failed to sanitize JSON",
					})
					return
				}

				safeBytes, err := json.Marshal(cleaned)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"error":   "sanitize_failed",
						"message": "failed to re-encode JSON",
					})
					return
				}

				// Replace request body with sanitized JSON
				c.Request.Body = io.NopCloser(bytes.NewReader(safeBytes))
				// Update Content-Length header (optional, harmless to set)
				c.Request.ContentLength = int64(len(safeBytes))
				c.Request.Header.Set("Content-Length", strconv.FormatInt(int64(len(safeBytes)), 10))
			}
		}

		c.Next()
	}
}

// --- helpers ---

func clean(data any) (any, error) {
	switch v := data.(type) {
	case map[string]any:
		for key, value := range v {
			sanitizedKey := sanitizeString(key)
			sanitizedVal, err := sanitizeValue(value)
			if err != nil {
				return nil, err
			}
			if sanitizedKey != key {
				delete(v, key)
				v[sanitizedKey] = sanitizedVal
			} else {
				v[key] = sanitizedVal
			}
		}
		return v, nil
	case []any:
		for i, item := range v {
			sanitized, err := sanitizeValue(item)
			if err != nil {
				return nil, err
			}
			v[i] = sanitized
		}
		return v, nil
	default:
		return sanitizeValue(v)
	}
}

func sanitizeValue(value any) (any, error) {
	switch v := value.(type) {
	case string:
		return sanitizeString(v), nil
	case map[string]any:
		return clean(v)
	case []any:
		return clean(v)
	default:
		return v, nil
	}
}

func sanitizeString(value string) string {
	return bluemonday.UGCPolicy().Sanitize(value)
}
