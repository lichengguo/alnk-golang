package main

import "fmt"

// 自定义类型和类型别名的区别

// type后面跟的是类型
type myInt int    // 自定义类型
type youInt = int // 类型别名

func main() {
	var n myInt
	n = 100
	fmt.Println(n)              // 100
	fmt.Printf("%T %d\n", n, n) // main.myInt 100

	var m youInt
	m = 100
	fmt.Println(m)        // 100
	fmt.Printf("%T\n", m) // int

	var c rune
	c = '中'
	fmt.Println(c)        // 20013 unicode编码十进制
	fmt.Printf("%T\n", c) // int32

	c1 := '国'
	fmt.Println(c1)        // 22269 unicode编码十进制
	fmt.Printf("%T\n", c1) // int32
}
