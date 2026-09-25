# 🐧 go-sysinfo

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://golang.org)
[![Platform](https://img.shields.io/badge/Platform-Linux-FCC624?style=flat&logo=linux&logoColor=black)](https://kernel.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

A lightweight, high-performance Linux system information CLI tool written in Go from scratch with **zero external dependencies**.

`go-sysinfo` queries the Linux virtual filesystem (`/proc`) and low-level kernel syscalls directly to display real-time host, CPU, memory, and disk statistics.

---

## ✨ Features

- ⚡ **Zero Dependencies:** Built entirely with Go standard library (`os`, `bufio`, `syscall`, `runtime`).
- 🖥️ **Host & OS Detection:** Reads `/etc/os-release` and hostname directly without spawning sub-shells.
- ⏱️ **Formatted Uptime:** Converts `/proc/uptime` seconds into human-readable days, hours, and minutes.
- ⚙️ **CPU Details:** Parses `/proc/cpuinfo` for processor model, physical cores, and logical execution threads.
- 🧠 **Live Memory & Swap:** Real-time RAM and Swap usage calculations via `/proc/meminfo`.
- 💾 **Disk Usage:** Direct root filesystem disk calculation using `syscall.Statfs`.

---

## 📸 Example Output

```text
========================================
       🐧 go-sysinfo System Summary      
               Version 0.1.0               
========================================
🖥️  Host:     Nishchal
📦 OS:       Ubuntu 26.04.1 LTS
⏱️  Uptime:   11 hours, 45 mins

--- ⚙️  CPU ---
Model:       AMD Ryzen 7 5800H with Radeon Graphics
Cores:       8 physical / 16 logical threads

--- 🧠 Memory ---
RAM:         5.51 GiB / 13.00 GiB (42.4% used)
Swap:        0.86 GiB / 4.00 GiB (21.4% used)

--- 💾 Disk (/) ---
Usage:       269.50 GB / 467.35 GB (57.7% used)
========================================
```

---

## 🚀 Quick Start & Installation

### Option 1: Install Debian / Ubuntu Package (`.deb`)
You can build and install the native `.deb` package on Ubuntu or Debian:
```bash
# Build the package
./scripts/build-deb.sh

# Install onto your system
sudo dpkg -i dist/go-sysinfo_0.1.0_amd64.deb

# Run anywhere from your terminal
sysinfo
```

### Option 2: Run with Go
```bash
go run main.go
```

### Option 3: Build Executable Binary
```bash
# Compile binary
go build -o sysinfo main.go

# Run the compiled binary
./sysinfo
```

### Option 4: Install Globally with Go
```bash
go install github.com/Nishchal-ll/go-sysinfo@latest
```

---

## 📂 Project Architecture

```
go-sysinfo/
├── main.go               # Application entry point & formatted terminal presentation
├── go.mod                # Go module definition
├── PLAN.md               # Project roadmap and packaging milestones
├── README.md             # Project documentation
└── internal/
    └── sysinfo/
        ├── host.go       # Hostname, OS distro detection, and system uptime
        ├── cpu.go        # CPU model and core counts via /proc/cpuinfo
        ├── memory.go     # RAM & Swap stats via /proc/meminfo
        └── disk.go       # Filesystem usage via syscall.Statfs
```

---

## 🗺️ Roadmap

Check out [PLAN.md](PLAN.md) to follow upcoming features:
- [ ] JSON output flag (`--json`)
- [ ] Automated `.deb` and `.rpm` packaging with GoReleaser
- [ ] GitHub Actions CI/CD release pipeline

---

## 📄 License

This project is open source and available under the [MIT License](LICENSE).
