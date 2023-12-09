package utils

import (
	"log"
	"net"
)

func GetUserIP() []string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
		return []string{}
	}

	results := []string{}

	for _, address := range addrs {
		ipnet, ok := address.(*net.IPNet)
		if ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipString := ipnet.IP.String()
				// check if the IP is in the private range
				if ipString[:3] == "10." || ipString[:7] == "192.168" || ipString[:4] == "172." {
					results = append(results, ipString)
				}
			}
		}
	}

	if len(results) == 0 {
		return []string{}
	}

	return results
}
