package main

import "fmt"

// 使用值接收者和指针接收者的区别？
// 使用值接收者实现接口,结构体类型和结构体指针类型的变量都能存
// 使用指针接收者实现接口,只能存结构体指针类型的变量

/*
使用场景：
当需要在方法内部修改接收者的属性时，应该使用指针类型的接收者
当不需要在方法内部修改接收者的属性，并且希望在方法内部复制接收者时，可以使用值类型的接收者
一般来说，推荐在方法内部修改接收者的属性时使用指针类型的接收者
这样可以避免对接收者值的复制，提高性能并确保方法对结构体的修改能够直接影响到原始对象
*/

// Mover接口
type Mover interface {
	move()
}

// dog结构体
type dog struct{}

// 值接收者实现接口
func (d dog) move() {
	fmt.Println("狗会动")
}

// cat结构体
type cat struct{}

// 指针接收者实现接口
func (c *cat) move() {
	fmt.Println("猫会动")
}

// 实现接口的是dog类型
func main() {
	var x Mover
	var wangcai = dog{} // 旺财是dog类型
	x = wangcai         // x可以接收dog类型
	x.move()

	var fugui = &dog{} // 富贵是*dog类型
	x = fugui          // x可以接收*dog类型
	x.move()
	/*
		从上面的代码中我们可以发现，使用值接收者实现接口之后，
		不管是 dog结构体 还是 结构体指针*dog类型 的变量都可以赋值给该接口变量。
		因为Go语言中有对指针类型变量求值的语法糖，
		dog指针fugui内部会自动求值*fugui
	*/

	// 使用指针接收者实现接口,只能存结构体指针类型的变量
	var c Mover
	var mimi = &cat{} // 咪咪是cat类型
	c = mimi          // c可以接收*cat类型
	c.move()          // 猫会动
}
