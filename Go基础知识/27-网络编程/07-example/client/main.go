package main

import (
	"07-example/protocol"
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

//客户端

func main() {
	//1.与服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("连接服务端失败,error: ", err)
		return
	}
	//2.发送数据
	readerScreen := bufio.NewReader(os.Stdin)
	readerRcev := bufio.NewReader(conn)

	for {
		//给服务端发送消息
		fmt.Print("请说话: ")
		msg, _ := readerScreen.ReadString('\n')
		msg = strings.TrimSpace(msg)
		//退出
		if msg == "exit" {
			break
		}
		//不能发送空白消息
		if len(msg) == 0 {
			continue
		}
		//编码消息
		msgByte, err := protocol.Encode(msg)
		if err != nil {
			fmt.Println("编码失败, error: ", err)
			return
		}
		conn.Write(msgByte)

		//接收从服务端返回的消息
		//调用协议解码数据
		recvStr, err := protocol.Decode(readerRcev)
		if err == io.EOF {
			return
		}
		if err != nil {
			return
		}
		fmt.Println("服务端的回复: ", recvStr)

	}
	defer conn.Close()
}
