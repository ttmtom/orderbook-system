package users

import (
	"net/http"
)

type Controller struct {
	Router  http.Handler
	Service *Service
}

func (c *Controller) createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := c.Service.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func NewUserController(service *Service) *Controller {
	mux := http.NewServeMux()
	controller := &Controller{
		Router:  mux,
		Service: service,
	}

	mux.HandleFunc("/user", controller.createUser)
	return controller
}
