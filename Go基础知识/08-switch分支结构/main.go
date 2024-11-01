package main

import "fmt"

// switch
// 简化大量的判断（一个变量和具体的值作比较）
func main() {
	n := 2
	// if n == 1 {
	// 	fmt.Println("a")
	// }else if n == 2 {
	// 	fmt.Println("b")
	// }else if n == 3 {
	// 	fmt.Println("c")
	// }else if n == 4 {
	// 	fmt.Println("d")
	// }else if n == 5 {
	// 	fmt.Println("e")
	// }else {
	// 	fmt.Println("No")
	// }

	// 使用switch简化代码
	switch n {
	case 1:
		fmt.Println("a")
	case 2:
		fmt.Println("b")
	case 3:
		fmt.Println("c")
	case 4:
		fmt.Println("d")
	case 5:
		fmt.Println("e")
	default:
		fmt.Println("No")
	}

	// 多个匹配值
	switch n1 := 7; n1 {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n1)
	}

}
