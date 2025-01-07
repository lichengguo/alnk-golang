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
	panic("goroutine hello挂了")

}

func main() {
	go hello()
	for {
		fmt.Println("我是主goroutine,我还没挂")
		time.Sleep(time.Second * 1)
	}
}
