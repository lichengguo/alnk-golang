package main

import (
	"fmt"
	"time"
)

// 超时处理

func main() {
	now := time.Now()

	// 创建一个通道，用于接收结果
	ch := make(chan string)

	// 启动一个协程，执行任务
	go func() {
		// 模拟一个耗时的任务，这里设置了3秒
		time.Sleep(3 * time.Second)
		// 将任务结果发送到通道
		ch <- "任务完成"
	}()

	// 使用select语句等待任务结果，设置超时时间为2秒
	select {
	case res := <-ch: // 如果在2秒内收到结果，则进入该分支
		fmt.Println(res)
	case <-time.After(2 * time.Second): // 如果超过2秒还没有收到结果，则进入该分支
		fmt.Println("任务超时")
	}

	// 主函数继续执行其他操作
	fmt.Println("程序运行时间:", time.Since(now))
	fmt.Println("main函数结束")
}
