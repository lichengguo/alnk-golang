package main

import (
	"fmt"
)

// 结构体是值类型
// 值类型和引用类型
// 值类型:int、float、bool、string、数组、struct
// 引用类型:指针、slice、map、chan、接口

// 定义一个结构体
type person struct {
	name, Gender string
}

// go语言中 函数传参数 永远传的是拷贝
// 切片也是拷贝切片传递过去，但是传过去的拷贝的切片和原来的切片指向同一个底层数组，切片是引用类型
// 如果出现append()操作，可能会导致底层数组扩容，那么就会指向不同的底层数组了

func f(x person) {
	x.Gender = "女" //结构体是值类型，相当于把值拷贝了一份传递过来 修改的是副本的Gender
	fmt.Println("==", x)
}

func f2(x *person) {
	//(*x).Gender = "女" //根据内存地址找到那个变量，修改的就是原来的变量
	x.Gender = "女" //语法糖，自动根据指针找对应的变量
}

func main() {
	var p person
	p.name = "周林"
	p.Gender = "男"
	f(p)
	fmt.Println(p.Gender) //男
	f2(&p)                //ox1241ac3 指针 相当于拷贝了p这个变量的指针传递到了f2的函数
	fmt.Println(p.Gender) //女

	// 1.结构体指针
	var p2 = new(person) //new函数返回的是传递进去的类型的指针
	(*p2).name = "理想"
	p2.Gender = "保密"
	fmt.Printf("%T\n", p2)   //*main.person
	fmt.Printf("%p\n", p2)   //0xc00000c080 p2保存的值就是一个内存地址
	fmt.Printf("%p\n", &p2)  //0xc00000e030 取p2这个变量本身的内存地址
	fmt.Printf("%#v\n", p2)  //&main.person{name:"理想", Gender:"保密"}
	fmt.Printf("%#v\n", *p2) //main.person{name:"理想", Gender:"保密"}
	fmt.Println(p2)          //&{理想 保密}
	fmt.Println(*p2)         //{理想 保密}

	// 2.结构体指针初始化
	// 2.1 key-value初始化
	var p3 = &person{
		name: "元帅",
	}
	fmt.Printf("%#v\n", p3) //&main.person{name:"元帅", Gender:""}
	// 2.2 使用值列表的形式初始化，值的顺序和结构体定义时字段的顺序一致
	p4 := &person{
		"小王子",
		"男",
	}
	fmt.Printf("%#v\n", p4) //&main.person{name:"小王子", Gender:"男"}
}
