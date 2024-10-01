package web

import (
	"encoding/json"
	"net/http"

	"github.com/ankardo/CleanArch/internal/dto"
	"github.com/ankardo/CleanArch/internal/entity"
	"github.com/ankardo/CleanArch/internal/usecase"
	"github.com/ankardo/CleanArch/pkg/events"
	"github.com/go-chi/chi/v5"
)

type OrderCreateHandler struct {
	EventDispatcher   events.EventDispatcherInterface
	OrderRepository   entity.OrderRepositoryInterface
	OrderCreatedEvent events.EventInterface
}

func NewOrderCreateHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreatedEvent events.EventInterface,
) *OrderCreateHandler {
	return &OrderCreateHandler{
		EventDispatcher:   EventDispatcher,
		OrderRepository:   OrderRepository,
		OrderCreatedEvent: OrderCreatedEvent,
	}
}

type FindAllOrdersHandler struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewFindAllOrdersHandler(
	OrderRepository entity.OrderRepositoryInterface,
) *FindAllOrdersHandler {
	return &FindAllOrdersHandler{
		OrderRepository: OrderRepository,
	}
}

type FindOrderHandler struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewFindOrderHandler(
	OrderRepository entity.OrderRepositoryInterface,
) *FindOrderHandler {
	return &FindOrderHandler{
		OrderRepository: OrderRepository,
	}
}

func (h *OrderCreateHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var dto dto.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewCreateOrderUseCase(
		h.OrderRepository,
		h.OrderCreatedEvent,
		h.EventDispatcher,
	)
	output, err := createOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *FindAllOrdersHandler) Get(w http.ResponseWriter, r *http.Request) {
	findAllOrders := usecase.NewFindAllOrdersUseCase(h.OrderRepository)
	output, err := findAllOrders.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *FindOrderHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	findOrder := usecase.NewFindOrderUseCase(h.OrderRepository)
	output, err := findOrder.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
