package main

import (
	"bufio"
	"fmt"
	"os"
)

// 获取用户输入时如果有空格

// Scanln 如果有空格会报错
func useScan() {
	var s string
	fmt.Print("请输入内容:")
	fmt.Scanln(&s)
	fmt.Printf("你输入的内容是:%s\n", s)
}

// useBufio 能处理有空格的输入
func useBufio() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("请输入内容:")
	s, _ = reader.ReadString('\n')
	fmt.Printf("你输入的内容是:%s", s)
}

// Fprintln 往终端输入、往文件输出
func useFprintln() {
	fmt.Fprintln(os.Stdout, "这是一条日志") //往终端屏幕写

	fileObj, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer fileObj.Close()

	fmt.Fprintln(fileObj, "这是一条日志记录!") //往文件写
}

func main() {
	//useScan()
	//useBufio()
	useFprintln()
}
