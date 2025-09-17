
package models


type Dosage struct {
	ID   int64  `json:"id" db:"id" validate:"omitempty"`
	Code string `json:"code" db:"code" validate:"required,max=50"`
	Name string `json:"name" db:"name" validate:"required,max=100"`
}
func (m *Dosage) Validate() error { return validate.Struct(m) }
