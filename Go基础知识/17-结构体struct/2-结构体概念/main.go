package main

import "fmt"

// 结构体
// Go语言中的基础数据类型可以表示一些事物的基本属性，但是当我们想表达一个事物的全部或部分属性时
// 这时候再用单一的基本数据类型明显就无法满足需求了
// Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，这种数据类型叫结构体，英文名称struct

// Go语言内置的基础数据类型是用来描述一个值的，而结构体是用来描述一组值的

// Go语言中没有类的概念，也不支持类的继承等面向对象的概念
// Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性

// 结构体实例化
// 只有当结构体实例化时，才会真正地分配内存.也就是必须实例化后才能使用结构体的字段
// 结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型

// person 定义一个结构体
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	// 结构体实例化
	// 声明一个person类型的变量p
	var p person
	// 通过字段赋值
	// 示例1
	p.name = "tom"
	p.age = 9000
	p.gender = "男"
	p.hobby = []string{"篮球", "足球", "双色球"}
	fmt.Println(p)         // {tom 9000 男 [篮球 足球 双色球]}
	fmt.Printf("%#v\n", p) // main.person{name:"tom", age:9000, gender:"男", hobby:[]string{"篮球", "足球", "双色球"}}
	// 访问变量p的字段
	fmt.Printf("%T\n", p) // main.person
	fmt.Println(p.name)   // tom

	// 示例2
	var p2 person
	p2.name = "理想"
	p2.age = 18
	fmt.Printf("type:%T -- value:%v\n", p2, p2) // type:main.person（表示main包下的person类型) -- value:{理想 18  []}

	// 匿名结构体：多用于临时场景
	var s struct {
		x string
		y int
	}
	s.x = "heiheihei"
	s.y = 100
	fmt.Printf("type:%T value:%v\n", s, s) //type:struct { x string; y int }     value:{heiheihei 100}
}
