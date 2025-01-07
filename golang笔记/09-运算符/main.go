// 运算符
package main

import "fmt"

func main() {
	var (
		a = 5
		b = 2
	)

	// 算术运算符
	fmt.Println("--------------- 算术运算符 ---------------------")
	fmt.Println(a + b) // 7
	fmt.Println(a - b) // 3
	fmt.Println(a * b) // 10
	fmt.Println(a / b) // 2
	fmt.Println(a % b) // 1
	a++                // 单独的语句 --> a=a+1
	b--                // 单独的语句 --> b=b-1
	fmt.Println(a, b)  // 6 1

	// 关系运算符
	// Go语言是强类型语言，相同类型的变量才能比较
	fmt.Println("--------------- 关系运算符 ---------------------")
	fmt.Println(a == b) // false
	fmt.Println(a != b) // true
	fmt.Println(a >= b) // true
	fmt.Println(a <= b) // false
	fmt.Println(a > b)  // true
	fmt.Println(a < b)  // false

	// 字符也可以比较
	fmt.Println("--------------- 字符 ---------------------")
	c1 := '中'
	c2 := '国'
	fmt.Printf("c1: %c\n", c1) // 中
	fmt.Println("c1: ", c1)    // 20013
	fmt.Println("c2: ", c2)    // 22269
	fmt.Println(c1 < c2)       // true 底层是int32类型，所以可以比较
	fmt.Printf("c1: %T\n", c1) // int32

	// 逻辑运算符
	// 如果年龄大于18岁 并且 并且小于60岁
	fmt.Println("--------------- 逻辑运算符 ---------------------")
	age := 22
	if age > 18 && age < 60 {
		fmt.Println("上班族")
	} else {
		fmt.Println("不上班")
	}

	// 如果年龄小于18岁 或者 年龄大于60岁
	if age < 18 || age > 60 {
		fmt.Println("不上班")
	} else {
		fmt.Println("work")
	}

	// not取反，原来为真就假，原来为假就真
	isMarried := false
	fmt.Println(isMarried)
	fmt.Println(!isMarried)

	// 位运算：针对的是二进制数
	// 5的二进制表示：0101
	// 2的二进制表示：0010
	// &:按位与(两位均为1才为1)
	fmt.Println("--------------- 位运算 ---------------------")
	fmt.Println(5 & 2) // 0
	// |:按位或（两位有一个为1就为1）
	fmt.Println(5 | 2) // 7
	// ^:按位异或（两位不一样则为1）
	fmt.Println(5 ^ 2) // 7
	// <<:将二进制位左移指定位数
	fmt.Println(5 << 1)  // 1010 => 10
	fmt.Println(1 << 10) // 10000000000 => 1024
	// >>:将二进制位右移指定的位数
	fmt.Println(5 >> 2) // 1
	// var m = int8(1)      // 只能存8位
	// fmt.Println(m << 10) // 因为int8 只能存储8位，向左移10位的话，就位0了

	// 赋值运算符，用来给变量赋值的
	fmt.Println("--------------- 赋值运算符 ---------------------")
	var x int
	x = 10
	x += 1 // x = x + 1
	x -= 1 // x = x - 1
	x *= 2 // x = x * 2
	x /= 2 // x = x / 2
	x %= 2 // x = x % 2

	x <<= 2 // x = x << 2
	x &= 2  // x = x & 2
	x |= 3  // x = x | 3
	x ^= 4  // x = x ^ 4
	x >>= 2 // x = x >> 2

	// 有一堆数字，如果除了一个数字以外，其他数字都出现了两次，那么如何找到出现一次的数字？
	s := [5]int{1, 2, 3, 1, 2}
	fmt.Println(s[0] ^ s[1] ^ s[2] ^ s[3] ^ s[4]) //3
}
