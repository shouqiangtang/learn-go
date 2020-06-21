package main

import (
	"log"
	"net"
	"net/http"
	"strings"
)

// GetLocalIP4 gets local ip address.
func GetLocalIP4() (ip string) {
	interfaces, err := net.Interfaces()
	net.InterfaceAddrs()
	if err != nil {
		return
	}
	if len(interfaces) == 2 {
		for _, face := range interfaces {
			if strings.Contains(face.Name, "lo") {
				continue
			}
			addrs, err := face.Addrs()
			if err != nil {
				return
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						currIP := ipnet.IP.String()
						if !strings.Contains(currIP, ":") && currIP != "127.0.0.1" {
							ip = currIP
						}
					}
				}
			}
		}
	}
	for _, face := range interfaces {
		if strings.Contains(face.Name, "lo") {
			continue
		}
		addrs, err := face.Addrs()
		if err != nil {
			return
		}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					currIP := ipnet.IP.String()
					if !strings.Contains(currIP, ":") && currIP != "127.0.0.1" && isIntranetIpv4(currIP) {
						ip = currIP
					}
				}
			}
		}
	}

	return
}

func isIntranetIpv4(ip string) bool {
	if strings.HasPrefix(ip, "192.168.") ||
		strings.HasPrefix(ip, "169.254.") ||
		strings.HasPrefix(ip, "172.") ||
		strings.HasPrefix(ip, "10.30.") ||
		strings.HasPrefix(ip, "10.31.") {
		return true
	}
	return false
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := GetLocalIP4()
		if _, err := w.Write([]byte(ip)); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadGateway)
		}
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
