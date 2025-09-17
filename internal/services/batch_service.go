package services

import (
	"context"
	"errors"
	"strings"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jackc/pgx/v5/pgconn"

	"moh/models"
)

// Patch struct with simple collapse semantics:
// - Strings: "" keeps old
// - Int/Float: 0 keeps old (note: cannot set quantity/price to 0 via PATCH with this simple scheme)
type BatchPatch struct {
	DrugID      int64   `json:"drug_id"`
	BatchNumber string  `json:"batch_number"`
	MfgDate     string  `json:"mfg_date"` // "YYYY-MM-DD", "" to keep
	ExpDate     string  `json:"exp_date"` // "YYYY-MM-DD", "" to keep
	Quantity    int64   `json:"quantity"`
	Status      string  `json:"status"`
	Price       float64 `json:"price"`
}

// AddBatch inserts and returns a joined row (with drug & api names)
func AddBatch(ctx context.Context, db *pgxpool.Pool, in models.Batch) (models.BatchOut, error) {
	in.BatchNumber = strings.TrimSpace(in.BatchNumber)
	in.Status = strings.TrimSpace(in.Status)

	if err := in.Validate(); err != nil {
		msg, _ := models.FirstError(err)
		return models.BatchOut{}, errors.New(msg)
	}

	const q = `
WITH ins AS (
	INSERT INTO batch (drug_id, batch_number, mfg_date, exp_date, quantity, status, price)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id, drug_id, batch_number, mfg_date, exp_date, quantity, status, price
)
SELECT b.id, b.drug_id, d.brand_name AS drug_brand_name, a.name AS api_name,
       b.batch_number, b.mfg_date, b.exp_date, b.quantity, b.status, b.price
FROM ins b
JOIN drug d ON d.id = b.drug_id
JOIN api  a ON a.id = d.api_id
`
	var out models.BatchOut
	if err := pgxscan.Get(ctx, db, &out, q,
		in.DrugID, in.BatchNumber, in.MfgDate, in.ExpDate, in.Quantity, in.Status, in.Price,
	); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return models.BatchOut{}, errors.New("invalid foreign key (drug_id)")
		}
		return models.BatchOut{}, err
	}
	return out, nil
}

// ListBatches returns all batches with joined drug & api names
func ListBatches(ctx context.Context, db *pgxpool.Pool) ([]models.BatchOut, error) {
	const q = `
SELECT b.id, b.drug_id, d.brand_name AS drug_brand_name, a.name AS api_name,
       b.batch_number, b.mfg_date, b.exp_date, b.quantity, b.status, b.price
FROM batch b
JOIN drug d ON d.id = b.drug_id
JOIN api  a ON a.id = d.api_id
ORDER BY b.mfg_date DESC, b.id DESC
`
	var out []models.BatchOut
	if err := pgxscan.Select(ctx, db, &out, q); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateBatch updates using collapse semantics and returns joined row
func UpdateBatch(ctx context.Context, db *pgxpool.Pool, id int64, p BatchPatch) (models.BatchOut, error) {
	p.BatchNumber = strings.TrimSpace(p.BatchNumber)
	p.Status = strings.TrimSpace(p.Status)

	const q = `
WITH upd AS (
	UPDATE batch
	   SET drug_id      = COALESCE(NULLIF($2, 0), drug_id),
	       batch_number = COALESCE(NULLIF($3, ''), batch_number),
	       mfg_date     = COALESCE(NULLIF($4, '')::date, mfg_date),
	       exp_date     = COALESCE(NULLIF($5, '')::date, exp_date),
	       quantity     = COALESCE(NULLIF($6, 0), quantity),
	       status       = COALESCE(NULLIF($7, ''), status),
	       price        = COALESCE(NULLIF($8, 0), price)
	 WHERE id = $1
	 RETURNING id, drug_id, batch_number, mfg_date, exp_date, quantity, status, price
)
SELECT b.id, b.drug_id, d.brand_name AS drug_brand_name, a.name AS api_name,
       b.batch_number, b.mfg_date, b.exp_date, b.quantity, b.status, b.price
FROM upd b
JOIN drug d ON d.id = b.drug_id
JOIN api  a ON a.id = d.api_id
`
	var out models.BatchOut
	if err := pgxscan.Get(ctx, db, &out, q,
		id, p.DrugID, p.BatchNumber, p.MfgDate, p.ExpDate, p.Quantity, p.Status, p.Price,
	); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23503":
				return models.BatchOut{}, errors.New("invalid foreign key (drug_id)")
			}
		}
		return models.BatchOut{}, err
	}
	return out, nil
}
