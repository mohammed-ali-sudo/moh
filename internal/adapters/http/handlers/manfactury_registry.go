package handlers

import (
	"net/http"

	"moh/internal/services"
	"moh/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// =============== Authority Holder ===================

func AddAuthHolderHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in models.AuthorityHolder
		if err := decodeStrict(c, &in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.AddAuthorityHolder(c.Request.Context(), db, in)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "create_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, out)
	}
}

func ListAuthHoldersHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := services.GetAllAuthorityHolders(c.Request.Context(), db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "list_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

func UpdateAuthHolderHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parseIDParam(c, "id")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_id", "message": err.Error()})
			return
		}
		var patch models.AuthorityHolder
		if err := decodeStrict(c, &patch); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.UpdateAuthorityHolder(c.Request.Context(), db, id, patch)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "update_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, out)
	}
}

// =============== Manufacturing Site ================

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

func ListManufacturingSitesHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := services.GetAllManufacturingSites(c.Request.Context(), db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "list_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

func UpdateManufacturingSiteHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parseIDParam(c, "id")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_id", "message": err.Error()})
			return
		}
		var patch models.ManufacturingSite
		if err := decodeStrict(c, &patch); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.UpdateManufacturingSite(c.Request.Context(), db, id, patch)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "update_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, out)
	}
}

// =================== Marketing =====================

func AddMarketingAuthorizationHandler(db *pgxpool.Pool) gin.HandlerFunc {
	// kept legacy name; internally uses the `marketing` table/model
	return func(c *gin.Context) {
		var in models.Marketing
		if err := decodeStrict(c, &in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.AddMarketing(c.Request.Context(), db, in)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "create_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, out)
	}
}

func ListMarketingAuthorizationsHandler(db *pgxpool.Pool) gin.HandlerFunc {
	// kept legacy name; internally reads from `marketing`
	return func(c *gin.Context) {
		items, err := services.GetAllMarketing(c.Request.Context(), db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "list_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

func UpdateMarketingAuthorizationHandler(db *pgxpool.Pool) gin.HandlerFunc {
	// kept legacy name; internally updates `marketing`
	return func(c *gin.Context) {
		id, err := parseIDParam(c, "id")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_id", "message": err.Error()})
			return
		}
		var patch models.Marketing
		if err := decodeStrict(c, &patch); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.UpdateMarketing(c.Request.Context(), db, id, patch)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "update_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, out)
	}
}
