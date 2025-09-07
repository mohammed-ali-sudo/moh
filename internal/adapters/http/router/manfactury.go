package router

import (
	"moh/internal/adapters/http/handlers"
	middleware "moh/shared/middlewares"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ManufacturerRouter(database *pgxpool.Pool) *mux.Router {
	router := mux.NewRouter()

	// Public
	public := router.PathPrefix("/").Subrouter()
	public.HandleFunc("/manfactory", handlers.InventoryCreateHandler(database)).Methods("POST")

	// Private (kept as in your file—no routes attached yet)
	private := router.PathPrefix("/").Subrouter()
	private.Use(middleware.Protect)

	return router
}
