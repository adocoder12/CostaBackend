package services

import (
	"context"
	"log/slog"

	"github.com/adocoder12/Costabackend/internal/models"
	"github.com/adocoder12/Costabackend/internal/repository"
	"github.com/google/uuid"
)

type ApartmentService struct {
	repo   repository.ApartmentsRepositoryInterface
	logger *slog.Logger
}

func NewApartmentService(repo repository.ApartmentsRepositoryInterface, logger *slog.Logger) *ApartmentService {
	return &ApartmentService{
		repo:   repo,
		logger: logger,
	}
}

func (s *ApartmentService) CreateApartment(ctx context.Context, apt *models.Apartment) error {
	if apt.Status == "" {
		apt.Status = "Clean"
	}

	s.logger.Info("creating new apartment", "name", apt.Name)

	_, err := s.repo.Create(ctx, apt)
	return err
}

func (s *ApartmentService) GetApartmentByID(ctx context.Context, id uuid.UUID) (*models.Apartment, error) {
	apt, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Warn("apartment not found", "id", id, "error", err)
		return nil, err
	}
	return apt, nil
}

func (s *ApartmentService) GetApartments(ctx context.Context) ([]models.Apartment, error) {
	s.logger.Debug("fetching all apartments")
	return s.repo.GetAllApartments(ctx)
}

func (s *ApartmentService) UpdateApartment(ctx context.Context, apt *models.Apartment) error {
	s.logger.Info("updating apartment", "id", apt.ID)
	_, err := s.repo.Update(ctx, apt)
	return err
}

func (s *ApartmentService) DeleteApartment(ctx context.Context, id uuid.UUID) error {
	s.logger.Info("deleting apartment", "id", id)
	return s.repo.Delete(ctx, id)
}
