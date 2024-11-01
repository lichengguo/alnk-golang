package split_string

import (
	"fmt"
	"strings"
)

// Split:切割字符串
// example:
// abc, b --> [a c]
func Split(str string, sep string) []string {
	//var ret = make([]string, 0, 1)

	// 优化代码，初始化的时候指定长度和容量，避免在append的时候去动态扩容，影响性能
	var ret = make([]string, 0, strings.Count(str, sep)+1) //切片的make参数:类型、长度、容量

	index := strings.Index(str, sep)
	for index >= 0 {
		ret = append(ret, str[:index])
		str = str[index+len(sep):] //注意这不是切片,这是字符串切割
		index = strings.Index(str, sep)
	}
	ret = append(ret, str)

	// 为了让测试率达不到100%，只是试验而已，以后可以不用写这个if
	if index == -5 {
		fmt.Println("No!!!")
	}

	return ret
}

// 用来做性能比较测试的例子
// Fib是一个计算第n个斐波那契数的函数
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
