package main

import (
	"fmt"
	"time"
)

/*
[通道同步]
创建一个用于通知任务完成的通道done，并在worker函数内部模拟了一个耗时的任务
在worker函数中，任务执行完毕后，向done通道发送了一个完成信号
在main函数中启动了一个goroutine执行任务，并在主线程中等待从通道中接收到完成信号
一旦接收到完成信号，打印出 work done
*/

func worker(done chan bool) {
	fmt.Println("任务开始...")
	time.Sleep(time.Second)
	fmt.Println("任务完成!")

	// 向通道发送完成信号
	done <- true
}

func main() {
	// 创建一个用于通知任务完成的通道
	done := make(chan bool)

	// 启动一个 goroutine 执行任务
	go worker(done)

	// 等待任务完成，阻塞直到从通道中接收到完成信号
	<-done
	fmt.Println("main work done")
}
