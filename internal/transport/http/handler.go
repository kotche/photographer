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

// @Summary Создаёт нового фотографа
// @Tags Photographers
// @Accept json
// @Produce json
// @Param request body CreatePhotographerRequest true "Payload для создания фотографа"
// @Success 200 {object} CreatePhotographerResponse "ID созданного фотографа"
// @Failure 400 {string} text/plain
// @Failure 500 {string} text/plain
// @Router /photographers [post]
func (h *Handler) createPhotographerHandler(w http.ResponseWriter, r *http.Request) {
	var req CreatePhotographerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("json decode error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.service.CreatePhotographer(r.Context(), req.Name)
	if err != nil {
		log.Printf("create photographer error,: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	encodeResponse(w, CreatePhotographerResponse{ID: id})
}

// @Summary Возвращает список фотографов
// @Tags Photographers
// @Accept json
// @Produce json
// @Success 200 {array} domain.Photographer
// @Failure 500 {string} text/plain
// @Router /photographers [get]
func (h *Handler) getPhotographersHandler(w http.ResponseWriter, r *http.Request) {
	photographers, err := h.service.GetPhotographers(r.Context())
	if err != nil {
		log.Printf("get photographers error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	encodeResponse(w, photographers)
}

// @Summary Создаёт нового клиента
// @Tags Clients
// @Accept json
// @Produce json
// @Param request body CreateClientRequest true "Payload для создания клиента"
// @Success 200 {object} CreateClientResponse "ID созданного клиента"
// @Failure 400 {string} text/plain
// @Failure 500 {string} text/plain
// @Router /clients [post]
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	encodeResponse(w, CreateClientResponse{ID: id})
}

// @Summary Обновляет данные клиента
// @Tags Clients
// @Accept json
// @Produce json
// @Param id path int true "ID клиента"
// @Param request body UpdateClientRequest true "Payload для обновления клиента"
// @Success 200
// @Failure 400 {string} text/plain
// @Failure 500 {string} text/plain
// @Router /clients/{id} [put]
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// @Summary Удаляет клиента по ID
// @Tags Clients
// @Accept json
// @Produce json
// @Param id path int true "ID клиента"
// @Success 200
// @Failure 400 {string} text/plain
// @Failure 500 {string} text/plain
// @Router /clients/{id} [delete]
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// @Summary Возвращает список клиентов фотографа
// @Tags Clients
// @Accept json
// @Produce json
// @Param photographerID path int true "ID фотографа"
// @Success 200 {array} domain.Client
// @Failure 400 {string} text/plain
// @Failure 500 {string} text/plain
// @Router /clients/{photographerID} [get]
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

// @Summary Добавляет задолженность для клиента фотографа
// @Tags Financial
// @Accept json
// @Produce json
// @Param request body AddDebtRequest true "Payload для добавления задолженности"
// @Success 200
// @Failure 400 {string} text/plain
// @Failure 500 {string} text/plain
// @Router /debt [post]
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

// @Summary Получает список должников фотографа
// @Tags Financial
// @Accept json
// @Produce json
// @Param photographerID path int true "ID фотографа"
// @Success 200 {array} domain.Debt "Список задолженностей"
// @Failure 400 {string} text/plain
// @Failure 500 {string} text/plain
// @Router /debtors/{photographerID} [get]
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

// @Summary Добавляет оплату клиента фотографу
// @Tags Financial
// @Accept json
// @Produce json
// @Param request body AddPaymentRequest true "Payload для добавления оплаты"
// @Success 200
// @Failure 400 {string} text/plain
// @Failure 500 {string} text/plain
// @Router /payment [post]
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

// @Summary Получает детализированный список доходов фотографа
// @Tags Financial
// @Accept json
// @Produce json
// @Param photographerID path int true "ID фотографа"
// @Success 200 {object} GetIncomesResponse "Список платежей и общий доход"
// @Failure 400 {string} text/plain
// @Failure 500 {string} text/plain
// @Router /incomes/{photographerID} [get]
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
