package service

import (
	"context"
	"photographer/internal/domain"
)

type Repository interface {
	CreatePhotographer(ctx context.Context, name string) (domain.PhotographerID, error)
	GetPhotographers(ctx context.Context) ([]domain.Photographer, error)
	UpdateClient(ctx context.Context, id domain.ClientID, name string) error
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreatePhotographer(ctx context.Context, name string) (domain.PhotographerID, error) {
	return s.repo.CreatePhotographer(ctx, name)
}

func (s *Service) GetPhotographers(ctx context.Context) ([]domain.Photographer, error) {
	return s.repo.GetPhotographers(ctx)
}
func (s *Service) UpdateClient(ctx context.Context, id domain.ClientID, name string) error {
	return nil
}
