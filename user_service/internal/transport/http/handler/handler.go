package handler

import (
	"github.com/go-chi/chi"
	"net/http"
)

type UserHandler interface {
	GetUserByID(w http.ResponseWriter, r *http.Request)
	VerifyUserCredentials(w http.ResponseWriter, r *http.Request)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUserByID(w http.ResponseWriter, r *http.Request)
	RegisterRoutes(r chi.Router)
}
