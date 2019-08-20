package main

import (
	"github.com/Unknwon/goconfig"
	"net"
	"os"
	"fmt"
	"strconv"
)

func main() {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil{
		panic("Load config fail")
	}
	listenaddr, err := cfg.GetValue("udp-tester", "listenaddr")
	minport, err := cfg.Int("udp-tester", "minport")
	maxport, err := cfg.Int("udp-tester", "maxport")
	for i := minport;i <= maxport ;i++ {
		//fmt.Printf("Try to connect udp port %d\n", i)
		i := strconv.Itoa(i)
		service := listenaddr + ":" + i
		udpAddr, err := net.ResolveUDPAddr("udp4", service)
		checkError(err)

		conn, err := net.DialUDP("udp", nil, udpAddr)
		checkError(err)

		_, err = conn.Write([]byte("anything"))
		checkError(err)

		var buf [512]byte
		n, err := conn.Read(buf[0:])
		checkError(err)

		fmt.Println(string(buf[0:n]))
	}
	//os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR]  Fatal error ", err.Error())
		os.Exit(1)
	}
}