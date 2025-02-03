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
		insert into photographers (name, created_at)
		values ($1, NOW())
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
	rows, err := r.db.QueryContext(ctx, "SELECT id, name FROM photographers WHERE deleted_at is null")
	if err != nil {
		return nil, fmt.Errorf("failed to get photographers: %w", err)
	}

	var photographers []domain.Photographer
	for rows.Next() {
		var photographer domain.Photographer
		if err = rows.Scan(&photographer.ID, &photographer.Name); err != nil {
			return nil, fmt.Errorf("failed to scan note: %w", err)
		}
		photographers = append(photographers, photographer)
	}

	return photographers, nil
}
func (r *Repository) UpdateClient(ctx context.Context, id domain.ClientID, name string) error {
	return nil
}
