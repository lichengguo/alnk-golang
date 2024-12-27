package main

import (
	"fmt"
	"time"
)

/*
非阻塞通道
常规的通过通道发送和接收数据是阻塞的
然而可以使用带一个default子句的select来实现非阻塞的发送、接收，甚至是非阻塞的多路select
*/

func main() {
	// 创建一个字符串通道
	ch := make(chan string)

	// 启动一个并发任务，向通道发送数据
	go func() {
		time.Sleep(2 * time.Second) // 模拟任务处理时长
		ch <- "任务完成"
	}()

	// time.Sleep(3 * time.Second) // 主函数中等待一段时间

	// 使用select语句进行非阻塞通道操作
	select {
	case data := <-ch: // 从通道接收数据
		fmt.Println(data)
	default: // 没有收到数据
		fmt.Println("暂无任务结果")
	}

	// 主函数继续执行其他操作
	fmt.Println("主函数继续执行")
}
