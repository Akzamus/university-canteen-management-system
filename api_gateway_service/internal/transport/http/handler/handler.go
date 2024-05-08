package handler

import (
	"github.com/go-chi/chi"
	"net/http"
)

type UserHandler interface {
	GetUserByID(w http.ResponseWriter, r *http.Request)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUserByID(w http.ResponseWriter, r *http.Request)
	GetSelfInfo(w http.ResponseWriter, r *http.Request)
	RegisterRoutes(r chi.Router)
}

type AuthHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Authenticate(w http.ResponseWriter, r *http.Request)
	RegisterRoutes(r chi.Router)
}
