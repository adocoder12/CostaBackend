package repository

import (
	"context"

	"github.com/adocoder12/Costabackend/internal/model"
	"github.com/google/uuid"
)

type ApartmentsRepositoryInterface interface {
	GetAllApartments(ctx context.Context) ([]model.Apartment, error)
	GetByEmail(ctx context.Context, email string) (*model.Apartment, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Apartment, error)
	Create(ctx context.Context, user *model.Apartment) (*model.Apartment, error)
	Update(ctx context.Context, user *model.Apartment) (*model.Apartment, error)
	Delete(ctx context.Context, id uint) error
}
type CleanersRepositoryInterface interface{}
