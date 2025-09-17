package services

import (
	"context"
	"errors"
	"strings"

	"moh/models"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

/*
Conventions:
- Add*: trims, UPPERCASEs Code, validates, inserts, returns the row.
- GetAll*: returns all rows ordered by code.
- Update*: trims, UPPERCASEs Code; uses COALESCE(NULLIF($x,''), col) so
  passing "" keeps the existing DB value (no NULL writes).
- Unique violation (code) -> friendly error.
*/

// ============================= Dosage ========================================

func AddDosage(ctx context.Context, db *pgxpool.Pool, in models.Dosage) (models.Dosage, error) {
	in.Code = strings.ToUpper(strings.TrimSpace(in.Code))
	in.Name = strings.TrimSpace(in.Name)

	if err := in.Validate(); err != nil {
		msg, _ := models.FirstError(err)
		return models.Dosage{}, errors.New(msg)
	}

	const q = `
		INSERT INTO dosage (code, name)
		VALUES ($1, $2)
		RETURNING id, code, name
	`
	var out models.Dosage
	if err := pgxscan.Get(ctx, db, &out, q, in.Code, in.Name); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.Dosage{}, errors.New("code already exists")
		}
		return models.Dosage{}, err
	}
	return out, nil
}

func GetAllDosages(ctx context.Context, db *pgxpool.Pool) ([]models.Dosage, error) {
	const q = `SELECT id, code, name FROM dosage ORDER BY code ASC`
	var out []models.Dosage
	if err := pgxscan.Select(ctx, db, &out, q); err != nil {
		return nil, err
	}
	return out, nil
}

func UpdateDosage(ctx context.Context, db *pgxpool.Pool, id int64, patch models.Dosage) (models.Dosage, error) {
	code := strings.ToUpper(strings.TrimSpace(patch.Code))
	name := strings.TrimSpace(patch.Name)

	const q = `
		UPDATE dosage
		   SET code = COALESCE(NULLIF($2, ''), code),
		       name = COALESCE(NULLIF($3, ''), name)
		 WHERE id = $1
		RETURNING id, code, name
	`
	var out models.Dosage
	if err := pgxscan.Get(ctx, db, &out, q, id, code, name); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.Dosage{}, errors.New("code already exists")
		}
		return models.Dosage{}, err
	}
	return out, nil
}

// ============================== Route ========================================

func AddRoute(ctx context.Context, db *pgxpool.Pool, in models.Route) (models.Route, error) {
	in.Code = strings.ToUpper(strings.TrimSpace(in.Code))
	in.Name = strings.TrimSpace(in.Name)

	if err := in.Validate(); err != nil {
		msg, _ := models.FirstError(err)
		return models.Route{}, errors.New(msg)
	}

	const q = `
		INSERT INTO route (code, name)
		VALUES ($1, $2)
		RETURNING id, code, name
	`
	var out models.Route
	if err := pgxscan.Get(ctx, db, &out, q, in.Code, in.Name); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.Route{}, errors.New("code already exists")
		}
		return models.Route{}, err
	}
	return out, nil
}

func GetAllRoutes(ctx context.Context, db *pgxpool.Pool) ([]models.Route, error) {
	const q = `SELECT id, code, name FROM route ORDER BY code ASC`
	var out []models.Route
	if err := pgxscan.Select(ctx, db, &out, q); err != nil {
		return nil, err
	}
	return out, nil
}

func UpdateRoute(ctx context.Context, db *pgxpool.Pool, id int64, patch models.Route) (models.Route, error) {
	code := strings.ToUpper(strings.TrimSpace(patch.Code))
	name := strings.TrimSpace(patch.Name)

	const q = `
		UPDATE route
		   SET code = COALESCE(NULLIF($2, ''), code),
		       name = COALESCE(NULLIF($3, ''), name)
		 WHERE id = $1
		RETURNING id, code, name
	`
	var out models.Route
	if err := pgxscan.Get(ctx, db, &out, q, id, code, name); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.Route{}, errors.New("code already exists")
		}
		return models.Route{}, err
	}
	return out, nil
}

// ============================ Strength =======================================

func AddStrength(ctx context.Context, db *pgxpool.Pool, in models.Strength) (models.Strength, error) {
	in.Code = strings.ToUpper(strings.TrimSpace(in.Code))
	in.Name = strings.TrimSpace(in.Name)

	if err := in.Validate(); err != nil {
		msg, _ := models.FirstError(err)
		return models.Strength{}, errors.New(msg)
	}

	const q = `
		INSERT INTO strength (code, name)
		VALUES ($1, $2)
		RETURNING id, code, name
	`
	var out models.Strength
	if err := pgxscan.Get(ctx, db, &out, q, in.Code, in.Name); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.Strength{}, errors.New("code already exists")
		}
		return models.Strength{}, err
	}
	return out, nil
}

func GetAllStrengths(ctx context.Context, db *pgxpool.Pool) ([]models.Strength, error) {
	const q = `SELECT id, code, name FROM strength ORDER BY code ASC`
	var out []models.Strength
	if err := pgxscan.Select(ctx, db, &out, q); err != nil {
		return nil, err
	}
	return out, nil
}

func UpdateStrength(ctx context.Context, db *pgxpool.Pool, id int64, patch models.Strength) (models.Strength, error) {
	code := strings.ToUpper(strings.TrimSpace(patch.Code))
	name := strings.TrimSpace(patch.Name)

	const q = `
		UPDATE strength
		   SET code = COALESCE(NULLIF($2, ''), code),
		       name = COALESCE(NULLIF($3, ''), name)
		 WHERE id = $1
		RETURNING id, code, name
	`
	var out models.Strength
	if err := pgxscan.Get(ctx, db, &out, q, id, code, name); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.Strength{}, errors.New("code already exists")
		}
		return models.Strength{}, err
	}
	return out, nil
}

// AddAPI inserts exactly one API and returns it.
func AddAPI(ctx context.Context, db *pgxpool.Pool, in models.API) (models.API, error) {
	in.Name = strings.TrimSpace(in.Name) // simple clean-up

	if err := in.Validate(); err != nil {
		msg, _ := models.FirstError(err)
		return models.API{}, errors.New(msg)
	}

	const q = `
		INSERT INTO api (name, ispsychotropic)
		VALUES ($1, $2)
		RETURNING id, name, ispsychotropic
	`
	var out models.API
	if err := pgxscan.Get(ctx, db, &out, q, in.Name, in.IsPsychotropic); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.API{}, errors.New("name already exists")
		}
		return models.API{}, err
	}
	return out, nil
}

// GetAllAPIs lists all APIs ordered by name.
func GetAllAPIs(ctx context.Context, db *pgxpool.Pool) ([]models.API, error) {
	const q = `SELECT id, name, ispsychotropic FROM api ORDER BY name ASC`
	var out []models.API
	if err := pgxscan.Select(ctx, db, &out, q); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateAPI updates name and/or ispsychotropic.
// - name uses "collapse": pass "" to keep the old name.
// - ispsychotropic is optional: pass nil to keep the old value, or &true / &false to set it.
func UpdateAPI(ctx context.Context, db *pgxpool.Pool, id int64, name string, ispsychotropic *bool) (models.API, error) {
	name = strings.TrimSpace(name)

	const q = `
		UPDATE api
		   SET name = COALESCE(NULLIF($2, ''), name),
		       ispsychotropic = COALESCE($3, ispsychotropic)
		 WHERE id = $1
		RETURNING id, name, ispsychotropic
	`
	var out models.API
	if err := pgxscan.Get(ctx, db, &out, q, id, name, ispsychotropic); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.API{}, errors.New("name already exists")
		}
		return models.API{}, err
	}
	return out, nil
}
