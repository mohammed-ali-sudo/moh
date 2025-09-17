package handlers

import (
	"net/http"

	execs "moh/internal/adapters/grpc"
	gapi "moh/internal/gen"
	"moh/internal/services"
	"moh/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// POST /manufacturer/drug
func AddDrugHandler(db *pgxpool.Pool, ex *execs.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in models.Drug
		if err := decodeStrict(c, &in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}

		// 1) insert locally
		out, err := services.AddDrug(c.Request.Context(), db, in)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "create_failed", "message": err.Error()})
			return
		}

		// 2) map to proto and notify other microservice
		pb := toProtoDrugOut(out)
		confirm, err := ex.AddDrug(c.Request.Context(), pb)
		if err != nil {
			// You can still return 201 for the local insert and include failure info for the sync step.
			c.JSON(http.StatusCreated, gin.H{
				"data": out,
				"execs": gin.H{
					"confirmation": false,
					"message":      err.Error(),
				},
			})
			return
		}

		// 3) return both: your created object + confirmation from other service
		c.JSON(http.StatusCreated, gin.H{
			"data": out,
			"execs": gin.H{
				"confirmation": confirm.Confirmation,
				"message":      confirm.Message,
			},
		})
	}
}

// --- mapping helper (adjust field names if your DTO differs) ---

// wrap returns nil when empty => optional fields stay unset
func wrapInt64(p *int64) *wrapperspb.Int64Value {
	if p == nil {
		return nil
	}
	return wrapperspb.Int64(*p)
}

func wrapString(p *string) *wrapperspb.StringValue {
	if p == nil || *p == "" {
		return nil
	}
	return wrapperspb.String(*p)
}

// Assuming you have a read DTO that mirrors the proto (IDs + names + countries).
// Replace models.DrugOut with whatever type `out` actually is.
func toProtoDrugOut(x models.DrugOut) *gapi.DrugOut {
	return &gapi.DrugOut{
		Id:        x.ID,
		BrandName: x.BrandName,
		Dose:      x.Dose,

		ApiId:   x.APIID,
		ApiName: x.APIName,

		DosageId:   x.DosageID,
		DosageName: x.DosageName,

		RouteId:   x.RouteID,
		RouteName: x.RouteName,

		StrengthId:   x.StrengthID,
		StrengthName: x.StrengthName,

		AuthorityHolderId:      wrapInt64(x.AuthorityHolderID),
		AuthorityHolderName:    wrapString(x.AuthorityHolderName),
		AuthorityHolderCountry: wrapString(x.AuthorityHolderCountry),

		ManufacturingSiteId:      wrapInt64(x.ManufacturingSiteID),
		ManufacturingSiteName:    wrapString(x.ManufacturingSiteName),
		ManufacturingSiteCountry: wrapString(x.ManufacturingSiteCountry),

		MarketingId:      wrapInt64(x.MarketingID),
		MarketingName:    wrapString(x.MarketingName),
		MarketingCountry: wrapString(x.MarketingCountry),
	}
}

// GET /manufacturer/drug
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

// PATCH /manufacturer/drug/:id
func UpdateDrugHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parseIDParam(c, "id")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_id", "message": err.Error()})
			return
		}
		var patch services.DrugPatch
		if err := decodeStrict(c, &patch); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json", "message": err.Error()})
			return
		}
		out, err := services.UpdateDrug(c.Request.Context(), db, id, patch)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "update_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, out)
	}
}

func ListDrugsWithBatchesHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := services.ListDrugsWithBatches(c.Request.Context(), db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "list_failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}
