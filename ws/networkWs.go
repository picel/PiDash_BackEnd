package ws

import (
	"log"
	"net/http"
	"time"

	"picel.pidash/services"
)

func ServeNetWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		stats, err := services.GetNetStats()
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
