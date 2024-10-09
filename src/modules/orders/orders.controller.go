package orders

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Controller struct {
	Router  http.Handler
	Service *Service
}

func NewOrdersController(router *mux.Router, service *Service) *Controller {
	ordersRouter := router.PathPrefix("/orders").Subrouter()
	controller := &Controller{
		Router:  ordersRouter,
		Service: service,
	}

	//mux.HandleFunc("/")
	return controller
}
