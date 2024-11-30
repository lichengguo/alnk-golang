package main

import (
	"fmt"
)

// 布尔值
// 布尔类型变量的默认值为false
// Go语言中不允许将整型强制转换为布尔型
// 布尔型无法参与数值运算，也无法与其他类型进行转换

func main() {
	b1 := true
	var b2 bool // 默认是false
	fmt.Printf("b1类型: %T\n", b1)
	fmt.Printf("b2类型: %T, b2值: %v\n", b2, b2)
}
