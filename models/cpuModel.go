package models

type CPUInfo struct {
	CPU        string   `json:"cpu"`
	VendorID   string   `json:"vendorId"`
	Family     string   `json:"family"`
	Model      string   `json:"model"`
	Stepping   string   `json:"stepping"`
	PhysicalId string   `json:"physicalId"`
	CoreId     string   `json:"coreId"`
	Cores      string   `json:"cores"`
	ModelName  string   `json:"modelName"`
	Mhz        string   `json:"mhz"`
	CacheSize  string   `json:"cacheSize"`
	Flags      []string `json:"flags"`
	Microcode  string   `json:"microcode"`
}
