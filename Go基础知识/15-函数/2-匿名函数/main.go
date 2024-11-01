package main

import (
	"fmt"
)

/*
匿名函数：匿名函数就是没有函数名的函数

函数当然还可以作为返回值，但是在Go语言中函数内部不能再像之前那样定义函数了，只能定义匿名函数。

func(参数)(返回值){
    函数体
}

匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量或者作为立即执行函数

匿名函数多用于实现回调函数和闭包
*/

var f2 = func(x, y int) {
	fmt.Println(x + y)
}

func main() {
	// 函数内部没有办法声明带名字的函数，但是可以声明匿名函数
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(10, 20)

	//声明匿名函数，并且直接调用
	func(x, y int) {
		fmt.Println(x + y)
		fmt.Println("hello world")
	}(100, 200)

	f2(1, 2)
}
