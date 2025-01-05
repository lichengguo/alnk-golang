package main

import (
	"fmt"
	"sync"
	"time"
)

// 如何优雅的控制子goroutine退出？
// 通道方式
// 管道方式存在的问题：
// 1. 使用全局变量在跨包调用时不容易实现规范和统一，需要维护一个共用的channel

var wg sync.WaitGroup
var exitChan = make(chan bool, 1) // 退出子goroutine标志位

func f() {
	defer wg.Done()

FORLOOP:
	for {
		fmt.Println("ffffff")
		time.Sleep(time.Second * 1)
		select {
		case <-exitChan: // 等待接收上级通知
			break FORLOOP
		default:
		}
	}
}

func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5) // sleep 5秒以免程序过快退出

	// 如何通知子goroutine退出
	exitChan <- true // 给子goroutine发送退出信号
	close(exitChan)

	wg.Wait()

	fmt.Println("over")
}
