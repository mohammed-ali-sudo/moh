package handlers

import (
	"encoding/json"
	"net/http"

	"moh/internal/services"
	"moh/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

// AddStrengthHandler: POST { "code": "MG", "name": "Milligram (mg)" }
func AddStrengthHandler(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var in models.StrengthUnit
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		if err := dec.Decode(&in); err != nil {
			http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		created, err := services.AddStrength(r.Context(), db, in)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(created)
	}
}
