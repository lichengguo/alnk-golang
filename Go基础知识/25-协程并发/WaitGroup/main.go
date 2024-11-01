package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// WaitGroup goroutine的计数器

// 声明一个计数器变量
var wg sync.WaitGroup

func f() {
	rand.Seed(time.Now().UnixNano()) //保证每次执行的时候获取的随机数都不一样

	for i := 0; i < 5; i++ {
		r1 := rand.Int()    //int64
		r2 := rand.Intn(10) //0 <= r2 < 10
		fmt.Println(r1, r2)
	}
}

func f1(i int) {
	defer wg.Done()                                       //计数器 -1
	time.Sleep(time.Second * time.Duration(rand.Intn(3))) //睡 [0-3) 秒
	fmt.Println(i)
}

func main() {
	//f()

	//wg.Add(10) //也可以这样写，直接告诉计数器我要开启10个goroutine
	for i := 0; i < 10; i++ { //启动多个goroutine
		wg.Add(1) //每开启一个goroutine，计数器就自动+1
		go f1(i)
	}

	//time.Sleep(1 * time.Second) //现在还用time.sleep来等待 goroutine 结束就不好用了，不知道要等待多久 goroutine 才结束

	wg.Wait() //等待wg的计数器为0的时候就结束main函数
}
