package handlers

import (
	"net/http"

	"moh/internal/services"
	"moh/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AddDrugHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in models.Drug
		if err := decodeStrict(c, &in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.AddDrug(c.Request.Context(), db, in)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "create_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, out)
	}
}

func AddBatchHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in models.Batch
		if err := decodeStrict(c, &in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.AddBatch(c.Request.Context(), db, in)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "create_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, out)
	}
}

func AddDrugRegistrationHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in models.DrugRegistration
		if err := decodeStrict(c, &in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.AddDrugRegistration(c.Request.Context(), db, in)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "create_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, out)
	}
}

func AddDrugRegistrationSiteHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in models.DrugRegistrationSite
		if err := decodeStrict(c, &in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.AddDrugRegistrationSite(c.Request.Context(), db, in)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "create_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, out)
	}
}

func AddDrugRegistrationAuthHolderHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in models.DrugRegistrationAuthHolder
		if err := decodeStrict(c, &in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.AddDrugRegistrationAuthHolder(c.Request.Context(), db, in)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "create_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, out)
	}
}
