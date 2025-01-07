# 函数

## 函数的基本概念

```go
package main

import (
	"fmt"
)

// 函数
// 函数存在的意义:函数能够让代码结构更加清晰，更简洁，能够让代码复用
// 函数是一段代码的封装
// 把一段逻辑抽象出来封装到一个函数中，给他起个名字，每次用它的时候直接用函数名调用即可

// 函数的基本定义
// func 函数名称(参数变量名称1 参数类型, 参数变量名称2 参数类型, ...) (返回值变量名称1 返回值类型, 返回值变量名称2 返回值类型, ...)
func sum(x int, y int) (ret int) {
	return x + y
}

func main() {
  ret := sum(1,1)
  fmt.Println(ret)
}
```



## 函数的变种

```go
package main

import (
	"fmt"
)

// 函数的变种1:没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

// 函数的变种2：没有参数也没有返回值
func f2() {
	fmt.Println("f2")
}

// 函数的变种3：没有参数但是有返回值
func f3() int {
	ret := 3
	return ret
}

// 函数的变种4：函数返回值可以命名也可以不命名
// 命名的返回值就相当于在函数中声明一个变量
func f4(x int, y int) (ret int) {
	ret = x + y
	//return ret
	return //使用命名返回值可以省略return后面的变量名称
}

// 函数变种5：多个返回值
func f5() (int, string) {
	return 1, "abc"
}

// 函数变种6：多个参数的类型简写
func f6(x, y, z int, m, n string, i, j bool) int {
	return x + y
}

// 函数变种7：可变长的参数
// 可变长的参数必须放在函数参数的最后
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y) // y的类型是 []int 切片
}

// 注意：Go语言中函数没有默认参数这个概念

func main() {
	f7("下雷了")
	f7("下雨了", 1, 2, 3, 4, 5)
}
```

```go
package main

import "fmt"

// 函数:一段代码的封装

// 不带参数，不带返回值的
func f1() {
	fmt.Println("Hello！")
}

// 带参数，不带返回值的
func f2(name string) {
	fmt.Println("Hello", name)
}

// 带参数和返回值的函数
func f3(x int, y int) int {
	sum := x + y
	return sum
}

// 参数类型简写
func f4(x, y int) int {
	return x + y
}

// 可变参数
func f5(title string, y ...int) int {
	fmt.Println(y) // y是一个int类型的切片
	return 1
}

// 命名返回值
func f6(x, y int) (sum int) {
	sum = x + y // 如果使用命名的返回值，那么在函数中可以直接使用返回值变量
	return      // 如果使用命名的返回值,return后面可以省略返回值变量
}

// Go语言中支持多个返回值
func f7(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return
}

func main() {
	f1()
	f2("理想")
	f2("姬无命")
	fmt.Println(f3(100, 200)) // 调用函数

	ret := f3(100, 200)
	fmt.Println(ret)

	f5("lixiang", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// 在一个命名的函数中不能够再声明命名函数
	// func f8(){}
}
```



## defer

### defer基本定义

```go
package main

import "fmt"

// defer
// defer多用于函数结束之前释放资源（文件句柄，数据库连接，socket连接）

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("aaaaaaa") //defer把它后面的语句延迟到函数即将返回的时候在执行
	defer fmt.Println("bbbbbbb") //一个函数中可以有多个defer语句
	defer fmt.Println("ccccccc") //多个defer语句按照先进后出的顺序延迟执行
	fmt.Println("end")
}

// 帮助理解
// Go解释器从上往下执行
// 1.定义f1这个函数，并且定义返回值是int类型的x变量，默认值为0
// 2.fmt.Println(f1()) 先执行括号里的f1()函数，调用f1()这个函数
// 3.定义 defer func(x int){}(x)，并调用，但是由于defer特性，此时没有执行函数体里面的语句，目前x=0，x++ 是1
// 4.执行return语句，把5赋值给要返回的变量x，此时x=5 。 执行defer，里面的x=1，并且打印.执行RET返回
// 5.执行fmt.Println(5)语句打印
func f1() (x int) {
	defer func(x int) {
		x++
		fmt.Println(x) //1
	}(x) //0
	return 5
}

func main() {
	deferDemo()
	fmt.Println("------分割线----------")
	fmt.Println(f1())
}
```

### defer练习

```go
package main

import "fmt"

// Go语言中函数的return不是原子操作，在底层是分为两步来执行
// 第一步：返回值赋值
// defer
// 第二步：真正的RET返回
// 1.返回值赋值 2.defer 3.真正的return指令
// 函数中如果存在defer，那么defer执行的时机是在第一步和第二步之间
// 如果不能理解，可以参考这篇博文
// https://www.jianshu.com/p/79c029c0bd58

func main() {
	fmt.Println(f1()) //5
	fmt.Println(f2()) //6
	fmt.Println(f3()) //5
	fmt.Println(f4()) //5
	fmt.Println(f5()) //5
	fmt.Println(f6()) //6
}

// 匿名返回值
// 注意这里：返回变量是匿名的（举例比如是：niming这个变量），而不是x。
// 1.返回值赋值： x=5 把值5赋值给niming这个变量
// 2.defer：x++ x=6，这里是x这个变量++，而不是niming这个变量++
// 3.真正的return指令: 返回niming这个变量，此时niming是等于5的
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// 返回值命名和变量命名相同
// 1.返回值赋值：返回变量是x，命名的,此时x=5的
// 2.defer x++ ,x这个变量++，此时是6
// 3.真正的return指令:返回x这个变量，此时x是=6
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 //返回值变量是x
}

// 返回值命名和变量命名不相同
// 1.返回值赋值：返回变量是y，命名的，此时y=5
// 2.defer: x++  x=6
// 3.真正的return指令:返回y这个变量，此时y=5
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// 在golang中函数的参数默认为 按值传递，即在函数内部修改传入参数的值是函数外部传入值的 拷贝
// 如果想要使用引用传递，需要将传入的参数设置为 指针类型
// 如果传入的参数数据很大，建议使用指针类型，减少内存因拷贝参数而占用
func f4() (x int) {
	defer func(x int) {
		x++                  //改变的是函数中x的副本
		fmt.Println("==", x) //1
	}(x) //x=0 值传递
	return 5 //返回值 =x =5
}

//
func f5() (x int) {
	fmt.Println("f5 x:", x)
	defer func(x int) int {
		fmt.Println("f5 defer=", x)
		x++
		fmt.Println("f5 defer==", x)
		return x
	}(x) //x=0
	return 5
}

// 1.返回值=x=5 2.defer x=6 3.RET返回 x=6
// 传一个x的指针到匿名函数中
func f6() (x int) {
	fmt.Println("f6 x:", x)
	defer func(x *int) {
		fmt.Println("defer f6 x: ", *x)
		*x++
		fmt.Println("defer f6 xx: ", *x)
	}(&x)
	return 5
}
```

```go
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
```



## 变量分类

```go
package main

import (
	"fmt"
)

/*
变量分类：
1.全局变量：全局变量是定义在函数外部的变量，它在程序整个运行周期内都有效。 在函数中可以访问到全局变量
2.局部变量
	函数内定义的变量
	语句块定义的变量
如果局部变量和全局变量重名，优先访问局部变量

变量的作用域
1.先在函数内部查找
2.找不到就往函数的外面查找，一直找到全局
3.如果全局也找不到，那么就报错了

函数内部定义的变量只能在该函数内部使用
*/

// 定义一个全局变量
var x = 100

// 定义一个函数
func f1() {
	//x := 10
	name := "lixiang"
	fmt.Println(x, name)
}

func main() {
	f1()
	//fmt.Println(name) //函数内部定义的变量只能在改函数内部使用

	//语句块作用域
	if i := 10; i < 18 {
		fmt.Println("读书")
	}
	//不能使用if语句中i变量，if语句中的i变量只能在改if语句块中使用，同理还有for循环也一样
	//fmt.Println(i)
}
```



## 函数类型的变量

```go
package main

import "fmt"

// 1.函数类型的变量
// 定义函数类型
type calculation func(int, int) int //语句定义了一个calculation类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值

// 满足这个条件的函数都是calculation类型的函数
func add(x, y int) int {
	return x + y
}

func f1() {
	fmt.Println("shahe")
}

func f2() int {
	return 10
}

func f4(x, y int) int {
	return x + y
}

// 函数也可以作为参数的类型
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

// 函数还可以作为返回值
func ff(a, b int) int {
	return a + b
}

func f5(x func() int) func(int, int) int {
	return ff
}

func main() {
	// 1.函数类型的变量
	var c calculation
	c = add
	ret := c(1, 2)
	fmt.Println("ret:", ret) //ret: 3
	a := f1
	fmt.Printf("%T\n", a) //func()
	b := f2
	fmt.Printf("%T\n", b) //func() int

	// 2.函数作为参数
	f3(f2) //10
	f3(b)  //10

	// f3(f4) //这里f4不能作为参数传给f3，因为f4的类型为func(int, int) int，而f3可以接受的函数类型为func() int
	fmt.Printf("%T\n", f4) //func(int, int) int

	// 3.函数还可以作为返回值
	f7 := f5(f2)           //ff函数
	fmt.Printf("%T\n", f7) //func(int, int) int
}
```



## 匿名函数

```go
package main

import (
	"fmt"
)

/*
匿名函数：匿名函数就是没有函数名的函数

函数当然还可以作为返回值，但是在Go语言中函数内部不能再像之前那样定义函数了，只能定义匿名函数。

func(参数)(返回值){
    函数体
}

匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量或者作为立即执行函数

匿名函数多用于实现回调函数和闭包
*/

var f2 = func(x, y int) {
	fmt.Println(x + y)
}

func main() {
	// 函数内部没有办法声明带名字的函数，但是可以声明匿名函数
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(10, 20)

	//声明匿名函数，并且直接调用
	func(x, y int) {
		fmt.Println(x + y)
		fmt.Println("hello world")
	}(100, 200)

	f2(1, 2)
}
```



## 闭包

### 闭包概念

```go
package main

import "fmt"

// 闭包
// 闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境

// 功能需求：把f2当做参数传入到f1  f1(f2)
// 解题思路：直接传递肯定不行，可以借助一个闭包函数

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

// 闭包函数
func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		f(x, y)
	}
	return tmp
}

func main() {
	ret := f3(f2, 100, 200) //把原来需要传递两个int类型的参数 包装成一个不需要传递参数的函数
	fmt.Printf("%T\n", ret) //func()
	f1(ret)
}
```



```go
package main

import "fmt"

// 闭包是什么？
// 闭包是一个函数，这个函数包含了它外部作用域的一个变量

// 底层原理
// 1.函数可以作为返回值
// 2.函数内部查找变量的顺序，先在自己内部找，找不到就往外层找

func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	// 变量ret是一个函数,并且它引用了其外部作用域中的x变量，此时ret就是一个闭包。
	// 在ret的生命周期内，变量x也一直有效
	var ret = adder2(100)
	fmt.Println(ret(200)) //300
	fmt.Println(ret(300)) //600
}
```



### 闭包练习

```go
package main

import (
	"fmt"
	"strings"
)

// 利用闭包实现的功能：增加指定的后缀。不存在则增加，存在则不改动

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	jpgFunc := makeSuffixFunc(".jpg")
	fmt.Println(jpgFunc("test1"))     //test1.jpg
	fmt.Println(jpgFunc("test2.jpg")) //test2.jpg

	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(txtFunc("a.txt")) //a.txt
	fmt.Println(txtFunc("b"))     //b.txt
}
```



```go
package main

import "fmt"

// calc接收一个基础数值，返回两个函数
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

func main() {
	f1, f2 := calc(10)

	// 变量f1,f2是一个函数,并且它引用了其外部作用域中的base变量，此时f1,f2就是一个闭包。
	// 在f1,f2的生命周期内，变量base也一直有效
	fmt.Println(f1(1), f2(2)) //11 9 此时base=9
	fmt.Println(f1(3), f2(4)) //12 8 此时base=8
	fmt.Println(f1(5), f2(6)) //13 7
}
```



## 递归函数

```go
package main

import "fmt"

// 递归：函数自己调用自己
// 递归适合处理那种问题相同\问题的规模越来越小的场景
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
// n个台阶，一次可以走1步，也可以走2步，有多少种走法
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
	fmt.Println(ret)
}
```



## 异常处理

```go
package main

import "fmt"

// panic 和 recover
// 异常处理

// 注意
// 1.recover()必须搭配defer使用
// defer一定要在可能引发panic的语句之前定义。

func funcA() {
	fmt.Println("this is funcA")
}

func funcB() {
	// 刚刚打开数据库连接
	fmt.Println("开始连接数据库...")
	defer func() {
		err := recover()
		fmt.Println("err: ", err)
		fmt.Println("释放数据库连接...")
	}()
	panic("程序出现了严重的BUG!!!!") //程序崩溃退出
	fmt.Println("this is funcB")
}

func funcC() {
	fmt.Println("this is funcC")
}

func main() {
	funcA()
	funcB()
	funcC()
}
```

```go
package main

import (
	"fmt"
	"time"
)

// 不管程序中的哪个goroutine发生panic，如果没有recover处理，那么程序都会退出

func hello() {
	//defer func() {
	//	err := recover()
	//	fmt.Println(err)
	//}()
	for i := 0; i < 3; i++ {
		fmt.Println("hello")
		time.Sleep(time.Second * 1)
	}
	panic("hello挂了")

}

func main() {
	go hello()
	for {
		fmt.Println("我是主goroutine,我还没挂")
		time.Sleep(time.Second * 1)
	}
}
```



## 练习题

```go
package main

import (
	"fmt"
	"strings"
)

/*
你有5000枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币

写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/

var (
	coins        = 5000
	users        = []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	distribution = make(map[string]int, len(users))
	rule         = map[string]int{
		"e": 1,
		"i": 2,
		"o": 3,
		"u": 4,
	}
)

func main() {
	left := dispatchCoin()
	fmt.Println("剩下的金币个数：", left)
	// 输出结果循环打印
	for name, value := range distribution {
		fmt.Printf("姓名:%s \t\t金币:%d\n", name, value)
	}
}

func dispatchCoin() (left int) {
	// 1. 依次拿到每个人的名字
	// 2. 拿到一个人名根据分金币的规则去分金币,
	// 2.1 每个人分的金币数应该保存到 distribution 中
	// 2.2 还要记录下剩余的金币数
	// 3. 整个第2步执行完就能得到最终每个人分的金币数和剩余金币数

	// 循环用户名
	for _, name := range users {
		// 对每个用户进行规则循环，并且统计每个用户金币数量
		for ruleKey, ruleValue := range rule {
			count := strings.Count(strings.ToLower(name), ruleKey)
			if count > 0 {
				distribution[name] += count * ruleValue
				coins -= count * ruleValue
			}
		}
	}
	return coins
}

```
