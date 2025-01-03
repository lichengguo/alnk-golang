package main

import "fmt"

// 接口示例2
// 不管什么牌子的车都能跑

// 定义一个car接口类型
// 不管是什么结构体，只要有run方法都能是carer类型
type carer interface {
	run() // 只要实现了run方法的变量都是carer类型, 方法签名
}

// drive 定义一个函数
func drive(c carer) {
	c.run()
}

// falali 定义一个结构体
type falali struct {
	name string
}

// falali 的run方法
func (f falali) run() {
	fmt.Printf("%s的速度70\n", f.name)
}

// baoshijie 定义一个结构体
type baoshijie struct {
	name string
}

// baoshijie 的run方法
func (b baoshijie) run() {
	fmt.Printf("%s的速度100\n", b.name)
}

// su7
type su7 struct {
	name string
}

// su7 的run方法
func (s su7) run() {
	fmt.Printf("%s的速度200\n", s.name)
}

func main() {
	f := falali{"法拉利"}
	// f.run()
	drive(f)

	b := baoshijie{"保时捷"}
	drive(b)

	var c su7
	c.name = "su7"
	drive(c)
	// c.run()
}
