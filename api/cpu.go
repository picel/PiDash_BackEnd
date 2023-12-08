package api

import (
	"encoding/json"
	"net/http"

	"picel.pidash/services"
)

func GetCPUInfo(w http.ResponseWriter, r *http.Request) {
	cpuInfo, err := services.GetCPUInfo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cpuInfo)
}
