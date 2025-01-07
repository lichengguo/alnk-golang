package main

import (
	"fmt"
	"sync"
)

/*
sync.Once
在编程的很多场景下我们需要确保某些操作在高并发的场景下只执行一次，例如只加载一次配置文件、只关闭一次通道等

Go语言中的sync包中提供了一个针对只执行一次场景的解决方案sync.Once

sync.Once只有一个Do方法，其签名如下：
func (o *Once) Do(f func()) {}
*/

var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1) //关闭通道ch1
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break //这里不能用return，要用break结束循环，然后调用wg.Done()，不然会出现死锁
		}
		ch2 <- x * 2
	}

	// 多次关闭通道会报错
	//func() {
	//	close(ch2)
	//}()

	//确保只关闭一次通道ch2
	once.Do(func() {
		close(ch2)
	})
}

func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)

	wg.Add(3)
	go f1(a)
	go f2(a, b)
	go f2(a, b)
	wg.Wait()

	for ret := range b {
		fmt.Println(ret)
	}
}
