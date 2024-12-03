package httpapi

import (
	"net/http"

	"github.com/gmcriptobox/otus-go-final-project/internal/controller/httpapi/handler"
	"github.com/julienschmidt/httprouter"
)

type APIRouter struct {
	router        *httprouter.Router
	authHandler   *handler.AuthHandler
	bucketHandler *handler.BucketHandler
	listHandler   *handler.ListHandler
}

func NewAPIRouter(authHandler *handler.AuthHandler, bucketHandler *handler.BucketHandler,
	listHandler *handler.ListHandler,
) *APIRouter {
	return &APIRouter{
		router:        httprouter.New(),
		authHandler:   authHandler,
		bucketHandler: bucketHandler,
		listHandler:   listHandler,
	}
}

func (r *APIRouter) RegisterRoutes() {
	r.router.GET("/api/health", func(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		writer.WriteHeader(http.StatusOK)
	})

	r.router.POST("/api/auth", r.authHandler.TryLogin)
	r.router.DELETE("/api/bucket", r.bucketHandler.ResetBuckets)

	r.router.POST("/api/blacklist", r.listHandler.AddToBlackList)
	r.router.DELETE("/api/blacklist", r.listHandler.RemoveFromBlackList)

	r.router.POST("/api/whitelist", r.listHandler.AddToWhiteList)
	r.router.DELETE("/api/whitelist", r.listHandler.RemoveFromWhiteList)
}

func (r *APIRouter) GetRouter() *httprouter.Router {
	return r.router
}
