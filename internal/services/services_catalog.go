
package services

import (
	"context"
	"errors"
	"strings"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"moh/models"
)

// AddDosageForm creates a dosage_form.
func AddDosageForm(ctx context.Context, db *pgxpool.Pool, in models.DosageForm) (models.DosageForm, error) {
	in.ID = uuid.NewString()
	in.Code = strings.ToUpper(strings.TrimSpace(in.Code))
	in.Name = strings.TrimSpace(in.Name)

	if err := in.Validate(); err != nil {
		msg, _ := models.FirstError(err)
		return models.DosageForm{}, errors.New(msg)
	}

	const q = `
		INSERT INTO dosage_form (id, code, name)
		VALUES ($1, $2, $3)
		RETURNING id, code, name, created_at, updated_at
	`
	var out models.DosageForm
	if err := pgxscan.Get(ctx, db, &out, q, in.ID, in.Code, in.Name); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.DosageForm{}, errors.New("code already exists")
		}
		return models.DosageForm{}, err
	}
	return out, nil
}

// AddStrengthUnit creates a strength_unit.
func AddStrengthUnit(ctx context.Context, db *pgxpool.Pool, in models.StrengthUnit) (models.StrengthUnit, error) {
	in.ID = uuid.NewString()
	in.Code = strings.ToUpper(strings.TrimSpace(in.Code))
	in.Name = strings.TrimSpace(in.Name)

	if err := in.Validate(); err != nil {
		msg, _ := models.FirstError(err)
		return models.StrengthUnit{}, errors.New(msg)
	}

	const q = `
		INSERT INTO strength_unit (id, code, name)
		VALUES ($1, $2, $3)
		RETURNING id, code, name, created_at, updated_at
	`
	var out models.StrengthUnit
	if err := pgxscan.Get(ctx, db, &out, q, in.ID, in.Code, in.Name); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.StrengthUnit{}, errors.New("code already exists")
		}
		return models.StrengthUnit{}, err
	}
	return out, nil
}

// AddRouteOfAdmin creates a route_of_admin.
func AddRouteOfAdmin(ctx context.Context, db *pgxpool.Pool, in models.RouteOfAdmin) (models.RouteOfAdmin, error) {
	in.ID = uuid.NewString()
	in.Code = strings.ToUpper(strings.TrimSpace(in.Code))
	in.Name = strings.TrimSpace(in.Name)

	if err := in.Validate(); err != nil {
		msg, _ := models.FirstError(err)
		return models.RouteOfAdmin{}, errors.New(msg)
	}

	const q = `
		INSERT INTO route_of_admin (id, code, name)
		VALUES ($1, $2, $3)
		RETURNING id, code, name, created_at, updated_at
	`
	var out models.RouteOfAdmin
	if err := pgxscan.Get(ctx, db, &out, q, in.ID, in.Code, in.Name); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.RouteOfAdmin{}, errors.New("code already exists")
		}
		return models.RouteOfAdmin{}, err
	}
	return out, nil
}

// AddAPI creates an API (active ingredient).
func AddAPI(ctx context.Context, db *pgxpool.Pool, in models.API) (models.API, error) {
	in.ID = uuid.NewString()
	in.Name = strings.TrimSpace(in.Name)
	if s := strings.ToLower(strings.TrimSpace(string(in.Status))); s == "" {
		in.Status = models.APIStatusActive
	} else {
		in.Status = models.APIStatus(s)
	}

	if err := in.Validate(); err != nil {
		msg, _ := models.FirstError(err)
		return models.API{}, errors.New(msg)
	}

	const q = `
		INSERT INTO api (id, name, status)
		VALUES ($1, $2, $3)
		RETURNING id, name, status, created_at, updated_at
	`
	var out models.API
	if err := pgxscan.Get(ctx, db, &out, q, in.ID, in.Name, in.Status); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.API{}, errors.New("name already exists")
		}
		return models.API{}, err
	}
	return out, nil
}
