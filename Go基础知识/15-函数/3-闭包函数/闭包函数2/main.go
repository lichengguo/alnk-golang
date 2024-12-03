package main

import "fmt"

// 闭包是什么
// 闭包是一个函数 这个函数包含了它外部作用域的一个变量

// 底层原理
// 1 函数可以作为返回值
// 2 函数内部查找变量的顺序 先在自己内部找 找不到就往外层找

func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	// 变量ret是一个函数 并且它引用了其外部作用域中的x变量 此时ret就是一个闭包
	// 在ret的生命周期内 变量x也一直有效
	var ret = adder2(100)
	fmt.Println(ret(200)) // 300
	fmt.Println(ret(300)) // 600
}
