package main

import (
	"fmt"
)

// Go语言的字符有以下两种
// byte型，或者叫uint8类型，代表了ASCII码的一个字符
// rune类型，实际是一个int32，代表一个UTF-8字符, rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成
// 当需要处理中文、日文或者其他复合字符时，则需要用到rune类型
// 字符串底层是一个byte数组，所以可以和[]byte类型相互转换

func main() {
	fmt.Println("---------------- 字符字节 -------------------------------------")
	s := "Hello中国사샤"    // 非ASCII码一个字符等于3个字节，ASCII码一个字符等于1个字节
	fmt.Println(len(s)) // len()求得是byte字节的数量 17

	fmt.Println("---------------- rune切片 -------------------------------------")
	s1 := []rune(s)                 // 转换成rune切片
	fmt.Println("s1:", s1, len(s1)) // s1: [72 101 108 108 111 27801 27827 49324 49380] 9
	for i := 0; i < len(s1); i++ {
		fmt.Printf("R: %c\n", s1[i]) // %c 字符
	}

	fmt.Println("---------------- range -------------------------------------")
	// 下面这种方法会乱码
	// for i := 0; i < len(s); i++ {
	// 	//fmt.Println(s[i])
	// 	fmt.Printf("%c\n", s[i])
	// }

	// 使用range这种方式去遍历字符串，得到每个字符。
	// 不管有没有中文都能正常显示
	for index, c := range s {
		fmt.Printf("---%T\n", c)          // int32:字符的底层就是int32类型
		fmt.Printf("%d : %c\n", index, c) // 注意这里的index并不是按照1 2 3 4 这样+1上去的，是按照字符所在的字节位置，1个中文一般等于3个字节
	}

	fmt.Println("---------------- 字符串 -------------------------------------")
	// 字符串修改(字符串是不能修改的)
	s2 := "白萝卜"               // -> '白' '萝' '卜'
	s3 := []rune(s2)          // 把字符串强制转换成了一个rune切片
	fmt.Println("s3--: ", s3) // [30333 33821 21340]
	s3[0] = '红'
	fmt.Printf("s3: %v\n", s3) // s3: [32418 33821 21340]
	fmt.Println(string(s3))    // 把rune切片强制转换成字符串 红萝卜
	fmt.Println("s2:", s2)     // s2: 白萝卜

	// 双引号字符串，单引号字符
	c1 := "红"                           // 字符串string
	c2 := '红'                           // int32 又叫rune 字符
	fmt.Printf("c1:%T c2:%T\n", c1, c2) // 结果：c1:string c2:int32

	c3 := "h"                           // 字符串
	c4 := byte('h')                     // 字节 byte(uint8)  byte底层就是uint8类型
	fmt.Printf("c3:%T c4:%T\n", c3, c4) // c3:string c4:uint8
	fmt.Printf("c4:%d\n", c4)           // 104 ascii码 对应的字符就是h

	fmt.Println("---------------- 类型转换 -------------------------------------")
	// 类型转换
	// int --> float
	n1 := 10
	f := float64(n1)
	fmt.Println(f)          // 10
	fmt.Printf("f:%T\n", f) // f:float64

	// float --> int 会把小数部分去掉
	f1 := float64(10.1)
	n2 := int(f1)
	fmt.Println(n2) // 10
}
