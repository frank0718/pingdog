package util

import (
	"./ping" //当前目录
)

func PingCheck(host string) bool {
	LogInfo().Println("start check ip")
	pinger, err := ping.NewPinger(host)
	if err != nil {
		LogError().Println(err)
		panic(err)
	}
	pinger.Count = 10
	pinger.Run()                 // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats
	LogInfo().Println(stats)
	var down bool
	if stats.PacketsRecv == 0 {
		down = true
	} else {
		down = false
	}
	return down
}
