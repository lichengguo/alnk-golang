package main

import "fmt"

// defer
// defer多用于函数结束之前释放资源 文件句柄 数据库连接 socket连接

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("aaaaaaa") // defer把它后面的语句延迟到函数即将返回的时候在执行
	defer fmt.Println("bbbbbbb") // 一个函数中可以有多个defer语句
	defer fmt.Println("ccccccc") // 多个defer语句按照先进后出的顺序延迟执行
	fmt.Println("end")
}

// 帮助理解
// Go解释器从上往下执行
// 1 定义f1这个函数 返回值是int类型的x变量 默认值为0
// 2 fmt.Println(f1()) 先执行括号里的f1()函数 调用f1()这个函数
// 3 定义 defer func(x int){}(x) 并调用 但是由于defer特性 此时没有执行函数体里面的语句 目前x=0 x++ 是1
// 4 执行return语句 把5赋值给要返回的变量x 此时x=5 执行defer 里面的x=1 并且打印执行RET返回
// 5 执行fmt.Println(5)语句打印
func f1() (x int) {
	defer func(x int) {
		x++
		fmt.Println(x) // 1
	}(x) // 0
	return 5
}

func main() {
	// deferDemo()
	// fmt.Println("------分割线----------")
	fmt.Println(f1())
}

// start
// end
// ccccccc
// bbbbbbb
// aaaaaaa
// ------分割线----------
// 1
// 5
