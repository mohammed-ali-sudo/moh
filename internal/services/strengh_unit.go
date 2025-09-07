package services

import (
	"context"
	"errors"
	"strings"

	"moh/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

// AddStrength inserts a strength_unit row (code, name).
func AddStrength(ctx context.Context, db *pgxpool.Pool, in models.StrengthUnit) (models.StrengthUnit, error) {
	in.Code = strings.ToUpper(strings.TrimSpace(in.Code))
	in.Name = strings.TrimSpace(in.Name)

	if msg, ok := in.Validate(); !ok {
		return models.StrengthUnit{}, errors.New(msg)
	}

	const q = `INSERT INTO strength_unit (code, name) VALUES ($1, $2)`
	if _, err := db.Exec(ctx, q, in.Code, in.Name); err != nil {
		return models.StrengthUnit{}, err
	}
	return in, nil
}
