package sysinfo

import (
	"fmt"
	"syscall"
)

// DiskInfo holds root filesystem disk space details.
type DiskInfo struct {
	Path        string
	TotalBytes  uint64
	FreeBytes   uint64
	UsedBytes   uint64
	UsedPercent float64
}

// FormatDiskSize converts raw bytes into readable GB string.
func FormatDiskSize(bytes uint64) string {
	gb := float64(bytes) / 1024.0 / 1024.0 / 1024.0
	return fmt.Sprintf("%.2f GB", gb)
}

// GetDiskInfo calls the Linux Statfs syscall to inspect disk usage.
func GetDiskInfo(path string) (*DiskInfo, error) {
	var stat syscall.Statfs_t
	err := syscall.Statfs(path, &stat)
	if err != nil {
		return nil, fmt.Errorf("statfs failed for path %s: %w", path, err)
	}

	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bfree * uint64(stat.Bsize)
	avail := stat.Bavail * uint64(stat.Bsize)
	used := total - free

	var usedPercent float64
	if total > 0 {
		usedPercent = (float64(used) / float64(total)) * 100.0
	}

	return &DiskInfo{
		Path:        path,
		TotalBytes:  total,
		FreeBytes:   avail,
		UsedBytes:   used,
		UsedPercent: usedPercent,
	}, nil
}
