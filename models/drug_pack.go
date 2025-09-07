package models

import "time"

type DrugPack struct {
	ID        string     `json:"id"        db:"id"         validate:"required,uuid4"`
	DrugID    string     `json:"drug_id"   db:"drug_id"    validate:"required,uuid4"`
	GTIN      string     `json:"gtin"      db:"gtin"       validate:"omitempty,numeric,min=8,max=20"`
	PackSize  int        `json:"pack_size" db:"pack_size"  validate:"required,gt=0"`
	PackUnit  string     `json:"pack_unit" db:"pack_unit"  validate:"required,alphanumunicode,max=16"`
	CreatedAt *time.Time `json:"created_at" db:"created_at" validate:"omitempty"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at" validate:"omitempty"`
}

func (m DrugPack) Validate() (string, bool) { return FirstError(validate.Struct(m)) }
