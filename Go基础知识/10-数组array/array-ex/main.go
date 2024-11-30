package main

import "fmt"

// array数组练习题
// 1. 求数组[1, 3, 5, 7, 8]所有元素的和
// 2. 找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)

func main() {
	// 1
	a1 := [...]int{1, 3, 5, 7, 8}
	sum1 := 0
	for _, v := range a1 {
		sum1 += v
	}
	fmt.Println("sum1:", sum1) // 24

	// 2
	a2 := [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(a2); i++ {
		for j := i + 1; j < len(a2); j++ {
			if a2[i]+a2[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}
