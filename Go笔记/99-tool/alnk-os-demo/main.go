package main

import (
	"fmt"
	"os"
)

// 获取一个文件的基本信息 如文件大小、文件名称
// 获取文件基本信息

// f1 获取文件详细信息，例如文件大小、名字等
func f1() {
	// 1.打开文件
	fileObj, err := os.Open("main.go")
	if err != nil {
		fmt.Printf("open file failed. err:%s\n", err)
		return
	}
	fmt.Printf("%T\n", fileObj) //文件对象的类型 *os.File指针

	// 2.获取文件对象的详细信息
	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err:%s\n", err)
		return
	}
	fmt.Printf("文件大小是:[%d]B\n", fileInfo.Size()) //文件大小是:[601]B
	fmt.Printf("文件名称是:[%s]\n", fileInfo.Name())  //文件名称是:[main.go] 只会获取文件名称，不会获取文件路径
}

func main() {
	f1()
}
