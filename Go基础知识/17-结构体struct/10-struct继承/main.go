package main

import "fmt"

// 结构体模拟实现其他语言中的继承

type animal struct {
	name string
}

// 函数 func 函数名(参数) (返回值) {...}
// 方法 func (变量名 结构体) 函数名() 返回值() {...}
// 给animal实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

// 狗类 嵌套了动物结构体
type dog struct {
	feet   uint8
	animal // animal拥有的方法，dog此时也有了
}

//给dog实现一个汪汪汪的方法
func (d dog) wang() {
	fmt.Printf("%s在叫: 汪汪汪~\n", d.name)
}

func main() {
	d1 := dog{
		feet: 4,
		animal: animal{
			name: "周林",
		},
	}
	fmt.Println(d1)
	d1.wang()
	d1.move()
}
