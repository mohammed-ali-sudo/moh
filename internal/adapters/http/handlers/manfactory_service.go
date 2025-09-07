package handlers

import (
	"encoding/json"
	"net/http"

	"moh/internal/services"
	"moh/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ManufacturerCreateHandler(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		var m models.ManufacturerSite
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		if err := dec.Decode(&m); err != nil {
			http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Validate with the simple first-error approach
		if msg, ok := m.Validate(); !ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]any{
				"status": "error",
				"error":  msg,
			})
			return
		}

		created, err := services.AddManufacturerSite(r.Context(), db, m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(created)
	}
}

// keep your router wiring alias
func InventoryCreateHandler(db *pgxpool.Pool) http.HandlerFunc {
	return ManufacturerCreateHandler(db)
}
