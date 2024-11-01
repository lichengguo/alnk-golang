package main

import "fmt"

/*
[类型断言]
空接口可以存储任意类型的值，那我们如何获取其存储的具体数据呢

一个接口的值（简称接口值）是由一个 具体类型 和 具体类型的值 两部分组成的
这两部分分别称为 接口的动态类型 和 接口的动态值

想要判断空接口中的值这个时候就可以使用类型断言，其语法格式：
x.(T)
其中：
	x：表示类型为interface{}的变量
	T：表示断言x可能是的类型
该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败

因为空接口可以存储任意类型值的特点，所以空接口在Go语言中的使用十分广泛

重点：关于接口需要注意的是，只有当 [有两个或两个以上的具体类型] 必须以相同的方式进行处理时才需要定义接口。
不要为了接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时损耗
*/

// 类型断言1
func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println("传进来的是一个字符串: ", str)
	}
}

// 类型断言2
// x.(type)
// x表示类型为interface{}的变量
// type获取x变量的动态类型
func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("这是一个字符串:", t)
	case int:
		fmt.Println("这是一个int:", t)
	case int64:
		fmt.Println("这是一个int64:", t)
	case bool:
		fmt.Println("这是一个bool:", t)
	case []int:
		fmt.Println("这是一个slice:", t)
	case map[string]int:
		fmt.Println("是一个map[string]int:", t)
	case func():
		fmt.Printf("这是一个函数类型: %T", t)
	}
}

func main() {
	assign("100")
	assign2(true)
	assign2("啊哈哈")
	assign2(int64(200))
	assign2([]int{1, 2, 3})
	assign2(map[string]int{"a": 1})
	assign2(func() {})
}
