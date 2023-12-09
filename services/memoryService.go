package services

import (
	"github.com/shirou/gopsutil/v3/mem"

	"picel.pidash/models"
	"picel.pidash/utils"
)

func GetMemInfo() (string, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}

	// return total
	total := utils.ByteCountDecimal(memInfo.Total)
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
