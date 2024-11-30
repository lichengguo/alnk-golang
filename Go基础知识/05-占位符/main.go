package main

import (
	"fmt"
)

// 占位符

func main() {
	var n = 100
	fmt.Printf("%T\n", n) // 查看类型
	fmt.Printf("%v\n", n) // 查看值，%v:可以接受所有类型的值
	fmt.Printf("%b\n", n) // 二进制
	fmt.Printf("%d\n", n) // 整型
	fmt.Printf("%o\n", n) // 八进制
	fmt.Printf("%x\n", n) // 十六进制

	var s = "Hello 中国"
	fmt.Printf("字符串：%s\n", s)  // 字符串：Hello 中国
	fmt.Printf("字符串：%v\n", s)  // 字符串：Hello 中国
	fmt.Printf("字符串：%#v\n", s) // %#v按照类型显示,比如字符串的话打印的时候会有引号,字符串："Hello 中国"
}
