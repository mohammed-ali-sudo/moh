
// models/manufacturing_site.go
package models

type ManufacturingSite struct {
	ID      int64  `json:"id" db:"id" validate:"omitempty"`
	Name    string `json:"name" db:"name" validate:"required,max=150"`
	Country string `json:"country" db:"country" validate:"required,max=100"`
}
func (m *ManufacturingSite) Validate() error { return validate.Struct(m) }
