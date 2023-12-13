package models

type CPUInfo struct {
	CPU        string `json:"cpu"`
	VendorID   string `json:"vendorId"`
	Family     string `json:"family"`
	Stepping   string `json:"stepping"`
	PhysicalId string `json:"physicalId"`
	Cores      string `json:"cores"`
	ModelName  string `json:"modelName"`
	Mhz        string `json:"mhz"`
	CacheSize  string `json:"cacheSize"`
}

type CPUStats struct {
	CPUCount int       `json:"cpuCount"`
	Loads    []float64 `json:"loads"`
}
