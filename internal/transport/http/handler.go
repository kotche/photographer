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

	CreateClient(ctx context.Context, name string) (int, error)
	UpdateClient(ctx context.Context, id domain.ClientID, name string) error
	DeleteClient(ctx context.Context, id domain.ClientID) error
	GetClients(ctx context.Context, photographerID domain.PhotographerID) ([]domain.Client, error)

	CreatDebt(ctx context.Context, photographerID domain.PhotographerID, clientID domain.ClientID, amount int) error
	GetDebts(ctx context.Context, photographerID domain.PhotographerID) ([]domain.Debt, error)

	CreatePayment(ctx context.Context, photographerID domain.PhotographerID, clientID domain.ClientID, amount int) error
	GetPayments(ctx context.Context, photographerID domain.PhotographerID) ([]domain.Payment, error)
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
	router.HandleFunc("/debt", h.createDebtHandler).Methods("POST")                    // добавить запись о задолженности
	router.HandleFunc("/transaction", h.createTransactionHandler).Methods("POST")      // провести оплату/возврат с обновлением debt_client
	router.HandleFunc("/debtors/{photographerID}", h.getDebtorsHandler).Methods("GET") // список должников у фотографа
	router.HandleFunc("/incomes/{photographerID}", h.getIncomesHandler).Methods("GET") // операции и суммарный доход у фотографа

	return router
}

func (h *Handler) createPhotographerHandler(w http.ResponseWriter, r *http.Request) {
	type createPhotographerRequest struct {
		Name string `json:"name"`
	}

	var req createPhotographerRequest

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

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(id); err != nil {
		log.Printf("json encode error: %v", err)
	}
}

func (h *Handler) getPhotographersHandler(w http.ResponseWriter, r *http.Request) {
	photographers, err := h.service.GetPhotographers(r.Context())
	if err != nil {
		log.Printf("get photographers error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(photographers); err != nil {
		log.Printf("json encode error: %v", err)
	}
}

func (h *Handler) createClientHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) updateClientHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("convert id '%s' to int error: %v", vars["id"], err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var name string
	if err = json.NewDecoder(r.Body).Decode(&name); err != nil {
		log.Printf("json decode error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = h.service.UpdateClient(r.Context(), domain.ClientID(id), name); err != nil {
		log.Printf("update client error,: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(id); err != nil {
		log.Printf("json encode error: %v", err)
	}
}

func (h *Handler) deleteClientHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getClientsHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) createDebtHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getDebtorsHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) createTransactionHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getIncomesHandler(w http.ResponseWriter, r *http.Request) {

}
