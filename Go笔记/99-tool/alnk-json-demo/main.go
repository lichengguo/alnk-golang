package main

import (
	"encoding/json"
	"fmt"
)

// str json 和结构体转换
// 序列化 反序列化

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 反序列化 str-->struct
	s1 := `{"name":"周林", "age":9000}`
	fmt.Printf("%T\n", s1) //string
	var p person
	_ = json.Unmarshal([]byte(s1), &p)
	fmt.Printf("%#v\n", p) //main.person{Name:"周林", Age:9000}

	// 序列化 strcut --> str
	p1 := person{
		Name: "保德路",
		Age:  22,
	}
	strJson, _ := json.Marshal(&p1)
	fmt.Printf("%#v\n", string(strJson)) //"{\"name\":\"保德路\",\"age\":22}"
}
