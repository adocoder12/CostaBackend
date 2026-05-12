package repository

import "errors"

var (
	// ErrNotFound is returned when a requested record does not exist.
	ErrNotFound = errors.New("not found")

	// ErrConflict is returned when a unique constraint is violated.
	// e.g. duplicate license_number on apartments.
	// Handler maps this → 409.
	ErrConflict = errors.New("conflict")

	// ErrForbidden is returned when a user tries to access a resource
	// they do not own. Handler maps this → 403.
	ErrForbidden = errors.New("forbidden")
)
