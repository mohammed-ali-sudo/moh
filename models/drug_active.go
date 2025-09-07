package models

type DrugActive struct {
	DrugID        string  `json:"drug_id"        db:"drug_id"        validate:"required,uuid4"`
	INNID         string  `json:"inn_id"         db:"inn_id"         validate:"required,uuid4"`
	StrengthValue float64 `json:"strength_value" db:"strength_value" validate:"required,gt=0"`
	StrengthUnit  string  `json:"strength_unit"  db:"strength_unit"  validate:"required,uppercase,alphanum,max=8"`
}

func (m DrugActive) Validate() (string, bool) { return FirstError(validate.Struct(m)) }
