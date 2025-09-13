// internal/domain/models/drug.go
package models

import "time"

type Drug struct {
	ID             string     `json:"id" db:"id" validate:"omitempty,uuid4"`
	BrandName      string     `json:"brand_name" db:"brand_name" validate:"required,notblank,max=200"`
	DosageFormID   string     `json:"dosage_form_id" db:"dosage_form_id" validate:"required,uuid4"`
	RouteID        string     `json:"route_id" db:"route_id" validate:"required,uuid4"`
	StrengthUnitID string     `json:"strength_unit_id" db:"strength_unit_id" validate:"required,uuid4"`
	Dose           float64    `json:"dose" db:"dose" validate:"required,gt=0"`
	APIID          string     `json:"api_id" db:"api_id" validate:"required,uuid4"` // NEW
	CreatedAt      *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

func (m *Drug) Validate() error { return validate.Struct(m) }
