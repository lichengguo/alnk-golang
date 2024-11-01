package main

import "fmt"

// Go语言中的每一个变量都有自己的类型，并且变量必须经过声明才能开始使用
// 同一作用域内不支持重复声明,并且Go语言的变量声明后必须使用(全局声明的变量可以不必必须使用)
// Go语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作
// 每个变量会被初始化成其类型的默认值
// 例如：
// 整型和浮点型变量的默认值为0
// 字符串变量的默认值为空字符串
// 布尔型变量默认为false
// 切片、函数、指针变量的默认为nil

// Go语言中推荐使用驼峰式命名
var studentName string

// 声明变量(先声明，后使用)
// var name string
// var age int
// var isOk bool

// 批量声明变量
var (
	name string // ""
	age  int    // 0
	isOk bool   // false
)

func main() {
	name = "Alnk"
	age = 16
	isOk = true
	// var heiheihei string
	// heiheihei = "嘿嘿嘿"
	// Go语言中非全局变量声明后必须使用，不使用就编译不过去
	fmt.Print(isOk)               // 在终端打印内容，不会打印换行符
	fmt.Printf("name:%s\n", name) // %s：占位符，使用name这个变量的值去替换占位符
	fmt.Println(age)              // 打印以后，还会在后面追加一个换行符

	// 声明变量的同时赋值
	var s1 string = "tom"
	fmt.Println(s1)

	// 一次声明多个变量，并且初始化多个变量
	var name1, age1 = "Alnk", 20
	fmt.Println(name1, age1)

	// 类型推导，根据值判断该变量是什么类型
	var s2 = "20"
	fmt.Println(s2)

	// 简短变量声明，只能在函数内部使用
	s3 := "哈啊哈哈"
	fmt.Println(s3)

	// s1 := "10" //同一个作用域不能重复声明变量
	// 匿名变量是一个特殊的变量：_
}
