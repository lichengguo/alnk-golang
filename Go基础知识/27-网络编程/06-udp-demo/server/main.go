package main

import (
	"fmt"
	"net"
	"strings"
)

// udp

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("listen upd failed, error:", err)
		return
	}
	defer conn.Close()

	// 不需要建立连接直接收发数据
	var data [1024]byte
	for {
		n, addr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read from udp failed, error:", err)
			return
		}
		fmt.Println(string(data[:n])) //打印收到的数据

		// 返回数据
		sendString := strings.ToUpper(string(data[:]))
		conn.WriteToUDP([]byte(sendString), addr)
	}
}
