package main

import (
	"flag"
	"fmt"
	"os"
)

// 实现类似Linux中cat命令的简单功能

// cat
func cat(fileName string) {
	ret, err := os.ReadFile(fileName)
	if err != nil {
		panic("文件读取失败")
	}

	fmt.Println(string(ret))
}

func main() {
	// 1.解析命令行参数
	flag.Parse()

	// 2.依次读取命令行输入的文件名称
	// flag.NArg() 获取命令行输入的文件参数个数
	if flag.NArg() == 0 {
		fmt.Println("请输入文件名称!")
		return
	}
	for i := 0; i < flag.NArg(); i++ {
		cat(flag.Arg(i))
		fmt.Println()
	}
}
