package sysinfo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// MemoryInfo holds memory statistics in KB and percentages.
type MemoryInfo struct {
	TotalKB       uint64
	AvailableKB   uint64
	UsedKB        uint64
	UsedPercent   float64
	SwapTotalKB   uint64
	SwapFreeKB    uint64
	SwapUsedKB    uint64
	SwapPercent   float64
}

// FormatGiB converts KB into formatted "X.XX GiB" string.
func FormatGiB(kb uint64) string {
	gib := float64(kb) / 1024.0 / 1024.0
	return fmt.Sprintf("%.2f GiB", gib)
}

// GetMemoryInfo parses /proc/meminfo to calculate real live RAM usage.
func GetMemoryInfo() (*MemoryInfo, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return nil, fmt.Errorf("failed to open /proc/meminfo: %w", err)
	}
	defer file.Close()

	info := &MemoryInfo{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		key := strings.TrimSuffix(fields[0], ":")
		val, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			continue
		}

		switch key {
		case "MemTotal":
			info.TotalKB = val
		case "MemAvailable":
			info.AvailableKB = val
		case "SwapTotal":
			info.SwapTotalKB = val
		case "SwapFree":
			info.SwapFreeKB = val
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}

	// RAM calculation
	if info.TotalKB > 0 {
		info.UsedKB = info.TotalKB - info.AvailableKB
		info.UsedPercent = (float64(info.UsedKB) / float64(info.TotalKB)) * 100.0
	}

	// Swap calculation
	if info.SwapTotalKB > 0 {
		info.SwapUsedKB = info.SwapTotalKB - info.SwapFreeKB
		info.SwapPercent = (float64(info.SwapUsedKB) / float64(info.SwapTotalKB)) * 100.0
	}

	return info, nil
}
