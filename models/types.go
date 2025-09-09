
package models

type APIStatus string
const (
    APIStatusActive    APIStatus = "active"
    APIStatusInactive  APIStatus = "inactive"
    APIStatusWithdrawn APIStatus = "withdrawn"
    APIStatusBanned    APIStatus = "banned"
)

type RegistrationStatus string
const (
    RegistrationActive    RegistrationStatus = "active"
    RegistrationSuspended RegistrationStatus = "suspended"
    RegistrationExpired   RegistrationStatus = "expired"
    RegistrationWithdrawn RegistrationStatus = "withdrawn"
)

type BatchStatus string
const (
    BatchPlanned  BatchStatus = "planned"
    BatchReleased BatchStatus = "released"
    BatchOnHold   BatchStatus = "on_hold"
    BatchRecalled BatchStatus = "recalled"
    BatchExpired  BatchStatus = "expired"
    BatchSoldOut  BatchStatus = "sold_out"
    BatchInactive BatchStatus = "inactive"
)
