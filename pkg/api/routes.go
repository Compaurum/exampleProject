package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h Handler) initRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/cats", h.getCats).Methods(http.MethodGet)
	router.HandleFunc("/cats", h.createCat).Methods(http.MethodPost)
	router.HandleFunc("/cats/{name}", h.removeCat).Methods(http.MethodDelete)
	router.HandleFunc("/cats/{name}", h.getCat).Methods(http.MethodGet)

	h.Router = router
	return router
}
