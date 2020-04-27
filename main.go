package main

import (
	"encoding/json"
	"log"
	"os"

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
	// fmt.Printf("%+v\n", postObject)
	jsonAddToFile(postObject)
}

func jsonAddToFile(input interface{}) {
	var jsonData []byte
	jsonData, err := json.MarshalIndent(input, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	data := string(jsonData)
	// data = data + "," + "\n"
	data = data + ","
	// f, err := os.OpenFile("output.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
