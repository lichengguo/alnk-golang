package main

import "fmt"

// Go语言中函数的return不是原子操作 在底层是分为两步来执行
// 第一步：返回值赋值
// defer
// 第二步：真正的RET返回
// 1.返回值赋值 2.defer 3.真正的return指令
// 函数中如果存在defer，那么defer执行的时机是在第一步和第二步之间
// 如果不能理解，可以参考这篇博文
// https://www.jianshu.com/p/79c029c0bd58

func main() {
	fmt.Println("f1: ", f1()) // f1:  5
	fmt.Println("f2: ", f2()) // f2:  6
	fmt.Println("f3: ", f3()) // f3:  5
	// fmt.Println("f4: ", f4()) // f4:  5
	fmt.Println("f5: ", f5()) // f5:  5
	fmt.Println("f6: ", f6()) // f6:  6
}

// 匿名返回值
// 注意这里：返回变量是匿名的（举例比如是：niming这个变量），而不是x
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
	return 5 // 返回值变量是x
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

// 在go中函数的参数默认为 按值传递，即在函数内部修改传入参数的值是函数外部传入值的 拷贝
// 如果想要使用引用传递，需要将传入的参数设置为 指针类型
// 如果传入的参数数据很大，建议使用指针类型，减少内存因拷贝参数而占用
// func f4() (x int) {
// 	defer func(x int) {
// 		x++                  // 改变的是函数中x的副本
// 		fmt.Println("==", x) // 1
// 	}(x) // x=0 值传递
// 	return 5 // 返回值 =x =5
// }

// 在go中函数的参数默认为 按值传递，即在函数内部修改传入参数的值是函数外部传入值的 拷贝
// 如果想要使用引用传递，需要将传入的参数设置为 指针类型
// 如果传入的参数数据很大，建议使用指针类型，减少内存因拷贝参数而占用
func f5() (x int) {
	// 代码从上往下执行
	// 在函数中返回值返回值 (x int) 已经被定义了，只是没有复制，此时为0
	fmt.Println("f5 x:", x) // f5 x: 0
	x = 1
	defer func(x int) int {
		fmt.Println("f5 defer=", x) // f5 defer= 1
		x++
		fmt.Println("f5 defer==", x) // f5 defer== 2
		return x
	}(x) // x=1
	return 5
}

// 1.返回值=x=5
// 2.defer x=6
// 3.RET返回 x=6
// 传一个x的指针到匿名函数中
func f6() (x int) {
	fmt.Println("f6 x:", x)
	defer func(x *int) {
		fmt.Println("defer f6 x: ", *x)
		*x++
		fmt.Println("defer f6 xx: ", *x)
	}(&x) // 传的是指针，不是值拷贝，底层共用一个数据
	return 5
}
