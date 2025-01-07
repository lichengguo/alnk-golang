package main

import "fmt"

// 定义支付接口
type payer interface {
	pay()
}

// 定义支付函数
func pay(p payer) {
	p.pay()
}

// 定义结构体
type weixin struct {
	name string
}

// 定义结构体方法 实现接口中的方法
func (w weixin) pay() {
	fmt.Printf("尊敬的[%s], 欢迎使用微信支付~\n", w.name)
}

// 定义结构体
type alipay struct {
	name string
}

// 定义结构体方法 实现接口中的方法
func (a alipay) pay() {
	fmt.Printf("欢迎[%s], 使用阿里支付~\n", a.name)
}

func main() {
	var w weixin
	var a alipay
	w.name = "Alnk"
	a.name = "tom"
	pay(w)
	pay(a)

	// 实现接口以后就不用每次都这么写重复的代码了
	// p1 := weixin{"Alnk"}
	// p1.pay()
	// p2 := alipay{"tom"}
	// p2.pay()
}
