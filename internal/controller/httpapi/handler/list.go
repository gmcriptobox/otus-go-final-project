package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gmcriptobox/otus-go-final-project/internal/entity/request"
	"github.com/gmcriptobox/otus-go-final-project/internal/service"
	"github.com/gmcriptobox/otus-go-final-project/internal/validator"
	"github.com/julienschmidt/httprouter"
	"github.com/mailru/easyjson"
)

const (
	whiteList = iota
	blackList
)

type ListHandler struct {
	whiteListService *service.ListService
	blackListService *service.ListService
}

func NewListHandler(whiteListService *service.ListService, blackListService *service.ListService) *ListHandler {
	return &ListHandler{
		whiteListService: whiteListService,
		blackListService: blackListService,
	}
}

func (l *ListHandler) AddToWhiteList(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	l.addToList(rw, r, whiteList)
}

func (l *ListHandler) AddToBlackList(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	l.addToList(rw, r, blackList)
}

func (l *ListHandler) RemoveFromWhiteList(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	l.removeFromList(rw, r, whiteList)
}

func (l *ListHandler) RemoveFromBlackList(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	l.removeFromList(rw, r, blackList)
}

func (l *ListHandler) addToList(rw http.ResponseWriter, r *http.Request, list int) {
	var networkRequest request.NetworkRequest
	err := easyjson.UnmarshalFromReader(r.Body, &networkRequest)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if !validator.ValidateNetworkRequest(&networkRequest) {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if list == blackList {
		err = l.blackListService.Add(r.Context(), networkRequest.Network)
	} else {
		err = l.whiteListService.Add(r.Context(), networkRequest.Network)
	}

	if err != nil {
		if errors.Is(err, service.ErrNetworkAlreadyExists) {
			rw.WriteHeader(http.StatusConflict)
			return
		}
		fmt.Println("error while adding network to black list: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

func (l *ListHandler) removeFromList(rw http.ResponseWriter, r *http.Request, list int) {
	var networkRequest request.NetworkRequest
	err := easyjson.UnmarshalFromReader(r.Body, &networkRequest)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if !validator.ValidateNetworkRequest(&networkRequest) {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if list == blackList {
		err = l.blackListService.Remove(r.Context(), networkRequest.Network)
	} else {
		err = l.whiteListService.Remove(r.Context(), networkRequest.Network)
	}

	if err != nil {
		fmt.Println("error while removing network from black list: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}
