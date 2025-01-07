package main

import "fmt"

// 递归 函数自己调用自己
// 递归适合处理那种问题相同问题的规模越来越小的场景
// 递归一定要有一个明确的退出条件

// 计算n的阶乘
// 3! = 3*2*1   		=3*2!
// 4! = 4*3*2*1 		=4*3!
// 5! = 5*4*3*2*1	=5*4!
func f(n uint64) uint64 {
	if n <= 1 {
		return n
	}
	return n * f(n-1)
}

// 上台阶练习
// n个台阶 一次可以走1步 也可以走2步 有多少种走法
func taijie(n uint64) uint64 {
	if n == 1 {
		// 如果只有一个台阶那就一种走法
		return 1
	}

	if n == 2 {
		// 如果2个台阶那就2种走法
		return 2
	}

	return taijie(n-1) + taijie(n-2)
}

func main() {
	// fmt.Println(f(5))
	ret := taijie(6)
	fmt.Println(ret) // 13
}
