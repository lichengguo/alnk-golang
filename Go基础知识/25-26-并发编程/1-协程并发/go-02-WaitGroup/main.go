package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 想要等待多个协程完成
// 可以使用WaitGroup goroutine的计数器

// 声明一个计数器变量
var wg sync.WaitGroup

func f() {
	// rand.Seed(time.Now().UnixNano()) // 保证每次执行的时候获取的随机数都不一样[rand.Seed 1.20版本被弃用]
	rand.New(rand.NewSource(time.Now().UnixNano())) // 保证每次执行的时候获取的随机数都不一样

	for i := 0; i < 2; i++ {
		r1 := rand.Int()    // 产生的随机数范围 int64
		r2 := rand.Intn(10) // 产生的随机数范围 0 <= r2 < 10
		fmt.Println(r1, r2)
	}
}

func f1(i int) {
	fmt.Printf("任务[%d]开始\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("任务[%d]结束\n", i)
}

func main() {
	fmt.Println("-------- 随机数 -------------")
	f()

	fmt.Println("------- waitgroup ------------")
	for i := 0; i < 10; i++ { // 启动多个goroutine
		wg.Add(1) // 每开启一个goroutine，计数器就自动+1

		go func() {
			defer wg.Done() // 计数器-1 计数器尽量不要直接丢到f1()函数里面去，不好管理
			f1(i)
		}()
	}

	wg.Wait() // 等待wg的计数器为0的时候就结束main函数

	fmt.Println("main 所有任务完成")
}
