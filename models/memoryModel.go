package models

type MemStats struct {
	MemTotal     uint64 `json:"memTotal"`
	MemAvailable uint64 `json:"memAvailable"`
	MemUsed      uint64 `json:"memUsed"`
	MemFree      uint64 `json:"memFree"`
}
