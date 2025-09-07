package models

import "time"

// models/manufacturer_site.go
type ManufacturerSite struct {
	ID          string     `json:"id,omitempty"   db:"id"            validate:"omitempty,uuid4"`
	Name        string     `json:"name"           db:"name"          validate:"required,min=2,max=200"`
	CompanyName string     `json:"company_name"   db:"company_name"  validate:"required,min=2,max=200"`
	CountryISO  string     `json:"country_code"   db:"country_code"  validate:"required,alpha,len=2,uppercase"`
	Address1    string     `json:"address_line1"  db:"address_line1" validate:"required,min=2,max=200"`
	City        string     `json:"city"           db:"city"          validate:"required,min=2,max=120"`
	PostalCode  string     `json:"postal_code"    db:"postal_code"   validate:"omitempty,alphanumunicode,max=20"`
	CreatedAt   *time.Time `json:"created_at"     db:"created_at"    validate:"omitempty"`
	UpdatedAt   *time.Time `json:"updated_at"     db:"updated_at"    validate:"omitempty"`
}

func (m ManufacturerSite) Validate() (string, bool) { return FirstError(validate.Struct(m)) }
