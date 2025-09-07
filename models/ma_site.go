package models

import "time"

type MASite struct {
	MAID      string     `json:"ma_id"   db:"ma_id"   validate:"required,uuid4"`
	SiteID    string     `json:"site_id" db:"site_id" validate:"required,uuid4"`
	Role      MASiteRole `json:"role"    db:"role"    validate:"required,oneof=API FINISHED_DOSE PACKER RELEASE_SITE"`
	ValidFrom time.Time  `json:"valid_from" db:"valid_from" validate:"required"`
	ValidTo   *time.Time `json:"valid_to"   db:"valid_to"   validate:"omitempty"`
}

func (m MASite) Validate() (string, bool) { return FirstError(validate.Struct(m)) }
