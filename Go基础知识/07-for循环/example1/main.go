package main

import (
	"fmt"
	"unicode"
)

// 1. 编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型
// 2. 编写代码统计出字符串`"hello中国小王子"`中汉字的数量

func main() {

	// i1 := 10
	// f1 := 1.234
	// b1 := true
	// s1 := "hello中国"
	// fmt.Printf("%T %d\n", i1, i1)
	// fmt.Printf("%T %f\n", f1, f1)
	// fmt.Printf("%T %v\n", b1, b1)
	// fmt.Printf("%T %s\n", s1, s1)

	s := "he中国小王llo子"
	result := chineseCount(s)
	fmt.Println(result) // 5

	multiplicationTable()

}

// chineseCount 统计字符串中的汉字
func chineseCount(str1 string) int64 {
	var count int64
	for _, char := range str1 {
		if unicode.Is(unicode.Han, char) {
			count++
		}
	}
	return count
}

// multiplicationTable 九九乘法表
func multiplicationTable() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, j*i)
		}
		fmt.Println()
	}
}
