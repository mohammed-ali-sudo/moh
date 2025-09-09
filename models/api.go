
package models

import "time"

type API struct {
    ID        string     `json:"id" db:"id" validate:"omitempty,uuid4"`
    Name      string     `json:"name" db:"name" validate:"required,notblank,max=200"`
    Status    APIStatus  `json:"status" db:"status" validate:"required,oneof=active inactive withdrawn banned"`
    CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
    UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

func (m *API) Validate() error { return validate.Struct(m) }
