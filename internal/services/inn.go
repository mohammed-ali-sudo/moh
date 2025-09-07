package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"moh/models"
)

// AddINN inserts a new INN with a server-generated UUID.
func AddINN(ctx context.Context, db *pgxpool.Pool, in models.INN) (models.INN, error) {
	// Always generate server-side ID
	in.ID = uuid.NewString()

	// Validate after ID is set (uses your FirstError+validator)
	if msg, ok := in.Validate(); !ok {
		return models.INN{}, errors.New(msg)
	}

	const q = `
		INSERT INTO inn (id, name)
		VALUES ($1, $2)
		RETURNING created_at, updated_at
	`
	var createdAt, updatedAt time.Time
	if err := db.QueryRow(ctx, q, in.ID, in.Name).Scan(&createdAt, &updatedAt); err != nil {
		return models.INN{}, err
	}

	in.CreatedAt = &createdAt
	in.UpdatedAt = &updatedAt
	return in, nil
}
