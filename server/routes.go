package server

import (
	"github.com/istudko/go-project-template/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.Use(loggingMiddleware)

	routerV1 := router.PathPrefix("/api/v1").Subrouter()
	routerV1.HandleFunc("/ping", handler.PingHandler).Methods(http.MethodGet)
	routerV1.HandleFunc("/ticket", handler.CreateTicketHandler).Methods(http.MethodPost)
	return router
}
