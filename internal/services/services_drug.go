
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

// AddDrug creates a drug row.
func AddDrug(ctx context.Context, db *pgxpool.Pool, in models.Drug) (models.Drug, error) {
	in.ID = uuid.NewString()
	in.BrandName = strings.TrimSpace(in.BrandName)

	if err := in.Validate(); err != nil {
		msg, _ := models.FirstError(err)
		return models.Drug{}, errors.New(msg)
	}

	const q = `
		INSERT INTO drug (id, brand_name, dosage_form_id, route_id, strength_unit_id, dose)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, brand_name, dosage_form_id, route_id, strength_unit_id, dose, created_at, updated_at
	`
	var out models.Drug
	if err := pgxscan.Get(ctx, db, &out, q, in.ID, in.BrandName, in.DosageFormID, in.RouteID, in.StrengthUnitID, in.Dose); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23503":
				return models.Drug{}, errors.New("invalid foreign key")
			case "23505":
				return models.Drug{}, errors.New("drug already exists")
			}
		}
		return models.Drug{}, err
	}
	return out, nil
}

// AddBatch creates a batch row.
func AddBatch(ctx context.Context, db *pgxpool.Pool, in models.Batch) (models.Batch, error) {
	in.ID = uuid.NewString()

	if err := in.Validate(); err != nil {
		msg, _ := models.FirstError(err)
		return models.Batch{}, errors.New(msg)
	}

	const q = `
		INSERT INTO batch (id, drug_id, drug_registration_id, batch_number, mfg_date, expire_date, qty_in_batch, status, price, recall_reason)
		VALUES ($1, $2, NULLIF($3, ''), $4, $5, $6, $7, $8, $9, $10)
		RETURNING id, drug_id, drug_registration_id, batch_number, mfg_date, expire_date, qty_in_batch, status, price, recall_reason, created_at, updated_at
	`
	var out models.Batch
	if err := pgxscan.Get(ctx, db, &out, q,
		in.ID, in.DrugID, in.DrugRegistrationID, in.BatchNumber, in.MfgDate, in.ExpireDate, in.QtyInBatch, in.Status, in.Price, in.RecallReason,
	); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23503":
				return models.Batch{}, errors.New("invalid foreign key")
			case "23505":
				return models.Batch{}, errors.New("batch already exists")
			}
		}
		return models.Batch{}, err
	}
	return out, nil
}

// AddDrugAPI links a drug to an API with a strength.
func AddDrugAPI(ctx context.Context, db *pgxpool.Pool, in models.DrugAPI) (models.DrugAPI, error) {
	in.ID = uuid.NewString()

	if err := in.Validate(); err != nil {
		msg, _ := models.FirstError(err)
		return models.DrugAPI{}, errors.New(msg)
	}

	const q = `
		INSERT INTO drug_api (id, drug_id, api_id, strength_value, strength_unit_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, drug_id, api_id, strength_value, strength_unit_id
	`
	var out models.DrugAPI
	if err := pgxscan.Get(ctx, db, &out, q, in.ID, in.DrugID, in.APIID, in.StrengthValue, in.StrengthUnitID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23503":
				return models.DrugAPI{}, errors.New("invalid foreign key")
			case "23505":
				return models.DrugAPI{}, errors.New("link already exists")
			}
		}
		return models.DrugAPI{}, err
	}
	return out, nil
}

// AddDrugRegistration creates a drug_registration row.
func AddDrugRegistration(ctx context.Context, db *pgxpool.Pool, in models.DrugRegistration) (models.DrugRegistration, error) {
	in.ID = uuid.NewString()

	if err := in.Validate(); err != nil {
		msg, _ := models.FirstError(err)
		return models.DrugRegistration{}, errors.New(msg)
	}

	const q = `
		INSERT INTO drug_registration (id, drug_id, ma_id, registration_number, status, valid_from, valid_to, is_primary)
		VALUES ($1, $2, $3, NULLIF($4, ''), $5, $6, $7, $8)
		RETURNING id, drug_id, ma_id, registration_number, status, valid_from, valid_to, is_primary, created_at, updated_at
	`
	var out models.DrugRegistration
	if err := pgxscan.Get(ctx, db, &out, q,
		in.ID, in.DrugID, in.MAID, in.RegistrationNumber, in.Status, in.ValidFrom, in.ValidTo, in.IsPrimary,
	); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23503":
				return models.DrugRegistration{}, errors.New("invalid foreign key")
			case "23505":
				return models.DrugRegistration{}, errors.New("registration already exists")
			}
		}
		return models.DrugRegistration{}, err
	}
	return out, nil
}

// AddDrugRegistrationSite links a drug registration to a manufacturing site.
func AddDrugRegistrationSite(ctx context.Context, db *pgxpool.Pool, in models.DrugRegistrationSite) (models.DrugRegistrationSite, error) {
	in.ID = uuid.NewString()

	if err := in.Validate(); err != nil {
		msg, _ := models.FirstError(err)
		return models.DrugRegistrationSite{}, errors.New(msg)
	}

	const q = `
		INSERT INTO drug_registration_site (id, drug_registration_id, site_id, role)
		VALUES ($1, $2, $3, NULLIF($4, ''))
		RETURNING id, drug_registration_id, site_id, role
	`
	var out models.DrugRegistrationSite
	if err := pgxscan.Get(ctx, db, &out, q, in.ID, in.DrugRegistrationID, in.SiteID, in.Role); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23503":
				return models.DrugRegistrationSite{}, errors.New("invalid foreign key")
			case "23505":
				return models.DrugRegistrationSite{}, errors.New("link already exists")
			}
		}
		return models.DrugRegistrationSite{}, err
	}
	return out, nil
}

// AddDrugRegistrationAuthHolder links a drug registration to an auth holder.
func AddDrugRegistrationAuthHolder(ctx context.Context, db *pgxpool.Pool, in models.DrugRegistrationAuthHolder) (models.DrugRegistrationAuthHolder, error) {
	in.ID = uuid.NewString()

	if err := in.Validate(); err != nil {
		msg, _ := models.FirstError(err)
		return models.DrugRegistrationAuthHolder{}, errors.New(msg)
	}

	const q = `
		INSERT INTO drug_registration_auth_holder (id, drug_registration_id, auth_holder_id, role)
		VALUES ($1, $2, $3, NULLIF($4, ''))
		RETURNING id, drug_registration_id, auth_holder_id, role
	`
	var out models.DrugRegistrationAuthHolder
	if err := pgxscan.Get(ctx, db, &out, q, in.ID, in.DrugRegistrationID, in.AuthHolderID, in.Role); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23503":
				return models.DrugRegistrationAuthHolder{}, errors.New("invalid foreign key")
			case "23505":
				return models.DrugRegistrationAuthHolder{}, errors.New("link already exists")
			}
		}
		return models.DrugRegistrationAuthHolder{}, err
	}
	return out, nil
}
