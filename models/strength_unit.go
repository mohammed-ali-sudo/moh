
package models

import "time"

type StrengthUnit struct {
    ID        string     `json:"id" db:"id" validate:"omitempty,uuid4"`
    Code      string     `json:"code" db:"code" validate:"required,notblank,max=32"`
    Name      string     `json:"name" db:"name" validate:"required,notblank,max=120"`
    CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
    UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

func (m *StrengthUnit) Validate() error { return validate.Struct(m) }
