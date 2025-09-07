package models

import (
	"encoding/json"
	"time"
)

type OutboxRow struct {
	ID          int64           `json:"id"           db:"id"            validate:"required,gt=0"`
	Aggregate   string          `json:"aggregate"    db:"aggregate"     validate:"required,oneof=DRUG PACK BATCH AUTH INN"`
	AggregateID string          `json:"aggregate_id" db:"aggregate_id"  validate:"required,uuid4"`
	Op          string          `json:"op"           db:"op"            validate:"required,oneof=UPSERT DELETE"`
	Payload     json.RawMessage `json:"payload"      db:"payload"       validate:"omitempty"`
	CreatedAt   time.Time       `json:"created_at"   db:"created_at"    validate:"required"`
}

func (m OutboxRow) Validate() (string, bool) { return FirstError(validate.Struct(m)) }
