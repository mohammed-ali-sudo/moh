// internal/adapters/http/router/manufacturer.go
package router

import (
	"net/http"

	"moh/internal/adapters/http/handlers"

	execs "moh/internal/adapters/grpc"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// internal/adapters/http/router/manufacturer.go

// ManufacturerRouter registers all endpoints under /manufacturer.
func ManufacturerRouter(r *gin.Engine, db *pgxpool.Pool, ex *execs.Client) {
	g := r.Group("/manufacturer")

	// Health
	g.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "ok") })

	// CORS preflight (optional)
	g.OPTIONS("/*path", func(c *gin.Context) { c.Status(http.StatusNoContent) })

	// ===================== INN / API =====================
	// Preferred path
	g.POST("/api", handlers.AddAPIHandler(db))
	g.GET("/api", handlers.GetAllAPIsHandler(db))
	g.PATCH("/api/:id", handlers.UpdateAPIHandler(db))

	// Legacy alias to avoid breaking old clients
	g.POST("/inn", handlers.AddAPIHandler(db))
	g.GET("/inn", handlers.GetAllAPIsHandler(db))
	g.PATCH("/inn/:id", handlers.UpdateAPIHandler(db))

	// ===================== Route =========================
	g.POST("/route", handlers.AddRouteHandler(db))
	g.GET("/route", handlers.GetAllRoutesHandler(db))
	g.PATCH("/route/:id", handlers.UpdateRouteHandler(db))

	// ===================== Dosage ========================
	g.POST("/dosage", handlers.AddDosageHandler(db))
	g.GET("/dosage", handlers.GetAllDosagesHandler(db))
	g.PATCH("/dosage/:id", handlers.UpdateDosageHandler(db))

	// ===================== Strength ======================
	g.POST("/strength", handlers.AddStrengthHandler(db))
	g.GET("/strength", handlers.GetAllStrengthsHandler(db))
	g.PATCH("/strength/:id", handlers.UpdateStrengthHandler(db))

	// ===================== Authority Holder ===============
	g.POST("/auth-holder", handlers.AddAuthHolderHandler(db))
	g.GET("/auth-holder", handlers.ListAuthHoldersHandler(db))
	g.PATCH("/auth-holder/:id", handlers.UpdateAuthHolderHandler(db))

	// ===================== Marketing =====================
	// (Handler names kept as "MarketingAuthorization" for backward compatibility)
	g.POST("/marketing-authorization", handlers.AddMarketingAuthorizationHandler(db))
	g.GET("/marketing-authorization", handlers.ListMarketingAuthorizationsHandler(db))
	g.PATCH("/marketing-authorization/:id", handlers.UpdateMarketingAuthorizationHandler(db))

	// ===================== Manufacturing Site ============
	// Preferred
	g.POST("/manufacturing-site", handlers.AddManufacturingSiteHandler(db))
	g.GET("/manufacturing-site", handlers.ListManufacturingSitesHandler(db))
	g.PATCH("/manufacturing-site/:id", handlers.UpdateManufacturingSiteHandler(db))
	// Legacy alias
	g.POST("/manfactory", handlers.AddManufacturingSiteHandler(db))

	// ===================== Drug ==========================
	g.GET("/drug", handlers.ListDrugsHandler(db))
	g.POST("/drug", handlers.AddDrugHandler(db, ex))

	g.PATCH("/drug/:id", handlers.UpdateDrugHandler(db))
	g.GET("/drug-with-batches", handlers.ListDrugsWithBatchesHandler(db)) // <-- add this

	// ===================== Batch =========================
	g.POST("/batch", handlers.AddBatchHandler(db))
	g.GET("/batch", handlers.ListBatchesHandler(db))
	g.PATCH("/batch/:id", handlers.UpdateBatchHandler(db))

	// ===================== Legacy read-only lists (keep) ==

}
