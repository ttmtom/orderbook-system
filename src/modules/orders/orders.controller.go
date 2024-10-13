package orders

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type Controller struct {
	Router  http.Handler
	Service *Service
}

type CreateOrderRequest struct {
	UserId string  `json:"userId"`
	Market string  `json:"market"`
	Side   string  `json:"side"`
	Size   float64 `json:"size"`
}

type CreateOrderResponse struct {
	OrderId string `json:"orderId"`
}

func (c *Controller) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var payload CreateOrderRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order, err := c.Service.PlaceOrder(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	response := CreateOrderResponse{
		OrderId: order.ID,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Controller) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["orderId"]
	order, err := c.Service.getOrdersById(orderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewOrdersController(router *mux.Router, service *Service) *Controller {
	ordersRouter := router.PathPrefix("/orders").Subrouter()
	controller := &Controller{
		Router:  ordersRouter,
		Service: service,
	}

	ordersRouter.HandleFunc("/", controller.CreateOrder).Methods("POST")
	ordersRouter.HandleFunc("/{orderId}", controller.GetOrder).Methods("GET")
	return controller
}
