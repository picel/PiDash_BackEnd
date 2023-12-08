package api

import (
	"encoding/json"
	"net/http"

	"picel.pidash/services"
)

func GetMemInfo(w http.ResponseWriter, r *http.Request) {
	memInfo, err := services.GetMemInfo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(memInfo)
}
