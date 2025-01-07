package main

import (
	"fmt"
)

// 函数
// 函数存在的意义 函数能够让代码结构更加清晰 更简洁 能够让代码复用
// 函数是一段代码的封装
// 把一段逻辑抽象出来封装到一个函数中 给他起个名字 每次用它的时候直接用函数名调用即可

// 函数的基本定义

// func 函数名称(参数变量名称1 参数类型, 参数变量名称2 参数类型, ...) (返回值变量名称1 返回值类型, 返回值变量名称2 返回值类型, ...)
func sum(x int, y int) (ret int) {
	return x + y
}

// 函数的变种1:没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

// 函数的变种2：没有参数也没有返回值
func f2() {
	fmt.Println("f2")
}

// 函数的变种3：没有参数但是有返回值
func f3() int {
	ret := 3
	return ret
}

// 函数的变种4：函数返回值可以命名也可以不命名
// 命名的返回值就相当于在函数中声明一个变量
func f4(x int, y int) (ret int) {
	ret = x + y
	//return ret
	return //使用命名返回值可以省略return后面的变量名称
}

// 函数变种5：多个返回值
func f5() (int, string) {
	return 1, "abc"
}

// 函数变种6：多个参数的类型简写
func f6(x, y, z int, m, n string, i, j bool) int {
	fmt.Println(m, n, i, j)
	return x + y + z
}

// 函数变种7：可变长的参数
// 可变长的参数必须放在函数参数的最后
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y) // y的类型是 []int 切片
}

// 注意：Go语言中函数没有默认参数这个概念

func main() {
	f7("下雷了")
	f7("下雨了", 1, 2, 3, 4, 5)

	ret := sum(1, 1)
	fmt.Println(ret)
}
