package util

import (
	"net"
)

func GetSelfIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		LogError().Println(err)
		panic(err)
		// os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		// os.Exit(1)
	}
	var addr string
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				addr = ipnet.IP.String()
				break
				// os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}
	return addr
}
