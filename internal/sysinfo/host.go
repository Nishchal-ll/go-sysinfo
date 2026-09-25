package sysinfo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// HostInfo holds host metadata and uptime.
type HostInfo struct {
	Hostname string
	OSName   string
	Uptime   string
}

// GetHostInfo gathers hostname, Linux distribution, and uptime.
func GetHostInfo() (*HostInfo, error) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Unknown"
	}

	info := &HostInfo{
		Hostname: hostname,
		OSName:   "Linux",
		Uptime:   "Unknown",
	}

	// 1. Read /etc/os-release for Linux distro
	if file, err := os.Open("/etc/os-release"); err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "PRETTY_NAME=") {
				val := strings.TrimPrefix(line, "PRETTY_NAME=")
				info.OSName = strings.Trim(val, "\"")
				break
			}
		}
	}

	// 2. Read /proc/uptime
	if data, err := os.ReadFile("/proc/uptime"); err == nil {
		fields := strings.Fields(string(data))
		if len(fields) > 0 {
			if secs, err := strconv.ParseFloat(fields[0], 64); err == nil {
				total := int(secs)
				days := total / 86400
				hours := (total % 86400) / 3600
				minutes := (total % 3600) / 60

				if days > 0 {
					info.Uptime = fmt.Sprintf("%d days, %d hours, %d mins", days, hours, minutes)
				} else {
					info.Uptime = fmt.Sprintf("%d hours, %d mins", hours, minutes)
				}
			}
		}
	}

	return info, nil
}
