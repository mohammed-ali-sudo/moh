package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"moh/internal/services"
	"moh/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Reuse your shared helpers if already present
func decodeStrict(c *gin.Context, dst any) error {
	dec := json.NewDecoder(c.Request.Body)
	dec.DisallowUnknownFields()
	return dec.Decode(dst)
}
func parseIDParam(c *gin.Context, name string) (int64, error) {
	idStr := strings.TrimSpace(c.Param(name))
	if idStr == "" {
		return 0, errors.New("missing id")
	}
	return strconv.ParseInt(idStr, 10, 64)
}

// POST /manufacturer/batch
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

// GET /manufacturer/batch
func ListBatchesHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := services.ListBatches(c.Request.Context(), db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "list_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

// PATCH /manufacturer/batch/:id
func UpdateBatchHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parseIDParam(c, "id")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_id", "message": err.Error()})
			return
		}
		var patch services.BatchPatch
		if err := decodeStrict(c, &patch); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.UpdateBatch(c.Request.Context(), db, id, patch)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "update_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, out)
	}
}
