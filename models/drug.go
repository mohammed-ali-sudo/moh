package models

import "time"

type Drug struct {
	ID             string     `json:"id"               db:"id"                validate:"required,uuid4"`
	MoHCode        string     `json:"moh_code"         db:"moh_code"          validate:"required,alphanumunicode,max=64"`
	BrandName      string     `json:"brand_name"       db:"brand_name"        validate:"required,min=2,max=160"`
	DosageFormCode string     `json:"dosage_form_code" db:"dosage_form_code"  validate:"required,uppercase,alphanum,max=16"`
	RouteCode      string     `json:"route_code"       db:"route_code"        validate:"required,uppercase,alphanum,max=16"`
	AIFingerprint  string     `json:"ai_fingerprint"   db:"ai_fingerprint"    validate:"required,min=8,max=200"`
	IsPsychotropic bool       `json:"is_psychotropic"  db:"is_psychotropic"   validate:"omitempty"`
	IsControlled   bool       `json:"is_controlled"    db:"is_controlled"     validate:"omitempty"`
	IsAntibiotic   bool       `json:"is_antibiotic"    db:"is_antibiotic"     validate:"omitempty"`
	CreatedAt      *time.Time `json:"created_at"       db:"created_at"        validate:"omitempty"`
	UpdatedAt      *time.Time `json:"updated_at"       db:"updated_at"        validate:"omitempty"`
}

func (m Drug) Validate() (string, bool) { return FirstError(validate.Struct(m)) }
