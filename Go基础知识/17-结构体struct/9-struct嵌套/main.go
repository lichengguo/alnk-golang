package main

import "fmt"

// 结构体嵌套

type address struct {
	province string
	city     string
}

type workPlace struct {
	province string
	city     string
}

type person struct {
	name    string
	age     int
	address //匿名嵌套结构体
	//address:address // 命名嵌套结构体
	workPlace
}

type company struct {
	name string
	address
}

func main() {
	p1 := person{
		name: "周林",
		age:  90,
		address: address{
			province: "山东",
			city:     "威海",
		},
		workPlace: workPlace{
			province: "山东",
			city:     "济南",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.address.city)

	// 先在自己结构体找这个字段，找不到就去匿名嵌套的结构体中查找该字段
	// 如果匿名嵌套的结构体中有多个city，会报错 ambiguous selector p1.city
	//fmt.Println(p1.city)

	// 为了防止报错，可以用下面这种方法
	fmt.Println(p1.address.city)
	fmt.Println(p1.workPlace.city)
}
