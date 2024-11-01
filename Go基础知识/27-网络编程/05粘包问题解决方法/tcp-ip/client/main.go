package main

import (
	"fmt"
	"net"
	"tcp-ip/protocol"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for i := 0; i < 20; i++ {
		msg := "Hello, Hello. How are you?"
		// 调用协议编码数据
		b, err := protocol.Encode(msg)
		if err != nil {
			fmt.Println("encode failed, error: ", err)
		}
		conn.Write(b)
	}
}
