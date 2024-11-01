package main

import (
	"fmt"
	"sync"
	"time"
)

// 如何优雅的控制子goroutine退出？
// 全局变量方式

// 存在的问题：
// 1.使用全局变量在跨包调用时不容易统一
// 2.如果f中再启动goroutine，就不太好控制了
var wg sync.WaitGroup
var notify bool // 标志位，控制子goroutine退出

func f() {
	defer wg.Done()

	for {
		fmt.Println("fffffffff")
		time.Sleep(time.Second)
		if notify {
			break
		}
	}
}

func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5) //sleep 5秒以免程序过快退出
	// 如何通知子goroutine退出
	notify = true //修改全局变量实现子goroutine的退出
	wg.Wait()
}
