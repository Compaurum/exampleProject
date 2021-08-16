package api

import (
	"example/pkg/storage"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router  *mux.Router
	storage storage.Storage
}

func NewHandler(s storage.Storage) *Handler {
	h := &Handler{
		storage: s,
	}
	h.Router = h.initRoutes()
	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// auth user
	start := time.Now()

	h.Router.ServeHTTP(w, req)

	// measure time
	fmt.Printf("request time is %v \n", time.Now().Sub(start))
}
