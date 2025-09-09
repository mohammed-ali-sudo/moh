
package models

import (
    "encoding/json"
    "time"
)

type Batch struct {
    ID                 string          `json:"id" db:"id" validate:"omitempty,uuid4"`
    DrugID             string          `json:"drug_id" db:"drug_id" validate:"required,uuid4"`
    DrugRegistrationID string          `json:"drug_registration_id,omitempty" db:"drug_registration_id" validate:"omitempty,uuid4"`
    BatchNumber        string          `json:"batch_number" db:"batch_number" validate:"required,notblank,max=120"`
    MfgDate            time.Time       `json:"mfg_date" db:"mfg_date" validate:"required"`
    ExpireDate         time.Time       `json:"expire_date" db:"expire_date" validate:"required"`
    QtyInBatch         int64           `json:"qty_in_batch" db:"qty_in_batch" validate:"gte=0"`
    Status             BatchStatus     `json:"status" db:"status" validate:"required,oneof=planned released on_hold recalled expired sold_out inactive"`
    Price              float64         `json:"price" db:"price" validate:"gte=0"`
    RecallReason       json.RawMessage `json:"recall_reason,omitempty" db:"recall_reason" validate:"omitempty,json"`
    CreatedAt          *time.Time      `json:"created_at,omitempty" db:"created_at"`
    UpdatedAt          *time.Time      `json:"updated_at,omitempty" db:"updated_at"`
}

func (m *Batch) Validate() error { return validate.Struct(m) }
