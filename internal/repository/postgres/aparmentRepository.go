package postgres

import (
	"context"

	"github.com/adocoder12/Costabackend/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ApartmentRepository struct {
	Pool *pgxpool.Pool
}

func (r *ApartmentRepository) Create(ctx context.Context, apt *models.Apartment) (*models.Apartment, error) {
	query := `
        INSERT INTO apartments (
            name, address, status, license_number, 
            cadastral_ref, door_code, notes, owner_id
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id, created_at, updated_at`

	err := r.Pool.QueryRow(ctx, query,
		apt.Name,
		apt.Address,
		apt.Status,
		apt.LicenseNumber,
		apt.CadastralRef,
		apt.DoorCode,
		apt.Notes,
		apt.OwnerID,
	).Scan(&apt.ID, &apt.CreatedAt, &apt.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return apt, nil
}

// Satisfying the rest of the interface stubs
func (r *ApartmentRepository) GetAllApartments(ctx context.Context) ([]models.Apartment, error) {
	return nil, nil
}

func (r *ApartmentRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Apartment, error) {
	return nil, nil
}

func (r *ApartmentRepository) Update(ctx context.Context, apt *models.Apartment) (*models.Apartment, error) {
	return nil, nil
}

func (r *ApartmentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
