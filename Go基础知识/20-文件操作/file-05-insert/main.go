package main

import (
	"fmt"
	"io"
	"os"
)

// 往文件中插入内容

func f2() {
	// 1.打开要操作的文件
	f, err := os.OpenFile("./test.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer f.Close()

	// 2.因为没有办法直接在文件中插入内容，所以要借助一个临时文件
	tmpFile, err := os.OpenFile("./test.tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("create tmp file failed, err:%v\n", err)
		return
	}
	defer tmpFile.Close()

	// 3.读取源文件部分内容写入临时文件
	var ret [2]byte // 读取了2个字节
	n, err := f.Read(ret[:])
	if err != nil {
		fmt.Printf("read from file failed, err:%v\n", err)
		return
	}

	// 4.读取的部分源文件内容写入临时文件
	tmpFile.Write(ret[:n])

	// 5.写入要插入的内容到临时文件
	var s []byte
	s = []byte{'a', 'b', 'c'}
	tmpFile.Write(s)

	// 6.紧接着把源文件后续的所有内容写入临时文件
	var x [1024]byte
	for {
		n, err := f.Read(x[:])
		if err == io.EOF {
			tmpFile.Write(x[:n])
			break
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v\n", err)
			return
		}
		tmpFile.Write(x[:n])
	}

	// 7.把临时文件改名为源文件
	os.Rename("./test.tmp", "./test.txt")

}

func main() {
	f2()
}
