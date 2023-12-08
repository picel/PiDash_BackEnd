package models

type GPUInfo struct {
	ProductName   string `json:"productName"`
	DriverVersion string `json:"driverVersion"`
	TotalMemory   string `json:"totalMemory"`
	MaxClocks     Clocks `json:"maxClocks"`
}

type GPUStats struct {
	MemoryUsage MemoryUsage `json:"memoryUsage"`
	Utilization Utilization `json:"utilization"`
	Temperature string      `json:"temperature"`
	Power       Power       `json:"power"`
	Clocks      Clocks      `json:"clocks"`
}

/// sub structs

type Clocks struct {
	GraphicsClock string `json:"graphicsClock"`
	SmClock       string `json:"smClock"`
	MemClock      string `json:"memClock"`
	VideoClock    string `json:"videoClock"`
}

type MemoryUsage struct {
	Total    string `json:"total"`
	Reserved string `json:"reserved"`
	Used     string `json:"used"`
	Free     string `json:"free"`
}

type Utilization struct {
	GPU    string `json:"gpu"`
	Memory string `json:"memory"`
}

type Power struct {
	Usage string `json:"usage"`
	Limit string `json:"limit"`
}
