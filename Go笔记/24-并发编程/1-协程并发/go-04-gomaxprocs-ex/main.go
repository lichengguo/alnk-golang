package main

import "runtime"

func task() {
	for {
	}
}

func main() {
	// 设置只使用1个cpu 这里不会跑满所有的cpu 如图2.png
	// 如果不设置这个值 默认会跑满所有的cpu 如图1.png
	runtime.GOMAXPROCS(1)

	go task()
	go task()
	go task()
	go task()
	go task()
	go task()
	go task()
	go task()
	select {}
}
