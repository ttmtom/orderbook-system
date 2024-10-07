package users

import (
	"net/http"
)

func userRoute(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		CreateUser()
	}
}

func InitController() {
	http.HandleFunc("/users", userRoute)

}
