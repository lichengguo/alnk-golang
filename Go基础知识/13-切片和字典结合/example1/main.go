package main

import (
	"fmt"
	"strings"
)

/*
1.写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1
*/

func main() {
	s1 := "how do you do"
	s2 := strings.Split(s1, " ") //返回 []string 切片

	//定义一个map并且初始化
	var m1 = make(map[string]int, 100) //初始化以后int类型的零值为0

	////循环s2这个切片
	//for _, key := range s2 {
	//	_, ok := m1[key]
	//	//如果map中没有这个key，那么该key的值为1，如果存在该key，则加1
	//	if !ok {
	//		m1[key] = 1
	//	} else {
	//		m1[key] += 1
	//	}
	//}
	//
	////循环m1这个map打印
	//for k, v := range m1 {
	//	fmt.Printf("%s=%d\n", k, v)
	//}

	for _, key := range s2 {
		//m1初始化以后,值int类型的零值为0
		//m1[key] = m1[key] + 1
		//m1[key] += 1
		m1[key]++
	}
	fmt.Println(m1)
}
