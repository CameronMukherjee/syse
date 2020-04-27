package main

import "./getos"

func main() {
	getos.GetUserDetails()      //Works without errors
	getos.GetProcessorDetails() //Works without errors
	getos.GetOS()               //Works without errors
	getos.GetUptime()           //Works without errors
	getos.GetVirtualMemUsage()  //Works without errors
	getos.GetCPUUsage()         //Works without errors
}

// getos.GetDiskSpace("/")     //Works without errors
