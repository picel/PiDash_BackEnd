package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"picel.pidash/api"
	"picel.pidash/interfaces"
	"picel.pidash/utils"
	"picel.pidash/ws"
)

var Version string

func main() {
	if Version == "" {
		Version = "dev"
	}

	var ipAddrs []string = utils.GetUserIP()
	interfaces.ShowIPToast(ipAddrs)

	go func() {
		interfaces.SetUpSystray(ipAddrs, Version)
	}()

	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()
	wsRouter := router.PathPrefix("/ws").Subrouter()

	api.SetUpRoutes(apiRouter)
	ws.SetUpRoutes(wsRouter)

	log.Fatal(http.ListenAndServe(":8080", router))
}
