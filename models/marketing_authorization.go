package models

import "time"

type MarketingAuthorization struct {
	ID         string     `json:"id"          db:"id"           validate:"required,uuid4"`
	DrugID     string     `json:"drug_id"     db:"drug_id"      validate:"required,uuid4"`
	AuthNumber string     `json:"auth_number" db:"auth_number"  validate:"required,max=80"`
	Status     MAStatus   `json:"status"      db:"status"       validate:"required,oneof=ACTIVE SUSPENDED REVOKED EXPIRED"`
	ValidFrom  time.Time  `json:"valid_from"  db:"valid_from"   validate:"required"`
	ValidTo    *time.Time `json:"valid_to"    db:"valid_to"     validate:"omitempty"`
	CreatedAt  *time.Time `json:"created_at"  db:"created_at"   validate:"omitempty"`
	UpdatedAt  *time.Time `json:"updated_at"  db:"updated_at"   validate:"omitempty"`
}

func (m MarketingAuthorization) Validate() (string, bool) { return FirstError(validate.Struct(m)) }
