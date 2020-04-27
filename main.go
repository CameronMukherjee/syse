package main

import (
	"./getinternet"
	"./getos"
)

func main() {
	getos.GetUserDetails()      //Works without errors
	getos.GetProcessorDetails() //Works without errors
	getos.GetOS()               //Works without errors
	getos.GetUptime()           //Works without errors
	getos.GetVirtualMemUsage()  //Works without errors
	getos.GetCPUUsage()         //Works without errors
	getinternet.GetLocalIP()    //Works without errors
	getinternet.GetPublicIP()   //Works without errors - returning []byte
}

// getos.GetDiskSpace("/")     //Works without errors
