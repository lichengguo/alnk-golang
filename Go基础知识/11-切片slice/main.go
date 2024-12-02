package main

import (
	"fmt"
)

// 切片
// 是一个拥有相同类型元素的可变长度的序列
// 它是基于数组类型做的一层封装，它非常灵活，支持自动扩容
// 切片是一个引用类型，它的内部结构包含地址、长度和容量
// 切片一般用于快速地操作一块数据集合
// 内置的len()函数求切片长度，内置的cap()函数求切片的容量
// 切片的底层就是一个数组，所以我们可以基于数组通过切片表达式得到切片
// 判断切片是否为空 请始终使用 len(s) == 0 来判断，而不应该使用s == nil来判断 初始化以后的切片 != nil

func main() {
	// 切片的定义
	fmt.Println("------------- 切片的定义 --------------------")
	var s1 []int                  // 定义一个存放int类型元素的切片
	var s2 []string               // 定义一个存放string类型元素的切片
	fmt.Println(s1, s2)           // [] []
	fmt.Println(s1 == nil)        // true
	fmt.Println(s2 == nil)        // true
	fmt.Println(len(s1), len(s2)) // 0 0

	// 切片初始化
	fmt.Println("------------- 切片初始化 --------------------")
	s1 = []int{1, 2, 3}
	s2 = []string{"沙河", "张江", "平山村"}
	fmt.Println(s1, s2)           // [1 2 3] [沙河 张江 平山村]
	fmt.Println(s1 == nil)        // false
	fmt.Println(s2 == nil)        // false
	fmt.Println(len(s1), len(s2)) // 3 3

	// 长度和容量
	fmt.Println("------------- 长度和容量 --------------------")
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1)) // len(s1):3 cap(s1):3
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2)) // len(s2):3 cap(s2):3

	// 由数组得到切片
	fmt.Println("------------- 由数组得到切片 --------------------")
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13} // 数组
	s3 := a1[0:4]                         // 基于一个数组切割，顾首不顾尾
	fmt.Println(s3)                       // [1 3 5 7]
	s4 := a1[1:6]
	fmt.Println(s4) // [3 5 7 9 11]
	s5 := a1[:4]
	s6 := a1[3:]
	s7 := a1[:]
	fmt.Println(s5) // [1 3 5 7]
	fmt.Println(s6) // [7 9 11 13]
	fmt.Println(s7) // [1 3 5 7 9 11 13]

	// 切片的容量是指底层数组的容量
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5)) // len(s5):4 cap(s5):7

	// 切片容量：底层数组从切片的第一个元素到最后的元素数量
	fmt.Printf("len(s6):%d cap(s6):%d\n", len(s6), cap(s6)) // len(s6):4 cap(s6):4

	// 切片再切割
	s8 := s6[3:]                                            // [13]
	fmt.Printf("len(s8):%d cap(s8):%d\n", len(s8), cap(s8)) // len(s8):1 cap(s8):1

	// 切片是引用类型，都指向了底层的一个数组
	// 修改底层数组的值，会影响切片
	fmt.Println("s6: ", s6) // [7 9 11 13]
	a1[6] = 1300            // 修改底层数组的值
	fmt.Println("s6: ", s6) // [7 9 11 1300]
	fmt.Println("s8: ", s8) // [1300]
	fmt.Println("a1: ", a1) // [1 3 5 7 9 11 1300]

	// 修改切片，会修改底层数组吗?
	// 如果只是修改值的话不涉及到扩容，是会修改原底层数组的
	s8[0] = 10000
	fmt.Println(s8)           // [10000]
	fmt.Println(a1)           // [1 3 5 7 9 11 10000]  a1数组发生改变
	s8 = append(s8, 20000)    // 此处已经产生了新的数组，扩容了，切片s8的容量是1，s8的切片不再指向原来的底层数组
	fmt.Printf("s8:%v\n", s8) // [10000 20000]
	fmt.Println(a1)           // [1 3 5 7 9 11 10000] a1数组不发生改变

	// append追加数组元素,如果底层数组容量不够的时候会扩容，产生新的数组
	fmt.Println("---------------- append追加数组元素 ----------------------------")
	a10 := [5]int{1, 2, 3}                                      // 已经初始化了, int类型没赋值,就是0
	fmt.Printf("a10: %v\n", a10)                                // a10: [1 2 3 0 0]
	fmt.Printf("a10 len: %d a10 cap: %d\n", len(a10), cap(a10)) // a10 len: 5 a10 cap: 5
	s10 := a10[:]                                               // 全切
	fmt.Printf("len(s10):%d cap(s10):%d\n", len(s10), cap(s10)) // len(s10):5 cap(s10):5
	// s10 = append(s10, 4)
	// fmt.Println(cap(s10)) //10
	s10[3] = 1 // 底层数组没扩容则直接修改数组的值, 如果扩容了则会生成一个新的数组
	// fmt.Println(a10, s10) // [1 2 3 0 0] [1 2 3 1 0 4] 如果没注释s10的append，证明Go已经把底层的数组换了
	fmt.Println(a10, s10) // [1 2 3 1 0] [1 2 3 1 0] 如果注释了s10的append，则s10[3]=1修改的是底层元素

	// make函数初始化切片
	fmt.Println("------------- make函数初始化切片 ------------------------------")
	var a = make([]string, 5, 10)         // 长度是5 容量是10
	fmt.Printf("cap:%d %#v\n", cap(a), a) // cap:10 []string{"", "", "", "", ""}
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i)) // 拼接成字符串
	}
	fmt.Printf("%#v\n", a)      // []string{"", "", "", "", "", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	fmt.Println(len(a), cap(a)) // 15 20
}
