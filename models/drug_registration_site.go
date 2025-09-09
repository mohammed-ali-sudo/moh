
package models

type DrugRegistrationSite struct {
    ID                 string `json:"id" db:"id" validate:"omitempty,uuid4"`
    DrugRegistrationID string `json:"drug_registration_id" db:"drug_registration_id" validate:"required,uuid4"`
    SiteID             string `json:"site_id" db:"site_id" validate:"required,uuid4"`
    Role               string `json:"role,omitempty" db:"role" validate:"omitempty,max=80"`
}

func (m *DrugRegistrationSite) Validate() error { return validate.Struct(m) }
