package main

import "fmt"

type Constraint interface {
    int | float64 | string // | []int
}

func Printer[T Constraint](vals ...T) {
    for _, v := range vals {
        fmt.Printf("%v ", v)
    }
    fmt.Println()
}

func main() {
    Printer(1, 2, 3)
    Printer(1.1, 2.2, 3.3)
    Printer("a", "b", "c")
    // Printer([]int{1,2}, []int{5, 6}, []int{8, 9})
}
