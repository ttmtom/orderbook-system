package users

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type CreateUserRequest struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
}

type CreateUserResponse struct {
	Data User `json:"data"`
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
	w.Header().Set("Content-Type", "application/json")
	response := CreateUserResponse{
		Data: user,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type UserInfo struct {
	User      User       `json:"user"`
	Positions []Position `json:"position"`
}

type UserInfoResponse struct {
	Data UserInfo `json:"data"`
}

func (c *Controller) userInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("Getting user with id %s", id)
	user, err := c.Service.GetUserById(id)

	if user == nil {
		log.Printf("User with id %s not found", id)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Printf("Error getting user with id %s: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := UserInfoResponse{
		Data: UserInfo{
			User: *user,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type Controller struct {
	Router  http.Handler
	Service *Service
}

func NewUserController(router *mux.Router, service *Service) *Controller {
	userRouter := router.PathPrefix("/users").Subrouter()
	controller := &Controller{
		Router:  userRouter,
		Service: service,
	}

	userRouter.HandleFunc("", controller.createUser).Methods("POST")
	userRouter.HandleFunc("/{id}", controller.userInfo).Methods("GET")

	return controller
}
