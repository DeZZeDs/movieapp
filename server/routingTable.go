package server

import (
	"github.com/gorilla/mux"
	"movieapp/server/handlers"
	"net/http"
)

func prepareRoutingTable() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/registration", handlers.RegistrationHandler).Methods(http.MethodPost)

	return router
}
