
package models

type DrugRegistrationAuthHolder struct {
    ID                 string `json:"id" db:"id" validate:"omitempty,uuid4"`
    DrugRegistrationID string `json:"drug_registration_id" db:"drug_registration_id" validate:"required,uuid4"`
    AuthHolderID       string `json:"auth_holder_id" db:"auth_holder_id" validate:"required,uuid4"`
    Role               string `json:"role,omitempty" db:"role" validate:"omitempty,max=80"`
}

func (m *DrugRegistrationAuthHolder) Validate() error { return validate.Struct(m) }
