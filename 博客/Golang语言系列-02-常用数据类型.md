## Go语言常用数据类型

**Go** 语言中有丰富的数据类型，除了基本的整型、浮点型、布尔型、字符串、byte/rune 之外，

还有数组、切片、函数、map、通道(channel)、结构体等。

Go语言的基本类型和其他语言大同小异。



### 整型

整型分为：有符号整型、无符号整型

其中比较特殊的整型数据类型有3个

+ uint 无符号整型，它在32位的操作系统上就是uint32，在64位的操作系统上就是uint64

+ int 有符号整型， 它在32位的操作系统上就是int32，在64位的操作系统上就是int64

+ uintptr 无符号整型，用于存放一个指针

  > 在使用 **int** 和 **uint** 类型时，不能假定他是32位或是64位的整型类型，而是要考虑程序可能运行在不同的平台上，导致数据位数不一样。所以，为了保持文件的结构不会受到不同的平台的影响，建议不要使用 **int** 和 **uint**

#### 有符号整形

+ int、int8、int16、int32、int64 都属于有符号整型的数据类型

#### 无符号整型

+ uint、uint8、uint16、uint32、uint64 都属于无符号整型的数据类型

#### 整型数据练习，数据类型转换

```go
package main

import "fmt"

func main() {
	// 十进制
	var i1 = 101
	fmt.Println("--- 十进制 ---")
	fmt.Printf("%d\n", i1) // 十进制
	fmt.Printf("%b\n", i1) // 十进制 -> 二进制
	fmt.Printf("%o\n", i1) // 十进制 -> 八进制
	fmt.Printf("%x\n", i1) // 十进制 -> 十六进制

	// 八进制
	i2 := 077
	fmt.Println("--- 八进制 ---")
	fmt.Printf("%d\n", i2) // 八进制 -> 十进制

	// 十六进制
	i3 := 0x123456f
	fmt.Println("--- 十六进制 ---")
	fmt.Printf("%d\n", i3) // 十六进制 -> 十进制
	fmt.Printf("%x\n", i3) // 十六进制

	// 查看变量类型
	fmt.Println("--- 查看变量类型 ---")
	fmt.Printf("%T\n", i3) // 如果不指定数据类型，那么默认则为int类型
	// 声明int8类型的变量
	i4 := int8(9) // 明确指定int8类型，否则默认就是int类型
	fmt.Printf("%T\n", i4)
}
```

```shell
lichengguo@lichengguodeMacBook-Pro % go run main.go
--- 十进制 ---
101
1100101
145
65
--- 八进制 ---
63
--- 十六进制 ---
19088751
123456f
--- 查看变量类型 ---
int
int8
```



### 浮点型

浮点型分为 **float32** 和 **float64** 两种类型

+ float32
+ float64

浮点类型比较简单，不多说，直接看代码

```go
package main

import (
	"fmt"
	"math"
)

//浮点数

func main() {
	// math.MaxFloat32 float32最大值
	// math.MaxFloat64 float64最大值
	fmt.Println("--- 1. 浮点型最大值 ---")
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	fmt.Println("---- 2.  -------")
	f1 := 1.23456189         // float64 类型
	fmt.Printf("%T\n", f1)   // 默认Go语言中的小数都是 float64 类型
	fmt.Printf("%f\n", f1)   // 1.234562 默认只打印小数点后六位，会采取进一法保留六位小数
	fmt.Printf("%.2f\n", f1) // 保留2位小数 1.23 会采用四舍五入保留2位小数

	fmt.Println("---- 3. float32 ----")
	f2 := float32(1.23456) // float32 类型
	fmt.Printf("%T\n", f2) // 显示声明 float32 类型
	// float32 类型的值不能直接赋值给 float64 类型的变量
	// f1 = f2 // 会报错

	fmt.Println("---- 4. ----")
	// 同类型的可以直接赋值
	f3 := 1.2 // float64 类型
	f1 = f3
	fmt.Println(f1)
	fmt.Printf("%T\n", f1)
	fmt.Printf("%T\n", f3)
}

```

执行结果

```shell
lichengguo@lichengguodeMacBook-Pro % go run main.go
--- 1. 浮点型最大值 ---
3.4028234663852886e+38
1.7976931348623157e+308
---- 2.  -------
float64
1.234562
1.23
---- 3. float32 ----
float32
---- 4. ----
1.2
float64
float64

```



### 布尔型

Go语言中以 **bool** 类型进行声明布尔型数据，布尔型数据只有 **true** 和 **false** 两个值

+ 布尔类型变量的默认值是 **false**
+ Go语言不允许将 **整型** 强制转换为 **布尔型**
+ 布尔型无法参与数值运算，也无法与其他类型进行转化

```go
package main

import (
	"fmt"
)

func main() {
	b1 := true
	var b2 bool // 默认是false
	fmt.Printf("%T\n", b1)
	fmt.Printf("%T value:%v\n", b2, b2)
}
```

执行结果

```shell
lichengguo@lichengguodeMacBook-Pro day01 % go run main.go
bool
bool value:false
```



### 字符串

Go语言中 **字符串** 是用 **双引号" "** 包裹的，用 **单引号' '** 包裹的是 **字符**，Go语言里的字符串内部实现使用的是**UTF8 **编码

+ ASCII编码中: 一个字符 **'A'** 占用1个字节(Byte)

+ UTF8编码中: 一个字符 **'A'** 占用1个字节(Byte)，一个汉字 **'中'** 一般占用3个字节(Byte)

+ > 小知识: 
  >
  > ​    1字节(Byte) = 8bit (8个二进制位)
  >
  > ​    1024字节(Byte) = 1KB



字符串和字符这两者之间的区别

**字符串**

+ 双引号
+ 字符串有一个或者多个 **字符** 组成
+ 字符串都是隐藏了一个结束符：`\0`

**字符**

+ 单引号
+ 往往只包含一个字符，转义字符除外，如`\n`

下面通过代码具体来看一下

```go
package main

import (
	"fmt"
)

func main() {
	// 字符
	fmt.Println("---- 1. 字符 ----")
	ch := 'a' // 简短声明变量并且赋值
	fmt.Println("ch =", ch)
	fmt.Printf("%T\n", ch)

	// 字符串
	fmt.Println("---- 2. 字符串 ----")
	s1 := "a"
	fmt.Println("s1 =", s1)
	fmt.Printf("%T\n", s1)
}
```

执行结果

> 从结果可以看出，字符类型本质上是一个**int32** 类型

```shell
lichengguo@lichengguodeMacBook-Pro day01 % go run main.go
---- 1. 字符 ----
ch = 97  // ASCII编码97对应的是小写字母a
int32
---- 2. 字符串 ----
s1 = a
string
```

**转义符 \\**

```go
package main

import (
	"fmt"
)

func main() {
	// \ 反斜杠是具有特殊含义的，应该告诉程序写的 \ 就是一个单纯的 \
	// path := "D:\\Go\\src\\studygo\\day01"
	// path := "'D:\\Go\\src\\studygo\\day01'"
	path := "\"D:\\Go\\src\\studygo\\day01\""
	fmt.Println(path)
}
```

**多行字符串**

```go
package main

import (
	"fmt"
)

func main() {
	// 多行字符串
	s2 := `
	aaa
bbb
		ccc
	`
	fmt.Println(s2)

	s3 := `D:\Go\src\code.oldboyedu.com\studygo\day01`
	fmt.Println(s3)
}
```

**字符串的一些常用操作**

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串相关操作
	// 1. 统计字节数
	fmt.Println("---- 1. 统计字节数 ----")
	s1 := "hello中国"
	// 注意 len函数 统计字符串的时候，是统计【字节数长度】，而不是字符的个数
	// 在utf8编码中，一个英文字符占用1个字节，一个中文字符一般占用3个字节
	fmt.Println(len(s1)) // 此处的11是字节数长度，而不是字符的个数

	// 2. 字符串拼接
	fmt.Println("---- 2. 字符串拼接 ----")
	name := "tom"
	world := "dsb"
	ss := name + world                         // 拼接方式1
	fmt.Println(ss)                            // tomdsb
	ss1 := fmt.Sprintf("%s - %s", name, world) // 拼接方式2
	fmt.Println(ss1)                           // tom - dsb

        // 3. strings包相关操作
	fmt.Println("---- 3. strings包相关操作 ----")
	
	// 3.1 分割
	fmt.Println("---- 3.1 分割 ----")
	s3 := `D:\Go\src\studygo\day01`
	ret := strings.Split(s3, "\\") // 注意这里的 \\ 前面的反斜杠是为了不让后面的反斜杠具有特殊意义
	fmt.Println(ret)               // [D: Go src studygo day01]

	// 3.2 包含
	fmt.Println("---- 3.2 包含 ----")
	fmt.Println(strings.Contains(ss, "dsb")) // true

	// 3.3 前缀
	fmt.Println("---- 3.3 前缀 ----")
	fmt.Println(strings.HasPrefix(ss, "tom")) // true

	// 3.4 后缀
	fmt.Println("---- 3.4 后缀 ----")
	fmt.Println(strings.HasSuffix(ss, "dsb")) // true

	// 3.5 查找字符串出现的索引
	fmt.Println("---- 3.5 查找字符串出现的索引 ----")
	s5 := "abcdeb"
	fmt.Println(strings.Index(s5, "c"))       // 2 从0开始计数
	fmt.Println(strings.LastIndex(s5, "eb"))  // 4

	// 单独的字母、汉字、符号并且用单引号括起来的表示一个字符
	// 双引号的是字符串了
	fmt.Println("-----------------------------")
	s6 := 'c'
	fmt.Printf("s6: %T\n", s6)  // int32
	s7 := '中'
	fmt.Printf("s7: %T\n", s7)  // int32  
	s8 := "c"
	fmt.Printf("s8: %T\n", s8)  // string

	// 3.6 切片拼接
	fmt.Println("---- 3.6 切片拼接 ----")
	fmt.Println(strings.Join(ret, "+"))  // D:+Go+src+studygo+day01
}

```

```shell
lichengguo@lichengguodeMacBook-Pro day01 % go run main.go 
---- 1. 统计字节数 ----
11
---- 2. 字符串拼接 ----
tomdsb
tom - dsb
---- 3. strings包相关操作 ----
---- 3.1 分割 ----
[D: Go src studygo day01]
---- 3.2 包含 ----
true
---- 3.3 前缀 ----
true
---- 3.4 后缀 ----
true
---- 3.5 查找字符串出现的索引 ----
2
4
-----------------------------
s6: int32
s7: int32
s8: string
---- 3.6 切片拼接 ----
D:+Go+src+studygo+day01

```



### byte和rune

Go语言中的 **字符** 有两种数据类型

+ uint8类型，或者叫做 byte 类型，代表了ASCII编码的一个字符
+ rune类型，代表了UTF-8编码的一个字符

当需要处理中文、日文等其他复合字符时，则需要用到`rune`类型。`rune`类型实际上就是一个`int32`类型

**遍历字符串**

```go
package main

import (
	"fmt"
)

// Go语言里的字符串内部实现使用的是UTF8编码

func main() {
	s := "Hello中国"
	n := len(s)    // len 的结果是字节的长度，而不是字符的数量
	fmt.Println(n) // 11

	// 遍历字符串
	// Go语言里的字符串内部实现使用的是UTF8编码，而rune类型，代表了UTF8编码的一个[字符]
	s1 := []rune(s)                 // 转换成rune类型的切片
	fmt.Println("s1:", s1, len(s1)) // s1: [72 101 108 108 111 228 184 173 229 155 189] 7
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%c", s1[i]) // %c 字符
	}
	fmt.Println()

	// 下面这种方法会乱码
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}
	fmt.Println()

	// 使用range这种方式去遍历字符串，得到每个字符。
	// 不管有没有中文都能正常显示
	for index, c := range s {
		// 注意这里的index并不是按照1 2 3 4 这样+1上去的，是按照字符所在的字节位置，1个中文一般等于3个字节
		fmt.Printf("数据类型:%T - index:%d -- 字符:%c\n", c, index, c)
	}
}

```

输出

```shell
lichengguo@lichengguodeMacBook-Pro day01 % go run test.go
11
s1: [72 101 108 108 111 20013 22269] 7
Hello中国
Helloä¸­å½
数据类型:int32 - index:0 -- 字符:H
数据类型:int32 - index:1 -- 字符:e
数据类型:int32 - index:2 -- 字符:l
数据类型:int32 - index:3 -- 字符:l
数据类型:int32 - index:4 -- 字符:o
数据类型:int32 - index:5 -- 字符:中
数据类型:int32 - index:8 -- 字符:国
```



### 类型转换

```
package main

import (
	"fmt"
)

func main() {
	// 类型转换
	// int --> float
	fmt.Println("---- int --> float ----")
	n1 := 10
	f := float64(n1)
	fmt.Println(f)        // 10
	fmt.Printf("%T\n", f) // f:float64

	// float --> int 会把小数部分去掉
	fmt.Println("---- float --> int ----")
	f1 := float64(10.1)
	n2 := int(f1)
	fmt.Println(n2) // 10

	// string <--> rune
	fmt.Println("---- string <--> []rune ----")
	s1 := "tom"
	s2 := []rune(s1)
	fmt.Println(s2) // [116 111 109] ASCII编码对照表

	s3 := "中国good"
	s4 := []rune(s3)       // string --> []rune
	fmt.Println(s4)        // [20013 22269 103 111 111 100]
	fmt.Printf("%T\n", s4) // 数据类型 []int32
	s5 := string(s4)       // []rune --> string
	fmt.Println(s5)        // 中国good

	// string <--> byte
	fmt.Println("---- string <--> []byte ----")
	s6 := "你好hello"
	s7 := []byte(s6)
	fmt.Println(s7)        // [228 189 160 229 165 189 104 101 108 108 111] 一个中文字符占3个字节
	fmt.Printf("%T\n", s7) // []uint8

	// 其他
	//s8 := []byte{'A', 1, 3, 'B', '汉'}  byte里面不能存中文，会报错 constant 27721 overflows byte
	fmt.Println("---- other ----")
	s8 := []byte{'A', 1, 3, 'B'}
	fmt.Println(s8)        // [65 1 3 66]
	fmt.Printf("%T\n", s8) // []uint8
	s9 := []rune{'中', 'A', 1, 3}
	fmt.Println(s9)        // [20013 65 1 3]
	fmt.Printf("%T\n", s9) // []int32

}

```

输出

```shell
lichengguo@lichengguodeMacBook-Pro day01 % go run main.go
---- int --> float ----
10
float64
---- float --> int ----
10
---- string <--> []rune ----
[116 111 109]
[20013 22269 103 111 111 100]
[]int32
中国good
---- string <--> []byte ----
[228 189 160 229 165 189 104 101 108 108 111]
[]uint8
---- other ----
[65 1 3 66]
[]uint8
[20013 65 1 3]
[]int32
```
