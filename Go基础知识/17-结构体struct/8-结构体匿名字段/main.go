package main

import "fmt"

// 匿名字段
// 字段比较少也比较简单的场景
// 不常用！！！

// 匿名字段结构体
type person struct {
	string
	int
}

func main() {
	p1 := person{
		"tom",
		9000,
	}
	fmt.Println(p1)        // {tom 9000}
	fmt.Println(p1.string) // tom
	fmt.Println(p1.int)    // 9000
}
