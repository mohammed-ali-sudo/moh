
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

// AddAuthHolder creates an authorization holder.
func AddAuthHolder(ctx context.Context, db *pgxpool.Pool, in models.AuthHolder) (models.AuthHolder, error) {
	in.ID = uuid.NewString()
	in.Name = strings.TrimSpace(in.Name)
	in.RegistrationNumber = strings.TrimSpace(in.RegistrationNumber)

	if err := in.Validate(); err != nil {
		msg, _ := models.FirstError(err)
		return models.AuthHolder{}, errors.New(msg)
	}

	const q = `
		INSERT INTO auth_holder (id, name, registration_number)
		VALUES ($1, $2, NULLIF($3, ''))
		RETURNING id, name, registration_number, created_at, updated_at
	`
	var out models.AuthHolder
	if err := pgxscan.Get(ctx, db, &out, q, in.ID, in.Name, in.RegistrationNumber); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.AuthHolder{}, errors.New("auth holder already exists")
		}
		return models.AuthHolder{}, err
	}
	return out, nil
}

// AddMarketingAuthorization creates a marketing authorization (MA).
func AddMarketingAuthorization(ctx context.Context, db *pgxpool.Pool, in models.MarketingAuthorization) (models.MarketingAuthorization, error) {
	in.ID = uuid.NewString()
	in.Name = strings.TrimSpace(in.Name)
	in.Country = strings.ToUpper(strings.TrimSpace(in.Country))

	if err := in.Validate(); err != nil {
		msg, _ := models.FirstError(err)
		return models.MarketingAuthorization{}, errors.New(msg)
	}

	const q = `
		INSERT INTO marketing_authorization (id, name, country)
		VALUES ($1, $2, $3)
		RETURNING id, name, country, created_at, updated_at
	`
	var out models.MarketingAuthorization
	if err := pgxscan.Get(ctx, db, &out, q, in.ID, in.Name, in.Country); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.MarketingAuthorization{}, errors.New("marketing authorization already exists")
		}
		return models.MarketingAuthorization{}, err
	}
	return out, nil
}

// AddManufacturingSite creates a manufacturing site.
func AddManufacturingSite(ctx context.Context, db *pgxpool.Pool, in models.ManufacturingSite) (models.ManufacturingSite, error) {
	in.ID = uuid.NewString()
	in.Name = strings.TrimSpace(in.Name)
	in.Country = strings.ToUpper(strings.TrimSpace(in.Country))

	if err := in.Validate(); err != nil {
		msg, _ := models.FirstError(err)
		return models.ManufacturingSite{}, errors.New(msg)
	}

	const q = `
		INSERT INTO manufacturing_site (id, name, country)
		VALUES ($1, $2, $3)
		RETURNING id, name, country, created_at, updated_at
	`
	var out models.ManufacturingSite
	if err := pgxscan.Get(ctx, db, &out, q, in.ID, in.Name, in.Country); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.ManufacturingSite{}, errors.New("manufacturing site already exists")
		}
		return models.ManufacturingSite{}, err
	}
	return out, nil
}
