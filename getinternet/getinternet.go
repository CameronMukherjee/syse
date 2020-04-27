package getinternet

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

type IP struct {
	LocalIP  net.IP `Json:"Local IP"`
	PublicIP string `Json:"Public IP"`
}

// GetLocalIP returns local IP address of machine
func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

// GetPublicIP returns the machines public IP address
func GetPublicIP() string {
	resp, err := http.Get("https://api.myip.com")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	return string([]byte(body))
}
