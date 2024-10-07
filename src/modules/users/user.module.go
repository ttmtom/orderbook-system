package users

import "net/http"

type UserModule struct {
	router http.Handler
	user   *user.Service
}
