
package models

import "time"

type ManufacturingSite struct {
    ID        string     `json:"id" db:"id" validate:"omitempty,uuid4"`
    Name      string     `json:"name" db:"name" validate:"required,notblank,max=200"`
    Country   string     `json:"country" db:"country" validate:"required,alpha,uppercase,len=2"`
    CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
    UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

func (m *ManufacturingSite) Validate() error { return validate.Struct(m) }
