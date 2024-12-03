package handler

import (
	"net/http"

	"github.com/gmcriptobox/otus-go-final-project/internal/entity/request"
	"github.com/gmcriptobox/otus-go-final-project/internal/service"
	"github.com/gmcriptobox/otus-go-final-project/internal/validator"
	"github.com/julienschmidt/httprouter"
	"github.com/mailru/easyjson"
)

type BucketHandler struct {
	service *service.Authorization
}

func NewBucketHandler(service *service.Authorization) *BucketHandler {
	return &BucketHandler{
		service: service,
	}
}

func (h *BucketHandler) ResetBuckets(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var bucketResetRequest request.BucketResetRequest

	err := easyjson.UnmarshalFromReader(r.Body, &bucketResetRequest)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if !validator.ValidateBucketResetRequest(&bucketResetRequest) {
		rw.WriteHeader(http.StatusBadRequest)
	}

	h.service.ResetBuckets(bucketResetRequest)
	rw.WriteHeader(http.StatusOK)
}
