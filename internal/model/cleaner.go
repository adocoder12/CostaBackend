package model

import (
	"time"

	"github.com/google/uuid"
)

type Cleaner struct {
	ID        uuid.UUID
	Name      string
	Phone     int64
	Email     string
	Address   string
	Zone      string
	Active    bool
	Verified  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
