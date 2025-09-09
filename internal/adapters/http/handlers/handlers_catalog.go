
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"moh/internal/services"
	"moh/models"
)

// helper decode with DisallowUnknownFields
func decodeStrict(c *gin.Context, dst any) error {
	dec := json.NewDecoder(c.Request.Body)
	dec.DisallowUnknownFields()
	return dec.Decode(dst)
}

func AddDosageFormHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in models.DosageForm
		if err := decodeStrict(c, &in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.AddDosageForm(c.Request.Context(), db, in)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "create_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, out)
	}
}

func AddStrengthUnitHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in models.StrengthUnit
		if err := decodeStrict(c, &in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.AddStrengthUnit(c.Request.Context(), db, in)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "create_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, out)
	}
}

func AddRouteOfAdminHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in models.RouteOfAdmin
		if err := decodeStrict(c, &in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.AddRouteOfAdmin(c.Request.Context(), db, in)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "create_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, out)
	}
}

func AddAPIHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in models.API
		if err := decodeStrict(c, &in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.AddAPI(c.Request.Context(), db, in)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "create_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, out)
	}
}
