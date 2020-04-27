package getinternet

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
)

// IP struct for local and public IP
type IP struct {
	LocalIP  net.IP `Json:"Local IP"`
	PublicIP string `Json:"Public IP"`
}

type GETReq struct {
	ip      string
	country string
	cc      string
}

// GetIPs returns local and public IP
func GetIPs() IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	localIP := localAddr.IP

	resp, err := http.Get("https://api.myip.com")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)

	return IP{
		LocalIP:  localIP,
		PublicIP: result["ip"],
	}
}

// // GetLocalIP returns local IP address of machine
// func GetLocalIP() net.IP {
// 	conn, err := net.Dial("udp", "8.8.8.8:80")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer conn.Close()

// 	localAddr := conn.LocalAddr().(*net.UDPAddr)
// 	return localAddr.IP
// }

// // GetPublicIP returns the machines public IP address
// func GetPublicIP() string {
// 	resp, err := http.Get("https://api.myip.com")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	return string([]byte(body))
// }
