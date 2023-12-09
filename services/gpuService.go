package services

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	"picel.pidash/models"
)

func GetGPUInfo() ([]models.GPUInfo, error) {
	// get product name, driver version, total memory by calling nvidia-smi
	cmd := exec.Command("nvidia-smi", "--query-gpu=name,driver_version,memory.total", "--format=csv,noheader,nounits")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
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
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
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
		MaxClock: models.Clock{
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
	cmd := exec.Command("nvidia-smi", "--query-gpu=memory.used,memory.total,memory.free,memory.reserved,utilization.gpu,utilization.memory,temperature.gpu,power.draw,power.limit", "--format=csv,noheader,nounits")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	Output, err := cmd.Output()
	if err != nil {
		return models.GPUStats{}, err
	}
	// trim \r\n
	trimmed := strings.Trim(string(Output), "\r\n")
	results := strings.Split(string(trimmed), ", ")
	// convert all data of results to int
	intConverted := make([]int, 7)
	floatConverted := make([]float64, 2)
	for i, val := range results {
		if i == 7 || i == 8 {
			floatVal, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return models.GPUStats{}, err
			}
			floatConverted[i-7] = floatVal
		} else {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				return models.GPUStats{}, err
			}
			intConverted[i] = intVal
		}
	}

	memoryUsage := models.MemoryUsage{
		Used:     intConverted[0],
		Total:    intConverted[1],
		Free:     intConverted[2],
		Reserved: intConverted[3],
	}

	utilization := models.Utilization{
		GPU:    intConverted[4],
		Memory: intConverted[5],
	}
	var temperature int = intConverted[6]

	power := models.Power{
		Usage: floatConverted[0],
		Limit: floatConverted[1],
	}

	// get clocks
	cmd = exec.Command("nvidia-smi", "-q", "-d", "CLOCK")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
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
		Clock: models.Clock{
			GraphicsClock: clocks[0],
			SmClock:       clocks[1],
			MemClock:      clocks[2],
			VideoClock:    clocks[3],
		},
	}

	return gpuStats, nil
}
