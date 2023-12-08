package ws

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func SetUpRoutes(router *mux.Router) {
	router.HandleFunc("/cpu", ServeCPUWs)
	router.HandleFunc("/mem", ServeMemWs)
	router.HandleFunc("/net", ServeNetWs)
	router.HandleFunc("/gpu", ServeGPUWs)
}
