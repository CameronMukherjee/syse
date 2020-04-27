package main

import "./getos"

func main() {
	getos.GetUserDetails()      //Works without errors
	getos.GetProcessorDetails() //Works without errors
	getos.GetDiskSpace("/")     //Works without errors
	// getos.GetPCI()
}
