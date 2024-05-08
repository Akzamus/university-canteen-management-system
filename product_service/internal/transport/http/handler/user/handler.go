package user

import (
	"encoding/json"
	"net/http"

	"github.com/Akzamus/university-canteen-management-system/product_service/internal/service"
	def "github.com/Akzamus/university-canteen-management-system/product_service/internal/transport/http/handler"
	responseUtils "github.com/Akzamus/university-canteen-management-system/product_service/internal/utils/response"
	"github.com/Akzamus/university-canteen-management-system/product_service/pkg/dto"
	"github.com/go-chi/chi"
)

var _ def.ProductHandler = (*handler)(nil)

type handler struct {
	ProductService service.ProductService
}

func NewHandler(ProductService service.ProductService) *handler {
	return &handler{
		ProductService: ProductService,
	}
}

func (h *handler) RegisterRoutes(r chi.Router) {
	r.Route("/api/v1/products", func(r chi.Router) {
		r.Get("/", h.GetAllProduct)
		r.Post("/", h.CreateProduct)
		r.Route("/{productUUID}", func(r chi.Router) {
			r.Get("/", h.GetProductByID)
			r.Put("/", h.UpdateProduct)
			r.Delete("/", h.DeleteProductByID)
		})
	})
}

func (h *handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	productUUID := chi.URLParam(r, "productUUID")
	response, err := h.ProductService.GetUserByID(r.Context(), productUUID)

	if err != nil {
		responseUtils.RespondWithError(w, http.StatusNotFound, "Product not found")
		return
	}

	responseUtils.RespondWithJSON(w, http.StatusOK, response)
}

func (h *handler) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	response, err := h.ProductService.GetAllProduct(r.Context())
	if err != nil {
		responseUtils.RespondWithError(w, http.StatusInternalServerError, "Failed to get products")
		return
	}

	responseUtils.RespondWithJSON(w, http.StatusOK, response)
}

func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var request dto.ProductRequestDto
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responseUtils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	response, err := h.ProductService.CreateProduct(r.Context(), request)
	if err != nil {
		responseUtils.RespondWithError(w, http.StatusInternalServerError, "Failed to create product")
		return
	}

	responseUtils.RespondWithJSON(w, http.StatusCreated, response)
}

func (h *handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productUUID := chi.URLParam(r, "productUUID")

	var request dto.UserRequestDto
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responseUtils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	response, err := h.ProductService.UpdateProduct(r.Context(), request, productUUID)
	if err != nil {
		responseUtils.RespondWithError(w, http.StatusInternalServerError, "Failed to update product")
		return
	}

	responseUtils.RespondWithJSON(w, http.StatusOK, response)
}

func (h *handler) DeleteProductByID(w http.ResponseWriter, r *http.Request) {
	productUUID := chi.URLParam(r, "productUUID")

	err := h.ProductService.DeleteUserByID(r.Context(), productUUID)
	if err != nil {
		responseUtils.RespondWithError(w, http.StatusInternalServerError, "Failed to delete product")
		return
	}

	responseUtils.RespondWithJSON(w, http.StatusNoContent, nil)
}
