package main

import (
	"fmt"
)

// 整型
// 有符号整型：int、int8、int16、int32、int64
// 无符号整型：uint、uint8、uint16、uint32、uint64

// 比较特殊的三个
// uint: 32位操作系统上就是uint32，64位操作系统上就是uint64
// int:  32位操作系统上就是int32，64位操作系统上就是int64
// uintptr: 无符号整型，用于存放一个指针

// 注意：在使用int和uint类型时，不能假定它是32位或64位的整型，而是考虑int和uint可能在不同平台上的差异
// 注意事项：获取对象的长度的内建len()函数返回的长度可以根据不同平台的字节长度进行变化
// 实际使用中，切片或map的元素数量等都可以用int来表示
// 在涉及到二进制传输、读写文件的结构描述时，为了保持文件的结构不会受到不同编译目标平台字节长度的影响，不要使用int和uint

// Go语言中有丰富的数据类型，整型、浮点型、布尔型、字符串、数组、切片、结构体、函数、map、通道（channel）等
// Go语言的基本类型和其他语言大同小异

func main() {
	// 十进制
	var i1 = 101
	fmt.Printf("%d\n", i1) // 十进制
	fmt.Printf("%b\n", i1) // 十进制 -> 二进制
	fmt.Printf("%o\n", i1) // 十进制 -> 八进制
	fmt.Printf("%x\n", i1) // 十进制 -> 十六进制

	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)

	// 十六进制
	i3 := 0x123456f
	fmt.Printf("%d\n", i3)
	fmt.Printf("%x\n", i3)

	// 查看变量类型
	fmt.Printf("%T\n", i3)

	// 声明int8类型的变量
	i4 := int8(9) // 明确指定int8类型，否则默认就是int类型
	fmt.Printf("%T\n", i4)

}
