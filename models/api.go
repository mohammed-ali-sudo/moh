package models

// models/api.go

type API struct {
	ID             int64  `json:"id" db:"id" validate:"omitempty"`
	Name           string `json:"name" db:"name" validate:"required,max=50"`
	IsPsychotropic bool   `json:"ispsychotropic" db:"ispsychotropic"`
}

func (m *API) Validate() error { return validate.Struct(m) }
