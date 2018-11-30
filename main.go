package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"./src/util"
)

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func checkdown(t time.Time) {
	util.LogInfo().Println(t, Servers)
	for ip, name := range Servers {
		go func(ip, name string) {
			if ip != util.GetSelfIp() {
				if util.PingCheck(ip) {
					sender := Servers[util.GetSelfIp()]
					util.Sendmail(fmt.Sprint(sender+"@vps.com"), Mailto, fmt.Sprint("send from: "+sender+"\n"+name+"\n"+ip+": server is down"))
				} else {
					util.LogInfo().Println("server " + name + " is ok")
				}
			}
		}(ip, name)
	}
}

type Config struct {
	Mailto  string            `json: "mailto"`
	Servers map[string]string `json: "server"`
}

var Mailto string
var Servers map[string]string

func initConfig() {
	_c, _ := ioutil.ReadFile("config.json")
	var c Config
	json.Unmarshal(_c, &c)
	Mailto = c.Mailto
	Servers = c.Servers
}

func main() {
	util.LoggerInit()
	initConfig()

	util.LogInfo().Println("info log ")
	util.LogWarning().Println("warn log ")
	util.LogError().Println("error log")

	doEvery(3600*time.Second, checkdown)
}
