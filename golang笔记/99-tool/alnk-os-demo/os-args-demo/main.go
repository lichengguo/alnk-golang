package main

import (
	"fmt"
	"os"
)

// os.Args
// 简单的获取命令行参数，使用os.Args来获取
// os.Args是一个[]string
// os.Args是一个存储命令行参数的字符串切片，它的第一个元素是执行文件的名称本身

func main() {
	for index, arg := range os.Args {
		fmt.Printf("args[%d]=%v\n", index, arg)
	}

}

/*
% ./02os.Args_demo args1 args2
args[0]=./02os.Args_demo
args[1]=args1
args[2]=args2
*/
