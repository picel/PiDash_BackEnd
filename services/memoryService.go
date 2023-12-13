package services

import (
	"github.com/shirou/gopsutil/v3/mem"

	"picel.pidash/models"
)

func GetMemInfo() (*models.MemInfo, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	// return total as JSON string
	total := &models.MemInfo{
		Total: memInfo.Total,
	}
	return total, nil
}

func GetMemStats() (*models.MemStats, error) {
	memStats, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	// return stats
	stats := &models.MemStats{
		MemTotal:     memStats.Total,
		MemAvailable: memStats.Available,
		MemUsed:      memStats.Used,
		MemFree:      memStats.Free,
	}
	return stats, nil
}
