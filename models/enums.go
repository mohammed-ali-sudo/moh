package models

type BatchStatus string

const (
	BatchReleased   BatchStatus = "RELEASED"
	BatchQuarantine BatchStatus = "QUARANTINE"
	BatchRecalled   BatchStatus = "RECALLED"
	BatchExpired    BatchStatus = "EXPIRED"
)

type MAStatus string

const (
	MAActive    MAStatus = "ACTIVE"
	MASuspended MAStatus = "SUSPENDED"
	MARevoked   MAStatus = "REVOKED"
	MAExpired   MAStatus = "EXPIRED"
)

type MARole string

const (
	RolePrimary    MARole = "PRIMARY"
	RoleCoHolder   MARole = "CO_HOLDER"
	RoleLocalAgent MARole = "LOCAL_AGENT"
)

type MASiteRole string

const (
	SiteAPI          MASiteRole = "API"
	SiteFinishedDose MASiteRole = "FINISHED_DOSE"
	SitePacker       MASiteRole = "PACKER"
	SiteReleaseSite  MASiteRole = "RELEASE_SITE"
)
