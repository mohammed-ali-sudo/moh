package router

import (
	"net/http"

	"moh/internal/adapters/http/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ManufacturerRouter registers routes under /manufacturer on the given Gin engine.
// Mirrors your mux endpoints:
//   - GET  /manufacturer/ping
//   - POST /manufacturer/inn
//   - POST /manufacturer/route
//   - POST /manufacturer/dosage
//   - POST /manufacturer/strength
//   - POST /manufacturer/manfactory
func ManufacturerRouter(r *gin.Engine, db *pgxpool.Pool) {
	group := r.Group("/manufacturer")

	// Public
	group.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// Private (attach auth middleware to 'group' if needed)
	group.POST("/inn", handlers.AddAPIHandler(db))                      // INN => API
	group.POST("/route", handlers.AddRouteOfAdminHandler(db))           // Route of admin
	group.POST("/dosage", handlers.AddDosageFormHandler(db))            // Dosage form
	group.POST("/strength", handlers.AddStrengthUnitHandler(db))        // Strength unit
	group.POST("/manfactory", handlers.AddManufacturingSiteHandler(db)) // Manufacturing site
}
