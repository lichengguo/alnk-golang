package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student) // 初始化map类型

	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	// 这里stu是个变量，在内存中内存地址是不变的
	// 改变的是stu所对应的值而已
	// m["小王子"] = stu的地址
	// m["娜扎"] = stu的地址
	// m["大王八"] = stu的地址
	// stu地址里面的值最后是{name:"大王八",age:9000}
	for _, stu := range stus {
		fmt.Println("stu:", stu)
		fmt.Printf("stu:%p\n", &stu)
		m[stu.name] = &stu
		fmt.Printf("%#v\n", m)
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name, "=>", v.age)
	}
}
