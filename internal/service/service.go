package service

import (
	"context"
	"photographer/internal/domain"
)

type Repository interface {
	CreatePhotographer(ctx context.Context, name string) (domain.PhotographerID, error)
	GetPhotographers(ctx context.Context) ([]domain.Photographer, error)

	CreateClient(ctx context.Context, photographerID domain.PhotographerID, name string) (domain.ClientID, error)
	UpdateClient(ctx context.Context, id domain.ClientID, name string) error
	DeleteClient(ctx context.Context, id domain.ClientID) error
	GetClients(ctx context.Context, photographerID domain.PhotographerID) ([]domain.Client, error)

	CreatDebt(ctx context.Context, photographerID domain.PhotographerID, clientID domain.ClientID, amount int) error
	GetDebts(ctx context.Context, photographerID domain.PhotographerID) ([]domain.Debt, error)

	CreatePayment(ctx context.Context, photographerID domain.PhotographerID, clientID domain.ClientID, amount int) error
	GetPayments(ctx context.Context, photographerID domain.PhotographerID) ([]domain.Payment, error)
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

func (s *Service) CreateClient(ctx context.Context, photographerID domain.PhotographerID, name string) (domain.ClientID, error) {
	return s.repo.CreateClient(ctx, photographerID, name)
}

func (s *Service) UpdateClient(ctx context.Context, id domain.ClientID, name string) error {
	return s.repo.UpdateClient(ctx, id, name)
}

func (s *Service) DeleteClient(ctx context.Context, id domain.ClientID) error {
	return s.repo.DeleteClient(ctx, id)
}

func (s *Service) GetClients(ctx context.Context, photographerID domain.PhotographerID) ([]domain.Client, error) {
	return s.repo.GetClients(ctx, photographerID)
}

func (s *Service) CreatDebt(ctx context.Context, photographerID domain.PhotographerID, clientID domain.ClientID, amount int) error {
	return s.repo.CreatDebt(ctx, photographerID, clientID, amount)
}

func (s *Service) GetDebts(ctx context.Context, photographerID domain.PhotographerID) ([]domain.Debt, error) {
	return s.repo.GetDebts(ctx, photographerID)
}

func (s *Service) CreatePayment(ctx context.Context, photographerID domain.PhotographerID, clientID domain.ClientID, amount int) error {
	return s.repo.CreatePayment(ctx, photographerID, clientID, amount)
}

func (s *Service) GetPayments(ctx context.Context, photographerID domain.PhotographerID) ([]domain.Payment, error) {
	return s.repo.GetPayments(ctx, photographerID)
}
