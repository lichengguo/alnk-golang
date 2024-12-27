package main

import (
	"fmt"
)

// 方法

// 标识符：变量名 函数名 类型名 方法名
// Go语言中如果标识符首字母是大写的 就表示对外部包可见

// dog 这是一个狗的结构体
type dog struct {
	name string
}

// newDog 狗构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法是作用于特定类型的函数
// 接收者：表示的是调用该方法的具体类变量 多用类型名首字符小写表示
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪\n", d.name)
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
	d1 := newDog("tom")
	d1.wang()

	p1 := newPerson("alnk", 18)
	fmt.Println(p1.age) // 18
	p1.guonian()
	fmt.Println(p1.age) // 18
	p1.zhenguonian()
	fmt.Println(p1.age) // 19
	p1.dream()          // 不上班也能挣钱！
}
