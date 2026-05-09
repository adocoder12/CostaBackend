package model

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID          uuid.UUID
	ApartmentID uuid.UUID
	GuestName   string
	CheckIn     time.Time
	CheckOut    time.Time
	Status      string // upcoming, active, completed, cancelled
	Registered  bool   // RD 933/2021 compliance flag
	GuestCount  int
	Notes       *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
