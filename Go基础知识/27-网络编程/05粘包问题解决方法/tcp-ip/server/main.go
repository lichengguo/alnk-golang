package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"tcp-ip/protocol"
)

// 粘包问题

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		// 调用协议解码数据
		recvStr, err := protocol.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			return
		}
		fmt.Println("收到客户端的数据: ", recvStr)
	}
}

func main() {
	// 1.
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()

	// 2.
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// 3.
		go process(conn)
	}
}
