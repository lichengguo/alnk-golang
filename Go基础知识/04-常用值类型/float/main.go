package main

import (
	"fmt"
	"math"
)

// 浮点数

func main() {
	// math.MaxFloat32 // float32最大值
	// math.MaxFloat64 // float64最大值
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	f1 := 1.23456
	fmt.Printf("%T\n", f1)   // 默认Go语言中的小数都是float64类型
	fmt.Printf("%f\n", f1)   // 1.234560 默认小数点后6位
	fmt.Printf("%.2f\n", f1) // 保留2位小数 1.23

	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2) // 显示声明float32类型

	// float32 类型的值不能直接赋值给 float64 类型的变量
	// f1 = f2

	f3 := 1.2
	f1 = f3
	fmt.Println(f1)
}
