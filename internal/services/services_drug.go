package services

import (
	"context"
	"errors"
	"strings"

	"moh/models"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AddDrug(ctx context.Context, db *pgxpool.Pool, in models.Drug) (models.Drug, error) {
	in.ID = uuid.NewString()
	in.BrandName = strings.TrimSpace(in.BrandName)

	if err := in.Validate(); err != nil {
		msg, _ := models.FirstError(err)
		return models.Drug{}, errors.New(msg)
	}

	const q = `
        INSERT INTO public.drugs
            (id, brand_name, dosage_form_id, route_id, strength_unit_id, dose, api_id)
        VALUES
            ($1, $2, $3, $4, $5, $6, $7)
        RETURNING
            id, brand_name, dosage_form_id, route_id, strength_unit_id, dose, api_id, created_at, updated_at
    `
	var out models.Drug
	if err := pgxscan.Get(ctx, db, &out, q,
		in.ID, in.BrandName, in.DosageFormID, in.RouteID, in.StrengthUnitID, in.Dose, in.APIID,
	); err != nil {
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

func ListDosageForms(ctx context.Context, db *pgxpool.Pool) ([]models.DosageForm, error) {
	const q = `SELECT id, code, name, created_at, updated_at
	           FROM public.dosage_forms ORDER BY code`
	var out []models.DosageForm
	return out, pgxscan.Select(ctx, db, &out, q)
}

func ListStrengthUnits(ctx context.Context, db *pgxpool.Pool) ([]models.StrengthUnit, error) {
	const q = `SELECT id, code, name, created_at, updated_at
	           FROM public.strength_units ORDER BY code`
	var out []models.StrengthUnit
	return out, pgxscan.Select(ctx, db, &out, q)
}

func ListRoutesOfAdmin(ctx context.Context, db *pgxpool.Pool) ([]models.RouteOfAdmin, error) {
	const q = `SELECT id, code, name, created_at, updated_at
	           FROM public.routes_of_admin ORDER BY code`
	var out []models.RouteOfAdmin
	return out, pgxscan.Select(ctx, db, &out, q)
}

func ListAPIs(ctx context.Context, db *pgxpool.Pool) ([]models.API, error) {
	const q = `SELECT id, name, status, created_at, updated_at
	           FROM public.apis ORDER BY lower(name)`
	var out []models.API
	return out, pgxscan.Select(ctx, db, &out, q)
}

func ListManufacturingSites(ctx context.Context, db *pgxpool.Pool) ([]models.ManufacturingSite, error) {
	const q = `SELECT id, name, country, created_at, updated_at
	           FROM public.manufacturing_sites ORDER BY lower(name), country`
	var out []models.ManufacturingSite
	return out, pgxscan.Select(ctx, db, &out, q)
}

func ListAuthHolders(ctx context.Context, db *pgxpool.Pool) ([]models.AuthHolder, error) {
	const q = `SELECT id, name, registration_number, created_at, updated_at
	           FROM public.auth_holders ORDER BY lower(name)`
	var out []models.AuthHolder
	return out, pgxscan.Select(ctx, db, &out, q)
}

func ListMarketingAuthorizations(ctx context.Context, db *pgxpool.Pool) ([]models.MarketingAuthorization, error) {
	const q = `SELECT id, name, country, created_at, updated_at
	           FROM public.marketing_authorizations ORDER BY lower(name), country`
	var out []models.MarketingAuthorization
	return out, pgxscan.Select(ctx, db, &out, q)
}

// ===== Domain tables =====

func ListDrugs(ctx context.Context, db *pgxpool.Pool) ([]models.Drug, error) {
	// Make sure models.Drug has APIID field: `json:"api_id" db:"api_id"`
	const q = `SELECT id, brand_name, dosage_form_id, route_id, strength_unit_id, dose, api_id, created_at, updated_at
	           FROM public.drugs ORDER BY lower(brand_name)`
	var out []models.Drug
	return out, pgxscan.Select(ctx, db, &out, q)
}

func ListBatches(ctx context.Context, db *pgxpool.Pool) ([]models.Batch, error) {
	const q = `SELECT id, drug_id, drug_registration_id, batch_number, mfg_date, expire_date,
	              qty_in_batch, status, price, recall_reason, created_at, updated_at
	           FROM public.batches ORDER BY expire_date DESC, batch_number`
	var out []models.Batch
	return out, pgxscan.Select(ctx, db, &out, q)
}

func ListDrugRegistrations(ctx context.Context, db *pgxpool.Pool) ([]models.DrugRegistration, error) {
	const q = `SELECT id, drug_id, ma_id, registration_number, status, valid_from, valid_to,
	              is_primary, created_at, updated_at
	           FROM public.drug_registrations ORDER BY valid_to DESC`
	var out []models.DrugRegistration
	return out, pgxscan.Select(ctx, db, &out, q)
}

func ListDrugRegistrationSites(ctx context.Context, db *pgxpool.Pool) ([]models.DrugRegistrationSite, error) {
	const q = `SELECT id, drug_registration_id, site_id, role
	           FROM public.drug_registration_sites ORDER BY role NULLS LAST`
	var out []models.DrugRegistrationSite
	return out, pgxscan.Select(ctx, db, &out, q)
}

func ListDrugRegistrationAuthHolders(ctx context.Context, db *pgxpool.Pool) ([]models.DrugRegistrationAuthHolder, error) {
	const q = `SELECT id, drug_registration_id, auth_holder_id, role
	           FROM public.drug_registration_auth_holders ORDER BY role NULLS LAST`
	var out []models.DrugRegistrationAuthHolder
	return out, pgxscan.Select(ctx, db, &out, q)
}
