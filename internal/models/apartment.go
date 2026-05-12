package models

import (
	"time"

	"github.com/google/uuid"
)

// Apartment mirrors the apartments table exactly.
// No JSON tags — models are never serialised directly to HTTP responses.
// No computed fields like IsMine — those live in the DTO layer.
type Apartment struct {
	ID            uuid.UUID  `db:"id"`
	Name          string     `db:"name"`
	Address       string     `db:"address"`
	Status        string     `db:"status"`
	LicenseNumber *string    `db:"license_number"`
	CadastralRef  *string    `db:"cadastral_ref"`
	DoorCode      *string    `db:"door_code"`
	NextCheckIn   *time.Time `db:"next_check_in"`
	CheckoutDate  *time.Time `db:"checkout_date"`
	GuestName     *string    `db:"guest_name"`
	Notes         *string    `db:"notes"`
	OwnerID       *uuid.UUID `db:"owner_id"`
	CreatedAt     time.Time  `db:"created_at"`
	UpdatedAt     time.Time  `db:"updated_at"`
}

// Status constants — used in service layer for defaults and validation.
const (
	StatusClean       = "clean"
	StatusDirty       = "dirty"
	StatusInProgress  = "in_progress"
	StatusMaintenance = "maintenance"
	StatusBlocked     = "blocked"
)
