package main

// 在函数中使用泛型约束

import "fmt"

type Constraint interface {
    // 只允许以下的几种数据类型传参
    int | float64 | string // | []int
}

func Printer[T Constraint](vals ...T) {
    for _, v := range vals {
        fmt.Printf("%v \n", v)
    }
    fmt.Println()
}

func main() {
    Printer(1, 2, 3)
    Printer(1.1, 2.2, 3.3, 4.4)
    Printer("a", "b", "c")
    // Printer([]int{1,2}, []int{5, 6}, []int{8, 9})
}
