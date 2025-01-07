package main

import "fmt"

// point 指针
// go语言中不会直接进行指针的运行 因此只要会使用 & * 基本上就够了

// make和new的区别
// make和new都是用来申请内存的
// new很少用 一般用来给基本数据类型申请内存 string int返回的是对应类型的指针(*string、*int)
// make是用来给slice map chan申请内存的 make函数返回的的是对应的这三个类型本身

func main() {
	fmt.Println("------------- & 取内存地址 ------------")
	n := 18
	p := &n               // &n取n变量在内存中的地址，并且赋值给p变量
	fmt.Println(p)        // 0xc00001a070 n变量在内存中的地址
	fmt.Println(&n)       // 0xc00001a070 n变量在内存中的地址
	fmt.Println(*p)       // 取值 18
	fmt.Println(*(&n))    // 18
	fmt.Printf("%T\n", p) // *int：int类型的指针
	fmt.Println(&p)       // 0xc00000e028 p变量自己在内存中的地址
	fmt.Println(*(&p))    // 0xc00001a070 取p变量指向的值，即n的内存地址 18

	fmt.Println("---------- *:根据地址取值 -----------------")
	m := *p                // 这个是取出p变量中存放的n的内存地址所指向的值
	fmt.Println(m)         // 18
	fmt.Printf("%T\n", m)  // int
	pn := &p               // &p p变量自己在内存中的地址
	fmt.Printf("%p\n", pn) // 0xc00008c018 p变量自己在内存中的地址

	fmt.Println("----------------------------------")
	var a1 *int     // nil pointer 未申请内存地址
	fmt.Println(a1) // <nil>

	var a2 = new(int) // new函数申请一块内存地址
	fmt.Println(a2)   // 0xc00001a078 这是一块内存地址
	fmt.Println(*a2)  // 获取这块内存地址指向的值是多少 //0
	*a2 = 100         // 给这块内存地址赋值
	fmt.Println(*a2)  // 100
	fmt.Println(&a2)  // 0xc0000ae028 a2在内存中的地址
}
