package main

import (
	"fmt"
)

// 切片的copy

func main() {
	a1 := []int{1, 3, 5}
	aa := a1[:1]
	fmt.Println(cap(aa), len(aa), aa) // 容量是3 长度是1 [1]
	a2 := a1                          // 赋值 a2 a1 指向的是同一个底层数组
	
	var a3 = make([]int, 3, 3)
	copy(a3, a1)            // copy 底层数组复制了一份，a3 和 a1 指向的不是同一个底层数组
	fmt.Println(a1, a2, a3) // [1 3 5] [1 3 5] [1 3 5]
	a1[0] = 100
	fmt.Println(a1, a2, a3) // [100 3 5] [100 3 5] [1 3 5]

	// 将a1中的索引为1的 3 这个元素删掉。 ... 由于切片没有直接删除元素的方法，所以可以采用这种方法
	a1 = append(a1[:1], a1[2:]...) // a1[:1] 容量是3，append 1 个元素的时候，底层数组没有发生变化
	fmt.Println(a1)                // [100 5]
	fmt.Println(cap(a1))           // 3
	fmt.Println(a2)                // [100 5 5]

	x1 := [...]int{1, 3, 5} // 数组
	fmt.Println("===", x1)
	s1 := x1[:] // 数组经过[L:M]以后，可以得到切片
	fmt.Println(s1, len(s1), cap(s1))
	// 1.切片不保存具体的值
	// 2.切片对应一个底层数组
	// 3.底层数组都是占用一块连续的内存
	fmt.Printf("%p\n", &s1[0])
	s1 = append(s1[:1], s1[2:]...)    // s1[:1] 容量是3，append 1 个元素的时候，底层数组没有发生变化
	fmt.Printf("%p\n", &s1[0])        // Go语言中不存在指针操作，只需要记住两个符号：&:取地址 *:根据地址取值
	fmt.Println(s1, len(s1), cap(s1)) // [1 5] 2 3
	fmt.Println(x1)                   // [1 5 5]
	s1[0] = 100
	fmt.Println(x1) // [100 5 5]
}
