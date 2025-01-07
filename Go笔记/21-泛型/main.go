package main

// 在Go 1.18版本中,泛型正式被引入

import "fmt"

// 定义一个泛型函数,用于交换两个值
func Swap[T any](a, b *T) {
    *a, *b = *b, *a
}

func main() {
    // 交换两个整数
    x, y := 1, 2
    fmt.Printf("x=%d, y=%d\n", x, y)
    Swap(&x, &y)
    fmt.Printf("x=%d, y=%d\n", x, y)

    // 交换两个字符串
    a, b := "hello", "world"
    fmt.Printf("a=%s, b=%s\n", a, b)
    Swap(&a, &b)
    fmt.Printf("a=%s, b=%s\n", a, b)
}
