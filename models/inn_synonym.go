package models

type INNSynonym struct {
	INNID   string `json:"inn_id"  db:"inn_id"  validate:"required,uuid4"`
	Synonym string `json:"synonym" db:"synonym" validate:"required,min=1,max=120"`
}

func (m INNSynonym) Validate() (string, bool) { return FirstError(validate.Struct(m)) }
