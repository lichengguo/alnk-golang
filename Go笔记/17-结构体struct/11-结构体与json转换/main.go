package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与json
// 反射

// 1.序列化：把go语言中的结构体变量 ---> json格式的字符串
// 2.反序列化：json格式的字符串 ---> go语言中能够识别的结构体变量

// 这里的字段名称开头要大写，是因为要把这个字段名称传入到第三方的包中，如果小写就不能暴露字段了
// 序列化和反序列化以后的字段如果需要是小写的，则需要后面的json标记
// 如果序列化和反序列化后都是大写开头，那么不需要json标记也可以
type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age"`
}

// type person struct {
// 	Name string
// 	Age  int
// }

func main() {
	// 序列化
	p1 := person{
		Name: "tom",
		Age:  9000,
	}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("序列化失败, err:%v", err)
		return
	}
	fmt.Printf("%v\n", string(b))

	// 反序列化
	//str := `{"Name":"理想", "Age":18}`
	str := `{"name":"李逵", "age":18}`
	var p2 person
	err = json.Unmarshal([]byte(str), &p2)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v", err)
		return
	}
	fmt.Printf("%#v\n", p2)
}
