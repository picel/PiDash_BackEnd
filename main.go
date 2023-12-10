package main

import (
	"log"
	"net/http"
	"os"

	"github.com/getlantern/systray"
	"github.com/go-toast/toast"
	"github.com/gorilla/mux"
	"picel.pidash/api"
	"picel.pidash/middlewares"
	"picel.pidash/utils"
	"picel.pidash/ws"
)

var ipAddrs []string
var Version string

func main() {
	if Version == "" {
		Version = "dev"
	}

	go func() {
		systray.Run(onReady, nil)
	}()

	ipAddrs = utils.GetUserIP()
	if len(ipAddrs) == 0 {
		toastMessage := toast.Notification{
			AppID:   "PiDash",
			Title:   "PiDash",
			Message: "No IP address found",
		}
		toastMessage.Push()
		log.Fatal("No IP address found")
		os.Exit(1)
	} else {
		stringBuilder := "Server running on one of the following addresses: "
		for _, ipAddr := range ipAddrs {
			stringBuilder += ipAddr + ":8080, "
		}
		toastMessage := toast.Notification{
			AppID:   "PiDash",
			Title:   "PiDash",
			Message: stringBuilder[:len(stringBuilder)-2],
		}
		toastMessage.Push()
	}

	router := mux.NewRouter()

	// add cors headers of mux router
	router.Use(middlewares.CorsHeader)

	apiRouter := router.PathPrefix("/api").Subrouter()
	wsRouter := router.PathPrefix("/ws").Subrouter()

	api.SetUpRoutes(apiRouter)
	ws.SetUpRoutes(wsRouter)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func onReady() {
	systray.SetTitle("PiDash")
	systray.SetTooltip("PiDash")

	mRunning := systray.AddMenuItem("PiDash Server "+Version, "")
	mRunning.Disable()

	systray.AddSeparator()

	for _, ipAddr := range ipAddrs {
		mIP := systray.AddMenuItem(ipAddr+":8080", "Open PiDash in browser")
		mIP.Disable()
	}

	mQuit := systray.AddMenuItem("Quit", "Quit PiDash")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
		os.Exit(0)
	}()
}
