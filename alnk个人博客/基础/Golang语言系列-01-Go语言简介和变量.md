## Go语言简介

Go（又称Golang）是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。

罗伯特·格瑞史莫（Robert Griesemer），罗勃·派克（Rob Pike）及肯·汤普逊（Ken Thompson）于2007年9月开始设计Go，稍后Ian Lance Taylor、Russ Cox加入项目。

Go是基于Inferno操作系统所开发的。Go于2009年11月正式宣布推出，成为开放源代码项目，并在Linux及Mac OS X平台上进行了实现，后来追加了Windows系统下的实现。

在2016年，Go被软件评价公司TIOBE 选为“TIOBE 2016 年最佳语言”。 目前，Go每半年发布一个二级版本（即从a.x升级到a.y）

> 以上内容摘自[百度百科](https://baike.baidu.com/item/go/953521?fr=aladdin)



## Go开发环境安装

Go语言开发环境配置，网络上已经有很多的教程，建议使用最新的Go版本，并使用 **Go model** 进行Go包管理

[Go官网](https://golang.google.cn/dl/)

[Mac环境]()

[Windows环境]()



## Go语言的第一个程序

可以在任意目录下面创建一个 **main.go** 文件，文件内容如下

> 建议做好目录规划，方便以后的学习

```go
package main // 包名

// 导入其他包
// 这里可以导入内置包和第三方包
import (
	"fmt"
)

// 程序的入口函数
func main() {
	fmt.Println("hello world!")
}
```

执行方式：在 **main.go** 同级目录下运行下面的命令

```shell
go run main.go
```

结果如下

```shell
lichengguo@lichengguodeMacBook-Pro ~ % ll main.go
-rw-r--r--  1 lichengguo  staff  74 11  4 16:16 main.go
lichengguo@lichengguodeMacBook-Pro ~ % go run main.go
hello world
```



## 变量

Go语言中每一个变量都有自己的类型，并且变量必须要先声明后使用。同一 [作用域]() 内不支持重复声明变量，并且Go语言的变量声明后必须使用，否则会报错（全局声明的变量可以不必必须使用）。



Go语言推荐使用**驼峰命名方式**，如  **studentName** 学生姓名

```go
package main

import (
	"fmt"
)

func main() {
	var studentName string   // 声明一个变量
        studentName = "tom"      // 给变量赋值
        fmt.Println(studentName)
}
```



Go语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作，每个变量都会被初始化成其类型的默认值。例如：

- 整型 和 浮点型 变量的默认值为 **0**
- 字符串变量的默认值为空字符串
- 布尔型变量的默认值为 **false**
- 切片、函数、指针等变量的默认值为 **nil**

```go
package main

import "fmt"

// 程序的入口函数
func main() {
	// 批量声明
	var (
		age  int    		      // 整型
		name string 		      // 字符串
		isOK bool   		      // 布尔值
                testFunc func()               // 函数
	)
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(isOK)
        fmt.Println(testFunc)
}
```

执行结果

```shell
lichengguo@lichengguodeMacBook-Pro % go run main.go
0
          // 注意这一行的结果是空字符串
false
<nil>
```



### 其他的变量声明方式

```go
package main

import "fmt"

func main() {
	// 1. 声明变量的同时给赋值
	var name string = "tom01"
	fmt.Println(name)

	// 2. 一次声明多个变量，并且初始化多个变量
	var name1, age1 = "tom02", 20
	fmt.Println(name1, age1)

	// 3. 类型推导，根据值判断该变量是什么类型
	var s2 = "20"
	fmt.Println(s2)

        // 4. 简短变量声明，只能在函数内部使用 (比较常用)
	s3 := "Hello Go"
	fmt.Println(s3)

	// 匿名变量是一个特殊的变量：_ （在局部声明赋值以后，也可以不被使用，而不报错）
}
```

执行结果

```shell
lichengguo@lichengguodeMacBook-Pro % go run main.go
tom01
tom02 20
20
Hello Go
```

> 注意上面提到的变量都是在函数内部声明并且使用的 **局部变量**



**全局变量**

```go
package main

import "fmt"

// 全局变量
var (
	// 批量声明变量
	name01 string
	age01  int
)

func main() {
	// 局部变量 在函数内部声明的变量
	var name02 string
	var age02 int
	fmt.Println(name02, age02)
}
```

执行结果

> 从执行结果可以看出：全局变量声明以后不必必须使用，局部变量声明以后必须要使用

```shell
lichengguo@lichengguodeMacBook-Pro % go run main.go
 0 // 注意这前面是有空字符串的
```



## 常量

常量：定义了常量以后不能修改，常量在定义的时候必须赋值。常量在程序运行期间不会改变



### 声明常量

```go
package main

import (
	"fmt"
)

// 声明一个常量
const pi = 3.14

// 批量声明常量
const (
	statusOk = 200
	notFoud  = 404
)

// 批量声明常量时，如果某一行声明后没有赋值，默认就和上一行一致
const (
	n1 = 100
	n2
	n3
)

func main() {
	fmt.Println("--- 声明一个常量 ---")
	fmt.Println(pi)

	fmt.Println("--- 批量声明常量 ---")
	fmt.Println(statusOk)
	fmt.Println(notFoud)

	fmt.Println("--- 批量声明常量 默认 ---")
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)
}
```

执行结果

```shell
lichengguo@lichengguodeMacBook-Pro % go run main.go
--- 声明一个常量 ---
3.14
--- 批量声明常量 ---
200
404
--- 批量声明常量 默认 ---
n1: 100
n2: 100
n3: 100
```



### 常量计数器

**iota** 是Go语言的常量计数器，只能在常量的表达式中使用

**iota** 在const关键字出现时将被重置为**0**，const中每  **新增一行** 常量声明将使iota计数一次 **+1**

使用**iota**能简化定义，在定义枚举时很有用

```go
package main

import (
	"fmt"
)

const (
	a1 = iota //0
	a2        //1
	a3        //2
)

const (
	b1 = iota //0
	b2 = iota //1
	_  = iota //2
	b3 = iota //3
)

// 插队
const (
	c1 = iota //0   iota=0
	c2 = 100  //100 iota=1
	c3 = iota //2   iota=2
	c4        //3   iota=3
)

// 多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 // d1:1 d2:2
	d3, d4 = iota + 1, iota + 2 // d3:2 d4:3
)

// 定义量级
const (
	_  = iota
	KB = 1 << (10 * iota) // 1左移10位，二进制10000000000 转化为十进制1024
	MB = 1 << (10 * iota) // 1左移20位 100000000000000000000
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println("----- a -----")
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	fmt.Println("----- b -----")
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)

	fmt.Println("----- 插队 -----")
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)

	fmt.Println("----- 多个常量声明在一行 -----")
	fmt.Println(d1, d2)
	fmt.Println(d3, d4)

	fmt.Println("----- 定义量级 -----")
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
}

```

执行结果

```shell
lichengguo@lichengguodeMacBook-Pro % go run main.go
----- a -----
0
1
2
----- b -----
0
1
3
----- 插队 -----
0
100
2
3
----- 多个常量声明在一行 -----
1 2
2 3
----- 定义量级 -----
1024
1048576
1073741824
1099511627776
1125899906842624

```
