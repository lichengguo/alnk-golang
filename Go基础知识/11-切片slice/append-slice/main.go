package main

import (
	"fmt"
)

// append()为切片追加元素
// 可能会导致数组扩容，从而让切片指向新的数组

func main() {
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1)) // s1=[北京 上海 深圳] len(s1)=3 cap(s1)=3
	// s1[3] = "guangzhou" // panic: runtime error: index out of range [3] with length 3

	// 调用append函数必须用原来的切片变量接收返回值
	// append追加元素，原来的底层数组放不下的时候，Go就会把底层数组换一个
	// 必须用变量接收append的返回值
	s1 = append(s1, "广州")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1)) // s1=[北京 上海 深圳 广州] len(s1)=4 cap(s1)=6

	s1 = append(s1, "杭州", "成都")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1)) // s1=[北京 上海 深圳 广州 杭州 成都] len(s1)=6 cap(s1)=6

	ss := []string{"武汉", "西安", "苏州"}
	s1 = append(s1, ss...)                                            // 表示拆开(打散)，一个一个的元素添加到别的切片
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1)) // s1=[北京 上海 深圳 广州 杭州 成都 武汉 西安 苏州] len(s1)=9 cap(s1)=12

	//
	// 关于append删除切片中的某个元素
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	ss1 := a1[:]
	// 删除索引为1的那个值3
	ss1 = append(ss1[:1], ss1[2:]...) // 底层没有产生新的数组
	fmt.Println(len(ss1[:1]), cap(ss1[:1]))
	fmt.Println(len(ss1), cap(ss1))
	fmt.Println(ss1) // [1 5 7 9 11 13 15 17]
	fmt.Println(a1)  // [1 5 7 9 11 13 15 17 17]

}
