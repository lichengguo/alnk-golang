package main

import "fmt"

// 定义一个泛型结构体
type Stack[T any] struct {
    data []T
}

// 为结构体定义Push方法
func (s *Stack[T]) Push(v T) {
    s.data = append(s.data, v)
}

// 为结构体定义Pop方法
func (s *Stack[T]) Pop() T {
    if len(s.data) == 0 {
        var zero T
        return zero
    }
    v := s.data[len(s.data)-1]
    s.data = s.data[:len(s.data)-1]
    return v
}

func main() {
    // 创建一个整数Stack
    intStack := Stack[int]{}
    intStack.Push(1)
    intStack.Push(2)
    intStack.Push(3)
    fmt.Println(intStack.Pop()) // 输出: 3
    fmt.Println(intStack.Pop()) // 输出: 2

    // 创建一个字符串Stack
    stringStack := Stack[string]{}
    stringStack.Push("hello")
    stringStack.Push("world")
    fmt.Println(stringStack.Pop()) // 输出: world
    fmt.Println(stringStack.Pop()) // 输出: hello
}
