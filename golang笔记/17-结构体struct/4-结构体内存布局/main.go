package main

import (
	"fmt"
	"unsafe"
)

// 结构体内存布局
// 结构体占用一块连续的内存空间

type x struct {
	a int8 // 8bit -> 1byte 字节
	b int8
	c int8
}

func main() {
	m := x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}
	fmt.Printf("%p\n", &(m.a))    // 0xc0000b4002
	fmt.Printf("%p\n", &(m.b))    // 0xc0000b4003
	fmt.Printf("%p\n", &(m.c))    // 0xc0000b4004
	fmt.Println(unsafe.Sizeof(m)) // 3

	// 空结构体
	// 空结构体是不占用空间的
	var v struct{}
	fmt.Println(unsafe.Sizeof(v)) // 0
}
