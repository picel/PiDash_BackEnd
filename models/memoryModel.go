package models

type MemStats struct {
	MemTotal     string `json:"memTotal"`
	MemAvailable string `json:"memAvailable"`
	MemUsed      string `json:"memUsed"`
	MemFree      string `json:"memFree"`
}
