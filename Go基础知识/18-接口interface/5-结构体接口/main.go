package main

import "fmt"

// 一个类型可以实现多个接口
// 同一个结构体可以实现多个接口
// 接口还可以嵌套

// animal接口
type animal interface {
	mover // 嵌套mover
	eater // 嵌套eater
}

// mover接口
type mover interface {
	move()
}

// eater接口
type eater interface {
	eat(string)
}

// cat结构体
type cat struct {
	name string
	feet int8
}

// cat实现了mover接口
func (c *cat) move() {
	fmt.Println("走猫步...")
}

// cat实现了eater接口
func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

func main() {
	c1 := cat{
		name: "tom",
		feet: 4,
	}

	var a1 animal // 定义一个animal类型的变量

	a1 = &c1      // 将cat赋值给animal类型的变量
	a1.move()     // 走猫步...
	a1.eat("小黄鱼") // 猫吃小黄鱼...

	var m1 mover
	m1 = &c1
	m1.move() // 走猫步...

	var e1 eater
	e1 = &c1
	e1.eat("鱼仔") // 猫吃鱼仔...
}
