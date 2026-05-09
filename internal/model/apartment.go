package model

import (
	"time"

	"github.com/google/uuid"
)

type Apartment struct {
	ID            uuid.UUID
	Name          string
	Address       string
	Status        string // clean, dirty, in_progress, maintenance
	LicenseNumber string
	CadastralRef  *string
	DoorCode      *string
	NextCheckIn   *time.Time
	CheckoutDate  *time.Time
	GuestName     *string
	Notes         *string
	OwnerID       uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
