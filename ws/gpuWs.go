package ws

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"picel.pidash/services"
)

func ServeGPUWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		stats, err := services.GetGPUStats()
		fmt.Println(stats)
		if err != nil {
			log.Println(err)
			return
		}

		if err := conn.WriteJSON(stats); err != nil {
			log.Println(err)
			return
		}

		time.Sleep(time.Second * 1)
	}
}
