
package models

import "time"

type AuthHolder struct {
    ID                 string     `json:"id" db:"id" validate:"omitempty,uuid4"`
    Name               string     `json:"name" db:"name" validate:"required,notblank,max=200"`
    RegistrationNumber string     `json:"registration_number,omitempty" db:"registration_number" validate:"omitempty,max=100"`
    CreatedAt          *time.Time `json:"created_at,omitempty" db:"created_at"`
    UpdatedAt          *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

func (m *AuthHolder) Validate() error { return validate.Struct(m) }
