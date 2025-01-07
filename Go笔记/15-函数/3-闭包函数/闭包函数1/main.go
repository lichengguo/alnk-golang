package main

import "fmt"

// 闭包
// 闭包指的是一个函数和与其相关的引用环境组合而成的实体
// 简单来说 闭包=函数+引用环境
// 功能需求 把f2当做参数传入到f1  f1(f2)
// 解题思路 直接传递肯定不行 可以借助一个闭包函数

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

// 闭包函数
func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		f(x, y)
	}
	return tmp
}

func main() {
	ret := f3(f2, 100, 200) // 把原来需要传递两个int类型的参数 包装成一个不需要传递参数的函数
	fmt.Printf("%T\n", ret) // func()
	f1(ret) // this is f1 this is f2 300
}
