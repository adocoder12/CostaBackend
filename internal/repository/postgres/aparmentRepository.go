package postgres

import (
	"context"
	"fmt"

	"github.com/adocoder12/Costabackend/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ApartmentRepository struct {
	Pool *pgxpool.Pool
}

func (r *ApartmentRepository) CreateApartment(ctx context.Context, apt *models.Apartment) (*models.Apartment, error) {
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

func (r *ApartmentRepository) GetApartments(ctx context.Context, viewerID uuid.UUID) ([]models.Apartment, []bool, error) {
	query := `
       SELECT 
          id, name, address, status, license_number, 
          cadastral_ref, door_code, next_check_in, checkout_date, 
          guest_name, notes, owner_id, created_at, updated_at,
          -- This ensures we ALWAYS get true or false, never NULL
          CASE 
            WHEN owner_id IS NOT NULL AND owner_id = $1 THEN true 
            ELSE false 
          END AS is_mine_flag
       FROM apartments
       ORDER BY created_at DESC`

	rows, err := r.Pool.Query(ctx, query, viewerID)
	if err != nil {
		return nil, nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var apartments []models.Apartment
	var isMineFlags []bool

	for rows.Next() {
		var a models.Apartment
		var isMine bool // Go bool defaults to false

		err := rows.Scan(
			&a.ID, &a.Name, &a.Address, &a.Status, &a.LicenseNumber,
			&a.CadastralRef, &a.DoorCode, &a.NextCheckIn, &a.CheckoutDate,
			&a.GuestName, &a.Notes, &a.OwnerID, &a.CreatedAt, &a.UpdatedAt,
			&isMine, // The 15th column
		)
		if err != nil {
			return nil, nil, fmt.Errorf("scan failed: %w", err)
		}

		apartments = append(apartments, a)
		isMineFlags = append(isMineFlags, isMine)
	}

	return apartments, isMineFlags, nil
}
func (r *ApartmentRepository) GetApartmentByID(ctx context.Context, id uuid.UUID) (*models.Apartment, error) {
	query := `SELECT 
			id, name, address, status, license_number, 
			cadastral_ref, door_code, next_check_in, checkout_date, 
			guest_name, notes, owner_id, created_at, updated_at
		FROM apartments 
		WHERE id = $1`
	row := r.Pool.QueryRow(ctx, query, id)
	var apartment models.Apartment
	err := row.Scan(
		&apartment.ID,
		&apartment.Name,
		&apartment.Address,
		&apartment.Status,
		&apartment.LicenseNumber,
		&apartment.CadastralRef,
		&apartment.DoorCode,
		&apartment.NextCheckIn,
		&apartment.CheckoutDate,
		&apartment.GuestName,
		&apartment.Notes,
		&apartment.OwnerID,
		&apartment.CreatedAt,
		&apartment.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to scan apartment row: %w", err)
	}
	return &apartment, nil
}

func (r *ApartmentRepository) UpdateApartment(ctx context.Context, apt *models.Apartment) (*models.Apartment, error) {
	query := `
		UPDATE apartments
		SET 
			name = $1, address = $2, status = $3, license_number = $4,
			cadastral_ref = $5, door_code = $6, next_check_in = $7,
			checkout_date = $8, guest_name = $9, notes = $10, 
			owner_id = $11, updated_at = NOW()
		WHERE id = $12
		RETURNING updated_at`

	err := r.Pool.QueryRow(ctx, query,
		apt.Name, apt.Address, apt.Status, apt.LicenseNumber,
		apt.CadastralRef, apt.DoorCode, apt.NextCheckIn,
		apt.CheckoutDate, apt.GuestName, apt.Notes,
		apt.OwnerID, apt.ID,
	).Scan(&apt.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to update apartment: %w", err)
	}

	return apt, nil
}

func (r *ApartmentRepository) DeleteApartment(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM apartments WHERE id = $1`

	res, err := r.Pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete apartment: %w", err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("apartment not found")
	}

	return nil

}
