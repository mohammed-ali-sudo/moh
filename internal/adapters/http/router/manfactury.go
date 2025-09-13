// internal/adapters/http/router/manufacturer.go
package router

import (
	"net/http"

	"moh/internal/adapters/http/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ManufacturerRouter registers all endpoints under /manufacturer.
func ManufacturerRouter(r *gin.Engine, db *pgxpool.Pool) {
	g := r.Group("/manufacturer")

	// Health
	g.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "ok") })

	// CORS preflight (optional)
	g.OPTIONS("/*path", func(c *gin.Context) { c.Status(http.StatusNoContent) })

	// ===== Master (POST) =====
	g.POST("/inn", handlers.AddAPIHandler(db))                // INN / API
	g.POST("/route", handlers.AddRouteOfAdminHandler(db))     // Route
	g.POST("/dosage", handlers.AddDosageFormHandler(db))      // Dosage form
	g.POST("/strength", handlers.AddStrengthUnitHandler(db))  // Strength unit
	g.POST("/auth-holder", handlers.AddAuthHolderHandler(db)) // Authorization holder
	g.POST("/marketing-authorization", handlers.AddMarketingAuthorizationHandler(db))
	g.POST("/manfactory", handlers.AddManufacturingSiteHandler(db)) // legacy alias
	g.POST("/manufacturing-site", handlers.AddManufacturingSiteHandler(db))

	// ===== Domain (POST) =====
	g.POST("/drug", handlers.AddDrugHandler(db))
	g.POST("/drug-registration", handlers.AddDrugRegistrationHandler(db))
	g.POST("/drug-registration/site", handlers.AddDrugRegistrationSiteHandler(db))
	g.POST("/drug-registration/auth-holder", handlers.AddDrugRegistrationAuthHolderHandler(db))
	g.POST("/batch", handlers.AddBatchHandler(db))

	// ===== Simple list (GET) =====
	g.GET("/inn", handlers.ListAPIsHandler(db))
	g.GET("/route", handlers.ListRoutesOfAdminHandler(db))
	g.GET("/dosage", handlers.ListDosageFormsHandler(db))
	g.GET("/strength", handlers.ListStrengthUnitsHandler(db))
	g.GET("/auth-holder", handlers.ListAuthHoldersHandler(db))
	g.GET("/marketing-authorization", handlers.ListMarketingAuthorizationsHandler(db))
	g.GET("/manufacturing-site", handlers.ListManufacturingSitesHandler(db))

	g.GET("/drug", handlers.ListDrugsHandler(db))
	g.GET("/registration", handlers.ListDrugRegistrationsHandler(db))
	g.GET("/registration/site", handlers.ListDrugRegistrationSitesHandler(db))
	g.GET("/registration/holder", handlers.ListDrugRegistrationAuthHoldersHandler(db))
	g.GET("/batch", handlers.ListBatchesHandler(db))
}
