package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"picel.pidash/api"
	"picel.pidash/ws"
)

func main() {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()
	wsRouter := router.PathPrefix("/ws").Subrouter()

	api.SetUpRoutes(apiRouter)
	ws.SetUpRoutes(wsRouter)

	log.Fatal(http.ListenAndServe(":8080", router))
}
