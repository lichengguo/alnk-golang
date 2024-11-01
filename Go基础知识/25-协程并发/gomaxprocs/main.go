package main

import (
	"fmt"
	"sync"
)

/*
[GOMAXPROCS]
Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码
默认值是机器上的CPU核心数
例如在一个8核心的机器上，调度器会把Go代码同时调度到8个OS线程上

Go1.5版本之前，默认使用的是单核心执行
Go1.5版本之后，默认使用全部的CPU逻辑核心数

Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数

Go语言中的操作系统线程和goroutine的关系：
	一个操作系统线程对应用户态多个goroutine。
	go程序可以同时使用多个操作系统线程。
	goroutine和OS线程是多对多的关系，即m:n
*/

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		fmt.Printf("A:%d ", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		fmt.Printf("B:%d ", i)
	}
}

func main() {
	/*
		默认为CPU的逻辑核心数，跑满整个CPU，见 图pc2 的结果可以知道，是多个线程同时执行，所以打印的次序乱了
		由于只有一个终端输出，所以多个线程打印的时候，会争抢资源，打印次序就混乱

		当设置为1的时候，就只使用1个线程，此时程序是串行的，打印也是有顺序的，见 图pc1
	*/

	//runtime.GOMAXPROCS(1)
	//runtime.GOMAXPROCS(2)
	//fmt.Println(runtime.NumCPU()) //获取本机物理机线程个数 4

	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
