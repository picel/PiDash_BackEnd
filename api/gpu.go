package api

import (
	"encoding/json"
	"net/http"

	"picel.pidash/services"
)

func GetGPUInfo(w http.ResponseWriter, r *http.Request) {
	gpuInfo, err := services.GetGPUInfo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(gpuInfo)
}
