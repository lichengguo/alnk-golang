package main

import (
	"fmt"
	"sync"
	"time"
)

/*
worker pool（goroutine池）
在工作中我们通常会使用可以指定启动的goroutine数量–worker pool模式，控制goroutine的数量，防止goroutine泄漏和暴涨
*/

var wg sync.WaitGroup

// worker 处理任务的goroutine
func worker(id int, jobs <-chan int, result chan<- int) {
	defer wg.Done()
	for j := range jobs {
		//fmt.Printf("goroutine:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("goroutine:%d end job:%d\n", id, j)
		result <- j * 2
	}

}

func main() {
	// 1.声明2个通道并初始化
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	// 2.开启3个goroutine 模拟goroutine池
	wg.Add(3)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, result)
	}

	// 3.往jobs通道写入5个任务内容 模拟任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) //关闭jobs通道

	// 4.等待goroutine结束
	wg.Wait()

	// 5.从通道中取值
	for a := 1; a <= 5; a++ {
		fmt.Println(<-result)
	}
}
