package users

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type Controller struct {
	Router  http.Handler
	Service *Service
}

type CreateUserRequest struct {
	ID string `json:"id"`
}

type CreateUserResponse struct {
	Status int  `json:"status"`
	Data   User `json:"data"`
}

func (c *Controller) createUser(w http.ResponseWriter, r *http.Request) {
	var payload CreateUserRequest
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

	user := User{
		ID: payload.ID,
	}
	if err := c.Service.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	response := CreateUserResponse{
		Data:   user,
		Status: 201,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Controller) userInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("Getting user with id %s", id)
}

type CreateAccountRequest struct {
	Currency string `json:"currency"`
}

type CreateAccountResponse struct {
	Status int     `json:"status"`
	Data   Account `json:"data"`
}

func (c *Controller) createAccount(w http.ResponseWriter, r *http.Request) {
	var payload CreateAccountRequest
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
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("Creating account for user with id %s", id)
	account := Account{
		Currency: payload.Currency,
		UserID:   id,
		Amount:   0,
	}

	if err := c.Service.CreateAccount(account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	response := CreateAccountResponse{
		Data:   account,
		Status: 201,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewUserController(router *mux.Router, service *Service) *Controller {
	userRouter := router.PathPrefix("/users").Subrouter()
	controller := &Controller{
		Router:  userRouter,
		Service: service,
	}

	userRouter.HandleFunc("", controller.createUser).Methods("POST")
	userRouter.HandleFunc("/{id}", controller.userInfo).Methods("GET")
	userRouter.HandleFunc("/{id}/accounts", controller.createAccount).Methods("POST")

	return controller
}
