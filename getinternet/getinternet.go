package getinternet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

// PublicIP storing information about the public ip address
type PublicIP struct {
	ip      string
	country string
	cc      string
}

// GetLocalIP returns local IP address of machine
func GetLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.IP)
	return string(localAddr.IP)
}

// GetPublicIP returns the machines public IP address
func GetPublicIP() PublicIP {
	resp, err := http.Get("https://api.myip.com")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	var response PublicIP
	json.Unmarshal([]byte(body), &response)
	fmt.Println(response)
	return response
}
