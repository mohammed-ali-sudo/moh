package models

import "time"

type Batch struct {
	ID          int64     `json:"id" db:"id" validate:"omitempty"`
	DrugID      int64     `json:"drug_id" db:"drug_id" validate:"required,gt=0"`
	BatchNumber string    `json:"batch_number" db:"batch_number" validate:"required,notblank,max=100"`
	MfgDate     time.Time `json:"mfg_date" db:"mfg_date" validate:"required"`
	ExpDate     time.Time `json:"exp_date" db:"exp_date" validate:"required,gtfield=MfgDate"`
	Quantity    int64     `json:"quantity" db:"quantity" validate:"gte=0"`
	Status      string    `json:"status" db:"status" validate:"required,oneof=planned released on_hold recalled expired sold_out inactive"`
	Price       float64   `json:"price" db:"price" validate:"gte=0"`
}

func (m *Batch) Validate() error { return validate.Struct(m) }

// Read DTO with joined names
type BatchOut struct {
	ID            int64     `json:"id" db:"id"`
	DrugID        int64     `json:"drug_id" db:"drug_id"`
	DrugBrandName string    `json:"drug_brand_name" db:"drug_brand_name"`
	APIName       string    `json:"api_name" db:"api_name"`
	BatchNumber   string    `json:"batch_number" db:"batch_number"`
	MfgDate       time.Time `json:"mfg_date" db:"mfg_date"`
	ExpDate       time.Time `json:"exp_date" db:"exp_date"`
	Quantity      int64     `json:"quantity" db:"quantity"`
	Status        string    `json:"status" db:"status"`
	Price         float64   `json:"price" db:"price"`
}
