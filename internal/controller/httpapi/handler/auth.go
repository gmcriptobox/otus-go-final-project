package handler

import (
	"fmt"
	"net/http"

	"github.com/gmcriptobox/otus-go-final-project/internal/entity/request"
	"github.com/gmcriptobox/otus-go-final-project/internal/entity/response"
	"github.com/gmcriptobox/otus-go-final-project/internal/service"
	"github.com/gmcriptobox/otus-go-final-project/internal/validator"
	"github.com/julienschmidt/httprouter"
	"github.com/mailru/easyjson"
)

type AuthHandler struct {
	service *service.Authorization
}

func NewAuthHandler(service *service.Authorization) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) TryLogin(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rw.Header().Set("Content-Type", "application/json")

	var authRequest request.AuthRequest
	err := easyjson.UnmarshalFromReader(r.Body, &authRequest)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if !validator.ValidateAuthRequest(&authRequest) {
		rw.WriteHeader(http.StatusBadRequest)
	}

	isAuthorized, err := h.service.Authorize(r.Context(), authRequest)
	if err != nil {
		fmt.Println("error while authorizing: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	authResponse := response.AuthResponse{
		Ok: isAuthorized,
	}

	body, err := easyjson.Marshal(authResponse)
	if err != nil {
		fmt.Println("error while marshalling response: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = rw.Write(body)
	if err != nil {
		fmt.Println("error while writing response: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}
