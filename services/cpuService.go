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
			CPU:        i,
			VendorID:   c.VendorID,
			Family:     c.Family,
			Stepping:   c.Stepping,
			PhysicalId: c.PhysicalID,
			Cores:      c.Cores,
			ModelName:  c.ModelName,
			Clock:      c.Mhz,
			CacheSize:  c.CacheSize,
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
