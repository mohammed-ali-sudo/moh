package handlers

import (
	"encoding/json"
	"net/http"

	"moh/internal/services"
	"moh/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
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

func ListDosageFormsHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := services.ListDosageForms(c.Request.Context(), db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "list_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

func ListStrengthUnitsHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := services.ListStrengthUnits(c.Request.Context(), db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "list_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

func ListRoutesOfAdminHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := services.ListRoutesOfAdmin(c.Request.Context(), db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "list_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

func ListAPIsHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := services.ListAPIs(c.Request.Context(), db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "list_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

func ListManufacturingSitesHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := services.ListManufacturingSites(c.Request.Context(), db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "list_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

func ListAuthHoldersHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := services.ListAuthHolders(c.Request.Context(), db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "list_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

func ListMarketingAuthorizationsHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := services.ListMarketingAuthorizations(c.Request.Context(), db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "list_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

func ListDrugsHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := services.ListDrugs(c.Request.Context(), db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "list_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

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

func ListDrugRegistrationsHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := services.ListDrugRegistrations(c.Request.Context(), db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "list_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

func ListDrugRegistrationSitesHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := services.ListDrugRegistrationSites(c.Request.Context(), db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "list_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

func ListDrugRegistrationAuthHoldersHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := services.ListDrugRegistrationAuthHolders(c.Request.Context(), db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "list_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}
