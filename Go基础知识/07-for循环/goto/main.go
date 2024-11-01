package main

import (
	"fmt"
)

// goto 跳转到指定标签 -不建议使用

func main() {
	// 跳出多层for循环，用标记位
	flag := false
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				flag = true
				break // 跳出内层的for循环
			}
			fmt.Printf("%v-%c\n", i, j)
		}
		if flag {
			break // 跳出外层的for循环
		}
	}

	// Go语言中默认使用utf8编码
	// 'a':这是字符，utf8编码底层是int32类型(又可以叫做rune类型)
	for k := 'a'; k <= 'z'; k++ {
		fmt.Printf("%c", k)
	}
	fmt.Println()

	// goto+label 实现跳出多层for循环（不建议使用）
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto XX // 跳到我指定的那个标签
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
XX: // label标签
	fmt.Println("over")
}
