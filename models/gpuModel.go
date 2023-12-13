package models

type GPUInfo struct {
	ProductName   string  `json:"productName"`
	DriverVersion string  `json:"driverVersion"`
	TotalMemory   float64 `json:"totalMemory"`
	MaxClock      Clock   `json:"maxClock"`
}

type GPUStats struct {
	MemoryUsage MemoryUsage `json:"memoryUsage"`
	Utilization Utilization `json:"utilization"`
	Temperature int         `json:"temperature"`
	Power       Power       `json:"power"`
	Clock       Clock       `json:"clock"`
}

/// sub structs

type Clock struct {
	GraphicsClock string `json:"graphicsClock"`
	SmClock       string `json:"smClock"`
	MemClock      string `json:"memClock"`
	VideoClock    string `json:"videoClock"`
}

type MemoryUsage struct {
	Total    int `json:"total"`
	Reserved int `json:"reserved"`
	Used     int `json:"used"`
	Free     int `json:"free"`
}

type Utilization struct {
	GPU    int `json:"gpu"`
	Memory int `json:"memory"`
}

type Power struct {
	Usage float64 `json:"usage"`
	Limit float64 `json:"limit"`
}
