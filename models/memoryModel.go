package models

type MemInfo struct {
	Total uint64 `json:"total"`
}

type MemStats struct {
	MemTotal     uint64 `json:"memTotal"`
	MemAvailable uint64 `json:"memAvailable"`
	MemUsed      uint64 `json:"memUsed"`
	MemFree      uint64 `json:"memFree"`
}
