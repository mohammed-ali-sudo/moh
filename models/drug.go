package models

type Drug struct {
	ID                  int64  `json:"id" db:"id" validate:"omitempty"`
	BrandName           string `json:"brand_name" db:"brand_name" validate:"required,notblank,max=150"`
	APIID               int64  `json:"api_id" db:"api_id" validate:"required,gt=0"`
	DosageID            int64  `json:"dosage_id" db:"dosage_id" validate:"required,gt=0"`
	RouteID             int64  `json:"route_id" db:"route_id" validate:"required,gt=0"`
	StrengthID          int64  `json:"strength_id" db:"strength_id" validate:"required,gt=0"`
	Dose                string `json:"dose" db:"dose" validate:"required,notblank,max=50"`
	AuthorityHolderID   *int64 `json:"authority_holder_id,omitempty" db:"authority_holder_id" validate:"omitempty,gt=0"`
	ManufacturingSiteID *int64 `json:"manufacturing_site_id,omitempty" db:"manufacturing_site_id" validate:"omitempty,gt=0"`
	MarketingID         *int64 `json:"marketing_id,omitempty" db:"marketing_id" validate:"omitempty,gt=0"`
}

func (m *Drug) Validate() error { return validate.Struct(m) }

// DrugOut is a read DTO with names joined from referenced tables.
// internal/domain/models/drug.go (append fields to DrugOut)
type DrugOut struct {
	ID        int64  `json:"id" db:"id"`
	BrandName string `json:"brand_name" db:"brand_name"`
	Dose      string `json:"dose" db:"dose"`

	APIID   int64  `json:"api_id" db:"api_id"`
	APIName string `json:"api_name" db:"api_name"`

	DosageID   int64  `json:"dosage_id" db:"dosage_id"`
	DosageName string `json:"dosage_name" db:"dosage_name"`

	RouteID   int64  `json:"route_id" db:"route_id"`
	RouteName string `json:"route_name" db:"route_name"`

	StrengthID   int64  `json:"strength_id" db:"strength_id"`
	StrengthName string `json:"strength_name" db:"strength_name"`

	AuthorityHolderID      *int64  `json:"authority_holder_id,omitempty" db:"authority_holder_id"`
	AuthorityHolderName    *string `json:"authority_holder_name,omitempty" db:"authority_holder_name"`
	AuthorityHolderCountry *string `json:"authority_holder_country,omitempty" db:"authority_holder_country"`

	ManufacturingSiteID      *int64  `json:"manufacturing_site_id,omitempty" db:"manufacturing_site_id"`
	ManufacturingSiteName    *string `json:"manufacturing_site_name,omitempty" db:"manufacturing_site_name"`
	ManufacturingSiteCountry *string `json:"manufacturing_site_country,omitempty" db:"manufacturing_site_country"`

	MarketingID      *int64  `json:"marketing_id,omitempty" db:"marketing_id"`
	MarketingName    *string `json:"marketing_name,omitempty" db:"marketing_name"`
	MarketingCountry *string `json:"marketing_country,omitempty" db:"marketing_country"`
}

type DrugWithBatches struct {
	DrugOut
	Batches []BatchOut `json:"batches"`
}
