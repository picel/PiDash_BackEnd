package models

type CPUInfo struct {
	CPU        int     `json:"cpu"`
	VendorID   string  `json:"vendorId"`
	Family     string  `json:"family"`
	Stepping   int32   `json:"stepping"`
	PhysicalId string  `json:"physicalId"`
	Cores      int32   `json:"cores"`
	ModelName  string  `json:"modelName"`
	Clock      float64 `json:"clock"`
	CacheSize  int32   `json:"cacheSize"`
}

type CPUStats struct {
	CPUCount int       `json:"cpuCount"`
	Loads    []float64 `json:"loads"`
}
