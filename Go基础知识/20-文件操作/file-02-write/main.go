package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
[文件写入操作]
os.OpenFile()函数能够以指定模式打开文件，从而实现文件写入相关功能
func OpenFile(name string, flag int, perm FileMode) (*File, error) {}
其中：
	name：要打开的文件名
	flag：打开文件的模式
	perm：权限

模式有以下几种：
	os.O_WRONLY	只写
	os.O_CREATE	创建文件
	os.O_RDONLY	只读
	os.O_RDWR	读写
	os.O_TRUNC	清空
	os.O_APPEND	追加

Write和WriteString
	Write：写入字节切片数据
	WriteString：直接写入字符串数据
*/

// 1.os.OpenFile() 打开文件写内容
func writeDemo1() {
	// 1.打开文件
	// 只写，如果文件不存在则创建，如果存在则清空
	fileObj, err := os.OpenFile("./test1.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644) 
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	// 2.关闭文件
	defer fileObj.Close()

	// 3.write 方法写入
	n, err := fileObj.Write([]byte("zhoulin mengbi le 嘛!\n"))
	if err != nil {
		fmt.Printf("write file failed, err:%v\n", err)
		return
	}
	fmt.Println("n:", n) //写入了多少个字节

	// 4.writeString 方法写入
	n, err = fileObj.WriteString("周林解释不了!\n")
	if err != nil {
		fmt.Printf("write file failed, err:%v\n", err)
		return
	}
	fmt.Println("n:", n) //写入了多少个字节
}

// 2.bufio.NewWriter() 写入文件
func writeDemo2() {
	// 1.打开文件
	fileObj, err := os.OpenFile("./test2.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644) // 追加写入，不存在则创建
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}

	// 2.关闭文件
	defer fileObj.Close() //延迟关闭文件

	// 3.写入文件
	wr := bufio.NewWriter(fileObj) //创建一个写的对象
	wr.WriteString("hello沙河\n")    //写到缓存中 WriteString
	wr.Write([]byte("难受啊\n"))      //写到缓存中 Write
	wr.Flush()                     //将缓存中的内容写入文件
}

// 3.ioutil.WriteFile 写入文件
func writeDemo3() {
	str := "hello沙河\n"
	err := os.WriteFile("./test3.txt", []byte(str), 0644) // 这种写入方式会清空被写入文件之前的数据
	if err != nil {
		fmt.Printf("write file failed, err:%v\n", err)
		return
	}
}

func main() {
	writeDemo1()
	writeDemo2()
	writeDemo3()
}
