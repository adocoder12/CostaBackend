package dto

import (
	"github.com/adocoder12/Costabackend/internal/models"
	"github.com/google/uuid"
	"time"
)

type CreateApartmentRequest struct {
	Name          string  `json:"name" binding:"required"`
	Address       string  `json:"address" binding:"required"`
	LicenseNumber string  `json:"license_number" binding:"required"`
	CadastralRef  *string `json:"cadastral_ref"`
	DoorCode      *string `json:"door_code"`
	Notes         *string `json:"notes"`
}

type ApartmentResponse struct {
	ID            uuid.UUID  `json:"id"`
	Name          string     `json:"name"`
	Address       string     `json:"address"`
	Status        string     `json:"status"`
	LicenseNumber string     `json:"license_number"`
	NextCheckIn   *time.Time `json:"next_check_in,omitempty"`
	GuestName     string     `json:"guest_name,omitempty"`
	DoorCode      *string    `json:"door_code,omitempty"` // Masked in service
	IsMine        bool       `json:"is_mine"`             // Calculated in DB/Service
}

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

func FromModel(m *models.Apartment) ApartmentResponse {
	res := ApartmentResponse{
		ID:          m.ID,
		Name:        m.Name,
		Address:     m.Address,
		Status:      m.Status,
		NextCheckIn: m.NextCheckIn,
		// ADDED THESE:
		IsMine:   m.IsMine,   // Pass the flag from the model
		DoorCode: m.DoorCode, // Pass the (potentially nil) door code
	}

	if m.LicenseNumber != nil {
		res.LicenseNumber = *m.LicenseNumber
	}

	if m.GuestName != nil {
		res.GuestName = *m.GuestName
	}

	return res
}

func FromModelList(models []models.Apartment) []ApartmentResponse {
	result := make([]ApartmentResponse, 0, len(models))
	for _, m := range models {
		result = append(result, FromModel(&m))
	}
	return result
}
