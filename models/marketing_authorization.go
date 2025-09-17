
// models/marketing.go
package models

type Marketing struct {
	ID      int64  `json:"id" db:"id" validate:"omitempty"`
	Name    string `json:"name" db:"name" validate:"required,max=150"`
	Country string `json:"country" db:"country" validate:"required,max=100"`
}
func (m *Marketing) Validate() error { return validate.Struct(m) }
