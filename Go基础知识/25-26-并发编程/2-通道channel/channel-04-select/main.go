package main

import (
	"fmt"
	"time"
)

/*
[select多路复用 通道选择器]
在某些场景下需要同时从多个通道接收数据
通道在接收数据时，如果没有数据接收可能将会发生阻塞
Go内置了select关键字，可以同时响应多个通道的操作
将协程、通道和选择器结合，是Go的一个强大特性

select的使用类似于switch语句，它有一系列case分支和一个默认的分支
每个case会对应一个通道的通信（接收或发送）过程
select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句

使用select语句能提高代码的可读性
可处理一个或多个channel的发送/接收操作
如果多个case同时满足，select会随机选择一个
对于没有case的select{}会一直等待，可用于阻塞main函数
*/

func f1() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println("x:", x)
		case ch <- i:
		}
	}
	/*
		结果:0 2 4 6 8
		分析:
		第一次i=0时候，通道可以放进数据，所以走case x := <-ch:
		第二次i=1时候，通道不可以放进数据，所以走case ch <- i:
		依次类推
	*/
}

func main() {
	// f1()

	// 创建两个通道，用于模拟两个并发的任务
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 启动第一个并发任务
	go func() {
		time.Sleep(2 * time.Second) // 模拟任务执行时间
		ch1 <- "任务1完成"              // 将任务1的结果发送到通道ch1
	}()

	// 启动第二个并发任务
	go func() {
		time.Sleep(1 * time.Second) // 模拟任务执行时间
		ch2 <- "任务2完成"              // 将任务2的结果发送到通道ch2
	}()

	// 记录当前时间
	now := time.Now()

	// 使用通道选择器选择首先到达的任务
	for i := 0; i < 2; i++ {
		select {
		case result := <-ch1: // 从通道ch1接收数据
			fmt.Println("ch1接收的结果:", result)
		case result := <-ch2: // 从通道ch2接收数据
			fmt.Println("ch2接收的结果:", result)
		}
	}

	// 主函数继续执行其他操作
	useTime := time.Since(now)
	fmt.Println("程序耗时时间:", useTime)
}
