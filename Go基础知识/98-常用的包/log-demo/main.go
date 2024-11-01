package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// 内置log包

func f1() {
	// 1.打开文件
	f, err := os.OpenFile("./test1.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //追加写入
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	// 2.延迟关闭文件
	defer f.Close()

	// 3.设置log输出位置，这里设置为文件，默认会输出到终端
	log.SetOutput(f) //往文件输出

	// 4.测试循环写入内容到日志文件
	for {
		log.Println("这是一条测试的日志f1") //不会打印到终端，会写入到文件
		time.Sleep(time.Second * 3)
	}
}

func f2() {
	log.Println("这是一条很普通的日志") //2020/09/10 10:16:02 这是一条很普通的日志

	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v) //2020/09/10 10:16:02 这是一条很普通的日志。

	// Fatal系列函数会在写入日志信息后调用os.Exit(1)
	log.Fatalln("这是一条会触发fatal的日志") //2020/09/10 10:16:28 这是一条会触发fatal的日志

	// Panic系列函数会在写入日志信息后panic。
	log.Panic("这是一条会触发Panic的日志")
}

func main() {
	//f1()
	f2()
}
