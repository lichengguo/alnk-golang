package main

import "fmt"

/*
https://www.liwenzhou.com/posts/Go/12_interface/
[接口]
接口（interface）定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节

在Go语言中接口（interface）是一种类型，一种抽象的类型，interface是一组method(方法)的集合

接口做的事情就像是定义一个协议（规则），只要一台机器有洗衣服和甩干的功能，我就称它为洗衣机。不关心属性（数据），只关心行为（方法）。
为了保护你的Go语言职业生涯，请牢记接口（interface）是一种类型

[为什么要使用接口？]
type Cat struct{}
func (c Cat) Say() string { return "喵喵喵" }

type Dog struct{}
func (d Dog) Say() string { return "汪汪汪" }

func main() {
	c := Cat{}
	fmt.Println("猫:", c.Say())
	d := Dog{}
	fmt.Println("狗:", d.Say())
}
上面的代码中定义了猫和狗，然后它们都会叫，你会发现 main 函数中明显有重复的代码，
如果我们后续再加上猪、青蛙等动物的话，我们的代码还会一直重复下去。
那我们能不能把它们当成“能叫的动物”来处理呢？
Go语言中为了解决类似上面的问题，就设计了接口这个概念。
接口区别于我们之前所有的具体类型，接口是一种抽象的类型。
当你看到一个接口类型的值时，你不知道它是什么，唯一知道的是通过它的方法能做什么。

Go语言提倡面向接口编程
每个接口由数个方法组成，接口的定义格式如下
type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
其中：
	接口名：使用type将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加er，
           如有写操作的接口叫Writer，有字符串功能的接口叫Stringer等。
           接口名最好要能突出该接口的类型含义。
	方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
	参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略


[实现接口的条件]
一个对象只要 全部实现 了接口中的方法，那么就实现了这个接口。换句话说，接口就是一个需要实现的方法列表
如果接口中有没被实现的方法，那么就会报错.
结构体实现接口中的方法，要么全部实现，要么就一个别实现，否则报错

[接口类型变量]
那实现了接口有什么用呢？
接口类型变量能够存储所有实现了该接口的实例。 例如上面的示例中，Sayer类型的变量能够存储dog和cat类型的变量。
*/

// 示例
// 由于猫、狗、人都能叫
// 所以定义一个能叫的接口
type speaker interface {
	speak() //只要实现了speak方法的变量都是speaker类型, 方法签名
}

//定义结构体
type cat struct{}

type dog struct{}

type person struct{}

// 定义相对应的方法，实现接口
func (c cat) speak() {
	fmt.Println("喵喵喵~")
}

func (d dog) speak() {
	fmt.Println("旺旺旺~")
}

func (p person) speak() {
	fmt.Println("啊啊啊~")
}

// 定义一个函数
func da(x speaker) {
	// 接收一个参数,传进来什么,我就打什么
	x.speak() // 挨打了就要叫
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person

	da(c1)
	da(d1)
	da(p1)

	var ss speaker // 定义一个接口类型:speaker 的变量:ss
	ss = c1
	fmt.Printf("%T\n", ss) //main.cat
	ss = d1
	fmt.Printf("%T\n", ss) //main.dog
	ss = p1
	fmt.Printf("%T\n", ss) //main.person
}
