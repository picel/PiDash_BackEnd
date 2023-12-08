package services

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"picel.pidash/models"
)

func GetGPUInfo() ([]models.GPUInfo, error) {
	// get product name, driver version, total memory by calling nvidia-smi
	cmd := exec.Command("nvidia-smi", "--query-gpu=name,driver_version,memory.total", "--format=csv,noheader,nounits")
	Output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	pInfo := strings.Split(string(Output), ", ")
	if len(pInfo) != 3 {
		return nil, errors.New("Error parsing GPU info")
	}
	for val := range pInfo {
		pInfo[val] = strings.Trim(pInfo[val], "\r\n")
	}

	// get MaxClocks
	cmd = exec.Command("nvidia-smi", "-q", "-d", "CLOCK")
	clockInfo, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// find "Max Clocks" and then get the next 4 lines
	var maxClocks []string
	clockInfos := strings.Split(string(clockInfo), "\n")
	for i, line := range clockInfos {
		if strings.Contains(line, "Max Clocks") {
			for j := i + 1; j < i+5; j++ {
				// remove all spaces
				tmp := strings.ReplaceAll(clockInfos[j], " ", "")
				// split by ":"
				tmpList := strings.Split(tmp, ":")
				maxClocks = append(maxClocks, strings.Trim(tmpList[1], "\r"))
			}
			break
		}
	}

	// parse pInfo
	gpuInfo := models.GPUInfo{
		ProductName:   string(pInfo[0]),
		DriverVersion: string(pInfo[1]),
		TotalMemory:   string(pInfo[2]),
		MaxClocks: models.Clocks{
			GraphicsClock: string(maxClocks[0]),
			SmClock:       string(maxClocks[1]),
			MemClock:      string(maxClocks[2]),
			VideoClock:    string(maxClocks[3]),
		},
	}

	return []models.GPUInfo{gpuInfo}, nil
}

func GetGPUStats() (models.GPUStats, error) {
	// memoryUsage, power
	cmd := exec.Command("nvidia-smi", "--query-gpu=memory.used,memory.total,memory.free,memory.reserved,power.draw,power.limit", "--format=csv,noheader")
	Output, err := cmd.Output()
	if err != nil {
		return models.GPUStats{}, err
	}
	// trim \r\n
	trimmed := strings.Trim(string(Output), "\r\n")
	results := strings.Split(string(trimmed), ", ")
	memoryUsage := models.MemoryUsage{
		Used:     string(results[0]),
		Total:    string(results[1]),
		Free:     string(results[2]),
		Reserved: string(results[3]),
	}
	power := models.Power{
		Usage: string(results[4]),
		Limit: string(results[5]),
	}

	// utilization, temperature
	cmd = exec.Command("nvidia-smi", "--query-gpu=utilization.gpu,utilization.memory,temperature.gpu", "--format=csv,noheader,nounits")
	Output, err = cmd.Output()
	if err != nil {
		return models.GPUStats{}, err
	}
	trimmed = strings.Trim(string(Output), "\r\n")
	results = strings.Split(string(trimmed), ", ")
	utilization := models.Utilization{
		GPU:    string(results[0]),
		Memory: string(results[1]),
	}
	temperature := string(results[2])

	fmt.Println(memoryUsage, utilization, temperature, power)

	// get clocks
	cmd = exec.Command("nvidia-smi", "-q", "-d", "CLOCK")
	clockInfo, err := cmd.Output()
	if err != nil {
		return models.GPUStats{}, err
	}

	// find "Clocks" and then get the next 4 lines
	var clocks []string
	clockInfos := strings.Split(string(clockInfo), "\n")
	for i, line := range clockInfos {
		if strings.Contains(line, "Clocks") {
			for j := i + 1; j < i+5; j++ {
				// remove all spaces
				tmp := strings.ReplaceAll(clockInfos[j], " ", "")
				// split by ":"
				tmpList := strings.Split(tmp, ":")
				clocks = append(clocks, strings.Trim(tmpList[1], "\r"))
			}
			break
		}
	}

	gpuStats := models.GPUStats{
		MemoryUsage: memoryUsage,
		Utilization: utilization,
		Temperature: temperature,
		Power:       power,
		Clocks: models.Clocks{
			GraphicsClock: clocks[0],
			SmClock:       clocks[1],
			MemClock:      clocks[2],
			VideoClock:    clocks[3],
		},
	}

	return gpuStats, nil
}
