package models

import "time"

type MAHolder struct {
	MAID      string     `json:"ma_id"     db:"ma_id"     validate:"required,uuid4"`
	HolderID  string     `json:"holder_id" db:"holder_id" validate:"required,uuid4"`
	Role      MARole     `json:"role"      db:"role"      validate:"required,oneof=PRIMARY CO_HOLDER LOCAL_AGENT"`
	ValidFrom time.Time  `json:"valid_from" db:"valid_from" validate:"required"`
	ValidTo   *time.Time `json:"valid_to"   db:"valid_to"   validate:"omitempty"`
}

func (m MAHolder) Validate() (string, bool) { return FirstError(validate.Struct(m)) }
