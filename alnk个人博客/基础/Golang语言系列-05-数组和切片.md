# 数组和切片

## 数组

### 概念

数组是同一种数据类型元素的集合；数组的长度必须是常量，并且长度是数组类型的一部分，一旦定义，长度不能变

例如：[5]int 和 [10]int 是不同的数组类型

使用时可以修改数组成员，但是数组大小长度不可变化



### 数组的初始化

```go
package main

import (
	"fmt"
)

func main() {
	var a1 [3]bool
	var a2 [4]bool
	var a122 [3]func()
	fmt.Printf("a1:%T a2:%T a122:%T\n", a1, a2, a122) //a1:[3]bool a2:[4]bool a122:[3]func()

	// 数组的初始化
	// 如果不初始化：默认元素都是零值（布尔值：false，整型和浮点型都是0，字符串为空:""，函数为nil）
        // a1:[false false false] a2:[false false false false] a122:[<nil> <nil> <nil>]
	fmt.Printf("a1:%v a2:%v a122:%v\n", a1, a2, a122) 
  
	// 初始化的方式
	// 1.初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1) //[true true true]

	// 2.初始化方式2：根据初始值自动推断数组的长度是多少
	a10 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(a10) //[1 2 3 4 5 6 7 8]

	// 3.初始化方式3：根据索引来初始化
	// a3 := [...]int{0:1, 4:2}  //索引为0的值为1，索引为4的值为2
	a3 := [5]int{0: 1, 4: 2} //索引为0的值为1，索引为4的值为2
	fmt.Println(a3)          //[1 0 0 0 2]
}
```



数组的遍历

```go
package main

import (
	"fmt"
)

func main() {
        // 数组的遍历
	citys := [...]string{"北京", "上海", "深圳"}
	// 1.根据索引遍历
	// 用 len() 函数统计数组的长度然后在进行 for 循环是可以的。这里和string类型的遍历有差别
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	// 2.for range遍历
	for _, v := range citys {
		fmt.Println(v)
	}
}
```



### 多维数组

```go
package main

import (
	"fmt"
)

func main() {
        // 多维数组
        // 注意:多维数组只有第一层可以使用 [...]int{1, 2, 3} 来让编译器推导数组长度
	// [ [1 2] [3 4] [5 6] ]
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11) //[[1 2] [3 4] [5 6]]
	a12 := [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a12) //[[1 2] [3 4] [5 6]]

	// 多维数组遍历
	// 1 2 3 4 5 6
	for _, v1 := range a11 {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
}
```



### 数组是值类型

```go
package main

import (
	"fmt"
)

func main() {
        // 数组是值类型
	// 赋值和传参会复制整个数组，因此改变副本的值，不会改变原来数组本身的值
	fmt.Println("------ 数组是值类型 ------")
	b1 := [3]int{1, 2, 3} //[1 2 3]
	b2 := b1              //[1 2 3] Ctrl+C Ctrl+V 相当于把world文档从文件夹A拷贝到文件夹B
	b2[0] = 100           //b2:[100 2 3]
	//b1[0] = 2 //数组可以修改值，但是长度不能变
	fmt.Println(b1, b2) //b1:[1 2 3]  b2:[100 2 3]

	// 数组支持 “==“、”!=” 操作符，因为内存总是被初始化过的
	arr1 := [...]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3}
	fmt.Println(arr1 == arr2) //true
}


```



### 指针数组和数组的指针

```go
package main

import (
	"fmt"
)

func main() {
        // 指针数组
	// [n]*T表示指针数组
	s1 := "abc"
	s2 := &s1
	s3 := &s1
	arr3 := [...]*string{s2, s3}
	fmt.Println(arr3) //[0xc00008e240 0xc00008e240]

	// 数组的指针
        // *[n]T表示数组的指针
	arr4 := [...]int{1, 2}
	op := &arr4
	fmt.Println(*op)          //[1 2]
	fmt.Printf("%T\n", *op)   //[2]int
	fmt.Printf("%p\n", op)    //0xc0000180e0  op的值是数组arr4这个变量的指针
	fmt.Printf("%p\n", &op)   //0xc00000e030  op自己在内存中的地址
	fmt.Printf("%p\n", &arr4) //0xc0000180e0  op的值是数组arr4这个变量的指针
}
```



### 数组练习

```go
package main

import "fmt"

// array数组练习题
// 1.求数组[1, 3, 5, 7, 8]所有元素的和
// 2.找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)

func main() {
	// 1
	a1 := [...]int{1, 3, 5, 7, 8}
	sum1 := 0
	for _, v := range a1 {
		sum1 += v
	}
	fmt.Println("sum1:", sum1)

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

```





## 切片

### 概念

切片是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容

切片是一个引用类型(数组是一个值类型)，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合

内置的`len()`函数求切片长度，内置的`cap()`函数求切片的容量

切片的底层就是一个数组，所以我们可以基于数组通过切片表达式得到切片

判断切片是否为空 请始终使用`len(s) == 0`来判断，而不应该使用s == nil来判断 初始化以后的切片 != nil

```go
package main

import (
	"fmt"
)

func main() {
	// 切片的定义
	var s1 []int                  //定义一个存放int类型元素的切片
	var s2 []string               //定义一个存放string类型元素的切片
	fmt.Println(s1, s2)           //[] []
	fmt.Println(s1 == nil)        //true
	fmt.Println(s2 == nil)        //true
	fmt.Println(len(s1), len(s2)) // 0 0
  
        // 切片初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"沙河", "张江", "平山村"}
	fmt.Println(s1, s2)           //[1 2 3] [沙河 张江 平山村]
	fmt.Println(s1 == nil)        //false
	fmt.Println(s2 == nil)        //false
	fmt.Println(len(s1), len(s2)) // 3 3
  
        // 长度和容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1)) //len(s1):3 cap(s1):3
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2)) //len(s2):3 cap(s2):3
  
        // 由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13} //数组
	s3 := a1[0:4]                         //基于一个数组切割，顾首不顾尾
	fmt.Println(s3)                       //[1 3 5 7]
	s4 := a1[1:6]
	fmt.Println(s4) //[3 5 7 9 11]
	s5 := a1[:4]
	s6 := a1[3:]
	s7 := a1[:]
	fmt.Println(s5) //[1 3 5 7]
	fmt.Println(s6) //[7 9 11 13]
	fmt.Println(s7) //[1 3 5 7 9 11 13]
        // 切片的容量：是指底层数组的容量
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5)) // len(s5):4 cap(s5):7
	// 切片容量：底层数组从切片的第一个元素到最后的元素数量
	fmt.Printf("len(s6):%d cap(s6):%d\n", len(s6), cap(s6)) // len(s6):4 cap(s6):4
	// 切片再切割
	s8 := s6[3:]                                            //[13]
	fmt.Printf("len(s8):%d cap(s8):%d\n", len(s8), cap(s8)) //len(s8):1 cap(s8):1
  
        // 切片是引用类型，都指向了底层的一个数组
	// 修改底层数组的值，会影响切片
	fmt.Println("s6: ", s6) //[7 9 11 13]
	a1[6] = 1300            //修改底层数组的值
	fmt.Println("s6: ", s6) //[7 9 11 1300]
	fmt.Println("s8: ", s8) //[1300]
	fmt.Println("a1: ", a1) //[1 3 5 7 9 11 1300]
  
        // 修改切片，会修改底层数组吗?如果只是修改值的话不涉及到扩容，是会修改原底层数组的
	s8[0] = 10000
	fmt.Println(s8)           //[10000]
	fmt.Println(a1)           //[1 3 5 7 9 11 10000] a1数组发生改变
	s8 = append(s8, 20000)    //此处已经产生了新的数组，切片s8的容量是1，s8的切片不再指向原来的底层数组
	fmt.Printf("s8:%v\n", s8) //[10000 20000]
	fmt.Println(a1)           //[1 3 5 7 9 11 10000] a1数组不发生改变
  
        // append追加数组元素,如果底层数组容量不够的时候会扩容，产生新的数组
	a10 := [5]int{1, 2, 3}
	s10 := a10[:]
	fmt.Printf("len(s10):%d cap(s10):%d\n", len(s10), cap(s10)) //len(s10):5 cap(s10):5
	//s10 = append(s10, 4)
	//fmt.Println(cap(s10)) //10
	s10[3] = 1
	//fmt.Println(a10, s10) //[1 2 3 0 0] [1 2 3 1 0 4] 如果没注释s10的append，证明Go已经把底层的数组换了
	fmt.Println(a10, s10) //[1 2 3 1 0] [1 2 3 1 0] 如果注释了s10的append，则s10[3]=1修改的是底层元素
}
```



### make函数初始化切片

```go
package main

import (
	"fmt"
)

func main() {
        // make函数初始化切片
	var a = make([]string, 5, 10) // make(类型，长度，容量)
	fmt.Printf("cap:%d %#v\n", cap(a), a) //cap:10 []string{"", "", "", "", ""}
	
        // append()
        for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i)) // 拼接成字符串
	}
	fmt.Printf("%#v\n", a)
        //[]string{"", "", "", "", "", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
  
	fmt.Println(len(a), cap(a)) //15 20
  
  
        // 初始化
	s1 := make([]int, 5, 10)                                          //make(类型，长度，容量)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1)) //s1=[0 0 0 0 0] len(s1)=5 cap(s1)=10

	s2 := make([]int, 0, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s2, len(s2), cap(s2)) //s1=[] len(s1)=0 cap(s1)=10
  
        // 切片的赋值
	s3 := []int{1, 3, 5}
	s4 := s3            //s3和s4都指向了同一个底层数组
	fmt.Println(s3, s4) //[1 3 5] [1 3 5]
	s3[0] = 1000
	fmt.Println(s3, s4) //[1000 3 5] [1000 3 5]
}
```



### 切片的遍历

```go
package main

import (
	"fmt"
)

func main() {
        // 切片的遍历
        s3 := []int{1, 3, 5}
  
	// 1.索引遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}

	// 2.for range循环
	for _, v := range s3 {
		fmt.Println(v)
	}
}
```



### 切片的append()

```go
package main

import (
	"fmt"
)

// append()为切片追加元素
// 可能会导致数组扩容，从而让切片指向新的数组

func main() {
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1)) //s1=[北京 上海 深圳] len(s1)=3 cap(s1)=3
	//s1[3] = "guangzhou" //panic: runtime error: index out of range [3] with length 3

	// 调用append函数必须用原来的切片变量接收返回值
	// append追加元素，原来的底层数组放不下的时候，Go就会把底层数组换一个
	// 必须用变量接收append的返回值
	s1 = append(s1, "广州")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1)) 
        //s1=[北京 上海 深圳 广州] len(s1)=4 cap(s1)=6

	s1 = append(s1, "杭州", "成都")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1)) 
        //s1=[北京 上海 深圳 广州 杭州 成都] len(s1)=6 cap(s1)=6

	ss := []string{"武汉", "西安", "苏州"}
	s1 = append(s1, ss...)                                            
        //表示拆开(打散)，一个一个的元素添加到别的切片
	
        fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1)) 
        //s1=[北京 上海 深圳 广州 杭州 成都 武汉 西安 苏州] len(s1)=9 cap(s1)=12
}
```



### 切片的拷贝

```go
package main

import (
	"fmt"
)

// 切片的copy

func main() {
	a1 := []int{1, 3, 5}
	aa := a1[:1]
	fmt.Println(cap(aa), len(aa), aa) // 容量是3 长度是1 [1]
	a2 := a1                          //赋值 a2 a1 指向的是同一个底层数组
	var a3 = make([]int, 3, 3)
	copy(a3, a1)            //copy 底层数组复制了一份，a3 和 a1 指向的不是同一个底层数组
	fmt.Println(a1, a2, a3) //[1 3 5] [1 3 5] [1 3 5]
	a1[0] = 100
	fmt.Println(a1, a2, a3) //[100 3 5] [100 3 5] [1 3 5]

	// 将a1中的索引为1的 3 这个元素删掉。 ... 由于切片没有直接删除元素的方法，所以可以采用这种方法
	a1 = append(a1[:1], a1[2:]...) // a1[:1] 容量是3，append 1 个元素的时候，底层数组没有发生变化
	fmt.Println(a1)                // [100 5]
	fmt.Println(cap(a1))           // 3
	fmt.Println(a2)                // [100 5 5]

	x1 := [...]int{1, 3, 5} //数组
	fmt.Println("===", x1)
	s1 := x1[:] //数组经过[L:M]以后，可以得到切片
	fmt.Println(s1, len(s1), cap(s1))
	// 1.切片不保存具体的值
	// 2.切片对应一个底层数组
	// 3.底层数组都是占用一块连续的内存
	fmt.Printf("%p\n", &s1[0])
	s1 = append(s1[:1], s1[2:]...)    //s1[:1] 容量是3，append 1 个元素的时候，底层数组没有发生变化
	fmt.Printf("%p\n", &s1[0])        //Go语言中不存在指针操作，只需要记住两个符号：&:取地址 *:根据地址取值
	fmt.Println(s1, len(s1), cap(s1)) //[1 5] 2 3
	fmt.Println(x1)                   //[1 5 5]
	s1[0] = 100
	fmt.Println(x1) //[100 5 5]
}

```



### 切片练习

```go
package main

import (
	"fmt"
	"sort"
)

// 切片练习题

func main() {
	var a = make([]int, 5, 10) //初始化切片，长度5，容量10
	fmt.Println(a)             //[0 0 0 0 0]
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)      //[0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
	fmt.Println(cap(a)) //20

	var a1 = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a1[:]) //对切片进行排序
	fmt.Println(a1)  //[1 3 7 8 9]

	// 要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断
	var aa []int           //只定义，没初始化
	fmt.Println(aa == nil) //true
	var bb []int           //定义
	bb = []int{}           //初始化后就不为nil了，但是还是空的切片
	fmt.Println(bb == nil) //false
	fmt.Println(aa, bb)    //[] []
	if len(bb) == 0 {
		fmt.Println("空数组")
	}
}

```

```go
package main

import "fmt"

//切片练习

func main() {
	//ex1
	//s1 := make([]int, 4, 100)
	//s1[0] = 1
	//s1[1] = 2
	//fmt.Println(s1) //[1 2 0 0]
	//s2 := s1[:]     //再切片，切片的切片，底层数组还是同一个
	////s2 := s1 //赋值 底层数组是同一个
	//s2[0] = 100
	//fmt.Println(s1)               //[100 2 0 0]
	//fmt.Println(s2)               //[100 2 0 0]
	//fmt.Println(cap(s1), cap(s2)) //100 100

	//ex2
	//s1 := make([]int, 4, 4)
	//s1[0] = 1
	//s1[1] = 2
	//fmt.Println(s1) //[1 2 0 0]
	//s2 := s1[:]     //切片的再切片 len:4 cap:4
	//
	////append()扩容以后会形成新的数组
	//s2 = append(s2, 100) //形成新的底层数组 [1 2 0 0 100] len:4 cap:8
	//
	//fmt.Println(s1)               //[1 2 0 0]
	//fmt.Println(s2)               //[1 2 0 0 100]
	//fmt.Println(cap(s1), cap(s2)) //4 8

	//ex3
	//s1 := make([]int, 4, 4)
	//s1[0] = 1
	//s1[1] = 2
	//fmt.Println(s1)                         //[1 2 0 0]
	//s2 := s1[:2]                            //[1 2] len:2 cap:4
	//fmt.Println("==", s2, len(s2), cap(s2)) //[1 2] 2 4
	//s2 = append(s2, 100)                    //append()这里没有扩容，因为长度是2，容量是4，所以不会形成新的底层数组,此时底层数组是:[1 2 100]
	//fmt.Println(s1)                         //[1 2 100 0]
	//fmt.Println(s2)                         //[1 2 100] 这里s2的长度是3，不是4哦，所以是 [1 2 100]
	//fmt.Println(cap(s1), cap(s2))           //4 4
	//fmt.Println(len(s1), len(s2))           //4 3

	//ex4
	//var s1 = [...]int{1, 2, 3, 4, 5}
	//s2 := s1[2:] //[3 4 5] len:3 cap:3
	//
	//s2 = append(s2, 100) //扩容了，产生新的底层数组 [3 4 5 100] len:4 cap:6
	//
	//fmt.Println(s1)               //[1 2 3 4 5]
	//fmt.Println(s2)               //[3 4 5 100]
	//fmt.Println(cap(s1), cap(s2)) //5 6
	//fmt.Println(len(s1), len(s2)) //5 4

	//ex5
	//var arr = [...]int{1, 2, 3, 4, 5}
	//slice := arr[:2]                  //[1 2] len:2 cap:5
	//slice = append(slice, 6, 7)       //没扩容 切片是:[1 2 6 7] len:4 cap:5
	//slice[0] = 100                    //[100 2 6 7]
	//fmt.Println(arr)                  //[100 2 6 7 5]
	//fmt.Println(slice)                //[100 2 6 7]
	//fmt.Println(cap(arr), cap(slice)) //5 5
	//fmt.Println(len(arr), len(slice)) //5 4

	//ex6
	var arr = [...]int{1, 2, 3, 4, 5}
	slice := arr[:2] // [1 2] len:2 cap:5

	slice = append(slice, 6, 7, 8, 9, 10) //底层数组扩容了，切片是[1 2 6 7 8 9 10]
	slice[0] = 100                        //[100 2 6 7 8 9 10]

	fmt.Println(arr)                  //[1 2 3 4 5]
	fmt.Println(slice)                //[100 2 6 7 8 9 10]
	fmt.Println(cap(arr), cap(slice)) //5 10
	fmt.Println(len(arr), len(slice)) //5 7
}

```
