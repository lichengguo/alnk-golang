# Map数据类型和指针



## Map数据类型

### Map基本概念

```go
package main

import "fmt"

// map

// make()函数和new()函数的区别
// make和new都是用来申请内存的
// new很少用，一般用来给基本数据类型申请内存，`string`、`int`,返回的是对应类型的指针(*string、*int)
// make是用来给`slice`、`map`、`chan`申请内存的，make函数返回的的是对应的这三个类型本身

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)        //true 还没初始化，没有在内存中开辟空间
	m1 = make(map[string]int, 10) //要估算好改map的容量，避免在程序运行期间再动态扩容，影响性能
	fmt.Println(m1 == nil)        //false 已经在内存中开辟空间了

	m1["age"] = 18
	m1["salary"] = 2000
	fmt.Println(m1)
	fmt.Println(m1["salary"])

	// 不存在的键
	// 约定成俗用ok接收返回的布尔值
	fmt.Println(m1["tom"]) //如果不存在这个key，则拿到对应值类型的零值 0
	value, ok := m1["tom"]
	if !ok {
		fmt.Println("没有此key")
	} else {
		fmt.Println(value)
	}

	// map的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// 只遍历key
	for k := range m1 {
		fmt.Println(k)
	}

	// 只遍历value
	for _, v := range m1 {
		fmt.Println(v)
	}

	// 删除
	delete(m1, "salary")
	fmt.Println(m1)
	delete(m1, "tom") //删除不存在的key，程序也不会报错
}

```



### 练习

```go
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 让map按照key排序打印

func main() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 100) //定义一个map，key为string类型，值为int类型

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0-99的随机整数
		scoreMap[key] = value
	}
	fmt.Println(scoreMap, len(scoreMap))

	// 取出map中所有的key存入切片keys
	var keys = make([]string, 0, 100)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//fmt.Println(keys)
	//sort.Ints() //按照整型排序
	sort.Strings(keys) //按照字符串排序

	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

```



### map和slice切片组合练习

```go
package main

import (
	"fmt"
)

// map和slice组合

func main() {
	// 元素类型为map的切片
	var s1 = make([]map[string]int, 10, 10) //切片的元素类型为map，长度10，容量10，切片已经初始化，但是里面的map类型没有初始化
	fmt.Println(s1)                         //[map[] map[] map[] map[] map[] map[] map[] map[] map[] map[]]

	// s1[0]的map类型初始化
	s1[0] = make(map[string]int, 10)
	fmt.Printf("%T %d\n", s1[0], len(s1[0]))  //map[string]int 0
	s1[0]["沙河"] = 10                          //s1[0]表示map，s1[0]["沙河"]表示map的key为沙河
	fmt.Printf("s1:%T s1[0]:%T\n", s1, s1[0]) //s1:[]map[string]int  s1[0]:map[string]int 注意这里s1是切片;s1[0]是map
	fmt.Println(s1)                           //[map[沙河:10] map[] map[] map[] map[] map[] map[] map[] map[] map[]]

	// s1[1]的map初始化
	s1[1] = make(map[string]int)
	s1[1]["kobe"] = 24
	s1[1]["duke"] = 21
	fmt.Println(s1) //[map[沙河:10] map[duke:21 kobe:24] map[] map[] map[] map[] map[] map[] map[] map[]]

	// 值为切片类型的map
	var m1 = make(map[string][]int) //key是string类型，值是[]int int切片类型
	m1["北京"] = []int{1, 2, 3}
	m1["上海"] = []int{4, 5}
	fmt.Println(m1) //map[上海:[4 5] 北京:[1 2 3]]
}

```



## 指针

```go
package main

import "fmt"

// point 指针
// go语言中不会直接进行指针的运行，因此只要会使用 & * 基本上就够了

// make和new的区别
// make和new都是用来申请内存的
// new很少用，一般用来给基本数据类型申请内存，`string`、`int`,返回的是对应类型的指针(*string、*int)
// make是用来给`slice`、`map`、`chan`申请内存的，make函数返回的的是对应的这三个类型本身

func main() {
	// 1. &:取内存地址
	n := 18
	p := &n               //&n 取n变量在内存中的地址，并且赋值给p变量
	fmt.Println(p)        //0xc00001a070 n变量在内存中的地址
	fmt.Println(&n)       //0xc00001a070 n变量在内存中的地址
	fmt.Println(*p)       //取值 18
	fmt.Println(*(&n))    //18
	fmt.Printf("%T\n", p) //*int：int类型的指针
	fmt.Println(&p)       //0xc00000e028 p变量自己在内存中的地址
	fmt.Println(*(&p))    //0xc00001a070 取p变量指向的值，即n的内存地址

	// 2. *:根据地址取值
	m := *p                //这个是取出p变量中存放的n的内存地址所指向的值
	fmt.Println(m)         //18
	fmt.Printf("%T\n", m)  //int
	pn := &p               //&p p变量自己在内存中的地址
	fmt.Printf("%p\n", pn) //0xc00008c018 p变量自己在内存中的地址

	fmt.Println("----------------------------------")
	var a1 *int     //nil pointer 未申请内存地址
	fmt.Println(a1) //<nil>

	var a2 = new(int) //new函数申请一块内存地址
	fmt.Println(a2)   //0xc00001a078 这是一块内存地址
	fmt.Println(*a2)  //获取这块内存地址指向的值是多少 //0
	*a2 = 100         //给这块内存地址赋值
	fmt.Println(*a2)  //100
	fmt.Println(&a2)  //0xc0000ae028 a2在内存中的地址
}

```

```go
package main

import "fmt"

/*
[pointer指针]

Go语言中的指针不能进行偏移和运算，是安全指针

要搞明白Go语言中的指针需要先知道3个概念：指针地址、指针类型和指针取值


[Go语言中的指针]
任何程序数据载入内存后，在内存都有他们的地址，这就是指针.而为了保存一个数据在内存中的地址，我们就需要指针变量

比如，“永远不要高估自己”这句话是我的座右铭，我想把它写入程序中，程序一启动这句话是要加载到内存（假设内存地址0x123456），
我在程序中把这段话赋值给变量A，把内存地址赋值给变量B。
这时候变量B就是一个指针变量。
通过变量A和变量B都能找到我的座右铭

Go语言中的指针不能进行偏移和运算，因此Go语言中的指针操作非常简单，我们只需要记住两个符号：
&（取地址）和 *（根据地址取值）

总结： 取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值

变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
	对变量进行取地址（&）操作，可以获得这个变量的指针变量。 //a := 10  b := &a
	指针变量的值是指针地址。 //b的值是指针地址
	对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。  // *b  的值是10


[new和make]
在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。
而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。

要分配内存，就引出来今天的new和make
Go语言中new和make是内建的两个函数，主要用来分配内存

new
new是一个内置的函数，它的函数签名如下：
func new(Type) *Type
Type表示类型，new函数只接受一个参数，这个参数是一个类型
*Type表示类型指针，new函数返回一个指向该类型内存地址的指针
new函数不太常用，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值
func main() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(*b)       // false
}

make
make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，
因为这三种类型就是引用类型，所以就没有必要返回他们的指针了
make函数是无可替代的，我们在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作

new与make的区别
	二者都是用来做内存分配的
	make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身
	new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针
*/

func main() {
	var a int
	a = 100
	b := &a //b是获取了a的内存地址  指针
	//fmt.Printf("type  a:%T  type b:%T\n", a, b)
	//fmt.Println("b:", b) //0xc000018080

	//
	fmt.Printf("%p\n", &a) //a的内存地址或者说指针 	0xc000018080
	fmt.Printf("%p\n", b)  //a的内存地址或者说指针 	0xc000018080

	fmt.Printf("%#v\n", b) //b的值 (*int)(0xc000018080) b这个变量存储的值就是a的内存地址
	fmt.Printf("%p\n", &b) //b自己本身内存地址	0xc0000ae018
	fmt.Printf("%v\n", *b) //100 查找b这个变量的值（a的内存地址）   所指向的值
}

```

## 小练习

写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中 how=1 do=2 you=1

```go
package main

import (
	"fmt"
	"strings"
)


// 写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1

func main() {
	s1 := "how do you do"
	s2 := strings.Split(s1, " ") //返回 []string 切片

	// 定义一个map并且初始化
	var m1 = make(map[string]int, 100) //初始化以后int类型的零值为0

	for _, key := range s2 {
                //m1[key] = m1[key] + 1
		m1[key]++
	}
	fmt.Println(m1)
}

```


回文判断

```go
package main

import "fmt"

func main() {
	// 回文判断 字符串从左往右读和从右往左读是一样的，那么就是回文。
	// 上海自来水来自海上 s[0]  s[len(s)-1]
	// 山西运煤车煤运西山

	ss := "a山西运煤车煤运西山a"
	r := make([]rune, 0, 255) //初始化 []rune切片 长度0，容量255

	// 把字符串中的字符拿出来放到一个切片中
	for _, c := range ss {
		r = append(r, c)
	}
	// 此时的切片r: [97 23665 35199 36816 29028 36710 29028 36816 35199 23665 97]

	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Printf("字符串 [%s] 不是回文\n", ss)
			return
		}
	}

	fmt.Printf("字符串 [%s] 是回文\n", ss)
}

```
