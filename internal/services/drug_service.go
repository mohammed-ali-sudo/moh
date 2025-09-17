package services

import (
	"context"
	"errors"
	"strings"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"moh/models"
)

// DrugPatch is used for PATCH updates (collapse semantics).
// Strings: "" keeps old value. Int IDs: 0 keeps old value.
type DrugPatch struct {
	BrandName           string `json:"brand_name"`
	Dose                string `json:"dose"`
	APIID               int64  `json:"api_id"`
	DosageID            int64  `json:"dosage_id"`
	RouteID             int64  `json:"route_id"`
	StrengthID          int64  `json:"strength_id"`
	AuthorityHolderID   int64  `json:"authority_holder_id"`
	ManufacturingSiteID int64  `json:"manufacturing_site_id"`
	MarketingID         int64  `json:"marketing_id"`
}

// AddDrug inserts a new drug and returns it with joined names.
// Note: optional FKs are *int64. nil => NULL on insert.
func AddDrug(ctx context.Context, db *pgxpool.Pool, in models.Drug) (models.DrugOut, error) {
	in.BrandName = strings.TrimSpace(in.BrandName)
	in.Dose = strings.TrimSpace(in.Dose)

	if err := in.Validate(); err != nil {
		msg, _ := models.FirstError(err)
		return models.DrugOut{}, errors.New(msg)
	}

	const q = `
WITH ins AS (
	INSERT INTO drug (
		brand_name, api_id, dosage_id, route_id, strength_id, dose,
		authority_holder_id, manufacturing_site_id, marketing_id
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING id, brand_name, api_id, dosage_id, route_id, strength_id, dose,
	          authority_holder_id, manufacturing_site_id, marketing_id
)
SELECT d.id, d.brand_name, d.dose,
       d.api_id, a.name AS api_name,
       d.dosage_id, dg.name AS dosage_name,
       d.route_id, r.name AS route_name,
       d.strength_id, s.name AS strength_name,
       d.authority_holder_id, ah.name AS authority_holder_name,
       d.manufacturing_site_id, ms.name AS manufacturing_site_name,
       d.marketing_id, mkt.name AS marketing_name
FROM ins d
JOIN api a              ON a.id  = d.api_id
JOIN dosage dg          ON dg.id = d.dosage_id
JOIN route r            ON r.id  = d.route_id
JOIN strength s         ON s.id  = d.strength_id
LEFT JOIN authority_holder   ah  ON ah.id  = d.authority_holder_id
LEFT JOIN manufacturing_site ms  ON ms.id  = d.manufacturing_site_id
LEFT JOIN marketing           mkt ON mkt.id = d.marketing_id
`
	var out models.DrugOut
	if err := pgxscan.Get(ctx, db, &out, q,
		in.BrandName, in.APIID, in.DosageID, in.RouteID, in.StrengthID, in.Dose,
		in.AuthorityHolderID, in.ManufacturingSiteID, in.MarketingID,
	); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23503" {
				return models.DrugOut{}, errors.New("invalid foreign key (api/dosage/route/strength or org IDs)")
			}
		}
		return models.DrugOut{}, err
	}
	return out, nil
}

// ListDrugs returns all drugs joined with human-readable names.
func ListDrugs(ctx context.Context, db *pgxpool.Pool) ([]models.DrugOut, error) {
	const q = `
SELECT d.id, d.brand_name, d.dose,
       d.api_id, a.name AS api_name,
       d.dosage_id, dg.name AS dosage_name,
       d.route_id, r.name AS route_name,
       d.strength_id, s.name AS strength_name,
       d.authority_holder_id,
         ah.name    AS authority_holder_name,
         ah.country AS authority_holder_country,
       d.manufacturing_site_id,
         ms.name    AS manufacturing_site_name,
         ms.country AS manufacturing_site_country,
       d.marketing_id,
         mkt.name   AS marketing_name,
         mkt.country AS marketing_country
FROM drug d
JOIN api a              ON a.id  = d.api_id
JOIN dosage dg          ON dg.id = d.dosage_id
JOIN route r            ON r.id  = d.route_id
JOIN strength s         ON s.id  = d.strength_id
LEFT JOIN authority_holder   ah  ON ah.id  = d.authority_holder_id
LEFT JOIN manufacturing_site ms  ON ms.id  = d.manufacturing_site_id
LEFT JOIN marketing           mkt ON mkt.id = d.marketing_id
ORDER BY d.brand_name ASC, d.id ASC
`
	var out []models.DrugOut
	if err := pgxscan.Select(ctx, db, &out, q); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateDrug updates fields using collapse semantics and returns the joined row.
func UpdateDrug(ctx context.Context, db *pgxpool.Pool, id int64, patch DrugPatch) (models.DrugOut, error) {
	name := strings.TrimSpace(patch.BrandName)
	dose := strings.TrimSpace(patch.Dose)

	const q = `
WITH upd AS (
	UPDATE drug
	   SET brand_name = COALESCE(NULLIF($2, ''), brand_name),
	       dose       = COALESCE(NULLIF($3, ''), dose),
	       api_id     = COALESCE(NULLIF($4, 0), api_id),
	       dosage_id  = COALESCE(NULLIF($5, 0), dosage_id),
	       route_id   = COALESCE(NULLIF($6, 0), route_id),
	       strength_id= COALESCE(NULLIF($7, 0), strength_id),
	       authority_holder_id   = COALESCE(NULLIF($8, 0), authority_holder_id),
	       manufacturing_site_id = COALESCE(NULLIF($9, 0), manufacturing_site_id),
	       marketing_id          = COALESCE(NULLIF($10,0), marketing_id)
	 WHERE id = $1
	 RETURNING id, brand_name, api_id, dosage_id, route_id, strength_id, dose,
	           authority_holder_id, manufacturing_site_id, marketing_id
)
SELECT d.id, d.brand_name, d.dose,
       d.api_id, a.name AS api_name,
       d.dosage_id, dg.name AS dosage_name,
       d.route_id, r.name AS route_name,
       d.strength_id, s.name AS strength_name,
       d.authority_holder_id, ah.name AS authority_holder_name,
       d.manufacturing_site_id, ms.name AS manufacturing_site_name,
       d.marketing_id, mkt.name AS marketing_name
FROM upd d
JOIN api a              ON a.id  = d.api_id
JOIN dosage dg          ON dg.id = d.dosage_id
JOIN route r            ON r.id  = d.route_id
JOIN strength s         ON s.id  = d.strength_id
LEFT JOIN authority_holder   ah  ON ah.id  = d.authority_holder_id
LEFT JOIN manufacturing_site ms  ON ms.id  = d.manufacturing_site_id
LEFT JOIN marketing           mkt ON mkt.id = d.marketing_id
`
	var out models.DrugOut
	if err := pgxscan.Get(ctx, db, &out, q,
		id, name, dose,
		patch.APIID, patch.DosageID, patch.RouteID, patch.StrengthID,
		patch.AuthorityHolderID, patch.ManufacturingSiteID, patch.MarketingID,
	); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23503" {
				return models.DrugOut{}, errors.New("invalid foreign key (api/dosage/route/strength or org IDs)")
			}
		}
		return models.DrugOut{}, err
	}
	return out, nil
}

func ListDrugsWithBatches(ctx context.Context, db *pgxpool.Pool) ([]models.DrugWithBatches, error) {
	// 1) Fetch all drugs (joined with names)
	drugs, err := ListDrugs(ctx, db)
	if err != nil {
		return nil, err
	}
	if len(drugs) == 0 {
		return []models.DrugWithBatches{}, nil
	}

	// 2) Collect IDs
	ids := make([]int64, len(drugs))
	for i, d := range drugs {
		ids[i] = d.ID
	}

	// 3) Fetch all batches for these drugs (joined with drug/api names)
	const qb = `
SELECT b.id, b.drug_id, d.brand_name AS drug_brand_name, a.name AS api_name,
       b.batch_number, b.mfg_date, b.exp_date, b.quantity, b.status, b.price
FROM batch b
JOIN drug d ON d.id = b.drug_id
JOIN api  a ON a.id = d.api_id
WHERE b.drug_id = ANY($1)
ORDER BY b.mfg_date DESC, b.id DESC
`
	var batches []models.BatchOut
	if err := pgxscan.Select(ctx, db, &batches, qb, ids); err != nil {
		return nil, err
	}

	// 4) Group batches by drug_id
	bmap := make(map[int64][]models.BatchOut, len(drugs))
	for _, b := range batches {
		bmap[b.DrugID] = append(bmap[b.DrugID], b)
	}

	// 5) Build response
	out := make([]models.DrugWithBatches, len(drugs))
	for i, d := range drugs {
		out[i] = models.DrugWithBatches{
			DrugOut: d,
			Batches: bmap[d.ID],
		}
	}
	return out, nil
}
