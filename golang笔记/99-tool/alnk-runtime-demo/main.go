package main

import (
	"fmt"
	"path"
	"runtime"
)

// runtime.Caller(): 获取调用runtime.Caller所在函数的一些信息

func f(skip int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}

	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName) //函数名

	fmt.Println(file) //文件名 全路径

	fmt.Println(line) //行号

	fmt.Println(path.Base(file)) //文件名，不是全路径
}

func f1(skip int) {
	fmt.Println(skip)
	f(1)
}

func main() {
	f(0)  // 注意这里的参数，0表示f这个函数本身调用
	f(1)  // 注意这里的参数 1表示函数main调用了f函数
	f1(2) // 2 表示嵌套了2层函数才调用到f函数
}
