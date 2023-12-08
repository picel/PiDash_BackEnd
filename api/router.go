package api

import (
	"github.com/gorilla/mux"
)

func SetUpRoutes(router *mux.Router) {
	router.HandleFunc("/cpu", GetCPUInfo).Methods("GET")
	router.HandleFunc("/mem", GetMemInfo).Methods("GET")
	router.HandleFunc("/net", GetNetInfo).Methods("GET")
	router.HandleFunc("/gpu", GetGPUInfo).Methods("GET")
}
