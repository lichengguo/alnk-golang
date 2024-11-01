package main

import (
	"07-example/protocol"
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

//服务端

// processConn 与客户端通信
func processConn(conn net.Conn) {
	defer conn.Close()              //关闭连接
	reader := bufio.NewReader(conn) //缓存指针
	for {
		//接收消息
		//调用协议解码
		recvStr, err := protocol.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			return
		}
		fmt.Println("收到客户端的数据: ", recvStr)

		//回复消息
		recvStr = strings.Replace(recvStr, "?", "!", -1)
		recvStr = strings.Replace(recvStr, "？", "!", -1)
		recvStr = strings.Replace(recvStr, "你", "我", -1)
		recvStr = strings.Replace(recvStr, "吗", "", -1)

		//调用协议编码
		recvByte, err := protocol.Encode(recvStr)
		if err != nil {
			fmt.Println("编码失败, error:", err)
		}

		conn.Write(recvByte) //发送消息给客户端
	}
}

func main() {
	//1.启动监听本地端口服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("start tcp server on 127.0.0.1:20000 failed, error:", err)
		return
	}
	//2.等待客户端过来建立连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed, error:", err)
		}
		//3.响应客户端请求，与客户端通信
		go processConn(conn)
	}
}
