package main

import (
	"fmt"
	"log"
)

/*
Go语言标准库log介绍

log包定义了Logger类型，该类型提供了一些格式化输出的方法。
本包也提供了一个预定义的“标准”logger，可以通过调用函数Print系列(Print|Printf|Println）、
Fatal系列（Fatal|Fatalf|Fatalln）、和Panic系列（Panic|Panicf|Panicln）来使用，比自行创建一个logger对象更容易使用
*/

func example1() {
	v := "很普通的"
	log.Println("这是一条很普通的日志。")
	log.Println("这是一条很普通的日志。")
	log.Printf("这是一条%s日志。\n", v)

	//log.Fatalln("这是一条会触发Fatal的日志。")
	//log.Panicln("这是一条会触发Panic的日志。")

	fmt.Println("结束!")
	/*
		logger会打印每条日志信息的日期、时间，默认输出到系统的标准错误。
		Fatal系列函数会在写入日志信息后调用os.Exit(1)。
		Panic系列函数会在写入日志信息后panic。
	*/
}

func main() {
	example1()
}
