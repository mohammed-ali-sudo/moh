package models

import (
	"encoding/json"
	"time"
)

type AuditLog struct {
	ID        int64           `json:"id"        db:"id"         validate:"required,gt=0"`
	Actor     string          `json:"actor"     db:"actor"      validate:"required,min=1,max=120"`
	Action    string          `json:"action"    db:"action"     validate:"required,min=1,max=80"`
	Entity    string          `json:"entity"    db:"entity"     validate:"required,oneof=DRUG PACK BATCH AUTH INN"`
	EntityID  string          `json:"entity_id" db:"entity_id"  validate:"required,uuid4"`
	Before    json.RawMessage `json:"before"    db:"before"     validate:"omitempty"`
	After     json.RawMessage `json:"after"     db:"after"      validate:"required"`
	CreatedAt time.Time       `json:"created_at" db:"created_at" validate:"required"`
}

func (m AuditLog) Validate() (string, bool) { return FirstError(validate.Struct(m)) }
