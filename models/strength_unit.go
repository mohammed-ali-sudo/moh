package models

type StrengthUnit struct {
	Code string `json:"code" db:"code" validate:"required,uppercase,alphanum,max=8"`
	Name string `json:"name" db:"name" validate:"required,min=1,max=40"`
}

func (m StrengthUnit) Validate() (string, bool) { return FirstError(validate.Struct(m)) }
