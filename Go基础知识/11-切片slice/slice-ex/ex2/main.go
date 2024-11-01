package main

import (
	"fmt"
	"sort"
)

// 切片练习题

func main() {
	var a = make([]int, 5, 10) // 初始化切片，长度5，容量10
	fmt.Println(a)             // [0 0 0 0 0]
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)      // [0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
	fmt.Println(cap(a)) // 20

	var a1 = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a1[:]) // 对切片进行排序
	fmt.Println(a1)  // [1 3 7 8 9]

	// 要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断
	var aa []int           // 只定义，没初始化
	fmt.Println(aa == nil) // true
	var bb []int           // 定义
	bb = []int{}           // 初始化后就不为nil了，但是还是空的切片
	fmt.Println(bb == nil) // false
	fmt.Println(aa, bb)    // [] []
	if len(bb) == 0 {
		fmt.Println("空数组")
	}
}
