package http_handler

import "photographer/internal/domain"

type (
	CreatePhotographerRequest struct {
		Name string `json:"name" example:"Alice"`
	}

	CreatePhotographerResponse struct {
		ID domain.PhotographerID `json:"id" example:"1"`
	}

	CreateClientRequest struct {
		PhotographerID domain.PhotographerID `json:"photographer_id" example:"1"`
		Name           string                `json:"name" example:"Alice"`
	}

	CreateClientResponse struct {
		ID domain.ClientID `json:"id" example:"1"`
	}

	UpdateClientRequest struct {
		Name string `json:"name" example:"Alice Updated"`
	}

	AddDebtRequest struct {
		PhotographerID int `json:"photographer_id" example:"1"`
		ClientID       int `json:"client_id" example:"2"`
		Amount         int `json:"amount" example:"500"`
	}

	AddPaymentRequest struct {
		PhotographerID int `json:"photographer_id" example:"1"`
		ClientID       int `json:"client_id" example:"2"`
		Amount         int `json:"amount" example:"500"`
	}

	GetIncomesResponse struct {
		Payments []domain.Payment `json:"payments"`
		Total    int              `json:"total" example:"10000"`
	}
)
