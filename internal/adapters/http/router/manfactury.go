package router

import (
	"net/http"

	"moh/internal/adapters/http/handlers"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

// New returns a mux with ABSOLUTE paths (no subrouters, no StripPrefix).
// Matches Postman collection:
//   - POST /manufacturer/inn
//   - POST /manufacturer/route
//   - POST /manufacturer/dosage
//   - POST /manufacturer/strength
//   - POST /manufacturer/manfactory
//   - GET  /manufacturer/ping
func ManufacturerRouter(database *pgxpool.Pool) *mux.Router {
	r := mux.NewRouter()

	// --- Public routes (no Protect) ---
	public := r.PathPrefix("/").Subrouter()
	public.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}).Methods("GET")

	// --- Private routes (with Protect) ---
	private := r.PathPrefix("/").Subrouter()

	private.HandleFunc("/inn", handlers.AddINNHandler(database)).Methods("POST")
	private.HandleFunc("/route", handlers.AddRouteHandler(database)).Methods("POST")
	private.HandleFunc("/dosage", handlers.AddDosageHandler(database)).Methods("POST")
	private.HandleFunc("/strength", handlers.AddStrengthHandler(database)).Methods("POST")
	private.HandleFunc("/manfactory", handlers.InventoryCreateHandler(database)).Methods("POST")

	return r
}
