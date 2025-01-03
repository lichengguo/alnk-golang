package main

import (
	"fmt"
	"sync"
)

// channel 练习
// 1.启动一个goroutine，生成100个数发送到ch1
// 2.启动一个goroutine，从ch1中取值，计算其平方放到ch2中
// 3.在main中，从ch2取值打印出来

var wg sync.WaitGroup
var once sync.Once

// 生成100个数发送到ch1
func f1(ch1 chan<- int) { // ch1 chan<- int  只能往ch1里面写入
	for i := 1; i <= 100; i++ {
		ch1 <- i
	}
	close(ch1) // 关闭通道以后，还可以读取数据，但是不能写入了
}

// 计算其平方放到ch2中
func f2(ch1 <-chan int, ch2 chan<- int) { // ch1 <-chan int 只能从ch1里面读取, ch2 chan<- int 只能往ch2里面写入
	for {
		// 没有关闭的通道，最后不会返回false，会一直hang住，然后导致死锁
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	once.Do(func() { close(ch2) }) // 如果f2使用多个goroutine加快计算，那么确保关闭通道的操作只执行一次
	// close(ch2)
}

func main() {
	// 创建两个通道
	a := make(chan int, 100)
	b := make(chan int, 100)

	// 启动goroutine,执行任务
	wg.Add(4) // 4个goroutine

	go func() { // 生成100个数发送到ch1
		defer wg.Done()
		f1(a)
	}()

	for i := 0; i < 3; i++ { // 启动3个goroutine，从ch1中取值，计算其平方放到ch2中
		go func() {
			defer wg.Done()
			f2(a, b)
		}()
	}

	wg.Wait() // 等待所有goroutine执行完毕

	// 用range读取通道的时候，需要关闭通道，不然会出现死锁
	for ret := range b {
		fmt.Println(ret)
	}
}
