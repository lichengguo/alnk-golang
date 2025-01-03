package main

import "fmt"

func main() {
	// 编译后再执行 不然找不到Add函数
	result := Add(1, 2)
	fmt.Println(result)
}
