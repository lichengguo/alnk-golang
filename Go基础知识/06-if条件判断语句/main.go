package main

import "fmt"

// if判断

func main() {
	age := 18
	if age > 18 {
		fmt.Println("aaa")
	} else {
		fmt.Println("bbb")
	}

	// 多个判断条件
	age = 19
	if age > 35 {
		fmt.Println("a")
	} else if age > 18 {
		fmt.Println("b")
	} else {
		fmt.Println("c")
	}

	// 作用域
	// age变量此时只在if条件判断语句中生效
	if age := 19; age > 18 {
		fmt.Println("a")
	} else {
		fmt.Println("b")
	}
}
