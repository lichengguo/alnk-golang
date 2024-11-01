package main

import "fmt"

// 1.函数类型的变量
// 定义函数类型
type calculation func(int, int) int //语句定义了一个calculation类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值

// 满足这个条件的函数都是calculation类型的函数
func add(x, y int) int {
	return x + y
}

func f1() {
	fmt.Println("shahe")
}

func f2() int {
	return 10
}

func f4(x, y int) int {
	return x + y
}

// 函数也可以作为参数的类型
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

// 函数还可以作为返回值
func ff(a, b int) int {
	return a + b
}

func f5(_ func() int) func(int, int) int {
	return ff
}

func main() {
	// 1.函数类型的变量
	var c calculation
	c = add
	ret := c(1, 2)
	fmt.Println("ret:", ret) //ret: 3
	a := f1
	fmt.Printf("%T\n", a) //func()
	b := f2
	fmt.Printf("%T\n", b) //func() int

	// 2.函数作为参数
	f3(f2) //10
	f3(b)  //10

	// f3(f4) //这里f4不能作为参数传给f3，因为f4的类型为func(int, int) int，而f3可以接受的函数类型为func() int
	fmt.Printf("%T\n", f4) //func(int, int) int

	// 3.函数还可以作为返回值
	f7 := f5(f2)           //ff函数
	fmt.Printf("%T\n", f7) //func(int, int) int
}
