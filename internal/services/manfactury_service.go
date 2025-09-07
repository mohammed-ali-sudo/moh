package services

import (
	"context"
	"errors"
	"time"

	"moh/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AddManufacturerSite(ctx context.Context, db *pgxpool.Pool, m models.ManufacturerSite) (models.ManufacturerSite, error) {
	// Always generate a fresh UUID (ignore any client-provided id)
	m.ID = uuid.NewString()

	// Validate after ID is set
	if msg, ok := m.Validate(); !ok {
		return models.ManufacturerSite{}, errors.New(msg)
	}

	const q = `
		INSERT INTO manufacturer_site (
			id, name, company_name, country_code, address_line1, city, postal_code
		) VALUES ($1,$2,$3,$4,$5,$6,$7)
		RETURNING created_at, updated_at
	`

	var createdAt, updatedAt time.Time
	if err := db.QueryRow(ctx, q,
		m.ID, m.Name, m.CompanyName, m.CountryISO, m.Address1, m.City, m.PostalCode,
	).Scan(&createdAt, &updatedAt); err != nil {
		return models.ManufacturerSite{}, err
	}

	m.CreatedAt = &createdAt
	m.UpdatedAt = &updatedAt
	return m, nil
}