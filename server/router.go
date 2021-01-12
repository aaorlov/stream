package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

// NewRouter create router
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		HandleWS(w, r)
	})

	return router
}
