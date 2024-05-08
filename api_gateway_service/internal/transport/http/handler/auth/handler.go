package auth

import (
	"encoding/json"
	"github.com/Akzamus/university-canteen-management-system/api_gateway_service/internal/service"
	def "github.com/Akzamus/university-canteen-management-system/api_gateway_service/internal/transport/http/handler"
	responseUtils "github.com/Akzamus/university-canteen-management-system/api_gateway_service/internal/utils/response"
	"github.com/Akzamus/university-canteen-management-system/api_gateway_service/pkg/dto"
	"github.com/go-chi/chi"
	"net/http"
)

var _ def.AuthHandler = (*handler)(nil)

type handler struct {
	authService service.AuthService
}

func NewHandler(authService service.AuthService) *handler {
	return &handler{
		authService: authService,
	}
}

func (h *handler) RegisterRoutes(r chi.Router) {
	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/register", h.Register)
		r.Post("/authenticate", h.Authenticate)
	})
}

func (h *handler) Register(w http.ResponseWriter, r *http.Request) {
	var request dto.AuthRequestDto
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responseUtils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.authService.Register(r.Context(), request)
	if err != nil {
		responseUtils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseUtils.RespondWithJSON(w, http.StatusCreated, response)
}

func (h *handler) Authenticate(w http.ResponseWriter, r *http.Request) {
	var request dto.AuthRequestDto
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responseUtils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.authService.Authenticate(r.Context(), request)
	if err != nil {
		responseUtils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseUtils.RespondWithJSON(w, http.StatusOK, response)
}
