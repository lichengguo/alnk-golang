package main

import "fmt"

/*
[为什么要使用channel]
单纯地将函数并发执行是没有意义的，函数与函数间需要交换数据才能体现并发执行函数的意义
虽然可以使用共享内存进行数据交换，但是共享内存在不同的goroutine中容易发生竞态问题
为了保证数据交换的正确性，必须使用互斥量对内存进行加锁，这种做法势必造成性能问题


[channel] 别名:管道、通道
通道channel是一种类型，一种引用类型，通道类型的空值是nil
通道channel必须要使用 make函数 初始化以后才能使用 和slice map一样
Go语言的并发模型是CSP，提倡 通过通信共享内存; 而不是 通过共享内存而实现通信
如果说goroutine是Go程序并发的执行体，那么channel就是它们之间的连接
channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制
Go 语言中的通道（channel）是一种特殊的类型
通道像一个传送带或者队列，总是遵循先入先出的规则，保证收发数据的顺序
每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型


[channel操作]
1.发送：发送数字1到ch1通道  ch1 <- 1  写入
2.接收：x读取通道ch1中的数字1  x := <- ch1 读取
3.关闭：close(ch1)
*/

func main() {
	fmt.Println("------- 无缓冲通道 -----")
	// 创建一个无缓冲的通道
	ch := make(chan int)

	// 启动一个goroutine将值发送到通道
	go func() {
		// 如果没有设置缓冲区，那么必须先有读取者，才能往通道中写入数据
		// 不然此处会hang住，所以使用goroutine能够让程序继续往下运行
		ch <- 10
	}()

	// 从通道中接收值并打印
	value := <-ch
	fmt.Printf("接收来自ch通道的值: %d\n", value)

	fmt.Println("------- 有缓冲通道 -----")
	// 创建一个有缓冲的通道，容量为3
	buffCh := make(chan string, 3)

	// 将值发送到有缓冲的通道
	// buffCh <- "hello"
	// buffCh <- "from"
	// buffCh <- "chan"
	// buffCh <- "alnk" // 第4个值会hang住，写在main主程序里会死锁
	go func() {
		buffCh <- "hello"
		buffCh <- "from"
		buffCh <- "chan"
		buffCh <- "alnk" // 第4个值会hang住
	}()

	// 从有缓冲的通道接收值并打印
	fmt.Println(<-buffCh)
	fmt.Println(<-buffCh)
	fmt.Println(<-buffCh)
}
