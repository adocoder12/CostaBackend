package dto

import (
	"github.com/adocoder12/Costabackend/internal/models"
	"github.com/google/uuid"
	"time"
)

// CreateApartmentRequest is what we receive from the frontend
type CreateApartmentRequest struct {
	Name          string  `json:"name" binding:"required"`
	Address       string  `json:"address" binding:"required"`
	LicenseNumber string  `json:"license_number" binding:"required"`
	CadastralRef  *string `json:"cadastral_ref"`
	DoorCode      *string `json:"door_code"`
	Notes         *string `json:"notes"`
}

// ApartmentResponse is what we send back to the frontend
type ApartmentResponse struct {
	ID            uuid.UUID  `json:"id"`
	Name          string     `json:"name"`
	Address       string     `json:"address"`
	Status        string     `json:"status"`
	LicenseNumber string     `json:"license_number"`
	NextCheckIn   *time.Time `json:"next_check_in,omitempty"`
	GuestName     string     `json:"guest_name,omitempty"`
}

func (req *CreateApartmentRequest) ToModel() *models.Apartment {
	return &models.Apartment{
		Name:          req.Name,
		Address:       req.Address,
		LicenseNumber: &req.LicenseNumber, // Convert to pointer for the model
		CadastralRef:  req.CadastralRef,
		DoorCode:      req.DoorCode,
		Notes:         req.Notes,
	}
}

// FromModel creates a response DTO from a database model(Response)
func FromModel(m *models.Apartment) ApartmentResponse {
	res := ApartmentResponse{
		ID:          m.ID,
		Name:        m.Name,
		Address:     m.Address,
		Status:      m.Status,
		NextCheckIn: m.NextCheckIn,
	}

	// Handle LicenseNumber pointer to string conversion safely
	if m.LicenseNumber != nil {
		res.LicenseNumber = *m.LicenseNumber
	}

	if m.GuestName != nil {
		res.GuestName = *m.GuestName
	}

	return res
}
