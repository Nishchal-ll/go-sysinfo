package main

import (
	"fmt"

	"github.com/Nishchal-ll/go-sysinfo/internal/sysinfo"
)

const version = "0.1.0"

func main() {
	// CLI Banner
	fmt.Println("========================================")
	fmt.Println("       🐧 go-sysinfo System Summary      ")
	fmt.Printf("               Version %s               \n", version)
	fmt.Println("========================================")

	// 1. Host & OS
	host, err := sysinfo.GetHostInfo()
	if err == nil {
		fmt.Printf("🖥️  Host:     %s\n", host.Hostname)
		fmt.Printf("📦 OS:       %s\n", host.OSName)
		fmt.Printf("⏱️  Uptime:   %s\n", host.Uptime)
	}

	// 2. CPU Info
	cpu, err := sysinfo.GetCPUInfo()
	if err == nil {
		fmt.Println("\n--- ⚙️  CPU ---")
		fmt.Printf("Model:       %s\n", cpu.ModelName)
		fmt.Printf("Cores:       %d physical / %d logical threads\n", cpu.PhysicalCores, cpu.LogicalCores)
	}

	// 3. Live RAM & Swap
	mem, err := sysinfo.GetMemoryInfo()
	if err == nil {
		fmt.Println("\n--- 🧠 Memory ---")
		fmt.Printf("RAM:         %s / %s (%.1f%% used)\n",
			sysinfo.FormatGiB(mem.UsedKB),
			sysinfo.FormatGiB(mem.TotalKB),
			mem.UsedPercent,
		)
		if mem.SwapTotalKB > 0 {
			fmt.Printf("Swap:        %s / %s (%.1f%% used)\n",
				sysinfo.FormatGiB(mem.SwapUsedKB),
				sysinfo.FormatGiB(mem.SwapTotalKB),
				mem.SwapPercent,
			)
		}
	}

	// 4. Disk Usage
	disk, err := sysinfo.GetDiskInfo("/")
	if err == nil {
		fmt.Println("\n--- 💾 Disk (/) ---")
		fmt.Printf("Usage:       %s / %s (%.1f%% used)\n",
			sysinfo.FormatDiskSize(disk.UsedBytes),
			sysinfo.FormatDiskSize(disk.TotalBytes),
			disk.UsedPercent,
		)
	}

	fmt.Println("========================================")
}
