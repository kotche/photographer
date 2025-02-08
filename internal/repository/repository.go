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
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, created_at at time zone 'Europe/Moscow' FROM photographers")
	if err != nil {
		return nil, fmt.Errorf("failed to get photographers: %w", err)
	}

	var photographers []domain.Photographer
	for rows.Next() {
		var photographer domain.Photographer
		if err = rows.Scan(&photographer.ID, &photographer.Name, &photographer.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan photographer: %w", err)
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
	query := `update clients set name = $1, updated_at = now() where id = $2`

	if _, err := r.db.ExecContext(ctx, query, name, id); err != nil {
		return fmt.Errorf("failed to update client: %w", err)
	}

	return nil
}

func (r *Repository) DeleteClient(ctx context.Context, id domain.ClientID) error {
	query := `update clients set deleted_at = now() where id = $1`

	if _, err := r.db.ExecContext(ctx, query, id); err != nil {
		return fmt.Errorf("failed to delete client: %w", err)
	}

	return nil
}

func (r *Repository) GetClients(ctx context.Context, photographerID domain.PhotographerID) ([]domain.Client, error) {
	query := `
		select id, photographer_id, name, 
		       created_at at time zone 'Europe/Moscow', 
		       updated_at at time zone 'Europe/Moscow', 
		       deleted_at at time zone 'Europe/Moscow'
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
			return nil, fmt.Errorf("failed to scan client: %w", err)
		}
		clients = append(clients, client)
	}

	return clients, nil
}

func (r *Repository) AddDebt(ctx context.Context, photographerID domain.PhotographerID, clientID domain.ClientID, amount int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("AddDebt: failed to start transaction: %w", err)
	}

	defer func() {
		_ = tx.Rollback()
	}()

	currentDebt, err := getDebt(ctx, tx, photographerID, clientID)
	if err != nil {
		return err
	}

	if err = addDebt(ctx, tx, photographerID, clientID, currentDebt+amount); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *Repository) GetDebts(ctx context.Context, photographerID domain.PhotographerID) ([]domain.Debt, error) {
	query := `
		select client_id, c.name, amount, occurred_at at time zone 'Europe/Moscow'
		from debts
		join public.clients c on debts.client_id = c.id
		where debts.photographer_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, photographerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get debts: %w", err)
	}

	var debts []domain.Debt
	for rows.Next() {
		var debt domain.Debt
		if err = rows.Scan(&debt.ClientID, &debt.ClientName, &debt.Amount, &debt.OccurredAt); err != nil {
			return nil, fmt.Errorf("failed to scan debts: %w", err)
		}
		debts = append(debts, debt)
	}

	return debts, nil
}

func (r *Repository) AddPayment(ctx context.Context, photographerID domain.PhotographerID, clientID domain.ClientID, amount int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("AddPayment: failed to start transaction: %w", err)
	}

	defer func() {
		_ = tx.Rollback()
	}()

	sumDebt, err := getDebt(ctx, tx, photographerID, clientID)
	if err != nil {
		return err
	}

	dif := sumDebt - amount

	if dif <= 0 {
		if err = deleteDebt(ctx, tx, photographerID, clientID); err != nil {
			return err
		}
	} else {
		if err = addDebt(ctx, tx, photographerID, clientID, dif); err != nil {
			return err
		}
	}

	if err = addPayment(ctx, tx, photographerID, clientID, amount); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *Repository) GetPayments(ctx context.Context, photographerID domain.PhotographerID) ([]domain.Payment, error) {
	query := `
		select client_id, amount, occurred_at at time zone 'Europe/Moscow'
		from payments
		where photographer_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, photographerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get payments: %w", err)
	}

	var payments []domain.Payment
	for rows.Next() {
		var payment domain.Payment
		if err = rows.Scan(&payment.ClientID, &payment.Amount, &payment.OccurredAt); err != nil {
			return nil, fmt.Errorf("failed to scan payment: %w", err)
		}
		payments = append(payments, payment)
	}

	return payments, nil
}

func (r *Repository) GetPaymentsTotal(ctx context.Context, photographerID domain.PhotographerID) (int, error) {
	query := `
		select coalesce(sum(amount),0)
		from payments
		where photographer_id = $1
	`

	var sum int
	if err := r.db.QueryRowContext(ctx, query, photographerID).Scan(&sum); err != nil {
		return 0, fmt.Errorf("failed to get payments total: %w", err)
	}

	return sum, nil
}

func addDebt(ctx context.Context, tx *sql.Tx, photographerID domain.PhotographerID, clientID domain.ClientID, amount int) error {
	query := `
		insert into debts (photographer_id, client_id, amount)
		values ($1, $2, $3)
		on conflict (photographer_id, client_id)
		do update set amount = excluded.amount;
	`

	if _, err := tx.ExecContext(ctx, query, photographerID, clientID, amount); err != nil {
		return fmt.Errorf("failed to create debt: %w", err)
	}

	return nil
}

func getDebt(ctx context.Context, tx *sql.Tx, photographerID domain.PhotographerID, clientID domain.ClientID) (int, error) {
	query := `
		select coalesce(sum(amount),0) from debts
		where photographer_id = $1 and client_id = $2
	`

	var sumDebt int
	err := tx.QueryRowContext(ctx, query, photographerID, clientID).Scan(&sumDebt)
	if err != nil {
		return 0, fmt.Errorf("failed to get debt: %w", err)
	}

	return sumDebt, nil
}

func deleteDebt(ctx context.Context, tx *sql.Tx, photographerID domain.PhotographerID, clientID domain.ClientID) error {
	query := `
		delete from debts
		where photographer_id = $1 and client_id = $2
	`

	if _, err := tx.ExecContext(ctx, query, photographerID, clientID); err != nil {
		return fmt.Errorf("failed to delete debt: %w", err)
	}

	return nil
}

func addPayment(ctx context.Context, tx *sql.Tx, photographerID domain.PhotographerID, clientID domain.ClientID, amount int) error {
	query := `
		insert into payments (photographer_id, client_id, amount)
		values ($1, $2, $3);
	`

	if _, err := tx.ExecContext(ctx, query, photographerID, clientID, amount); err != nil {
		return fmt.Errorf("failed to add payment: %w", err)
	}

	return nil
}
