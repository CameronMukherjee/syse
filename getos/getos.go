package getos

import (
	"log"
	"os/user"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"

	"github.com/shirou/gopsutil/host"

	"github.com/denisbrodbeck/machineid"
	"github.com/klauspost/cpuid"
)

// User struct, contains name, uid, username and home directory.
type User struct {
	HWID          string `Json:"eHWID"`
	GID           string `Json:"GID"`
	Name          string `Json:"Name"`
	UID           string `Json:"UID"`
	Username      string `Json:"Username"`
	HomeDirectory string `Json:"HomeDirectory"`
}

// CPUInfo contains information about the CPU
type CPUInfo struct {
	BrandName      string      `Json:"BrandName"`
	PhysicalCores  int         `Json:"PhysicalCores"`
	ThreadsPerCore int         `Json:"ThreadsPerCore"`
	LogicalCores   int         `Json:"LogicalCores"`
	Family         int         `Json:"Family"`
	Features       cpuid.Flags `Json:"Features"`
}

// DiskStatus contains information about local system disk space.
type DiskStatus struct {
	All  uint64 `Json:"all"`
	Used uint64 `Json:"used"`
	Free uint64 `Json:"free"`
}

// PCI returning vendorname and productName
type PCI struct {
	VendorName  string `Json:"VendorName"`
	ProductName string `Json:"ProductName"`
}

// OS returning Operating System Details
type OS struct {
	Platform       string `Json:"Platform"`
	Family         string `Json:"Famil"`
	ProductVersion string `Json:"ProductVersion"`
}

// Uptime of system.
type Uptime struct {
	Days    uint64 `Json:"Days"`
	Hours   uint64 `Json:"Hours"`
	Minutes uint64 `Json:"Minutes"`
	Seconds uint64 `Json:"Seconds"`
}

// VirtualMemory usage of system.
type VirtualMemory struct {
	Total       uint64  `Json:"Total"`
	Free        uint64  `Json:"Free"`
	UsedPercent float64 `Json:"UsedPercentage"`
}

// GetUserDetails returns the local user details logged into the computer.
func GetUserDetails() User {
	user, err := user.Current()
	if err != nil {
		log.Println(err)
	}
	id, err := machineid.ProtectedID("abigail123")
	if err != nil {
		log.Println(err)
	}

	return User{
		HWID:          id,
		GID:           user.Gid,
		Name:          user.Name,
		UID:           user.Uid,
		Username:      user.Username,
		HomeDirectory: user.HomeDir,
	}
}

// GetProcessorDetails returns pc specs.
func GetProcessorDetails() CPUInfo {
	return CPUInfo{
		BrandName:      cpuid.CPU.BrandName,
		PhysicalCores:  cpuid.CPU.PhysicalCores,
		ThreadsPerCore: cpuid.CPU.ThreadsPerCore,
		LogicalCores:   cpuid.CPU.LogicalCores,
		Family:         cpuid.CPU.Family,
		Features:       cpuid.CPU.Features,
	}
}

// GetOS returning OS struct including details of the operating system.
func GetOS() OS {
	platform, family, pversion, err := host.PlatformInformation()
	if err != nil {
		log.Println(err)
	}
	return OS{
		Platform:       platform,
		Family:         family,
		ProductVersion: pversion,
	}
}

// GetUptime of the system and returns a Uptime Struct
func GetUptime() Uptime {
	uptime, _ := host.Uptime()
	days := uptime / (60 * 60 * 24)
	hours := (uptime - (days * 60 * 60 * 24)) / (60 * 60)
	minutes := ((uptime - (days * 60 * 60 * 24)) - (hours * 60 * 60)) / 60
	return Uptime{
		Days:    days,
		Hours:   hours,
		Minutes: minutes,
		Seconds: uptime,
	}
}

// GetVirtualMemUsage returns the virtual memory usage of the system.
func GetVirtualMemUsage() VirtualMemory {
	v, err := mem.VirtualMemory()
	if err != nil {
		log.Println(err)
	}
	return VirtualMemory{
		Total:       v.Total,
		Free:        v.Free,
		UsedPercent: v.UsedPercent,
	}
}

// GetCPUUsage returns the cpu usage
func GetCPUUsage() []float64 {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Println(err)
	}
	return percent
}

// GetDisKSpace only works within Unix Systems

// GetDiskSpace returns diskspace details in struct based on path put in.
// func GetDiskSpace(path string) DiskStatus {
// 	fs := syscall.Statfs_t{}
// 	// Maybe change path into hardcoded C:// or \\
// 	err := syscall.Statfs(path, &fs)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	return DiskStatus{
// 		All:  fs.Blocks * uint64(fs.Bsize),
// 		Free: fs.Bfree * uint64(fs.Bsize),
// 		Used: (fs.Blocks * uint64(fs.Bsize)) - (fs.Bfree * uint64(fs.Bsize)),
// 	}
// }

// GetPCI only working on Linux

// GetPCI returns a list of PCI structs including vendorNames and productNames.
// func GetPCI() []PCI {
// 	// Does not work with MacOS.
// 	pci, err := ghw.PCI()
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	devices := pci.ListDevices()
// 	if len(devices) == 0 {
// 		log.Fatal("Error getting PCI Details")
// 	}

// 	var pciList = []PCI{}

// 	for _, device := range devices {
// 		vendor := device.Vendor
// 		vendorName := vendor.Name

// 		product := device.Product
// 		productName := product.Name

// 		newDevice := PCI{
// 			VendorName:  vendorName,
// 			ProductName: productName,
// 		}
// 		pciList = append(pciList, newDevice)
// 	}

// 	return pciList
// }
