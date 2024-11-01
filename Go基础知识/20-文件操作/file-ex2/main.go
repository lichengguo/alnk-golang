package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// 实现类似Linux中cat命令的简单功能

// cat
func cat(r *bufio.Reader) {
	for {
		// 这里可能会有bug 如果一行数据没有以 \n 结尾的话，那么读取不到
		//b, err := r.ReadBytes('\n') //注意是字符
		s, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		//fmt.Fprintf(os.Stdout, "%s", b)
		fmt.Fprintf(os.Stdout, s)
	}
}

func main() {
	// 1.1解析命令行参数
	flag.Parse()

	// 1.2依次读取命令行输入的文件名称
	// flag.NArg() 获取命令行输入的文件参数个数
	if flag.NArg() == 0 {
		fmt.Println("请输入文件名称!")
		return
	}
	for i := 0; i < flag.NArg(); i++ {
		//flag.Arg(i) 获取命令行的参数
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stdout, "reading from %s failed, err:%v\n", flag.Arg(i), err)
			return
		}

		// 1.3读取文件内容输出到终端
		cat(bufio.NewReader(f))
	}
}
