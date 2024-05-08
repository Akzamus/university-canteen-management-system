package user

import (
	"encoding/json"
	"github.com/Akzamus/university-canteen-management-system/user_service/internal/service"
	def "github.com/Akzamus/university-canteen-management-system/user_service/internal/transport/http/handler"
	responseUtils "github.com/Akzamus/university-canteen-management-system/user_service/internal/utils/response"
	"github.com/Akzamus/university-canteen-management-system/user_service/pkg/dto"
	"github.com/go-chi/chi"
	"net/http"
)

var _ def.UserHandler = (*handler)(nil)

type handler struct {
	userService service.UserService
}

func NewHandler(userService service.UserService) *handler {
	return &handler{
		userService: userService,
	}
}

func (h *handler) RegisterRoutes(r chi.Router) {
	r.Route("/api/v1/users", func(r chi.Router) {
		r.Get("/", h.GetAllUsers)
		r.Post("/", h.CreateUser)
		r.Post("/verify", h.VerifyUserCredentials)
		r.Route("/{userUUID}", func(r chi.Router) {
			r.Get("/", h.GetUserByID)
			r.Put("/", h.UpdateUser)
			r.Delete("/", h.DeleteUserByID)
		})
	})
}

func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userUUID := chi.URLParam(r, "userUUID")
	response, err := h.userService.GetUserByID(r.Context(), userUUID)

	if err != nil {
		responseUtils.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	responseUtils.RespondWithJSON(w, http.StatusOK, response)
}

func (h *handler) VerifyUserCredentials(w http.ResponseWriter, r *http.Request) {
	var request dto.UserCredentialsRequestDto
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responseUtils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	response, err := h.userService.VerifyUserCredentials(r.Context(), request)
	if err != nil {
		responseUtils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	responseUtils.RespondWithJSON(w, http.StatusOK, response)
}

func (h *handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	response, err := h.userService.GetAllUsers(r.Context())
	if err != nil {
		responseUtils.RespondWithError(w, http.StatusInternalServerError, "Failed to get users")
		return
	}

	responseUtils.RespondWithJSON(w, http.StatusOK, response)
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var request dto.UserRequestDto
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responseUtils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	response, err := h.userService.CreateUser(r.Context(), request)
	if err != nil {
		responseUtils.RespondWithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	responseUtils.RespondWithJSON(w, http.StatusCreated, response)
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userUUID := chi.URLParam(r, "userUUID")

	var request dto.UserRequestDto
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responseUtils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	response, err := h.userService.UpdateUser(r.Context(), request, userUUID)
	if err != nil {
		responseUtils.RespondWithError(w, http.StatusInternalServerError, "Failed to update user")
		return
	}

	responseUtils.RespondWithJSON(w, http.StatusOK, response)
}

func (h *handler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	userUUID := chi.URLParam(r, "userUUID")

	err := h.userService.DeleteUserByID(r.Context(), userUUID)
	if err != nil {
		responseUtils.RespondWithError(w, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	responseUtils.RespondWithJSON(w, http.StatusNoContent, nil)
}
