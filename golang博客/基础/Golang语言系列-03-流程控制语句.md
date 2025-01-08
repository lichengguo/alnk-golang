# Go语言流程控制语句

Go语言中最常用的流程控制语句有 `if` 和  `for` ，没有像Python中的while语句。另外，Go语言还有`switch`和`goto`语句，不过这两个主要是用来简化代码的，属于扩展类的流程控制，使用率没有`if`和`for`多。



## if语句

### if语句的基本格式

```go
if 表达式1 {
  分支1
} else if 表达式2 {
  分支2
} else {
  分支3
}

// 表达式1为true时，执行分支1
// 表达式1为false时。判断表达式2，为true则执行分支2
// 如果表达式1和表达式2都为false时，执行分支3
```



### if 单分支

```go
package main

import (
	"fmt"
)

func main() {
	// 声明一个变量
	alnkAge := 18

	if alnkAge >= 18 {
		// 条件满足走这个分支
		fmt.Println("年轻人")
	}

	// 代码从上到下执行
	fmt.Println("代码执行结束")
}
```

```shell
lichengguo@lichengguodeMacBook-Pro % go run main.go
年轻人
代码执行结束
```



### if-else 双分支

```go
package main

import (
	"fmt"
)

func main() {
	// 声明一个变量
	alnkAge := 17

	if alnkAge >= 18 {
		// 条件满足走这个分支
		fmt.Println("年轻人")
	} else {
		// 如果上一个条件不满足，则走这个分支
		fmt.Println("小孩子")
	}

	// 代码从上到下执行
	fmt.Println("代码执行结束")
}
```

```shell
lichengguo@lichengguodeMacBook-Pro % go run main.go
小孩子
代码执行结束
```



### if-else if-else 多分支

```go
package main

import (
	"fmt"
)

func main() {
	// 声明一个变量
	alnkAge := 17

	if alnkAge >= 35 {
		// 条件满足走这个分支
		fmt.Println("中年人")
	} else if alnkAge >= 18 {
		// 如果上一个条件不满足，则走这个分支
		fmt.Println("骚年")
	} else {
		// 如果前面2个条件都不满足，则走这个分支
		fmt.Println("学生")
	}

	// 代码从上到下执行
	fmt.Println("代码执行结束")
}
```

```shell
lichengguo@lichengguodeMacBook-Pro % go run main.go
学生
代码执行结束

```



### 其他的if情况

```go
package main

import (
	"fmt"
)

func main() {
	// 在if语句这一行先声明一个变量，然后在进行判断
	// 这里涉及到一个新知识点 变量作用域
	// 此时的 变量alnkAge 只在if语句这个作用域内有效
	if alnkAge := 18; alnkAge > 18 {
		fmt.Println("大于18岁")
	} else {
		fmt.Println("学生")
	}

	// fmt.Println(alnkAge) // 会报错 undefined: alnkAge
}
```

> 作用域：变量可以使用的范围

* 局部变量：函数内部、或者流程控制语句内部定义的变量，就叫做局部变量。局部变量在哪里定义，就只能在该范围内使用，如果在其他范围使用，则会报错

* 全局变量：在函数、流程控制语句等外部定义的变量，就叫做全局变量。所有范围都可以使用全局变量。

变量作用域查找顺序：遵循LEGB原则（和Python一样）

+ L : local 局部作用域
+ E : enclose 嵌套作用域(函数中包含一个函数)
+ G : global 全局作用域
+ B : built in 内置作用域

```go
package main

import (
	"fmt"
)

// 全局变量的定义
//name1 := "Alnk" // 全局变量的定义不支持简短声明
var name1 = "Alnk" // 定义一个全局变量

// 在函数内部定义一个局部变量
func test() {
  // 定义一个局部变量
	var name2 = "Alnk2"
	fmt.Println(name2)
}

func main() {
	fmt.Println(name1) // 在 main 函数内部使用全局变量

	// 在 main 函数中不能使用 test 函数定义的局部变量。因为不在同一个作用域
	//fmt.Println(name2)
}

```



## for语句

### for循环语句的基本格式

```go
for 初始语句; 条件表达式; 结束语句 {
  循环体语句
}

// 条件表达式为true时，会一直循环，知道条件表达式为false时，退出循环
```



### 基本循环格式

```go
package main

import (
	"fmt"
)

// 最常用的for循环格式

func main() {
	// 基本循环格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
```



### 异形格式1

```go
package main

import (
	"fmt"
)

// 这种格式不常用

func main() {
	// 复习一下作用域
	i := 5
	for ; i < 10; i++ {
		// 这里的for循环里面为什么能找到变量i呢？参考LEGB原则
		// 首先会在for这个代码块的局部作用域中寻找i变量
		// 如果没找到，则会去上一层的嵌套作用域查找 这里是main函数内部
		// 如果还没找到，则会去全局作用域查找
		// 最后再去内置作用局查找，如果还没找到，则会报错
		fmt.Println(i)
	}
}

```



### 异形格式2

```go
package main

import (
	"fmt"
)

func main() {
  // 不常用
	i := 5
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

```



### for死循环

```go
package main

import (
	"fmt"
)

// 比较常用
// 用来hang住程序，不让程序退出
// ctrl+c 退出程序

func main() {
	for  {
		fmt.Println("Alnk")
	}
}

```



### for range 循环 

```go
package main

import (
	"fmt"
)

// 常用

func main() {
	s := "hello，你好"
	for index, value := range s {
		fmt.Printf("%d - %c\n", index, value)
	}
}

```



### 跳出for循环

#### break

```go
package main

import (
	"fmt"
)

func main() {
	// 当i=5时就跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // 直接跳出整个for循环
		}
		fmt.Println(i)
	}
	fmt.Println("over")
}

```



#### continue

```go
package main

import (
	"fmt"
)

func main() {
	// 当i=5时，跳过此次for循环（不执行for循环内部的打印语句），
	// 继续下一次循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // 跳过本次循环，执行下一次循环
		}
		fmt.Println(i)
	}
	fmt.Println("over")
}

```



## switch语句

switch语句主要是用来简化大量的判断，一个变量和具体的值作比较



#### 单个值匹配

```go
package main

import (
	"fmt"
)

func main() {
	//n := 2
	//if n == 1 {
	//	fmt.Println("1")
	//} else if n == 2 {
	//	fmt.Println("2")
	//} else if n == 3 {
	//	fmt.Println("3")
	//} else if n == 4 {
	//	fmt.Println("4")
	//} else if n == 5 {
	//	fmt.Println("5")
	//} else {
	//	fmt.Println("No")
	//}

	// 可以看出上面的if-else分支太多，导致代码可读性变差
	n := 2
	switch n {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	default:
		fmt.Println("No")
	}
}
```



#### 多个值匹配

```go
package main

import (
	"fmt"
)

func main() {
	//多个匹配值
	n1 := 7
	switch n1 {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n1)
	}

	// 也可以在switch代码块中声明变量并且判断
	switch n2 := 8; n2 {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n1)
	}
}
```



 ## goto语句

goto语句用来跳转到指定的标签位置，不建议使用，会影响代码的可读性

如果要跳出多层for循环，用标记位

```go
package main

import (
	"fmt"
)

func main() {
	// 用标记位, 跳出多层for循环
	flag := false

	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				flag = true
				break // 跳出内层的for循环
			}
			fmt.Printf("%v-%c\n", i, j)
		}
		if flag {
			break // 跳出外层的for循环
		}
	}
}
```

```go
package main

import (
	"fmt"
)

func main() {
	// goto+label(标签) 实现跳出多层for循环(不建议使用)
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto XX // 跳到指定的那个标签
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
XX: // label(标签)
	fmt.Println("over")
}

```



## 实践练习

- 编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型
- 编写代码统计出字符串`"hello你好呀Alnk"` 中的汉字的数量
- 打印一个`九九乘法表`
```go
package main

import (
	"fmt"
	"unicode"
)

func main() {
	//1
	//i1 := 10
	//f1 := 1.234
	//b1 := true
	//s1 := "hello沙河"
	//fmt.Printf("%T %d\n", i1, i1)
	//fmt.Printf("%T %f\n", f1, f1)
	//fmt.Printf("%T %v\n", b1, b1)
	//fmt.Printf("%T %s\n", s1, s1)

	//2
	s := "hello你好呀Alnk"
	result := chineseCount(s)
	fmt.Println(result)

	//3 九九乘法表
	multiplicationTable()

}

func chineseCount(str1 string) (count int) {
	for _, char := range str1 {
		if unicode.Is(unicode.Han, char) {
			count++
		}
	}
	return
}

func multiplicationTable() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, j*i)
		}
		fmt.Println()
	}
}

```