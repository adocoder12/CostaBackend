package repository

import (
	"context"

	"github.com/adocoder12/Costabackend/internal/models"
	"github.com/google/uuid"
)

type ApartmentsRepositoryInterface interface {
	GetAllApartments(ctx context.Context) ([]models.Apartment, error)
	GetByEmail(ctx context.Context, email string) (*models.Apartment, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Apartment, error)
	Create(ctx context.Context, user *models.Apartment) (*models.Apartment, error)
	Update(ctx context.Context, user *models.Apartment) (*models.Apartment, error)
	Delete(ctx context.Context, id uint) error
}
type CleanersRepositoryInterface interface{}
