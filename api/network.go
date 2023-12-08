package api

import (
	"encoding/json"
	"net/http"

	"picel.pidash/services"
)

func GetNetInfo(w http.ResponseWriter, r *http.Request) {
	netInfo, err := services.GetNetInfo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(netInfo)
}
