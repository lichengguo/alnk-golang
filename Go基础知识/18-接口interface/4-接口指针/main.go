package main

import "fmt"

// 使用值接收者和指针接收者的区别？
// 使用值接收者实现接口,结构体类型和结构体指针类型的变量都能存
// 使用指针接收者实现接口,只能存结构体指针类型的变量

// Mover接口和一个dog结构体
type Mover interface {
	move()
}

type dog struct{}

//// 值接收者实现接口
//func (d dog) move() {
//	fmt.Println("狗会动")
//}
//// 实现接口的是dog类型
//func main() {
//	var x Mover
//	var wangcai = dog{} // 旺财是dog类型
//	x = wangcai         // x可以接收dog类型
//	x.move()
//	var fugui = &dog{} // 富贵是*dog类型
//	x = fugui          // x可以接收*dog类型
//	x.move()
//}
//
///*
//从上面的代码中我们可以发现，使用值接收者实现接口之后，
//不管是 dog结构体 还是 结构体指针*dog类型 的变量都可以赋值给该接口变量。
//因为Go语言中有对指针类型变量求值的语法糖，
//dog指针fugui内部会自动求值*fugui
//*/

// 指针接收者实现接口
func (d *dog) move() {
	fmt.Println("狗会动")
}

func main() {
	var x Mover
	//var wangcai = dog{} // 旺财是dog类型
	//x = wangcai         // x不可以接收dog类型
	var fugui = &dog{} // 富贵是*dog类型
	x = fugui          // x可以接收*dog类型
	x.move()
}

// 此时实现Mover接口的是*dog类型，所以不能给x传入dog类型的wangcai，此时x只能存储*dog类型的值fugui
