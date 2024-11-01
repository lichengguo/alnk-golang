package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
[Go语言文件操作] -- 读取文件

文件是什么？
计算机中的文件是存储在外部介质（通常是磁盘）上的数据集合，文件分为文本文件和二进制文件

[打开和关闭文件]
os.Open()函数能够打开一个文件，返回一个*File和一个err。对得到的文件实例调用close()方法能够关闭文件

[读取文件的三种方法]
1.Read方法定义如下：
	func (f *File) Read(b []byte) (n int, err error)
    它接收一个字节切片，返回读取的字节数和可能的具体错误，读到文件末尾时会返回0和io.EOF

2.bufio读取文件

3.ioutil读取整个文件，如果文件很大，使用这种方法是否会导致内存飙升？
*/

// 1.第一种读取文件的方法 os.Open() file.Read()
func readFromFile1() {
	fileObj, err := os.Open("./main.go") //底层调用的其实就是OpenFile函数，只不过Open函数更简单
	//fileObj, err := os.OpenFile("./main.go", os.O_RDONLY, 0)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}

	// 记得关闭文件
	defer fileObj.Close()

	// 读文件
	//var tmp = make([]byte, 128) //指定读的长度 切片
	var tmp [128]byte // 数组

	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v\n", err)
			return
		}
		//fmt.Printf("读了%d个字节\n", n)
		fmt.Printf("%s", string(tmp[:n])) //为了保持文本输出的格式，这里建议用Printf

		if n < 128 {
			return
		}
	}
}

// 2.bufio
// bufio是在file的基础上封装了一层API，支持更多的功能
func readFromFilebyBufio() {
	// 1.打开文件
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}

	// 2.关闭文件
	defer fileObj.Close()

	// 3.创建一个用来从文件中读内容的对象
	reader := bufio.NewReader(fileObj)

	// 4.循环读取文件内容
	for {
		line, err := reader.ReadString('\n') //这里可能会产生bug，如果一行的结尾不是\n的话
		if err == io.EOF {
			//fmt.Println("文件读完了！！！")
			return
		}
		if err != nil {
			fmt.Printf("read line failed, err:%v\n", err)
			return
		}
		fmt.Print(line)
	}
}

// 3.ioutil读取整个文件
func readFromFileByIoutil() {
	// ret, err := ioutil.ReadFile("./main.go")
	ret, err := os.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read file failed, err:%v\n", err)
		return
	}

	fmt.Println(string(ret))
}

func main() {
	//readFromFile1()
	//readFromFilebyBufio()
	readFromFileByIoutil()
}
