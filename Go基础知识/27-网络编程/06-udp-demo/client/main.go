package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("连接服务端失败, error: ", err)
		return
	}
	defer socket.Close()

	var reply [1024]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入内容: ")
		msg, _ := reader.ReadString('\n')
		socket.Write([]byte(msg))

		//收回复的数据
		n, _, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("recv reply msg failed, error:", err)
			return
		}
		fmt.Println("收到回复: ", string(reply[:n]))
	}
}
