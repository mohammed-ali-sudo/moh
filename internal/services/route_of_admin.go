package services

import (
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"moh/models"
)

// AddRoute inserts a route_of_admin row (code, name).
func AddRoute(ctx context.Context, db *pgxpool.Pool, in models.RouteOfAdmin) (models.RouteOfAdmin, error) {
	// normalize code to uppercase/trim
	in.Code = strings.ToUpper(strings.TrimSpace(in.Code))
	in.Name = strings.TrimSpace(in.Name)

	// validate
	if msg, ok := in.Validate(); !ok {
		return models.RouteOfAdmin{}, errors.New(msg)
	}

	const q = `INSERT INTO route_of_admin (code, name) VALUES ($1, $2)`
	if _, err := db.Exec(ctx, q, in.Code, in.Name); err != nil {
		return models.RouteOfAdmin{}, err
	}
	return in, nil
}
