package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"net"
	"os"
	"strconv"
	"time"
	_ "github.com/Unknwon/goconfig"
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
		fmt.Printf("Listen udp port %d\n", i)
		go func() {
			i := strconv.Itoa(i)
			service := listenaddr + ":" + i
			udpAddr, err := net.ResolveUDPAddr("udp4", service)
			checkError(err)

			conn, err := net.ListenUDP("udp", udpAddr)
			checkError(err)

			for {
				handleClient(conn, i)
			}
		}()
		time.Sleep(time.Duration(100)*time.Millisecond)
	}
	select {}

}

func handleClient(conn *net.UDPConn, port string) {

	var buf [1024]byte

	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}

	//daytime := time.Now().String()
	conninfo := "[INFO]  " + time.Now().String() + "  " + port +" Connect Sucess!!!"

	//conn.WriteToUDP([]byte(daytime), addr)
	conn.WriteToUDP([]byte(conninfo), addr)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}