package services

import (
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"moh/models"
)

// AddDosage inserts a dosage_form row (code, name).
func AddDosage(ctx context.Context, db *pgxpool.Pool, in models.DosageForm) (models.DosageForm, error) {
	in.Code = strings.ToUpper(strings.TrimSpace(in.Code))
	in.Name = strings.TrimSpace(in.Name)

	if msg, ok := in.Validate(); !ok {
		return models.DosageForm{}, errors.New(msg)
	}

	const q = `INSERT INTO dosage_form (code, name) VALUES ($1, $2)`
	if _, err := db.Exec(ctx, q, in.Code, in.Name); err != nil {
		return models.DosageForm{}, err
	}
	return in, nil
}
