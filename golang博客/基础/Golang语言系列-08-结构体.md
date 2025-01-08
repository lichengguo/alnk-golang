## 结构体

### 自定义类型和类型别名的区别

```go
package main

import "fmt"

// 自定义类型和类型别名的区别

// type后面跟的是类型
type myInt int    //自定义类型
type youInt = int //类型别名

func main() {
  // 自定义类型
	var n myInt
	n = 100
	fmt.Println(n)              //100
	fmt.Printf("%T %d\n", n, n) //main.myInt 100

  // 类型别名
	var m youInt
	m = 100
	fmt.Println(m)        //100
	fmt.Printf("%T\n", m) //int

  //
	var c rune
	c = '中'
	fmt.Println(c)        //20013 unicode编码十进制
  // unicode编码
  // 首先把20013转为十六进制0x4E2D
  // 然后把十六进制转为中文
  // unicode编码十六进制表示方式 \u4E2D
	fmt.Printf("%T\n", c) //int32

	c1 := '国'
	fmt.Println(c1)        //22269 unicode编码十进制
	fmt.Printf("%T\n", c1) //int32
}
```



### 结构体的定义和实例化

```go
type 结构体名 struct{
  字段1 字段1的类型
  字段2 字段2的类型
  ...
}
```

```go
package main

import "fmt"

/*
结构体
Go语言中的基础数据类型可以表示一些事物的基本属性，但是当我们想表达一个事物的全部或部分属性时，
这时候再用单一的基本数据类型明显就无法满足需求了
Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，这种数据类型叫结构体，英文名称struct

Go语言内置的基础数据类型是用来描述一个值的，而结构体是用来描述一组值的

Go语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。
Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性

结构体实例化
	只有当结构体实例化时，才会真正地分配内存.也就是必须实例化后才能使用结构体的字段
	结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型
*/

// 定义一个结构体
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	// 结构体实例化
	// 声明一个person类型的变量p
	var p person
	// 通过字段赋值
	// 示例1
	p.name = "周林"
	p.age = 9000
	p.gender = "男"
	p.hobby = []string{"篮球", "足球", "双色球"}
	fmt.Println(p)         //{周林 9000 男 [篮球 足球 双色球]}
	fmt.Printf("%#v\n", p) //main.person{name:"周林", age:9000, gender:"男", hobby:[]string{"篮球", "足球", "双色球"}}
	// 访问变量p的字段
	fmt.Printf("%T\n", p) //main.person
	fmt.Println(p.name)   //周林

	// 示例2
	var p2 person
	p2.name = "理想"
	p2.age = 18
	fmt.Printf("type:%T -- value:%v\n", p2, p2) //type:main.person（表示main包下的person类型) -- value:{理想 18  []}

	// 匿名结构体：多用于临时场景
	var s struct {
		x string
		y int
	}
	s.x = "heiheihei"
	s.y = 100
	fmt.Printf("type:%T value:%v\n", s, s) //type:struct { x string; y int }     value:{heiheihei 100}
}
```



### 结构体指针

```go
package main

import (
	"fmt"
)

// 结构体是值类型
// 值类型和引用类型 https://www.cnblogs.com/hq82/p/11072005.html
// 值类型:int、float、bool、string、数组、struct
// 引用类型:指针、slice、map、chan、接口

// 定义一个结构体
type person struct {
	name, gender string
}

// go语言中 函数传参数 永远传的是拷贝
// 切片也是拷贝切片传递过去，但是传过去的拷贝的切片和原来的切片指向同一个底层数组，切片是引用类型
// 如果出现append()操作，可能会导致底层数组扩容，那么就会指向不同的底层数组了

func f(x person) {
	x.gender = "女" //结构体是值类型，相当于把值拷贝了一份传递过来 修改的是副本的gender
	//fmt.Println("==", x)
}

func f2(x *person) {
	//(*x).gender = "女" //根据内存地址找到那个变量，修改的就是原来的变量
	x.gender = "女" //语法糖，自动根据指针找对应的变量
}

func main() {
	var p person
	p.name = "周林"
	p.gender = "男"
	f(p)
	fmt.Println(p.gender) //男
	f2(&p)                //ox1241ac3 指针 相当于拷贝了p这个变量的指针传递到了f2的函数
	fmt.Println(p.gender) //女

	// 1.结构体指针
	var p2 = new(person) //new函数返回的是传递进去的类型的指针
	(*p2).name = "理想"
	p2.gender = "保密"
	fmt.Printf("%T\n", p2)   //*main.person
	fmt.Printf("%p\n", p2)   //0xc00000c080 p2保存的值就是一个内存地址
	fmt.Printf("%p\n", &p2)  //0xc00000e030 取p2这个变量本身的内存地址
	fmt.Printf("%#v\n", p2)  //&main.person{name:"理想", gender:"保密"}
	fmt.Printf("%#v\n", *p2) //main.person{name:"理想", gender:"保密"}
	fmt.Println(p2)          //&{理想 保密}
	fmt.Println(*p2)         //{理想 保密}

	// 2.结构体指针初始化
	// 2.1 key-value初始化
	var p3 = &person{
		name: "元帅",
	}
	fmt.Printf("%#v\n", p3) //&main.person{name:"元帅", gender:""}
	// 2.2 使用值列表的形式初始化，值的顺序和结构体定义时字段的顺序一致
	p4 := &person{
		"小王子",
		"男",
	}
	fmt.Printf("%#v\n", p4) //&main.person{name:"小王子", gender:"男"}
}
```



### 结构体内存布局

```go
package main

import (
	"fmt"
	"unsafe"
)

// 结构体内存布局
// 结构体占用一块连续的内存空间

type x struct {
	a int8 // 8bit -> 1byte 字节
	b int8
	c int8
}

func main() {
	m := x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}
	fmt.Printf("%p\n", &(m.a))    //0xc0000b4002
	fmt.Printf("%p\n", &(m.b))    //0xc0000b4003
	fmt.Printf("%p\n", &(m.c))    //0xc0000b4004
	fmt.Println(unsafe.Sizeof(m)) //3

	// 空结构体
	// 空结构体是不占用空间的
	var v struct{}
	fmt.Println(unsafe.Sizeof(v)) //0
}
```



### 结构体构造函数

```go
package main

import "fmt"

// 构造函数

// 定义一个结构体
type person struct {
	name string
	age  int
}

// 定义一个结构体
type dog struct {
	name string
}

// 构造函数：约定成俗用new
// 当结构体比较大的时候尽量使用结构体指针，减少程序的内存开销
func newPerson(name string, age int) *person { //返回值是person类型的指针
	return &person{
		name: name,
		age:  age,
	}
}

func newDog(name string) *dog {
	return &dog{
		name: name,
	}
}

func main() {
	p1 := newPerson("园林", 18)
	p2 := newPerson("周林", 9000)
	fmt.Println(p1, p2) //&{园林 18} &{周林 9000}

	d1 := newDog("周林")
	fmt.Println(d1) //&{周林}
}
```



### 结构体方法

```go
方法是作用于特定类型的函数.

方法的定义:(万变不离其宗)

func (接收者变量 接收者类型) 方法名(参数) 返回值{
    // 方法体
}

接收者通常使用类型首字母的小写,不建议使用诸如this和self这样的.
```

```go
package main

import (
	"fmt"
)

// 方法

// 标识符：变量名 函数名 类型名 方法名
// Go语言中如果标识符首字母是大写的，就表示对外部包可见（暴露的，共有的）

// dog 这是一个狗的结构体
type dog struct {
	name string
}

// person 人的结构体
type person struct {
	name string
	age  int
}

// newPerson 人结构体构造函数
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

// newDog 狗构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法是作用于特定类型的函数
// 接收者：表示的是调用该方法的具体类变量，多用类型名首字符小写表示
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪\n", d.name)
}

// 使用值接收者：传递拷贝进去
func (p person) guonian() {
	p.age++
}

// 指针接收者：传内存地址（指针）进去
func (p *person) zhenguonian() {
	p.age++
}

func (p *person) dream() {
	fmt.Println("不上班也能挣钱！")
}

func main() {
	d1 := newDog("zhoulin")
	d1.wang()

	p1 := newPerson("元帅", 18)
	//p1.wang()
	fmt.Println(p1.age) //18
	p1.guonian()
	fmt.Println(p1.age) //18
	p1.zhenguonian()
	fmt.Println(p1.age) //19
  p1.dream() //不上班也能挣钱！
}
```

```go
package main

import "fmt"

// 给自定义类型加方法
// 不能给别的包里面的类型添加方法，只能给自己包里的类型添加方法

type myInt int

func (m myInt) hello() {
	fmt.Println("我是一个int")
}

func main() {
	m := myInt(100)
	m.hello()
}
```

```go
package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student) //初始化map类型

	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	// 这里stu是个变量，在内存中内存地址是不变的
	// 改变的是stu所对应的值而已
	// m["小王子"] = stu的地址；
	// m["娜扎"] = stu的地址；
	// m["大王八"] = stu的地址；
	// stu地址里面的值最后是{name:"大王八",age:9000}
	for _, stu := range stus {
		fmt.Println("stu:", stu)
		fmt.Printf("stu:%p\n", &stu)
		m[stu.name] = &stu
		fmt.Printf("%#v\n", m)
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name, "=>", v.age)
	}
}

/*
stu: {小王子 18}
stu:0xc0000a6020
stu: {娜扎 23}
stu:0xc0000a6020
stu: {大王八 9000}
stu:0xc0000a6020
小王子 => 大王八
娜扎 => 大王八
大王八 => 大王八
*/
```



### 结构体常见问题总结

```go
package main

import (
	"fmt"
)

// 结构体遇到的问题

// 问题1. myInt(100)是个啥
type myInt int

func (m myInt) hello() {
	fmt.Println("我是一个int")
}

type person struct {
	name string
	age  int
}

func main() {
	//声明一个int32类型的变量x，它的值是10
	//方法1：
	//var x int32
	//x = 10
	//方法2：
	//var x int32 = 10
	//方法3：
	//var x = int32(10)
	//方法4
	//x := int32(10)
	//fmt.Println(x)

	//声明一个myInt类型的变量m，它的值是100
	//方法1:
	//var m myInt
	//m = 100
	//方法2:
	//var m myInt = 100
	//方法3:
	//var m = myInt(100)
	//方法4:
	//m := myInt(100)
	//fmt.Println(m)
	//m.hello()

	// 问题2：结构体初始化
	// 方法1：直接赋值
	var p person //声明一个p变量，他的数据类型是person
	p.name = "元帅"
	p.age = 18
	fmt.Println(p)
	var p1 person
	p1.name = "周林"
	p1.age = 9000
	fmt.Println(p1)
	// 方法2：键值对初始化
	var p2 = person{
		name: "冠华",
		age:  15,
	}
	fmt.Println(p2)
	// 方法3： 值列表初始化
	var p3 = person{"理想", 100}
	fmt.Println(p3)

	// 问题3：为什么要有构造函数
	p4 := newPerson("tom", 100)
	fmt.Println(p4)
	p5 := newPerson("jerry", 99)
	fmt.Println(p5)
}

// 问题3：为什么要有构造函数
func newPerson(name string, age int) person {
	// 别人调用我，我能给他一个person类型的变量
	return person{
		name: name,
		age:  age,
	}
}

// 返回一个指针
//func newPerson(name string, age int) *person {
//	//别人调用我，我能给他一个person类型的变量
//	return &person{
//		name: name,
//		age:  age,
//	}
//}
```



### 结构体匿名字段

```go
package main

import "fmt"

// 匿名字段
// 字段比较少也比较简单的场景
// 不常用！！！

// 匿名字段结构体
type person struct {
	string
	int
}

func main() {
	p1 := person{
		"周玲",
		9000,
	}
	fmt.Println(p1)        //{周玲 9000}
	fmt.Println(p1.string) //周玲
	fmt.Println(p1.int)    //9000
}
```



### 结构体嵌套

```go
package main

import "fmt"

// 结构体嵌套

type address struct {
	province string
	city     string
}

type workPlace struct {
	province string
	city     string
}

type person struct {
	name    string
	age     int
	address //匿名嵌套结构体
	//address:address // 命名嵌套结构体
	workPlace
}

type company struct {
	name string
	address
}

func main() {
	p1 := person{
		name: "周林",
		age:  90,
		address: address{
			province: "山东",
			city:     "威海",
		},
		workPlace: workPlace{
			province: "山东",
			city:     "济南",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.address.city)

	// 先在自己结构体找这个字段，找不到就去匿名嵌套的结构体中查找该字段
	// 如果匿名嵌套的结构体中有多个city，会报错 ambiguous selector p1.city
	//fmt.Println(p1.city)

	// 为了防止报错，可以用下面这种方法
	fmt.Println(p1.address.city)
	fmt.Println(p1.workPlace.city)
}
```



### 结构体模拟实现继承

```go
package main

import "fmt"

// 结构体模拟实现其他语言中的继承

type animal struct {
	name string
}

// 函数 func 函数名(参数) (返回值) {...}
// 方法 func (变量名 结构体) 函数名() 返回值() {...}
// 给animal实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

// 狗类 嵌套了动物结构体
type dog struct {
	feet   uint8
	animal // animal拥有的方法，dog此时也有了
}

// 给dog实现一个汪汪汪的方法
func (d dog) wang() {
	fmt.Printf("%s在叫: 汪汪汪~\n", d.name)
}

func main() {
	d1 := dog{
		feet: 4,
		animal: animal{
			name: "周林",
		},
	}
	fmt.Println(d1)
	d1.wang()
	d1.move()
}
```



### 结构体与JSON

```go
package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与json
// 反射

// 1.序列化：把go语言中的结构体变量 ---> json格式的字符串
// 2.反序列化：json格式的字符串 ---> go语言中能够识别的结构体变量

// 这里的字段名称开头要大写，是因为要把这个字段名称传入到第三方的包中，如果小写就不能暴露字段了
// 但是序列化和反序列化以后，可能是小写的，所以需要后面的json标记
// 如果序列化和反序列化后都是大写开头，那么久不需要json标记也可以
type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age"`
}

//type person struct {
//	Name string
//	Age  int
//}

func main() {
	// 序列化
	p1 := person{
		Name: "周玲",
		Age:  9000,
	}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v", err)
		return
	}
	fmt.Printf("%v\n", string(b))

	// 反序列化
	//str := `{"Name":"理想", "Age":18}`
	str := `{"name":"理想", "age":18}`
	var p2 person
	err = json.Unmarshal([]byte(str), &p2)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v", err)
		return
	}
	fmt.Printf("%#v\n", p2)
}

```



### 案例

```go
package main

import (
	"fmt"
	"os"
)

/*
函数版学生管理系统
写一个系统能够查看，新增学生，删除学生
*/

var (
	allStudent map[int64]*student //声明全局变量并初始化，好让所有函数都能调用修改
)

type student struct {
	id   int64
	name string
}

//newStudent 是student类型的构造函数
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showAllstudent() {
	//把所有的学生都打印出来
	for k, v := range allStudent {
		fmt.Printf("学号:%d 姓名:%s\n", k, v.name)
	}
}

func addStudent() {
	//向allStudent中添加一个新的学生
	//1.创建一个新学生
	//1.1获取用户输入
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生学号: ")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名: ")
	fmt.Scanln(&name)
	//1.2造学生（调用student的构造函数)
	newStu := newStudent(id, name)
	//2.追加到allStudent这个map中
	allStudent[id] = newStu
}

func deleteStudent() {
	//1.请用户输入要删除的学生序号
	var (
		deleteID int64
	)
	fmt.Print("请输入学生学号: ")
	fmt.Scanln(&deleteID)
	//2.去allStudent这个map中根据学号删除对应的键值对
	delete(allStudent, deleteID)
}

func main() {
	allStudent = make(map[int64]*student, 48) //初始化，开辟内存空间，防止在运行过程中动态扩容影响性能

	for {
		//1.打印菜单
		fmt.Println("\t*******欢迎光临学生管理系统********")
		fmt.Println(`
		1.查看所有学生
		2.新增学生
		3.删除学生
		4.退出
		`)

		//2.等待用户选择要做什么
		fmt.Print("请输入编号: ")
		var choice string
		//var choice int //这里可以声明为整型或者字符串，建议整型，后面可以少写点代码
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%s这个选项！\n", choice)

		//3.执行对应的函数
		switch choice {
		case "1":
			showAllstudent()
		case "2":
			addStudent()
		case "3":
			deleteStudent()
		case "4":
			os.Exit(1)
		default:
			fmt.Println("输入有误")
		}
	}
}

```



### 练习

```go
package main

import (
	"fmt"
	"os"
)

/*
结构体方法版学生管理系统
写一个系统能够查看，新增学生，删除学生
*/

// 学生信息结构体
type student struct {
	id   int
	name string
}

// 存储所有学生结构体
type class struct {
	Map map[int]*student
}

// 增加
func (c *class) addStudent() {
	var (
		id   int
		name string
	)
	fmt.Printf("input ID:")
	fmt.Scanln(&id)
	fmt.Printf("input name:")
	fmt.Scanln(&name)
	stu := &student{
		id:   id,
		name: name,
	}

	c.Map[id] = stu
}

// 查看
func (c *class) showStudent() {
	//fmt.Println(c.Map)
	for index, value := range c.Map {
		fmt.Printf("ID:%d Name:%s\n", index, value.name)
	}
}

// 删除
func (c *class) delStudent() {
	var id int
	fmt.Printf("del ID:")
	fmt.Scanln(&id)
	delete(c.Map, id)
	c.showStudent()
}

func main() {
	// 初始化
	c := &class{}
	c.Map = make(map[int]*student, 60)

	for {
		fmt.Printf("\t\t--------system admin --------")
		fmt.Println(`
			1.add student
			2.show students
			3.del student
			4.exit
 		`)

		var choice int
		fmt.Print("input ID: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			c.addStudent()
		case 2:
			c.showStudent()
		case 3:
			c.delStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("ID Error!")
		}
	}
}

```
