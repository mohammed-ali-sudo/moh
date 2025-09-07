package models

import "time"

type AuthHolder struct {
	ID         string     `json:"id"           db:"id"            validate:"required,uuid4"`
	Name       string     `json:"name"         db:"name"          validate:"required,min=2,max=200"`
	RegNumber  string     `json:"reg_number"   db:"reg_number"    validate:"required,alphanumunicode,max=60"`
	CountryISO string     `json:"country_code" db:"country_code"  validate:"required,alpha,len=2,uppercase"`
	CreatedAt  *time.Time `json:"created_at"   db:"created_at"    validate:"omitempty"`
	UpdatedAt  *time.Time `json:"updated_at"   db:"updated_at"    validate:"omitempty"`
}

func (m AuthHolder) Validate() (string, bool) { return FirstError(validate.Struct(m)) }
