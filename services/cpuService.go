package services

import (
	"strconv"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"

	"picel.pidash/models"
)

func GetCPUInfo() ([]models.CPUInfo, error) {
	cpu, err := cpu.Info()
	if err != nil {
		return nil, err
	}

	cpuInfo := make([]models.CPUInfo, len(cpu))
	for i, c := range cpu {
		cpuInfo[i] = models.CPUInfo{
			CPU:        strconv.Itoa(int(c.CPU)),
			VendorID:   c.VendorID,
			Family:     c.Family,
			Model:      c.Model,
			Stepping:   strconv.Itoa(int(c.Stepping)),
			PhysicalId: c.PhysicalID,
			CoreId:     c.CoreID,
			Cores:      strconv.Itoa(int(c.Cores)),
			ModelName:  c.ModelName,
			Mhz:        strconv.FormatFloat(c.Mhz, 'f', 2, 64),
			CacheSize:  strconv.FormatUint(uint64(c.CacheSize), 10),
			Flags:      c.Flags,
			Microcode:  c.Microcode,
		}
	}

	return cpuInfo, nil
}

func GetCPUStats() (*models.CPUStats, error) {
	cpuLoad, err := cpu.Percent(time.Second, true)
	if err != nil {
		return nil, err
	}

	for i, load := range cpuLoad {
		cpuLoad[i], _ = strconv.ParseFloat(strconv.FormatFloat(load, 'f', 2, 64), 64)
	}

	cpuStats := &models.CPUStats{
		CPUCount: len(cpuLoad),
		Loads:    cpuLoad,
	}

	return cpuStats, nil
}
