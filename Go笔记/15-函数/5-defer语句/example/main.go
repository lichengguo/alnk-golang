package main

import "fmt"

// 写出下面打印的值

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b)) // 3
	a = 0
	defer calc("2", a, calc("20", a, b)) // 2
	b = 1
}

// 思路
// 1. a = 1
// 2. b = 2
// 3. defer这里他会把变量值确定下来传递给calc函数，然后延迟调用。所以应该是 calc("1", 1, 3) 。然后这里会最先打印："10" 1 2 3
// 4. a = 0
// 5. defer这里他会把变量值确定下来传递给calc函数，然后延迟调用。所以应该是 calc("2",0,2) 。然后这里会打印 "20" 0 2 2
// 6. b=1
// 7. 执行calc("2",0,2) 会打印 "2" 0 2 2
// 8. 执行calc("1", 1, 3) 会打印 "1" 1 3 4

// 打印
// "10" 1 2 3
// "20" 0 2 2
// "2" 0 2 2
// "1" 1 3 4
