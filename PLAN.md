# 🚀 go-sysinfo: Development Roadmap & Packaging Plan

A lightweight, zero-dependency Linux system information CLI written in Go from scratch.

---

## 🎯 Project Goals
- **Deep Linux Understanding:** Interact directly with `/proc` virtual filesystem and Linux syscalls.
- **Idiomatic Go:** Learn clean architecture, error handling, structs, interfaces, and unit testing.
- **Open-Source & CV Ready:** Production-grade code formatting, automated CI/CD releases, and `.deb` packaging for Ubuntu/Debian.

---

## 📋 Milestone Checklist

### Phase 1: Go Basics & Project Architecture
- [x] Understand Go fundamentals (`package`, `import`, structs, functions, error handling).
- [x] Establish package structure (`internal/sysinfo/`).

### Phase 2: Linux Metric Collectors (No external libraries)
- [x] **Memory (`/proc/meminfo`):** Total, Free, Available, Buffers/Cached, Used %, Swap.
- [x] **CPU (`/proc/cpuinfo`):** Model name, physical & logical core counts.
- [x] **Host / OS (`/etc/os-release`, syscalls):** Distro name, version, hostname.
- [x] **Uptime (`/proc/uptime`):** Raw seconds parsed into human-friendly format.
- [x] **Disk Usage (`syscall.Statfs`):** Total, free, and used space on root `/`.

### Phase 3: CLI Interface & Formats
- [ ] Command-line flags (e.g., `--json`, `--human`, `--compact`, `--version`).
- [ ] Clean terminal styling with ANSI colors and formatted ASCII tables.

### Phase 4: Quality & Testing
- [ ] Unit tests for string parsing and math conversions using Go standard `testing` package.
- [ ] Graceful error handling (handling non-Linux environments and permission restrictions).

### Phase 5: Build, Packaging & Distribution
- [x] Multi-architecture builds (`amd64`).
- [x] Debian / Ubuntu `.deb` packaging script (`scripts/build-deb.sh`).
- [ ] GitHub Actions CI workflow (linting, tests, build verification).
- [ ] Packaging with **GoReleaser** for automated GitHub release assets.
- [ ] Publish GitHub Release with pre-built binaries.
- [x] Polished `README.md` with badges, instructions, and architecture breakdown.
