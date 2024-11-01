package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// 客户端
/*
存在的问题:
1.每次只能接收128byte，超过128byte的消息会堵塞
*/

func main() {
	// 1.与服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("连接服务端失败,error: ", err)
		return
	}
	defer conn.Close()

	// 2.发送数据
	var tmp [128]byte
	reader := bufio.NewReader(os.Stdin) //创建一个缓冲区用来获取终端输入

	for {
		// 给服务端发送消息
		fmt.Print("请说话: ")
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("screen input failed, error: ", err)
			break
		}
		msg = strings.TrimSpace(msg)
		if msg == "" {
			continue //没消息内容直接不发给服务端
		}
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))

		// 接收从服务端返回的消息
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("get from server info failed, error: ", err)
			return
		}
		fmt.Println("服务端的回复: ", string(tmp[:n]))
	}
}
