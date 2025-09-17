package handlers

import (
	"net/http"

	"moh/internal/services"
	"moh/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ----------------------------------------------------------------------------
// DOSAGE
// ----------------------------------------------------------------------------

func AddDosageHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in models.Dosage
		if err := decodeStrict(c, &in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.AddDosage(c.Request.Context(), db, in)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "create_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, out)
	}
}

func GetAllDosagesHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := services.GetAllDosages(c.Request.Context(), db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "list_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

func UpdateDosageHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parseIDParam(c, "id")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_id", "message": err.Error()})
			return
		}
		var patch models.Dosage
		if err := decodeStrict(c, &patch); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.UpdateDosage(c.Request.Context(), db, id, patch)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "update_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, out)
	}
}

// ----------------------------------------------------------------------------
// ROUTE
// ----------------------------------------------------------------------------

func AddRouteHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in models.Route
		if err := decodeStrict(c, &in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.AddRoute(c.Request.Context(), db, in)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "create_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, out)
	}
}

func GetAllRoutesHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := services.GetAllRoutes(c.Request.Context(), db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "list_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

func UpdateRouteHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parseIDParam(c, "id")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_id", "message": err.Error()})
			return
		}
		var patch models.Route
		if err := decodeStrict(c, &patch); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.UpdateRoute(c.Request.Context(), db, id, patch)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "update_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, out)
	}
}

// ----------------------------------------------------------------------------
// STRENGTH
// ----------------------------------------------------------------------------

func AddStrengthHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in models.Strength
		if err := decodeStrict(c, &in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.AddStrength(c.Request.Context(), db, in)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "create_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, out)
	}
}

func GetAllStrengthsHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := services.GetAllStrengths(c.Request.Context(), db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "list_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

func UpdateStrengthHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parseIDParam(c, "id")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_id", "message": err.Error()})
			return
		}
		var patch models.Strength
		if err := decodeStrict(c, &patch); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.UpdateStrength(c.Request.Context(), db, id, patch)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "update_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, out)
	}
}

// ----------------------------------------------------------------------------
// API
// ----------------------------------------------------------------------------

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

func GetAllAPIsHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := services.GetAllAPIs(c.Request.Context(), db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "list_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

func UpdateAPIHandler(db *pgxpool.Pool) gin.HandlerFunc {
	type apiPatch struct {
		Name           string `json:"name"`
		IsPsychotropic *bool  `json:"ispsychotropic"`
	}
	return func(c *gin.Context) {
		id, err := parseIDParam(c, "id")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_id", "message": err.Error()})
			return
		}
		var patch apiPatch
		if err := decodeStrict(c, &patch); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.UpdateAPI(c.Request.Context(), db, id, patch.Name, patch.IsPsychotropic)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "update_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, out)
	}
}
