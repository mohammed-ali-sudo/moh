package models

type RouteOfAdmin struct {
	Code string `json:"code" db:"code" validate:"required,uppercase,alphanum,max=16"`
	Name string `json:"name" db:"name" validate:"required,min=2,max=80"`
}

func (m RouteOfAdmin) Validate() (string, bool) { return FirstError(validate.Struct(m)) }
