## 运算符

Go语言内置的运算符有

1. 算术运算符
2. 关系运算符
3. 逻辑运算符
4. 位运算符
5. 赋值运算符



### 算术运算符

```go
package main

import "fmt"

func main() {
	var (
		a = 5
		b = 2
	)

	// 算术运算符
	fmt.Println(a + b) // 加
	fmt.Println(a - b) // 减
	fmt.Println(a * b) // 乘
	fmt.Println(a / b) // 除
	fmt.Println(a % b) // 取余

	// ++(自增) 和 --(自减) 在Go语言中是单独的语句，并不是运算符
	// a++ // 单独的语句 --> a=a+1
	// b-- // 单独的语句 --> b=b-1
}

```

结果

```shell
lichengguo@lichengguodeMacBook-Pro 04operator % go run test.go 
7
3
10
2
1
```



### 关系运算符

```go
package main

import "fmt"

func main() {
	var (
		a = 5
		b = 2
	)

	// 关系运算符
	// Go语言是强类型语言，只有相同类型的变量才能比较
	fmt.Println(a == b) // 等于
	fmt.Println(a != b) // 不等于
	fmt.Println(a >= b) // 大于等于
	fmt.Println(a <= b) // 小于等于
	fmt.Println(a > b)  // 大于
	fmt.Println(a < b)  // 小于

	fmt.Println("---- 分隔符 ----")
	// 字符也可以比较
	c1 := '中'
	c2 := '国'
	fmt.Printf("==%c\n", c1)    // 中
	fmt.Println("=", c1)        // 20013
	fmt.Println("==", c2)       // 2 2269
	fmt.Println("===", c1 < c2) // 底层是int32类型，所以可以比较
	fmt.Printf("%T\n", c1)      // int32
}

```

结果

```shell
lichengguo@lichengguodeMacBook-Pro 04operator % go run test.go
false
true
true
false
true
false
---- 分隔符 ----
==中
= 20013
== 22269
=== true
int32

```



### 逻辑运算符

```go
package main

import "fmt"

func main() {
	// 逻辑运算符
	// && 逻辑AND运算符 如果两边的操作数都是 True，则为 True，否则为 False
	// || 逻辑OR运算符  如果两边的操作数有一个 True，则为 True，否则为 False
	// !  逻辑NOT运算符 如果条件为 True，则为 False，否则为 True。

	// 如果年龄大于18岁 并且 并且小于60岁 &&
	age := 22
	if age > 18 && age < 60 {
		fmt.Println("上班族")
	} else {
		fmt.Println("不上班")
	}

	// 如果年龄小于18岁 或者 年龄大于60岁 ||
	if age < 18 || age > 60 {
		fmt.Println("不上班")
	} else {
		fmt.Println("work")
	}

	// not取反，原来为真就假，原来为假就真 !
	isMarried := false
	fmt.Println(!isMarried)
}

```

结果

```shell
lichengguo@lichengguodeMacBook-Pro 04operator % go run test.go
上班族
work
true

```



### 位运算符

```go
package main

import "fmt"

func main() {
	// 位运算：针对的是二进制数
	// 5的二进制表示：0101
	// 2的二进制表示：0010

	// &:按位与(两位均为1才为1)
	fmt.Println(5 & 2)
	// |:按位或（两位有一个为1就为1）
	fmt.Println(5 | 2)
	// ^:按位异或（两位不一样则为1）
	fmt.Println(5 ^ 2)
	// <<:将二进制位左移指定位数
	fmt.Println(5 << 1)  // 将5左移1位 0101 => 10100 = 10
	fmt.Println(1 << 10) // 将1左移10位 1 => 10000000000 = 1024
	// >>:将二进制位右移指定的位数
	fmt.Println(5 >> 2)
	var m = int8(1)      // 只能存8位
	fmt.Println(m << 10) // 因为int8 只能存储8位，向左移10位的话，就位0了
}
```

```shell
lichengguo@lichengguodeMacBook-Pro 04operator % go run test.go
0
7
7
10
1024
1
0

```



### 赋值运算符

```go
package main

func main() {
	// 赋值运算符，用来给变量赋值的
	var x int
	x = 10 // 简单的赋值运算符，将一个表达式的值赋给一个左值
	x += 1 //x = x + 1 相加后再赋值
	x -= 1 //x = x - 1 相减后再赋值
	x *= 2 //x = x * 2 相乘后再赋值
	x /= 2 //x = x / 2 相除后再赋值
	x %= 2 //x = x % 2 求余后再赋值

	x <<= 2 //x = x << 2 左移后赋值
	x &= 2  //x = x & 2 按位与后赋值
	x |= 3  //x = x | 3 按位或后赋值
	x ^= 4  //x = x ^ 4 按位异或后赋值
	x >>= 2 //x = x >> 2 右移后赋值
}
```



### 小练习

有一堆数字，如果除了一个数字以外，其他数字都出现了两次，那么如何找到出现一次的数字？

例如：`123432155`

```go
package main

import (
	"fmt"
)

func main() {
	// ^ 按位异或（两位不一样则为1）
	s := [9]int64{1, 2, 3, 4, 3, 2, 1, 5, 5}
	fmt.Println(s[0] ^ s[1] ^ s[2] ^ s[3] ^ s[4] ^ s[5] ^ s[6] ^ s[7] ^ s[8]) // 4
}
```