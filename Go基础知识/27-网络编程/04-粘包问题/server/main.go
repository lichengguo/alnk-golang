package main

import (
	"fmt"
	"io"
	"net"
)

// 粘包问题

func process(conn net.Conn) {
	defer conn.Close()

	var buf [1024]byte
	for {
		n, err := conn.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		recvStr := string(buf[:n])
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
