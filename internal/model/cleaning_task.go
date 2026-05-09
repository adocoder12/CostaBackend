package model

import (
	"time"

	"github.com/google/uuid"
)

type CleaningTask struct {
	ID                uuid.UUID
	ApartmentID       uuid.UUID
	AssignedCleanerID *uuid.UUID
	Status            string // pending, in_progress, done
	Priority          string // normal, urgent
	Notes             *string
	PhotoURL          *string // S3 key
	ScheduledFor      time.Time
	CompletedAt       *time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
