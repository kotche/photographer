package http_handler

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"photographer/internal/domain"
	"strconv"
)

type Service interface {
	CreatePhotographer(ctx context.Context, name string) (domain.PhotographerID, error)
	GetPhotographers(ctx context.Context) ([]domain.Photographer, error)

	CreateClient(ctx context.Context, photographerID domain.PhotographerID, name string) (domain.ClientID, error)
	UpdateClient(ctx context.Context, id domain.ClientID, name string) error
	DeleteClient(ctx context.Context, id domain.ClientID) error
	GetClients(ctx context.Context, photographerID domain.PhotographerID) ([]domain.Client, error)

	AddDebt(ctx context.Context, photographerID domain.PhotographerID, clientID domain.ClientID, amount int) error
	GetDebts(ctx context.Context, photographerID domain.PhotographerID) ([]domain.Debt, error)

	AddPayment(ctx context.Context, photographerID domain.PhotographerID, clientID domain.ClientID, amount int) error
	GetPayments(ctx context.Context, photographerID domain.PhotographerID) ([]domain.Payment, int, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Handle() *mux.Router {
	// Маршруты
	router := mux.NewRouter()

	// Фотографы
	router.HandleFunc("/photographers", h.createPhotographerHandler).Methods("POST")
	router.HandleFunc("/photographers", h.getPhotographersHandler).Methods("GET")

	// Клиенты
	router.HandleFunc("/clients", h.createClientHandler).Methods("POST")
	router.HandleFunc("/clients/{id}", h.updateClientHandler).Methods("PUT")
	router.HandleFunc("/clients/{id}", h.deleteClientHandler).Methods("DELETE")
	router.HandleFunc("/clients/{photographerID}", h.getClientsHandler).Methods("GET")

	// Операции с денежными средствами
	router.HandleFunc("/debt", h.addDebtHandler).Methods("POST")                       // добавить сумму задолженности
	router.HandleFunc("/payment", h.addPaymentHandler).Methods("POST")                 // провести оплату с обновлением задолженности
	router.HandleFunc("/debtors/{photographerID}", h.getDebtorsHandler).Methods("GET") // список должников фотографа
	router.HandleFunc("/incomes/{photographerID}", h.getIncomesHandler).Methods("GET") // операции и суммарный доход у фотографа

	return router
}

func (h *Handler) createPhotographerHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("json decode error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.service.CreatePhotographer(r.Context(), req.Name)
	if err != nil {
		log.Printf("create photographer error,: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	encodeResponse(w, id)
}

func (h *Handler) getPhotographersHandler(w http.ResponseWriter, r *http.Request) {
	photographers, err := h.service.GetPhotographers(r.Context())
	if err != nil {
		log.Printf("get photographers error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	encodeResponse(w, photographers)
}

func (h *Handler) createClientHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PhotographerID domain.PhotographerID `json:"photographer_id"`
		Name           string                `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("json decode error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateClient(r.Context(), req.PhotographerID, req.Name)
	if err != nil {
		log.Printf("create client error,: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	encodeResponse(w, id)
}

func (h *Handler) updateClientHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("convert id '%s' to int error: %v", vars["id"], err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	type updateClientRequest struct {
		Name string `json:"name"`
	}

	var req updateClientRequest

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("json decode error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = h.service.UpdateClient(r.Context(), domain.ClientID(id), req.Name); err != nil {
		log.Printf("update client error,: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h *Handler) deleteClientHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("convert id '%s' to int error: %v", vars["id"], err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = h.service.DeleteClient(r.Context(), domain.ClientID(id)); err != nil {
		log.Printf("delete client error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h *Handler) getClientsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	photographerID, err := strconv.Atoi(vars["photographerID"])
	if err != nil {
		log.Printf("convert id '%s' to int error: %v", vars["id"], err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	clients, err := h.service.GetClients(r.Context(), domain.PhotographerID(photographerID))
	if err != nil {
		log.Printf("get clients error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	encodeResponse(w, clients)
}

func (h *Handler) addDebtHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PhotographerID int `json:"photographer_id"`
		ClientID       int `json:"client_id"`
		Amount         int `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("decode request body error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.AddDebt(r.Context(), domain.PhotographerID(req.PhotographerID), domain.ClientID(req.ClientID), req.Amount); err != nil {
		log.Printf("add debt error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) getDebtorsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	photographerID, err := strconv.Atoi(vars["photographerID"])
	if err != nil {
		log.Printf("convert id '%s' to int error: %v", vars["id"], err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	debts, err := h.service.GetDebts(r.Context(), domain.PhotographerID(photographerID))
	if err != nil {
		log.Printf("get debts error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	encodeResponse(w, debts)
}

func (h *Handler) addPaymentHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PhotographerID int `json:"photographer_id"`
		ClientID       int `json:"client_id"`
		Amount         int `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("decode request body error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.AddPayment(r.Context(), domain.PhotographerID(req.PhotographerID), domain.ClientID(req.ClientID), req.Amount); err != nil {
		log.Printf("add payment error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) getIncomesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	photographerID, err := strconv.Atoi(vars["photographerID"])
	if err != nil {
		log.Printf("convert id '%s' to int error: %v", vars["id"], err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payments, total, err := h.service.GetPayments(r.Context(), domain.PhotographerID(photographerID))
	if err != nil {
		log.Printf("get payments error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := struct {
		Payments []domain.Payment `json:"payments"`
		Total    int              `json:"total"`
	}{
		Payments: payments,
		Total:    total,
	}

	encodeResponse(w, resp)
}

func encodeResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("json encode error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
