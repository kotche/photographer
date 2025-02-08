package service

import (
	"context"
	"photographer/internal/domain"

	"golang.org/x/sync/errgroup"
)

type Repository interface {
	CreatePhotographer(ctx context.Context, name string) (domain.PhotographerID, error)
	GetPhotographers(ctx context.Context) ([]domain.Photographer, error)

	CreateClient(ctx context.Context, photographerID domain.PhotographerID, name string) (domain.ClientID, error)
	UpdateClient(ctx context.Context, id domain.ClientID, name string) error
	DeleteClient(ctx context.Context, id domain.ClientID) error
	GetClients(ctx context.Context, photographerID domain.PhotographerID) ([]domain.Client, error)

	AddDebt(ctx context.Context, photographerID domain.PhotographerID, clientID domain.ClientID, amount int) error
	GetDebts(ctx context.Context, photographerID domain.PhotographerID) ([]domain.Debt, error)

	AddPayment(ctx context.Context, photographerID domain.PhotographerID, clientID domain.ClientID, amount int) error
	GetPayments(ctx context.Context, photographerID domain.PhotographerID) ([]domain.Payment, error)
	GetPaymentsTotal(ctx context.Context, photographerID domain.PhotographerID) (int, error)
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

func (s *Service) AddDebt(ctx context.Context, photographerID domain.PhotographerID, clientID domain.ClientID, amount int) error {
	return s.repo.AddDebt(ctx, photographerID, clientID, amount)
}

func (s *Service) GetDebts(ctx context.Context, photographerID domain.PhotographerID) ([]domain.Debt, error) {
	return s.repo.GetDebts(ctx, photographerID)
}

func (s *Service) AddPayment(ctx context.Context, photographerID domain.PhotographerID, clientID domain.ClientID, amount int) error {
	return s.repo.AddPayment(ctx, photographerID, clientID, amount)
}

func (s *Service) GetPayments(ctx context.Context, photographerID domain.PhotographerID) ([]domain.Payment, int, error) {
	var (
		payments []domain.Payment
		total    int
	)

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		var err error
		payments, err = s.repo.GetPayments(ctx, photographerID)
		return err
	})

	eg.Go(func() error {
		var err error
		total, err = s.repo.GetPaymentsTotal(ctx, photographerID)
		return err
	})

	if err := eg.Wait(); err != nil {
		return nil, 0, err
	}

	return payments, total, nil
}
