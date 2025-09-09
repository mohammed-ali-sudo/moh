
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"moh/internal/services"
	"moh/models"
)

func AddAuthHolderHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in models.AuthHolder
		if err := decodeStrict(c, &in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.AddAuthHolder(c.Request.Context(), db, in)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "create_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, out)
	}
}

func AddMarketingAuthorizationHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in models.MarketingAuthorization
		if err := decodeStrict(c, &in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.AddMarketingAuthorization(c.Request.Context(), db, in)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "create_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, out)
	}
}

func AddManufacturingSiteHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in models.ManufacturingSite
		if err := decodeStrict(c, &in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.AddManufacturingSite(c.Request.Context(), db, in)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "create_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, out)
	}
}
