package main

import (
	"fmt"

	"./getinternet"
	"./getos"
)

// Final struct of all other structs
type Final struct {
	user     getos.User           `Json:"User"`
	cpuInfo  getos.CPUInfo        `Json:"CPU Info"`
	osInfo   getos.OS             `Json:"OS"`
	uptime   getos.Uptime         `Json:"Uptime"`
	virtMem  getos.VirtualMemory  `Json:"Virtual Memory"`
	cpuUsage []float64            `Json:"CPU Usage"`
	localIP  string               `Json:"Local IP"`
	publicIP getinternet.PublicIP `Json:"Public IP"`
}

func main() {
	finalSending := Final{
		user:     getos.GetUserDetails(),
		cpuInfo:  getos.GetProcessorDetails(),
		osInfo:   getos.GetOS(),
		uptime:   getos.GetUptime(),
		virtMem:  getos.GetVirtualMemUsage(),
		cpuUsage: getos.GetCPUUsage(),
		localIP:  getinternet.GetLocalIP(),
		publicIP: getinternet.GetPublicIP(),
	}
	fmt.Println(finalSending.localIP)
}

// Only works on UNIX
// getos.GetDiskSpace("/")     //Works without errors
// getos.GetPCI() 			   //Untested
