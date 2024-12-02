package main

import (
	"fmt"
	"strings"
)

//  写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1

func main() {
	s1 := "how do you do"
	s2 := strings.Split(s1, " ") // 返回 []string 切片

	// 定义一个map并且初始化
	var m1 = make(map[string]int, 100) // 初始化以后int类型的零值为0

	for _, key := range s2 {
		//m1初始化以后,值int类型的零值为0
		//m1[key] = m1[key] + 1
		//m1[key] += 1
		m1[key]++
	}

	fmt.Println(m1)         // map[do:2 how:1 you:1]
	fmt.Printf("%#v\n", m1) // map[string]int{"do":2, "how":1, "you":1}
}
