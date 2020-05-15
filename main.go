package main

import (
	"./getinternet"
	"./getos"
)

// Export struct of all other structs
type Export struct {
	User    getos.User          `Json:"User"`
	CPUInfo getos.CPUInfo       `Json:"CPU Info"`
	OSInfo  getos.OS            `Json:"OS"`
	Uptime  getos.Uptime        `Json:"Uptime"`
	VirtMem getos.VirtualMemory `Json:"Virtual Memory"`
	IPs     getinternet.IP      `Json:"IP Addresses"`
}

func main() {
	postObject := Export{
		User:    getos.GetUserDetails(),
		CPUInfo: getos.GetProcessorDetails(),
		OSInfo:  getos.GetOS(),
		Uptime:  getos.GetUptime(),
		VirtMem: getos.GetVirtualMemUsage(),
		IPs:     getinternet.GetIPs(),
	}
	getfiles.JSONAddToFile(postObject)
}

// Only works on UNIX
// getos.GetDiskSpace("/")     //Works without errors
// getos.GetPCI() 			   //Untested

// fmt.Printf("%+v\n", postObject)
