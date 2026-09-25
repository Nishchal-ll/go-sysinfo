package sysinfo

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// CPUInfo holds details about processor model and core count.
type CPUInfo struct {
	ModelName    string
	LogicalCores int
	PhysicalCores int
	SpeedMHz     float64
}

// GetCPUInfo reads /proc/cpuinfo to extract CPU architecture details.
func GetCPUInfo() (*CPUInfo, error) {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return nil, fmt.Errorf("failed to open /proc/cpuinfo: %w", err)
	}
	defer file.Close()

	info := &CPUInfo{
		LogicalCores: runtime.NumCPU(), // Standard Go runtime core count
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, ":") {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		switch key {
		case "model name":
			if info.ModelName == "" {
				info.ModelName = val
			}
		case "cpu cores":
			if info.PhysicalCores == 0 {
				if cores, err := strconv.Atoi(val); err == nil {
					info.PhysicalCores = cores
				}
			}
		case "cpu MHz":
			if info.SpeedMHz == 0 {
				if mhz, err := strconv.ParseFloat(val, 64); err == nil {
					info.SpeedMHz = mhz
				}
			}
		}
	}

	return info, nil
}
