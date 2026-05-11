package models

import (
	"time"

	"github.com/google/uuid"
)

type Apartment struct {
	ID            uuid.UUID  `json:"id"`
	Name          string     `json:"name"`
	Address       string     `json:"address"`
	Status        string     `json:"status"` // clean, dirty, etc.
	LicenseNumber string     `json:"license_number"`
	CadastralRef  *string    `json:"cadastral_ref,omitempty"`
	DoorCode      *string    `json:"door_code,omitempty"`
	NextCheckIn   *time.Time `json:"next_check_in,omitempty"`
	CheckoutDate  *time.Time `json:"checkout_date,omitempty"`
	GuestName     *string    `json:"guest_name,omitempty"`
	Notes         *string    `json:"notes,omitempty"`
	OwnerID       *uuid.UUID `json:"owner_id,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}
