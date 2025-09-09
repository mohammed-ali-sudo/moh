
package models

import "time"

type DrugRegistration struct {
    ID                 string             `json:"id" db:"id" validate:"omitempty,uuid4"`
    DrugID             string             `json:"drug_id" db:"drug_id" validate:"required,uuid4"`
    MAID               string             `json:"ma_id" db:"ma_id" validate:"required,uuid4"`
    RegistrationNumber string             `json:"registration_number,omitempty" db:"registration_number" validate:"omitempty,max=100"`
    Status             RegistrationStatus `json:"status" db:"status" validate:"required,oneof=active suspended expired withdrawn"`
    ValidFrom          time.Time          `json:"valid_from" db:"valid_from"`
    ValidTo            time.Time          `json:"valid_to" db:"valid_to"`
    IsPrimary          bool               `json:"is_primary" db:"is_primary"`
    CreatedAt          *time.Time         `json:"created_at,omitempty" db:"created_at"`
    UpdatedAt          *time.Time         `json:"updated_at,omitempty" db:"updated_at"`
}

func (m *DrugRegistration) Validate() error { return validate.Struct(m) }
