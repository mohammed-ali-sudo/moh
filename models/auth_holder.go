
// models/authority_holder.go
package models

type AuthorityHolder struct {
	ID      int64  `json:"id" db:"id" validate:"omitempty"`
	Name    string `json:"name" db:"name" validate:"required,max=150"`
	Country string `json:"country" db:"country" validate:"required,max=100"`
}
func (m *AuthorityHolder) Validate() error { return validate.Struct(m) }
