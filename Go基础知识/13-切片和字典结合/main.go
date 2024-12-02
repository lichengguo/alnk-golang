package main

import (
	"fmt"
)

// map和slice组合

func main() {
	// 元素类型为map的切片
	var s1 = make([]map[string]int, 10, 10) // 切片的元素类型为map 切片长度10 容量10
	fmt.Println(s1)                         // 切片已经初始化，但是里面的map类型没有初始化 [map[] map[] map[] map[] map[] map[] map[] map[] map[] map[]]

	// s1[0]的map类型初始化
	s1[0] = make(map[string]int, 10)
	fmt.Printf("%T %d\n", s1[0], len(s1[0]))  // map[string]int 0
	s1[0]["沙河"] = 10                          // s1[0]表示map，s1[0]["沙河"]表示map的key为沙河
	fmt.Printf("s1:%T s1[0]:%T\n", s1, s1[0]) // s1:[]map[string]int  s1[0]:map[string]int 注意这里s1是切片;s1[0]是map
	fmt.Println(s1)                           // [map[沙河:10] map[] map[] map[] map[] map[] map[] map[] map[] map[]]

	// s1[1]的map初始化
	s1[1] = make(map[string]int)
	s1[1]["kobe"] = 24
	s1[1]["duke"] = 21
	fmt.Println(s1) //[map[沙河:10] map[duke:21 kobe:24] map[] map[] map[] map[] map[] map[] map[] map[]]

	// 值为切片类型的map
	var m1 = make(map[string][]int) //key是string类型，值是[]int int切片类型
	m1["北京"] = []int{1, 2, 3}
	m1["上海"] = []int{4, 5}
	fmt.Println(m1) //map[上海:[4 5] 北京:[1 2 3]]
}
