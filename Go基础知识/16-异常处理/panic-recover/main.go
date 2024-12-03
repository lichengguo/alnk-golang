package main

import "fmt"

// panic 和 recover
// 异常处理

// 注意
// 1.recover()必须搭配defer使用
// defer一定要在可能引发panic的语句之前定义

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
	panic("程序出现了严重的BUG!!!!") // 程序崩溃退出
	// fmt.Println("this is funcB")
}

func funcC() {
	fmt.Println("this is funcC")
}

func main() {
	funcA()
	funcB()
	funcC()
}
