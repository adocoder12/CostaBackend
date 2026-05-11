package models

import (
	"time"

	"github.com/google/uuid"
)

type Guest struct {
	ID            uuid.UUID
	BookingID     uuid.UUID
	Name          string
	Surname       string
	IDType        string // dni, passport, nie
	IDNumber      string
	SupportNumber string
	Nationality   string // ISO 3166-1 alpha-2
	BirthDate     time.Time
	Residence     string
	SignatureURL  string // S3 key
	RegisteredAt  time.Time
	CreatedAt     time.Time
}
