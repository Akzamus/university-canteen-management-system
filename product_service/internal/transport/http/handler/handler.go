package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

type UserHandler interface {
	GetProductByID(w http.ResponseWriter, r *http.Request)
	GetAllProduct(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	DeleteProductByID(w http.ResponseWriter, r *http.Request)
	RegisterRoutes(r chi.Router)
}
