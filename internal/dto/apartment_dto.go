package dto

import (
	"time"

	"github.com/adocoder12/Costabackend/internal/models"
	"github.com/google/uuid"
)

// CreateApartmentRequest is the inbound payload for POST /apartments.
type CreateApartmentRequest struct {
	Name          string  `json:"name"           binding:"required"`
	Address       string  `json:"address"        binding:"required"`
	LicenseNumber string  `json:"license_number" binding:"required"`
	CadastralRef  *string `json:"cadastral_ref"`
	DoorCode      *string `json:"door_code"`
	Notes         *string `json:"notes"`
}

// UpdateApartmentRequest is the inbound payload for PUT /apartments/:id.
// All fields are optional — only set fields are updated.
type UpdateApartmentRequest struct {
	Name          *string `json:"name"`
	Address       *string `json:"address"`
	LicenseNumber *string `json:"license_number"`
	CadastralRef  *string `json:"cadastral_ref"`
	DoorCode      *string `json:"door_code"`
	Notes         *string `json:"notes"`
	Status        *string `json:"status"`
}

// ApartmentResponse is the outbound payload for all apartment endpoints.
// IsMine and DoorCode masking are applied in the service layer.
type ApartmentResponse struct {
	ID            uuid.UUID  `json:"id"`
	Name          string     `json:"name"`
	Address       string     `json:"address"`
	Status        string     `json:"status"`
	LicenseNumber string     `json:"license_number"`
	NextCheckIn   *time.Time `json:"next_check_in,omitempty"`
	GuestName     string     `json:"guest_name,omitempty"`
	DoorCode      *string    `json:"door_code,omitempty"` // nil if viewer is not the owner
	IsMine        bool       `json:"is_mine"`             // computed in service not in db
}

// ToModel maps a CreateApartmentRequest to an Apartment model for persistence.
func (req *CreateApartmentRequest) ToModel() *models.Apartment {
	return &models.Apartment{
		Name:          req.Name,
		Address:       req.Address,
		LicenseNumber: &req.LicenseNumber,
		CadastralRef:  req.CadastralRef,
		DoorCode:      req.DoorCode,
		Notes:         req.Notes,
	}
}

// FromModel maps an Apartment model to an ApartmentResponse DTO.
func FromModel(m *models.Apartment) ApartmentResponse {
	res := ApartmentResponse{
		ID:          m.ID,
		Name:        m.Name,
		Address:     m.Address,
		Status:      m.Status,
		NextCheckIn: m.NextCheckIn,
		DoorCode:    m.DoorCode,
	}

	if m.LicenseNumber != nil {
		res.LicenseNumber = *m.LicenseNumber
	}

	if m.GuestName != nil {
		res.GuestName = *m.GuestName
	}

	return res
}

// FromModelList maps a slice of Apartment models to a slice of ApartmentResponse DTOs.
func FromModelList(apts []models.Apartment) []ApartmentResponse {
	result := make([]ApartmentResponse, 0, len(apts))
	for _, m := range apts {
		result = append(result, FromModel(&m))
	}
	return result
}
