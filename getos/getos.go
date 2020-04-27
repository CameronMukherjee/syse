package getos

import (
	"log"
	"os/user"
	"syscall"

	"github.com/denisbrodbeck/machineid"
	"github.com/jaypipes/ghw"
	"github.com/klauspost/cpuid"
)

// User struct, contains name, uid, username and home directory.
type User struct {
	HWID          string `json:"eHWID"`
	Name          string `json:"Name"`
	UID           string `json:"UID"`
	Username      string `json:"Username"`
	HomeDirectory string `json:"HomeDirectory"`
}

// CPUInfo contains information about the CPU
type CPUInfo struct {
	BrandName      string      `json:"BrandName"`
	PhysicalCores  int         `json:"PhysicalCores"`
	ThreadsPerCore int         `json:"ThreadsPerCore"`
	LogicalCores   int         `json:"LogicalCores"`
	Family         int         `json:"Family"`
	Features       cpuid.Flags `json:"Features"`
}

// DiskStatus contains information about local system disk space.
type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

// PCI returning vendorname and productName
type PCI struct {
	VendorName  string `json:"VendorName"`
	ProductName string `json:"ProductName"`
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

// GetDiskSpace returns diskspace details in struct based on path put in.
func GetDiskSpace(path string) DiskStatus {
	fs := syscall.Statfs_t{}
	// Maybe change path into hardcoded C:// or \\
	err := syscall.Statfs(path, &fs)
	if err != nil {
		log.Println(err)
	}

	return DiskStatus{
		All:  fs.Blocks * uint64(fs.Bsize),
		Free: fs.Bfree * uint64(fs.Bsize),
		Used: (fs.Blocks * uint64(fs.Bsize)) - (fs.Bfree * uint64(fs.Bsize)),
	}
}

// GetPCI returns a list of PCI structs including vendorNames and productNames.
func GetPCI() []PCI {
	// Does not work with MacOS.
	pci, err := ghw.PCI()
	if err != nil {
		log.Println(err)
	}
	devices := pci.ListDevices()
	if len(devices) == 0 {
		log.Fatal("Error getting PCI Details")
	}

	var pciList = []PCI{}

	for _, device := range devices {
		vendor := device.Vendor
		vendorName := vendor.Name

		product := device.Product
		productName := product.Name

		newDevice := PCI{
			VendorName:  vendorName,
			ProductName: productName,
		}
		pciList = append(pciList, newDevice)
	}

	return pciList
}
