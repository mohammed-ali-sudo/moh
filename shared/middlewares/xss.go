package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

// XssValidator is a middleware that sanitizes all incoming request bodies,
// paths, and query parameters to prevent Cross-Site Scripting (XSS) attacks.
func XssValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("XSS validator is starting...")

		// ---- Sanitize Path ----
		r.URL.Path = saintiaizeString(r.URL.Path)

		// ---- Sanitize Query Parameters ----
		sanitizedQuery := url.Values{}
		for key, values := range r.URL.Query() {
			safeKey := saintiaizeString(key)
			for _, v := range values {
				safeVal := saintiaizeString(v)
				sanitizedQuery.Add(safeKey, safeVal)
			}
		}
		r.URL.RawQuery = sanitizedQuery.Encode()

		// ---- Sanitize JSON Body ----
		if r.Body != nil && r.ContentLength > 0 &&
			strings.Contains(r.Header.Get("Content-Type"), "application/json") {

			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "failed to read request body", http.StatusBadRequest)
				return
			}

			// Decode JSON into a generic structure
			var payload any
			if err := json.Unmarshal(bodyBytes, &payload); err != nil {
				http.Error(w, "invalid JSON body", http.StatusBadRequest)
				return
			}

			// Recursively clean
			cleaned, err := clean(payload)
			if err != nil {
				http.Error(w, "failed to sanitize JSON", http.StatusInternalServerError)
				return
			}

			// Re-encode sanitized JSON
			safeBytes, err := json.Marshal(cleaned)
			if err != nil {
				http.Error(w, "failed to re-encode JSON", http.StatusInternalServerError)
				return
			}

			// Replace body with sanitized JSON
			r.Body = io.NopCloser(bytes.NewBuffer(safeBytes))
			r.ContentLength = int64(len(safeBytes))
		}

		// Continue to next handler
		next.ServeHTTP(w, r)
	})
}

// Recursively clean any arbitrary structure (map/slice/string).
func clean(data any) (any, error) {
	switch v := data.(type) {
	case map[string]any:
		for key, value := range v {
			sanitizedKey := saintiaizeString(key)
			sanitizedVal, err := sanitiazeValue(value)
			if err != nil {
				return nil, err
			}
			// If key changed, move it
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
			sanitized, err := sanitiazeValue(item)
			if err != nil {
				return nil, err
			}
			v[i] = sanitized
		}
		return v, nil
	default:
		return sanitiazeValue(v)
	}
}

func sanitiazeValue(value any) (any, error) {
	switch v := value.(type) {
	case string:
		return saintiaizeString(v), nil
	case map[string]any:
		return clean(v)
	case []any:
		return clean(v)
	default:
		return v, nil
	}
}

func saintiaizeString(value string) string {
	return bluemonday.UGCPolicy().Sanitize(value)
}
