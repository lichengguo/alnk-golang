package main

import "fmt"

/*
[close() 关闭通道]
关于关闭通道需要注意的事情
只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道
(如果一个通道没有关闭，那么通道中数据读取到最后会出现死锁，如果关闭了则不会)
通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的，在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的

关闭后的通道有以下特点：
	对一个关闭的通道再发送值就会导致panic
	对一个关闭的通道进行接收会一直获取值直到通道为空
	对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值
	关闭一个已经关闭的通道会导致panic
*/

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	close(ch1)
	//for x := range ch1 {
	//	fmt.Println(x)
	//}
	<-ch1
	<-ch1

	x, ok := <-ch1
	fmt.Println(x, ok) //0 false 关闭通道以后，通道的值如果取完了，还可以再取，但是返回的是该类型的零值和一个false

	x, ok = <-ch1
	fmt.Println(x, ok)

	x, ok = <-ch1
	fmt.Println(x, ok)
}
