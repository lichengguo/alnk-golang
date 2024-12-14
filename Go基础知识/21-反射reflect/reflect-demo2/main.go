package main

import (
	"fmt"
	"reflect"
)

// r eflect 反射
/*
任意值通过reflect.TypeOf()获得反射对象信息后，如果它的类型是结构体，
可以通过反射值对象（reflect.Type）的NumField()和Field()方法获得结构体成员的详细信息。

reflect.Type中与获取结构体成员相关的的方法如下表所示
Field(i int) StructField										根据索引，返回索引对应的结构体字段的信息。
NumField() int													返回结构体成员字段数量。
FieldByName(name string) (StructField, bool)					根据给定字符串返回字符串对应的结构体字段的信息。
FieldByIndex(index []int) StructField							多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息。
FieldByNameFunc(match func(string) bool) (StructField,bool)		根据传入的匹配函数匹配需要的字段。
NumMethod() int													返回该类型的方法集中方法的数目
Method(int) Method												返回该类型方法集中的第i个方法
MethodByName(string)(Method, bool)								根据方法名返回该类型方法集中的方法
*/

type student struct {
	Name  string `json:"name" zhoulin:"嘿嘿嘿"`
	Score int    `json:"score" zhoulin:"哈哈哈"`
}

func main() {
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) //student struct

	//通过for循环遍历结构体的所有字段信息
	fmt.Println(t.NumField()) //2
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("zhoulin"))
		/*
			name:Name index:[0] type:string json tag:嘿嘿嘿
			name:Score index:[1] type:int json tag:哈哈哈
		*/

		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
		/*
			name:Name index:[0] type:string json tag:name
			name:Score index:[1] type:int json tag:score
		*/

	}

	//通过字段名获取指定结构体的字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("===%#v\n", scoreField)
		//===reflect.StructField{Name:"Score", PkgPath:"", Type:(*reflect.rtype)(0x10ae620), Tag:"json:\"score\" zhoulin:\"哈哈哈\"", Offset:0x10, Index:[]int{1}, Anonymous:false}

		fmt.Println(scoreField) //{Score  int json:"score" zhoulin:"哈哈哈" 16 [1] false}

		fmt.Printf("name:%s inde:%d type:%v json_tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
		//name:Score inde:[1] type:int json_tag:score

		//字段+字段值
		fmt.Printf("字段:[%s], 字段值:[%v]\n", scoreField.Name, stu1.Score) //字段:[Score], 字段值:[90]
	}
}
