package main

import "fmt"

// calc接收一个基础数值 返回两个函数
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

func main() {
	f1, f2 := calc(10)

	// 变量f1和f2是一个函数 并且它引用了其外部作用域中的base变量 此时f1和f2就是一个闭包
	// 在f1和f2的生命周期内 变量base也一直有效
	fmt.Println(f1(1), f2(2)) // 11 9 此时base=9
	fmt.Println(f1(3), f2(4)) // 12 8 此时base=8
	fmt.Println(f1(5), f2(6)) // 13 7
}
