package services

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/adocoder12/Costabackend/internal/dto"
	"github.com/adocoder12/Costabackend/internal/models"
	"github.com/adocoder12/Costabackend/internal/repository"
	"github.com/google/uuid"
)

// ApartmentServiceInterface is the contract the handler depends on.
// Handlers only see DTOs — never raw models.
type ApartmentServiceInterface interface {
	GetApartments(ctx context.Context, viewerID uuid.UUID) ([]dto.ApartmentResponse, error)
	GetApartmentByID(ctx context.Context, id, viewerID uuid.UUID) (*dto.ApartmentResponse, error)
	CreateApartment(ctx context.Context, req dto.CreateApartmentRequest) (dto.ApartmentResponse, error)
	UpdateApartment(ctx context.Context, id uuid.UUID, req dto.UpdateApartmentRequest) error
	DeleteApartment(ctx context.Context, id uuid.UUID) error
}

// ApartmentService implements ApartmentServiceInterface.
type ApartmentService struct {
	repo   repository.ApartmentsRepositoryInterface
	logger *slog.Logger
}

func NewApartmentService(repo repository.ApartmentsRepositoryInterface, logger *slog.Logger) *ApartmentService {
	return &ApartmentService{repo: repo, logger: logger}
}

func (s *ApartmentService) GetApartments(ctx context.Context, viewerID uuid.UUID) ([]dto.ApartmentResponse, error) {
	apts, flags, err := s.repo.GetApartments(ctx, viewerID)
	if err != nil {
		return nil, fmt.Errorf("apartment service - error get all: %w", err)
	}

	response := make([]dto.ApartmentResponse, 0, len(apts))
	for i, a := range apts {
		d := dto.FromModel(&a)
		d.IsMine = flags[i]

		// Security: hide door code from non-owners
		if !d.IsMine {
			d.DoorCode = nil
		}

		response = append(response, d)
	}

	return response, nil
}

// GetApartmentByID now correctly returns *dto.ApartmentResponse, not *models.Apartment.
func (s *ApartmentService) GetApartmentByID(ctx context.Context, id, viewerID uuid.UUID) (*dto.ApartmentResponse, error) {
	apt, err := s.repo.GetApartmentByID(ctx, id)
	if err != nil {
		s.logger.Warn("apartment not found", "id", id, "error", err)
		return nil, err // ErrNotFound propagates to handler
	}

	response := dto.FromModel(apt)

	// Determine IsMine for a single apartment fetch
	if apt.OwnerID != nil && *apt.OwnerID == viewerID {
		response.IsMine = true
	}

	// Security: hide door code from non-owners
	if !response.IsMine {
		response.DoorCode = nil
	}

	return &response, nil
}

func (s *ApartmentService) CreateApartment(ctx context.Context, req dto.CreateApartmentRequest) (dto.ApartmentResponse, error) {
	apt := req.ToModel()

	if apt.Status == "" {
		apt.Status = models.StatusClean
	}

	s.logger.Info("creating apartment", "name", apt.Name)

	created, err := s.repo.CreateApartment(ctx, apt)
	if err != nil {
		return dto.ApartmentResponse{}, fmt.Errorf("apartment service - create: %w", err)
	}

	response := dto.FromModel(created)
	response.IsMine = true // creator is always the owner

	return response, nil
}

// UpdateApartment applies partial updates from the request DTO to the existing model.
func (s *ApartmentService) UpdateApartment(ctx context.Context, id uuid.UUID, req dto.UpdateApartmentRequest) error {
	// Fetch current state first
	existing, err := s.repo.GetApartmentByID(ctx, id)
	if err != nil {
		return err // ErrNotFound propagates
	}

	// Apply only the fields that were provided (partial update)
	if req.Name != nil {
		existing.Name = *req.Name
	}
	if req.Address != nil {
		existing.Address = *req.Address
	}
	if req.Status != nil {
		existing.Status = *req.Status
	}
	if req.LicenseNumber != nil {
		existing.LicenseNumber = req.LicenseNumber
	}
	if req.CadastralRef != nil {
		existing.CadastralRef = req.CadastralRef
	}
	if req.DoorCode != nil {
		existing.DoorCode = req.DoorCode
	}
	if req.Notes != nil {
		existing.Notes = req.Notes
	}

	s.logger.Info("updating apartment", "id", id)

	if _, err := s.repo.UpdateApartment(ctx, existing); err != nil {
		return fmt.Errorf("apartment service - update: %w", err)
	}

	return nil
}

func (s *ApartmentService) DeleteApartment(ctx context.Context, id uuid.UUID) error {
	s.logger.Info("deleting apartment", "id", id)

	if err := s.repo.DeleteApartment(ctx, id); err != nil {
		return fmt.Errorf("apartment service - delete: %w", err)
	}

	return nil
}
