package main

import (
	"fmt"
	"time"
)

/*
[Go语言中的并发编程 goroutine]
博客地址: https://www.liwenzhou.com/posts/Go/14_concurrence/

[并发与并行]
并发：同一 时间段 内执行多个任务（你在用微信和两个女朋友聊天）
并行：同一 时刻   执行多个任务（你和你朋友都在用微信和女朋友聊天）

[goroutine]
Go语言的并发通过goroutine实现，goroutine类似于线程，属于用户态的线程，比内核态线程更轻量级，是由Go语言的运行时(runtime)调度的
Go程序会智能地将 goroutine 中的任务合理地分配给每个CPU。
Go语言之所以被称为现代化的编程语言，就是因为它在语言层面已经内置了调度和上下文切换的机制。
Go语言还提供channel在多个goroutine间进行通信。
Go语言中使用goroutine非常简单，只需要在调用函数的时候在前面加上go关键字，就可以为一个函数创建一个goroutine。
一个goroutine必定对应一个函数，可以创建多个goroutine去执行相同的函数

goroutine什么结束?
goroutine 对应的函数结束了，goroutine结束了。
main函数执行完了，由main函数创建的那些goroutine都结束了
*/

func hello(i int) {
	fmt.Println("hello ", i)
}

// 在程序启动时，Go程序就会为main()函数创建一个默认的goroutine
func main() {
	// 启动多个goroutine
	for i := 0; i < 10; i++ {
		go hello(i) //开启一个单独的goroutine去执行hello函数（任务）
	}
	fmt.Println("main")

	// 如果main函数结束，那么由main函数启动的goroutine也都结束了
	// 所以在这里等待1s，等其他的goroutine执行完毕在结束
	time.Sleep(time.Second)
}
