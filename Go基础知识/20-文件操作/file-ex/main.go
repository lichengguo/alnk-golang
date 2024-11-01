package main

import (
	"fmt"
	"io"
	"os"
)

// 借助io.copy() 实现一个拷贝文件函数

// CopyFile拷贝文件函数
func CopyFile(dstName, srcName string) (wrtten int64, err error) {
	// 以读方式打开
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", srcName, err)
		return
	}
	defer src.Close()

	// 以写|创建的方式打开目标文件
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644) //没有os.O_APPEND 会清空之前的文件
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", dstName, err)
		return
	}
	defer dst.Close()

	return io.Copy(dst, src) //调用io.Copy()拷贝内容
}

func main() {
	_, err := CopyFile("./dst.txt", "./src.txt")
	if err != nil {
		fmt.Println("copy file failed, err:", err)
		return
	}

	fmt.Println("copy done.")
}
