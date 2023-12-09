package services

import (
	"time"

	"github.com/shirou/gopsutil/v3/net"

	"picel.pidash/models"
)

func GetNetInfo() ([]models.NetInfo, error) {
	netInfo, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	netInfoList := make([]models.NetInfo, len(netInfo))
	for i, n := range netInfo {
		netInfoList[i] = models.NetInfo{
			Interface: n.Name,
			Mac:       n.HardwareAddr,
		}
	}

	return netInfoList, nil
}

func GetNetStats() (models.NetStats, error) {
	stats1, err := net.IOCounters(true)
	if err != nil {
		return models.NetStats{}, err
	}
	time.Sleep(time.Second * 1)
	stats2, err := net.IOCounters(true)
	if err != nil {
		return models.NetStats{}, err
	}

	// find the biggest difference between stats1 and stats2
	var maxDiffTx uint64
	var maxDiffRx uint64
	var maxDiffName string
	for _, stat1 := range stats1 {
		for _, stat2 := range stats2 {
			if stat1.Name == stat2.Name {
				diffTx := stat2.BytesSent - stat1.BytesSent
				diffRx := stat2.BytesRecv - stat1.BytesRecv
				if diffTx > maxDiffTx {
					maxDiffTx = diffTx
					maxDiffName = stat1.Name
				}
				if diffRx > maxDiffRx {
					maxDiffRx = diffRx
					maxDiffName = stat1.Name
				}
			}
		}
	}

	netStat := models.NetStats{
		Interface: maxDiffName,
		TxSpeed:   maxDiffTx,
		RxSpeed:   maxDiffRx,
	}

	return netStat, nil
}
