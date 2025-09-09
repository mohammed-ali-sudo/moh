
package models

type DrugAPI struct {
    ID             string  `json:"id" db:"id" validate:"omitempty,uuid4"`
    DrugID         string  `json:"drug_id" db:"drug_id" validate:"required,uuid4"`
    APIID          string  `json:"api_id" db:"api_id" validate:"required,uuid4"`
    StrengthValue  float64 `json:"strength_value" db:"strength_value" validate:"required,gt=0"`
    StrengthUnitID string  `json:"strength_unit_id" db:"strength_unit_id" validate:"required,uuid4"`
}

func (m *DrugAPI) Validate() error { return validate.Struct(m) }
