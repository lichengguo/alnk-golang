package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// Go语言实现TCP通信
// 服务端
/*
存在的问题:
1.每次只能接收128byte，超过128byte的消息会堵塞
2.服务端如果回复是空包，那么会堵塞
*/

// processConn 与客户端通信
func processConn(conn net.Conn) {
	defer conn.Close() //关闭连接

	var tmp [128]byte                   //元素为byte类型的数组
	reader := bufio.NewReader(os.Stdin) //创建一个缓冲区用来获取终端输入

	for {
		// 接收消息
		n, err := conn.Read(tmp[:]) //读取客户端传过来的数据
		if err != nil {
			fmt.Println("read from client info failed, error: ", err)
			return
		}
		fmt.Println("客户端的消息: ", string(tmp[:n])) //打印客户端发送过来的消息

		// 给客户端回复消息
		fmt.Print("请回复: ")
		msg, err := reader.ReadString('\n') //读取终端输入的内容，以换行结束
		if err != nil {
			fmt.Println("screen input failed, error: ", err)
			break
		}
		msg = strings.TrimSpace(msg) //去掉首尾空格
		if msg == "exit" {
			break //退出
		}
		conn.Write([]byte(msg)) //发送消息给客户端
	}
}

func main() {
	// 1.启动监听本地端口服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("start tcp server on 127.0.0.1:20000 failed, error:", err)
		return
	}

	// 2.等待客户端过来建立连接
	for {
		conn, err := listener.Accept() //阻塞，等待连接
		if err != nil {
			fmt.Println("accept failed, error:", err)
		}

		// 3.响应客户端请求，与客户端通信
		go processConn(conn) //一个客户端分配一个goroutine去处理
	}
}
