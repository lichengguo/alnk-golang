package main

import "fmt"

// 接口的实现

// 定义一个动物接口
type animal interface {
	move()
	eat(string)
}

// 定义一个猫的结构体
type cat struct {
	name string
	feet int8
}

// 猫结构体方法
func (c cat) move() {
	fmt.Println("走猫步...")
}

// 猫结构体方法
func (c cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

// 定义一个鸡的结构体
type chicken struct {
	feet int8
}

// 鸡结构体方法
func (c chicken) move() {
	fmt.Println("鸡动!")
}

// 鸡结构体方法
func (c chicken) eat(food string) {
	fmt.Printf("吃%s...\n", food)
}

func main() {
	var a1 animal          //定义一个接口类型的变量
	fmt.Printf("%T\n", a1) //<nil>

	// 定义一个cat类型的变量bc
	bc := cat{
		name: "淘气",
		feet: 4,
	}
	a1 = bc
	a1.eat("小黄鱼")          //猫吃小黄鱼...
	fmt.Printf("%T\n", a1) //main.cat
	fmt.Println(a1)        //{淘气 4}
	fmt.Println(bc)        //{淘气 4}

	kfc := chicken{feet: 2}
	a1 = kfc
	fmt.Printf("%T\n", a1) //main.chicken
	a1.eat("饲料")           //吃饲料...
}
