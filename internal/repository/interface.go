package repository

import (
	"context"

	"github.com/adocoder12/Costabackend/internal/models"
	"github.com/google/uuid"
)

type ApartmentsRepositoryInterface interface {
	GetApartments(ctx context.Context, viewerID uuid.UUID) ([]models.Apartment, []bool, error)
	GetApartmentByID(ctx context.Context, id uuid.UUID) (*models.Apartment, error)
	CreateApartment(ctx context.Context, apt *models.Apartment) (*models.Apartment, error)
	UpdateApartment(ctx context.Context, apt *models.Apartment) (*models.Apartment, error)
	DeleteApartment(ctx context.Context, id uuid.UUID) error
}

type CleanersRepositoryInterface interface{}
