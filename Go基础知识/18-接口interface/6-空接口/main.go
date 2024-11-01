package main

import "fmt"

// 空接口
// 空接口是指没有定义任何方法的接口
// 因此任何类型都实现了空接口
// 空接口类型的变量可以存储任意类型的变量

// interface:关键字
// interface{}:空接口类型

// 空接口的应用:
// 1.空接口作为函数的参数
// 2.空接口作为map的值

// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T --- value:%v\n", a, a)
}

func main() {
	// 空接口作为map的值
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "周林"
	m1["age"] = 9000
	m1["merried"] = true
	m1["hobby"] = [...]string{"唱", "跳", "rap"}
	fmt.Printf("%#v\n", m1) //map[string]interface {}{"age":9000, "hobby":[3]string{"唱", "跳", "rap"}, "merried":true, "name":"周林"}

	// 空接口作为函数参数
	show(false) //type:bool --- value:false
	show(nil)   //type:<nil> --- value:<nil>
	show(m1)    //type:map[string]interface {} --- value:map[age:9000 hobby:[唱 跳 rap] merried:true name:周林]

	// 空接口作为map的键和值
	var m2 map[interface{}]interface{}
	m2 = make(map[interface{}]interface{}, 5)
	m2[1] = "a"
	m2["b"] = true
	m2[true] = 1
	fmt.Printf("%#v\n", m2) //map[interface {}]interface {}{true:1, 1:"a", "b":true}
	show(m2)                //type:map[interface {}]interface {} --- value:map[true:1 1:a b:true]
}
