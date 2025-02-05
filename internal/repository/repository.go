package repository

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"photographer/internal/domain"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) CreatePhotographer(ctx context.Context, name string) (domain.PhotographerID, error) {
	query := `
		insert into photographers (name)
		values ($1)
		returning id
	`

	var id domain.PhotographerID
	err := r.db.QueryRowContext(ctx, query, name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create photographer: %w", err)
	}

	return id, nil
}

func (r *Repository) GetPhotographers(ctx context.Context) ([]domain.Photographer, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, created_at FROM photographers")
	if err != nil {
		return nil, fmt.Errorf("failed to get photographers: %w", err)
	}

	var photographers []domain.Photographer
	for rows.Next() {
		var photographer domain.Photographer
		if err = rows.Scan(&photographer.ID, &photographer.Name, &photographer.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan note: %w", err)
		}
		photographers = append(photographers, photographer)
	}

	return photographers, nil
}

func (r *Repository) CreateClient(ctx context.Context, photographerID domain.PhotographerID, name string) (domain.ClientID, error) {
	query := `
		insert into clients (photographer_id, name)
		values ($1, $2)
		returning id
	`

	var id domain.ClientID
	err := r.db.QueryRowContext(ctx, query, photographerID, name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create client: %w", err)
	}

	return id, nil
}

func (r *Repository) UpdateClient(ctx context.Context, id domain.ClientID, name string) error {
	query := `update clients set name = $1 where id = $2`

	if _, err := r.db.ExecContext(ctx, query, name, id); err != nil {
		return fmt.Errorf("failed to update client: %w", err)
	}

	return nil
}

func (r *Repository) DeleteClient(ctx context.Context, id domain.ClientID) error {
	return nil
}

func (r *Repository) GetClients(ctx context.Context, photographerID domain.PhotographerID) ([]domain.Client, error) {
	query := `
		select id, photographer_id, name, created_at, updated_at, deleted_at
		from clients
		where photographer_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, photographerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get clients: %w", err)
	}

	var clients []domain.Client
	for rows.Next() {
		var client domain.Client
		if err = rows.Scan(&client.ID, &client.PhotographerID, &client.Name,
			&client.CreatedAt, &client.UpdatedAt, &client.DeletedAt); err != nil {
			return nil, fmt.Errorf("failed to scan note: %w", err)
		}
		clients = append(clients, client)
	}

	return clients, nil
}

func (r *Repository) CreatDebt(ctx context.Context, photographerID domain.PhotographerID, clientID domain.ClientID, amount int) error {
	return nil
}

func (r *Repository) GetDebts(ctx context.Context, photographerID domain.PhotographerID) ([]domain.Debt, error) {
	return nil, nil
}

func (r *Repository) CreatePayment(ctx context.Context, photographerID domain.PhotographerID, clientID domain.ClientID, amount int) error {
	return nil
}

func (r *Repository) GetPayments(ctx context.Context, photographerID domain.PhotographerID) ([]domain.Payment, error) {
	return nil, nil
}
