package main

import (
	"encoding/json"
	"log"
	"os"

	"./getinternet"
	"./getos"
)

// Final struct of all other structs
type Final struct {
	user    getos.User          `Json:"User"`
	cpuInfo getos.CPUInfo       `Json:"CPU Info"`
	osInfo  getos.OS            `Json:"OS"`
	uptime  getos.Uptime        `Json:"Uptime"`
	virtMem getos.VirtualMemory `Json:"Virtual Memory"`
	getIPs  getinternet.IP      `Json:"IP Addresses"`
}

func main() {
	// postObject := Final{
	// 	user:    getos.GetUserDetails(),
	// 	cpuInfo: getos.GetProcessorDetails(),
	// 	osInfo:  getos.GetOS(),
	// 	uptime:  getos.GetUptime(),
	// 	virtMem: getos.GetVirtualMemUsage(),
	// 	getIPs:  getinternet.GetIPs(),
	// }
	// jsonAddToFile(postObject)
	jsonAddToFile(getos.GetUserDetails())
	jsonAddToFile(getos.GetProcessorDetails())
	jsonAddToFile(getos.GetOS())
	jsonAddToFile(getos.GetUptime())
	jsonAddToFile(getos.GetVirtualMemUsage())
	jsonAddToFile(getinternet.GetIPs())
}

func jsonAddToFile(input interface{}) {
	var jsonData []byte
	jsonData, err := json.MarshalIndent(input, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	data := string(jsonData)
	data = data + "," + "\n"
	f, err := os.OpenFile("output.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := f.WriteString(data); err != nil {
		log.Fatal(err)
	}
}

// Only works on UNIX
// getos.GetDiskSpace("/")     //Works without errors
// getos.GetPCI() 			   //Untested
