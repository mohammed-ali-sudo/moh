package models

import "time"

type INN struct {
	ID        string     `json:"id"         db:"id"          validate:"required,uuid4"`
	Name      string     `json:"name"       db:"name"        validate:"required,min=2,max=120"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"  validate:"omitempty"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"  validate:"omitempty"`
}

func (m INN) Validate() (string, bool) { return FirstError(validate.Struct(m)) }
