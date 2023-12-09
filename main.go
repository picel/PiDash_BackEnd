package main

import (
	"log"
	"net/http"
	"os"

	"github.com/getlantern/systray"
	"github.com/gorilla/mux"
	"picel.pidash/api"
	"picel.pidash/ws"
)

func main() {
	go func() {
		systray.Run(onReady, nil)
	}()

	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()
	wsRouter := router.PathPrefix("/ws").Subrouter()

	api.SetUpRoutes(apiRouter)
	ws.SetUpRoutes(wsRouter)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func onReady() {
	systray.SetTitle("PiDash")
	systray.SetTooltip("PiDash")

	mRunning := systray.AddMenuItem("PiDash Server Running", "")
	mRunning.Disable()

	mQuit := systray.AddMenuItem("Quit", "Quit PiDash")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
		os.Exit(0)
	}()
}
