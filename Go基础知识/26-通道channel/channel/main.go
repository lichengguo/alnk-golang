package main

import (
	"fmt"
	"sync"
)

/*
[为什么要使用channel?]
单纯地将函数并发执行是没有意义的，函数与函数间需要交换数据才能体现并发执行函数的意义
虽然可以使用共享内存进行数据交换，但是共享内存在不同的goroutine中容易发生竞态问题
为了保证数据交换的正确性，必须使用互斥量对内存进行加锁，这种做法势必造成性能问题

[channel] 别名:管道、通道
通道channel是一种类型，一种引用类型，通道类型的空值是nil
通道channel必须要使用 make函数 初始化以后才能使用。和slice，map一样

Go语言的并发模型是CSP，提倡通过 通信共享内存; 而不是 通过共享内存而实现通信

如果说goroutine是Go程序并发的执行体，那么channel就是它们之间的连接
channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制

Go 语言中的通道（channel）是一种特殊的类型
通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。
每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型

[channel操作]
1.发送：发送数字1到ch1通道  ch1 <- 1  写入
2.接收：x读取通道ch1中的数字1  x := <- ch1 读取
3.关闭：close(ch1)
*/

var a []int
var b chan int //需要指定通道中元素的类型
var wg sync.WaitGroup

// noBufChannel 无缓冲区的通道
func noBufChannel() {
	fmt.Println("无缓冲区通道b: ", b) //nil
	b = make(chan int)          //初始化，但是不带缓冲区。如果没有设置缓冲区，那么必须先有读取者，才能往通道中写入数据

	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b //从通道b中读取数据
		fmt.Println("无缓冲区 后台goroutine从通道b中取到了", x)
	}()

	b <- 10                           //10 写入到通道中。如果没有上面的匿名函数接收数据，那么此处会hang住
	fmt.Println("无缓冲区 10发送到通道b中了...") //后台goroutine从通道b中取到了 10

	wg.Wait()
}

// bufChannel 有缓冲区通道
func bufChannel() {
	fmt.Println(b)         //nil
	b = make(chan int, 10) //带有缓冲区的通道,容量为10

	b <- 10 //10写入到通道
	fmt.Println("10发送到通道b中了...")

	b <- 20 //20写入到通道
	fmt.Println("20发送到通道b中了...")

	x := <-b
	fmt.Println("从通道中取到了", x)

	x = <-b
	fmt.Println("从通道中取到了", x)

	close(b)
}

func main() {
	//noBufChannel()
	bufChannel()
}
