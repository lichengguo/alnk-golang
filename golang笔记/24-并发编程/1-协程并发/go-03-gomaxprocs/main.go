package main

import (
	"fmt"
	"runtime"
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
	一个操作系统线程对应用户态多个goroutine
	go程序可以同时使用多个操作系统线程
	goroutine和OS线程是多对多的关系，即m:n
*/

func main() {
	// runtime.GOMAXPROCS(1) // 设置程序使用的cpu核数
	// runtime.GOMAXPROCS(2) // 设置程序使用的cpu核数
	fmt.Println(runtime.NumCPU()) // 获取本机物理机线程个数 8
}
