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
func AddAuthorityHolder(ctx context.Context, db *pgxpool.Pool, in models.AuthorityHolder) (models.AuthorityHolder, error) {
	in.Name = strings.TrimSpace(in.Name)
	in.Country = strings.TrimSpace(in.Country)

	if err := in.Validate(); err != nil {
		msg, _ := models.FirstError(err)
		return models.AuthorityHolder{}, errors.New(msg)
	}

	const q = `
		INSERT INTO authority_holder (name, country)
		VALUES ($1, $2)
		RETURNING id, name, country
	`
	var out models.AuthorityHolder
	if err := pgxscan.Get(ctx, db, &out, q, in.Name, in.Country); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.AuthorityHolder{}, errors.New("name already exists")
		}
		return models.AuthorityHolder{}, err
	}
	return out, nil
}

func GetAllAuthorityHolders(ctx context.Context, db *pgxpool.Pool) ([]models.AuthorityHolder, error) {
	const q = `SELECT id, name, country FROM authority_holder ORDER BY name ASC`
	var out []models.AuthorityHolder
	if err := pgxscan.Select(ctx, db, &out, q); err != nil {
		return nil, err
	}
	return out, nil
}

func UpdateAuthorityHolder(ctx context.Context, db *pgxpool.Pool, id int64, patch models.AuthorityHolder) (models.AuthorityHolder, error) {
	name := strings.TrimSpace(patch.Name)
	country := strings.TrimSpace(patch.Country)

	const q = `
		UPDATE authority_holder
		   SET name = COALESCE(NULLIF($2, ''), name),
		       country = COALESCE(NULLIF($3, ''), country)
		 WHERE id = $1
		RETURNING id, name, country
	`
	var out models.AuthorityHolder
	if err := pgxscan.Get(ctx, db, &out, q, id, name, country); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.AuthorityHolder{}, errors.New("name already exists")
		}
		return models.AuthorityHolder{}, err
	}
	return out, nil
}

// ================= Manufacturing Site =================

func AddManufacturingSite(ctx context.Context, db *pgxpool.Pool, in models.ManufacturingSite) (models.ManufacturingSite, error) {
	in.Name = strings.TrimSpace(in.Name)
	in.Country = strings.TrimSpace(in.Country)

	if err := in.Validate(); err != nil {
		msg, _ := models.FirstError(err)
		return models.ManufacturingSite{}, errors.New(msg)
	}

	const q = `
		INSERT INTO manufacturing_site (name, country)
		VALUES ($1, $2)
		RETURNING id, name, country
	`
	var out models.ManufacturingSite
	if err := pgxscan.Get(ctx, db, &out, q, in.Name, in.Country); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.ManufacturingSite{}, errors.New("name already exists")
		}
		return models.ManufacturingSite{}, err
	}
	return out, nil
}

func GetAllManufacturingSites(ctx context.Context, db *pgxpool.Pool) ([]models.ManufacturingSite, error) {
	const q = `SELECT id, name, country FROM manufacturing_site ORDER BY name ASC`
	var out []models.ManufacturingSite
	if err := pgxscan.Select(ctx, db, &out, q); err != nil {
		return nil, err
	}
	return out, nil
}

func UpdateManufacturingSite(ctx context.Context, db *pgxpool.Pool, id int64, patch models.ManufacturingSite) (models.ManufacturingSite, error) {
	name := strings.TrimSpace(patch.Name)
	country := strings.TrimSpace(patch.Country)

	const q = `
		UPDATE manufacturing_site
		   SET name = COALESCE(NULLIF($2, ''), name),
		       country = COALESCE(NULLIF($3, ''), country)
		 WHERE id = $1
		RETURNING id, name, country
	`
	var out models.ManufacturingSite
	if err := pgxscan.Get(ctx, db, &out, q, id, name, country); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.ManufacturingSite{}, errors.New("name already exists")
		}
		return models.ManufacturingSite{}, err
	}
	return out, nil
}

// ================= Marketing =================

func AddMarketing(ctx context.Context, db *pgxpool.Pool, in models.Marketing) (models.Marketing, error) {
	in.Name = strings.TrimSpace(in.Name)
	in.Country = strings.TrimSpace(in.Country)

	if err := in.Validate(); err != nil {
		msg, _ := models.FirstError(err)
		return models.Marketing{}, errors.New(msg)
	}

	const q = `
		INSERT INTO marketing (name, country)
		VALUES ($1, $2)
		RETURNING id, name, country
	`
	var out models.Marketing
	if err := pgxscan.Get(ctx, db, &out, q, in.Name, in.Country); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.Marketing{}, errors.New("name already exists")
		}
		return models.Marketing{}, err
	}
	return out, nil
}

func GetAllMarketing(ctx context.Context, db *pgxpool.Pool) ([]models.Marketing, error) {
	const q = `SELECT id, name, country FROM marketing ORDER BY name ASC`
	var out []models.Marketing
	if err := pgxscan.Select(ctx, db, &out, q); err != nil {
		return nil, err
	}
	return out, nil
}

func UpdateMarketing(ctx context.Context, db *pgxpool.Pool, id int64, patch models.Marketing) (models.Marketing, error) {
	name := strings.TrimSpace(patch.Name)
	country := strings.TrimSpace(patch.Country)

	const q = `
		UPDATE marketing
		   SET name = COALESCE(NULLIF($2, ''), name),
		       country = COALESCE(NULLIF($3, ''), country)
		 WHERE id = $1
		RETURNING id, name, country
	`
	var out models.Marketing
	if err := pgxscan.Get(ctx, db, &out, q, id, name, country); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.Marketing{}, errors.New("name already exists")
		}
		return models.Marketing{}, err
	}
	return out, nil
}