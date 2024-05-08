package user

import (
	"encoding/json"
	"fmt"
	"github.com/Akzamus/university-canteen-management-system/api_gateway_service/internal/model"
	"github.com/Akzamus/university-canteen-management-system/api_gateway_service/internal/service"
	def "github.com/Akzamus/university-canteen-management-system/api_gateway_service/internal/transport/http/handler"
	responseUtils "github.com/Akzamus/university-canteen-management-system/api_gateway_service/internal/utils/response"
	"github.com/Akzamus/university-canteen-management-system/api_gateway_service/pkg/dto"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"net/http"
)

const (
	permissionDeniedText = "Permission denied"
)

var _ def.UserHandler = (*handler)(nil)

type handler struct {
	userService service.UserService
	jwtAuth     *jwtauth.JWTAuth
}

func NewHandler(userService service.UserService, jwtAuth *jwtauth.JWTAuth) *handler {
	return &handler{
		userService: userService,
		jwtAuth:     jwtAuth,
	}
}

func (h *handler) RegisterRoutes(r chi.Router) {
	r.Route("/api/v1/users", func(r chi.Router) {
		r.Use(jwtauth.Verifier(h.jwtAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/", h.GetAllUsers)
		r.Post("/", h.CreateUser)
		r.Get("/me", h.GetSelfInfo)
		r.Route("/{userUUID}", func(r chi.Router) {
			r.Get("/", h.GetUserByID)
			r.Put("/", h.UpdateUser)
			r.Delete("/", h.DeleteUserByID)
		})
	})
}

func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userUuid := chi.URLParam(r, "userUUID")
	_, claims, _ := jwtauth.FromContext(r.Context())

	if claims["role"] != string(model.AdminRole) && claims["userUuid"] != userUuid {
		responseUtils.RespondWithError(w, http.StatusForbidden, permissionDeniedText)
		return
	}

	response, err := h.userService.GetUserByID(r.Context(), userUuid)

	if err != nil {
		responseUtils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	responseUtils.RespondWithJSON(w, http.StatusOK, response)
}

func (h *handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())

	if claims["role"] != string(model.AdminRole) {
		responseUtils.RespondWithError(w, http.StatusForbidden, permissionDeniedText)
		return
	}

	response, err := h.userService.GetAllUsers(r.Context())
	if err != nil {
		responseUtils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseUtils.RespondWithJSON(w, http.StatusOK, response)
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())

	if claims["role"] != string(model.AdminRole) {
		responseUtils.RespondWithError(w, http.StatusForbidden, permissionDeniedText)
		return
	}

	var request dto.UserRequestDto
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responseUtils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.userService.CreateUser(r.Context(), request)
	if err != nil {
		responseUtils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseUtils.RespondWithJSON(w, http.StatusCreated, response)
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userUuid := chi.URLParam(r, "userUUID")
	_, claims, _ := jwtauth.FromContext(r.Context())

	if claims["role"] != string(model.AdminRole) && claims["userUuid"] != userUuid {
		responseUtils.RespondWithError(w, http.StatusForbidden, permissionDeniedText)
		return
	}

	var request dto.UserRequestDto
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responseUtils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.userService.UpdateUser(r.Context(), request, userUuid)
	if err != nil {
		responseUtils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseUtils.RespondWithJSON(w, http.StatusOK, response)
}

func (h *handler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	userUuid := chi.URLParam(r, "userUUID")
	_, claims, _ := jwtauth.FromContext(r.Context())

	if claims["role"] != string(model.AdminRole) && claims["userUuid"] != userUuid {
		responseUtils.RespondWithError(w, http.StatusForbidden, permissionDeniedText)
		return
	}

	err := h.userService.DeleteUserByID(r.Context(), userUuid)
	if err != nil {
		responseUtils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseUtils.RespondWithJSON(w, http.StatusNoContent, nil)
}

func (h *handler) GetSelfInfo(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())

	response, err := h.userService.GetUserByID(r.Context(), fmt.Sprintf("%v", claims["userUuid"]))

	if err != nil {
		responseUtils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	responseUtils.RespondWithJSON(w, http.StatusOK, response)
}
