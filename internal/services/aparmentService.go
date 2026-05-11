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

type ApartmentService struct {
	repo   repository.ApartmentsRepositoryInterface
	logger *slog.Logger
}

// ApartmentServiceInterface is the 'Contract' the Handler uses.
type ApartmentServiceInterface interface {
	GetApartments(ctx context.Context, viewerID uuid.UUID) ([]dto.ApartmentResponse, error)
	GetApartmentByID(ctx context.Context, id uuid.UUID) (*models.Apartment, error)
	CreateApartment(ctx context.Context, req dto.CreateApartmentRequest) (dto.ApartmentResponse, error)
	UpdateApartment(ctx context.Context, apt *dto.ApartmentResponse) error
	DeleteApartment(ctx context.Context, id uuid.UUID) error
}

func NewApartmentService(repo repository.ApartmentsRepositoryInterface, logger *slog.Logger) *ApartmentService {
	return &ApartmentService{
		repo:   repo,
		logger: logger,
	}
}

func (s *ApartmentService) GetApartmentByID(ctx context.Context, id uuid.UUID) (*models.Apartment, error) {
	apt, err := s.repo.GetApartmentByID(ctx, id)
	if err != nil {
		s.logger.Warn("apartment not found", "id", id, "error", err)
		return nil, err
	}
	return apt, nil
}

func (s *ApartmentService) GetApartments(ctx context.Context, viewerID uuid.UUID) ([]dto.ApartmentResponse, error) {
	// 1. Get raw models AND the is_mine flags from the Repo
	apts, flags, err := s.repo.GetApartments(ctx, viewerID)
	if err != nil {
		return nil, fmt.Errorf("service failed to fetch apartments: %w", err)
	}

	// 2. Map to DTOs and apply security masking
	response := make([]dto.ApartmentResponse, 0, len(apts))
	for i, a := range apts {
		d := dto.FromModel(&a)

		// Use the flag returned by the Repo SQL
		d.IsMine = flags[i]

		// Security: If not the owner, hide sensitive info
		if !d.IsMine {
			d.DoorCode = nil

		}

		response = append(response, d)
	}

	return response, nil
}

func (s *ApartmentService) CreateApartment(ctx context.Context, req dto.CreateApartmentRequest) (dto.ApartmentResponse, error) {
	// 1. Map Request DTO to Model
	apt := req.ToModel()

	// 2. Business Logic / Defaults
	if apt.Status == "" {
		apt.Status = models.StatusClean
	}

	s.logger.Info("creating new apartment", "name", apt.Name)

	// 3. Persist to Database
	createdApt, err := s.repo.CreateApartment(ctx, apt)
	if err != nil {
		return dto.ApartmentResponse{}, err
	}

	// 4. Return the Response DTO
	// The service now returns the "Presentation" version of the data
	response := dto.FromModel(createdApt)

	// Since the creator is the owner, we can explicitly set this
	response.IsMine = true

	return response, nil
}

func (s *ApartmentService) DeleteApartment(ctx context.Context, id uuid.UUID) error {
	s.logger.Info("deleting apartment", "id", id)
	return s.repo.DeleteApartment(ctx, id)
}
