package models

import "time"

type Batch struct {
	ID                 string      `json:"id"                    db:"id"                      validate:"required,uuid4"`
	DrugPackID         string      `json:"drug_pack_id"          db:"drug_pack_id"            validate:"required,uuid4"`
	ManufacturerSiteID string      `json:"manufacturer_site_id"  db:"manufacturer_site_id"    validate:"required,uuid4"`
	LotNo              string      `json:"lot_no"                db:"lot_no"                  validate:"required,min=1,max=64"`
	MfgDate            *time.Time  `json:"mfg_date"              db:"mfg_date"                validate:"omitempty"`
	ExpiryDate         time.Time   `json:"expiry_date"           db:"expiry_date"             validate:"required"`
	Status             BatchStatus `json:"status"                db:"status"                  validate:"required,oneof=RELEASED QUARANTINE RECALLED EXPIRED"`
	RecallReason       string      `json:"recall_reason"         db:"recall_reason"           validate:"omitempty,min=3,max=300"`
	QtyManufactured    int64       `json:"qty_manufactured"      db:"qty_manufactured"        validate:"required,gte=0"`
	CreatedAt          *time.Time  `json:"created_at"            db:"created_at"              validate:"omitempty"`
	UpdatedAt          *time.Time  `json:"updated_at"            db:"updated_at"              validate:"omitempty"`
}

func (m Batch) Validate() (string, bool) { return FirstError(validate.Struct(m)) }
