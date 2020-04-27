package main

import "./getos"

func main() {
	getos.GetUserDetails()
	getos.GetProcessorDetails()
	getos.GetDiskSpace("\\")
	getos.GetPCI()
}
