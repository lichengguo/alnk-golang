package main

import "runtime"

func task() {
	for {
	}
}

func main() {
	//runtime.GOMAXPROCS(1)  //如果不设置这个值,默认会跑满所有的cpu,见图1
	runtime.GOMAXPROCS(1) //设置只使用1个cpu,这里不会跑满所有的cpu
	go task()
	go task()
	go task()
	go task()
	select {}
}
