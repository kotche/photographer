package domain

import "time"

type Photographer struct {
	ID        PhotographerID `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
}

type Client struct {
	ID             ClientID       `json:"id"`
	Name           string         `json:"name"`
	PhotographerID PhotographerID `json:"photographer_id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      *time.Time     `json:"deleted_at"`
}

type Debt struct {
	ClientID   ClientID `json:"client_id"`
	ClientName string   `json:"client_name"`
	Amount     int      `json:"amount"`
	OccurredAt time.Time
}

type Payment struct {
	ClientID   ClientID `json:"client_id"`
	Amount     int      `json:"amount"`
	OccurredAt time.Time
}
