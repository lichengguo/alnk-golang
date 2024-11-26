package main

import (
	"fmt"
)

// 常量
// 定义了常量以后不能修改,常量在定义的时候必须赋值
// 在程序运行期间不会改变的量
const pi = 3.1415926

// 批量声明常量
const (
	statusOk = 200
	notFoud  = 404
)

// 批量声明常量时，如果某一行声明后没有赋值，默认就和上一行一致
const (
	n1 = 100
	n2
	n3
)

// iota
// iota是go语言的常量计数器，只能在常量的表达式中使用
// iota在const关键字出现时将被重置为0, const中每 [新增一行] 常量声明将使iota计数一次(+1)
// 使用iota能简化定义，在定义枚举时很有用
const (
	a1 = iota // 0
	a2        // 1
	a3        // 2
)

const (
	b1 = iota // 0
	b2 = iota // 1
	_  = iota // 2
	b3 = iota // 3
)

// 插队
const (
	c1 = iota // 0   iota=0
	c2 = 100  // 100 iota=1
	c3 = iota // 2 iota=2
	c4        // 3 iota=3
)

// 多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 // d1:1 d2:2
	d3, d4 = iota + 1, iota + 2 // d3:2 d4:3
)

// 定义量级
const (
	_  = iota
	KB = 1 << (10 * iota) // 1左移10位，二进制10000000000 转化为十进制1024
	MB = 1 << (10 * iota) // 1左移20位 100000000000000000000
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	// pi = 123 // 常量不能改变

	fmt.Println("--- 批量声明常量 ---")
	fmt.Println("statusOk: ", statusOk) // 200
	fmt.Println("notFoud: ", notFoud)   // 400
	fmt.Println("n1:", n1)              // 100
	fmt.Println("n2:", n2)              // 100
	fmt.Println("n3:", n3)              // 100

	fmt.Println("--- iota ---")
	fmt.Println("a1:", a1) // 0
	fmt.Println("a2:", a2) // 1
	fmt.Println("a3:", a3) // 2
	fmt.Println("b1:", b1) // 0
	fmt.Println("b2:", b2) // 1
	fmt.Println("b3:", b3) // 3

	fmt.Println("--- iota插队 ---")
	fmt.Println("c1:", c1) // c1: 0
	fmt.Println("c2:", c2) // c2: 100
	fmt.Println("c3:", c3) // c3: 2
	fmt.Println("c4:", c4) // c4: 3

	fmt.Println("-- 多个常量声明在一行 ---")
	fmt.Println("d1:", d1) // 1
	fmt.Println("d2:", d2) // 2
	fmt.Println("d3:", d3) // 2
	fmt.Println("d4:", d4) // 3

	fmt.Println("--- 定义量级 ---")
	fmt.Println("KB: ", KB) // 1024
	fmt.Println("MB: ", MB) // 1048576 1024*1024
	fmt.Println("GB: ", GB) // 1073741824
	fmt.Println("TB: ", TB) // 1099511627776
	fmt.Println("PB: ", PB) // 1125899906842624
}
